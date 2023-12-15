//
//  WorkoutListView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/26/23.
//

import SwiftUI
import SentrySwiftUI

struct CompetitionWorkoutListView: View {
  
  @EnvironmentObject var data: Datamodel
  
  @State private var showingImportView: Bool = false
  
  var forComp: UUID?
  
  var body: some View {
//    NavigationStack {
      Group {
        if data.myCompletedWorkouts.isEmpty {
          VStack {
            Spacer()
            HStack {
              Spacer()
              Text("No workouts yet, use your watch to start one!")
              Spacer()
            }
            Spacer()
          }
          .DemoiosBackgroundColor()
        } else {
          
          
          if forComp != nil && !data.myCompletedWorkouts.contains(where: { $0.workout_data?.competition?.id == forComp }) {
            VStack(alignment: .center) {
              Spacer()
              HStack {
                Spacer()
                Text("No workouts completed yet for this competition! Get out there!")
                  .DemoiosText(fontSize: .large)
                Spacer()
              }
              Spacer()
            }
//            .DemoiosBackgroundColor()
          } else {
            ScrollView {

              ForEach(data.myCompletedWorkouts.sorted(by: { $0.workout_data?.healthkit_workout_end_date ?? "" > $1.workout_data?.healthkit_workout_end_date ?? "" }).filter { $0.workout_data?.competition?.id == forComp }) { workout in
                NavigationLink(value: workout) {
                  
                  WorkoutListItem(workout: workout)
                    .padding(.horizontal)
                }
              }
              
              
            }
            .DemoiosBackgroundColor()
            .navigationDestination(for: Workout.self) { workout in
              WorkoutSummaryView(workout: workout)
            }
          }
          
          
        }
      }
      .navigationBarTitleDisplayMode(.inline)
    
    //            .sentryTrace("Workout List View - Top Level")
    
      .navigationTitle("Your Workouts")
    //      .toolbar {
      //        ToolbarItem {
      //          Button("Import") {
      //            showingImportView.toggle()
      //          }
      //        }
      //      }
      .task {
        data.health.requestHealthKitAuthorization { success in
          print("Got auth for healthkit? \(success)")
        }
      }
//    }
    .task {
      if data.myCompletedWorkouts.isEmpty {
        await data.api.getMyCompletedWorkouts()
        //                syncDemoiosWorkoutsWithAPI(data: data)
      }
    }
    .refreshable {
     await data.api.getMyCompletedWorkouts()
      //            syncDemoiosWorkoutsWithAPI(data: data)
    }
    .sheet(isPresented: $showingImportView) {
//      NavigationStack {
        WorkoutImportView()
//      }
      //                .environment(\.editMode, Binding.constant(EditMode.active))
    }
  }
}

struct WorkoutListView_Previews: PreviewProvider {
  static var previews: some View {
    CompetitionWorkoutListView(forComp: Competition.test_ended_joined.id)
      .environmentObject(Datamodel())
      .setupPreview()
    CompetitionWorkoutListView()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}

//
//  AddWorkoutSheet.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/1/23.
//

import SwiftUI

struct AddWorkoutSheet: View {
  @Environment(\.dismiss) var dismiss
  
  @EnvironmentObject var data: Datamodel
  
  @Binding var comp: Competition
  @Binding var leaderboardData: [CompetitionLeaderboardItem]
  @Binding var graphData: [CompetitionStatGraphDataPoint]
  @Binding var profileParticipants: [PartialProfile]
  @Binding var myStats: CompMyStats?
  
  //  @Binding var searchText: String
  //  @Binding var selectedResults: Set<PartialProfile>
  @State var applicableWorkouts: [Workout] = []
  
  @State var selectedWorkouts: [Workout] = []
  
  @State var loaded: Bool = false
  @State var err: APIError?
  
  var body: some View {
    //    NavigationStack {
    ScrollView {
      
      if err == nil && loaded {
        VStack {
          VStack(alignment: .leading) {
            Text("Select which workouts you would like to apply to this competition.")
              .DemoiosText(fontSize: .small, color: .graySix)
            Text("⚠️ Once you apply a workout, you cannot apply it to another competition!")
              .DemoiosText(fontSize: .small, color: .graySix)
          }
          Divider()
          Text("\(applicableWorkouts.count) applicable workouts")
            .DemoiosText(fontSize: .base, color: .graySix, weight: .bold)
          
          ForEach(applicableWorkouts.sorted(by: { ($0.workout_data?.healthkit_workout_end_date ?? Date.now.intoString) ?? "" > $1.workout_data?.healthkit_workout_end_date ?? Date.now.intoString ?? "" })) { w in
            
            Button {
              
              if selectedWorkouts.contains(where: { $0.id == w.id }) {
                // already in, REMOVE
                selectedWorkouts.removeAll(where: { $0.id == w.id })
              } else {
                // APPEND
                selectedWorkouts.append(w)
              }
              
            } label: {
              
              WorkoutListItem(workout: w)
                .overlay(alignment: .bottomTrailing) {
                  if selectedWorkouts.map({ $0.id }).contains(w.id) {
                    HStack {
                      Spacer()
                      Image(systemName: "circle.fill")
                    }
                    .DemoiosText(fontSize: .base, color: .orange)
                    .padding(.trailing)
                    .padding(.bottom)
                  }
                }
            }
            .padding(.horizontal)
            
          }
        }
      } else if loaded && err != nil {
        Text("Failed to load applicable workouts")
      } else {
        ProgressView()
      }
      
    }
    .navigationTitle("Select workouts to apply")
    .navigationBarTitleDisplayMode(.inline)
    .toolbar {
      ToolbarItem(placement: .primaryAction) {
        Button("Done") {
          
          if !selectedWorkouts.isEmpty {
            
            let ids = selectedWorkouts.compactMap { $0.workout_data?.id }
            
            Task {
              do {
                try await data.api.applyWorkoutsToComp(data: data, compId: comp.id, workoutIds: ids)
                
                triggerLocalNotification(title: "Workouts added!", body: "", background: .green, icon: .ok)
                
                let (members, gd, cl, ms) = await data.api.refreshCompData(comp.id)
                profileParticipants = members
                graphData = gd
                leaderboardData = cl
                myStats = ms
                await data.api.getMyCompletedWorkouts()
                
              } catch {
                triggerLocalNotification(title: "Failed to add workouts", body: "", background: .red, icon: .warning)
              }
            }
          }
          dismiss()
        }
      }
    }
    .padding(.bottom)
    .onDisappear {
      dispatchAsync {
        self.loaded = false
        self.applicableWorkouts = []
      }
    }
    .task {
      do {
        let workouts: [Workout] = try await data.api.getApplicableWorkoutsForCompetition(comp_id: comp.id)
        loaded = true
        applicableWorkouts = workouts
      } catch {
        
        loaded = true
        err = .invalidResponse
        triggerLocalNotification(title: "Failed to get Applicable workouts", body: "", background: .red, icon: .warning)
      }
      
    }
  }
  
  
}

struct AddWorkoutSheet_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    
    
    NavigationStack {
      //      AddWorkoutSheet(comp: $comp, leaderboardData: $leaderboardData, graphData: $graphData, profileParticipants: $profileParticipants, myStats: $myStats)
      AddWorkoutSheet(comp: .constant(.test_ended_joined), leaderboardData: .constant(CompetitionLeaderboardItem.testData), graphData: .constant(CompetitionStatGraphDataPoint.testData), profileParticipants: .constant([.default]), myStats: .constant(nil))
        .onAppear {
          
        }
        .setupPreview()
        .environmentObject(data)
    }
    .environment(\.colorScheme, .dark)
    
  }
}

//
//  WorkoutSummaryView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/26/23.
//

import SwiftUI
import CoreLocation

struct WorkoutSummaryView: View {
  @EnvironmentObject var data: Datamodel
  var state: WorkoutState = .summary
  @State var workout: Workout
  @State var otherUserWorkoutInstances: [WorkoutData] = []
  @State var graphData: [WorkoutGraphDataPoint] = []
  
  @State var workoutRouteLocations: [CLLocation] = []
  
  @State var doneGatheringRouteCoords: Bool = false
  
  var items: [GridItem] = Array(repeating: GridItem(.flexible()), count: 4)
  
  var body: some View {
    ScrollView {
      VStack {
        HStack {
          
          if let competition = workout.workout_data?.competition {
            HStack {
              VStack(alignment: .leading) {
                Text(competition.title)
                  .DemoiosText(fontSize: .base)
                
                Text("Competition")
                  .DemoiosText(fontSize: .small, color: .primarySeven, weight: .light)
              }
              .padding(.top, 2)
//              Spacer()
            }
          }
          
          Spacer()
          
          HStack {
            // MARK: - Weather group
            if let symbol = workout.workout_data?.weather?.symbol, let tempF = workout.workout_data?.weather?.temp_with_unit, let city = workout.workout_data?.weather?.location_city {
              HStack {
                Image(systemName: symbol)
                Text("\(tempF) in \(city)")
              }
            }
            // MARK: - weather group
          }
          
        }
        
        
        if let wkData = workout.workout_data {
          Divider()
          VStack(spacing: 0) {
            
                SectionTitle(title: "Stats")
            
            LazyVGrid(columns: items) {
              WorkoutDataStatItem(title: "Time", value: wkData.duration)
              WorkoutDataStatItem(title: "Distance", value: "\(wkData.miles)")
              WorkoutDataStatItem(title: "Avg HR", value: wkData.formattedHR)
              WorkoutDataStatItem(title: "Calories", value: wkData.calories)
            }
            
          }
        }
        
        Divider()
        VStack {
          if !otherUserWorkoutInstances.isEmpty {
            VStack(spacing: 8) {
              
              SectionTitle(title: "Invited Users Workouts")
              
              LazyHStack {
              ScrollView(.horizontal) {
               
                  ForEach(otherUserWorkoutInstances) { workoutDataInstance in
                    WorkoutInstanceCard(workoutData: workoutDataInstance)
                  }
                }
              }
              
            }
          }
          if !graphData.isEmpty {
            VStack(spacing: 8){
              SectionTitle(title: "Heartrate Data")
              WorkoutStatGraph(graphData: graphData)
            }
          }
          if !workoutRouteLocations.isEmpty && doneGatheringRouteCoords {
            MapView(route: workoutRouteLocations)
              .padding(.top)
          }
        }
      }
      .onAppear {
        workoutRouteQuery(for: .heartRate, fromWorkout: UUID(uuidString: workout.workout_data?.healthkit_workout_id ?? "")) { graphRes in
          dispatchAsync {
            self.graphData = graphRes
          }
          
        }
        
        healthStoreWorkoutRouteQuery(fromWorkout: UUID(uuidString: workout.workout_data?.healthkit_workout_id ?? "")) { locations in
          dispatchAsync {
            self.workoutRouteLocations.append(contentsOf: locations)
          }
        } finished: { done in
          if done {
            doneGatheringRouteCoords = true
          }
        }
      }
      .task {
        if getCurrentEnvType() != .preview {
          otherUserWorkoutInstances = await data.api.getWorkoutDataByWorkoutId(workoutId: workout.id)
        }
      }
      .padding(.horizontal)
      
    }
    .navigationTitle(workout.full_title)
    .DemoiosBackgroundColor()
  } // toolbar
}

struct WorkoutSummaryView_Previews: PreviewProvider {
  
  
  static var previews: some View {
    
    let data = Datamodel()
    
    var c1: Workout = .defaultWorkout
    
    NavigationStack {
      WorkoutSummaryView(workout: c1, otherUserWorkoutInstances: [WorkoutData.test, WorkoutData.test2], graphData: Workout.testGraphData)
        .environmentObject(WeatherManager())
        .environmentObject(data)
        .environmentObject(LocationManager())
    }
    .setupPreview()
    .previewDisplayName("With graph data")
    
    NavigationStack {
      WorkoutSummaryView(workout: c1, otherUserWorkoutInstances: [], graphData: [])
        .environmentObject(WeatherManager())
        .environmentObject(data)
        .environmentObject(LocationManager())
      
      
    }
    .setupPreview()
    .previewDisplayName("no graph data")
    
  }
}

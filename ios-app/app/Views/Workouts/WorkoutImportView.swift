//
//  WorkoutImportView.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/7/23.
//

import SwiftUI

struct WorkoutImportView: View {
  @Environment(\.dismiss) private var dismiss
  @Environment(\.editMode) private var editMode
  
  @EnvironmentObject var data: Datamodel
  
  @State private var importableWorkouts: [ImportableWorkout] = []
  
  @State private var selectedImportableWorkouts: [ImportableWorkout] = []
  @State var successes: Int = 0
  @State var failures: Int = 0
  @State var notiBody: String = ""
  var body: some View {
    // for each workout in the last month where the source is NOT Demoios, and the id from the workout is not in the list of
    //      NavigationStack {
    ScrollView {
      ForEach(importableWorkouts, id: \.self) { workout in
        Button {
          if selectedImportableWorkouts.map({ $0.healthkit_workout_id }).contains(workout.healthkit_workout_id) {
            selectedImportableWorkouts.removeAll(where: { $0.healthkit_workout_id == workout.healthkit_workout_id })
          } else {
            selectedImportableWorkouts.append(workout)
          }
        } label: {
          
          WorkoutImportListItem(workout: workout)
            .padding(.horizontal)
            .overlay(alignment: .bottomTrailing) {
              if selectedImportableWorkouts.map({ $0.healthkit_workout_id }).contains(workout.healthkit_workout_id) {
                Image(systemName: "circle.fill")
                  .foregroundColor(.orange)
                  .padding(.trailing)
              }
              
            }
        }
      }
    }
    .navigationTitle("Select workouts to import")
    .navigationBarTitleDisplayMode(.inline)
    .toolbar {
      Button("Finish") {
        if !selectedImportableWorkouts.isEmpty {
          Task {
            for workout in selectedImportableWorkouts {
              do {
                try await data.api.importWorkout(payload: workout.toDictionary())
                
                triggerLocalNotification(title: "Imported \(selectedImportableWorkouts.count)", body: "", background: .green)
              } catch {
                triggerLocalNotification(title: "Failed to import the workouts", body: "Please try again", background: .red)
              }
            } // importable workout loop
            
            await data.api.getMyCompletedWorkouts()
          } // task wrapper
        } // if selectedWorkouts!= 0
  
        dismiss()
      }
    }
    .onAppear {
      listImportableWorkouts(data: data) { (workouts, error) in
        guard let workouts = workouts, error == nil else {
          print("Could not list importable workouts in on appear")
          return
        }
        
        importableWorkouts = workouts
      }
    }
  }
}

struct WorkoutImportView_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutImportView()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}
//Task {
//
//  for workout in selectedImportableWorkouts {
//
//    data.api.importWorkout(payload: workout.toDictionary()) { (result: Result<Workout, APIError>) in
//
//
//      switch result {
//      case .success(_):
//        successes += 1
//      case .failure(_):
//        failures += 1
//      }
//      notiBody = "Successfully imported \(successes) workouts"
//      if failures != 0 {
//        notiBody += " and failed to import \(failures)."
//      }
//    }
//  }
//  dispatchAsync {
//    NotificationHandler.shared.new(title: "Finished importing workouts!", body: notiBody, background: .blue)
//    successes = 0
//    failures = 0
//  }
//  DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
//    data.api.getMyCompletedWorkouts()
//  }
//
//}

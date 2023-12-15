//
//  WorkoutImportListItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/7/23.
//

import SwiftUI

struct WorkoutImportListItem: View {
    var showLocationList = ["Walking", "Running", "Cycling"]
    
    let workout: ImportableWorkout
    
  var body: some View {
    SectionItem(fixedBy: .v) {
    HStack {
      if workout.source == "Demoios" {
        Logo(size: 50, logoColor: .white)
          .padding(.leading)
      } else {
        Image(systemName: "square.and.arrow.down")
          .foregroundColor(.white)
          .padding(.leading)
      }
      
      VStack(alignment: .leading) {
        Text((Consts.cardioActivityList.contains(workout.healthkit_workout_activity_type) ? "\(workout.location_type.capitalized) " : "") + "\(workout.healthkit_workout_activity_type)")
          .DemoiosText(fontSize: .base)
          .multilineTextAlignment(.leading)
        Text("Workout")
          .DemoiosText(fontSize: .xxs, color: .graySix)
      }
      
      Spacer()
      
      //            if workout.status == "active" {
      //
      //                Text("Active")
      //                    .foregroundColor(.primarySeven)
      //                    .DemoiosText(fontSize: .large)
      //            } else {
      // FINISHED WORKOUT
      HStack(alignment: .center) {
        VStack(alignment: .center) {
          // miles + calories for outdoor cardio, calories + HR for other
          if workout.location_type == "outdoor" && Consts.cardioActivityList.contains(workout.healthkit_workout_activity_type) {
            
            // outdoor cardio
            
            VStack(alignment: .center) {
              Text(workout.miles)
                .DemoiosText(fontSize: .small)
              Text("Distance")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
            
            VStack(alignment: .center) {
              
              Text(workout.formattedHR)
                .DemoiosText(fontSize: .small)
              Text("HR")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
          } else if Consts.cardioActivityList.contains(workout.healthkit_workout_activity_type) {
            // here we have an indoor cardio session
            
            VStack(alignment: .center) {
              
              Text(workout.calories)
                .DemoiosText(fontSize: .small)
              Text("Calories")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
            
            VStack(alignment: .center) {
              
              Text(workout.formattedHR)
                .DemoiosText(fontSize: .small)
              Text("HR")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
          } else {
            // indoor non cardio
            VStack(alignment: .center) {
              
              Text(workout.calories)
                .DemoiosText(fontSize: .small)
              Text("Calories")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
            
            VStack(alignment: .center) {
              
              Text(workout.formattedHR)
                .DemoiosText(fontSize: .small)
              Text("HR")
                .DemoiosText(fontSize: .xxs, color: .graySix)
            }
          }
        }
        
        VStack(alignment: .center) {
          VStack(alignment: .center) {
            Text(workout.healthkit_workout_end_date.intoDate?.timeAgo() ?? "")
              .DemoiosText(fontSize: .small)
            Text("Completed")
              .DemoiosText(fontSize: .xxs, color: .graySix)
          }
          
          VStack(alignment: .center) {
            Text(workout.duration)
              .DemoiosText(fontSize: .small)
            Text("Duration")
              .DemoiosText(fontSize: .xxs, color: .graySix)
          }
        }
      }
      .padding(.trailing)
      .padding(.trailing)
      
    }
  }
    }
}

struct WorkoutImportListItem_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutImportListItem(workout: .defaultWorkout)
        .setupPreview()
    }
}

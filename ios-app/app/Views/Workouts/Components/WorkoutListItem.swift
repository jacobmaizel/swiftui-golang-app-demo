//
//  WorkoutListItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/29/23.
//

import SwiftUI

struct WorkoutListItem: View {
  @EnvironmentObject var data: Datamodel
  
  var workout: Workout
  
  
  
  var body: some View {
    SectionItem(fixedBy: .v, cornerRadius: .large) {
      
      if let workoutData = workout.workout_data {
        HStack {
          VStack {
            
            HStack {
              VStack(alignment: .leading) {
                Text(workout.full_title)
                  .DemoiosText(fontSize: .large, weight: .bold)
                  .multilineTextAlignment(.leading)
                
                
                Text(workoutData.weather?.location_city ?? "")
                
                Text(workoutData.healthkit_workout_end_date.intoDate!, style: .date)
                  .DemoiosText(fontSize: .small)
                
                Text(workoutData.weather?.location_city ?? "")
                  .DemoiosText(fontSize: .small)
                
              }
              
              Spacer()
              
              
              
            }
           
            
            HStack {
              CircleImage(url: workout.leader.picture, size: 30)
       
                Text("Lead by \(workout.leader.id == data.profile.id ? "You" : workout.leader.full_name)")
                  .multilineTextAlignment(.leading)
                  .lineLimit(1)
                  .DemoiosText(fontSize: .small)
            
              Spacer()
              
            }
//            .padding([.leading, .bottom], 8)
            
          }
          
          Spacer()
          VStack {
       
            HStack {
              VStack(alignment: .leading) {
              
                  Text("Duration")
                  .DemoiosText(fontSize: .small, color: .graySix)
                
                if Consts.cardioActivityList.contains(where: { $0 == workout.healthkit_workout_activity_type.lowercased() }) {
                  Text("Miles")
                    .DemoiosText(fontSize: .small, color: .graySix)
                }
              }
              VStack(alignment: .leading) {
                Text(workoutData.duration)
                  .DemoiosText(fontSize: .small)
                if Consts.cardioActivityList.contains(where: { $0 == workout.healthkit_workout_activity_type.lowercased() }) {
                  Text("\(workoutData.miles)")
                    .DemoiosText(fontSize: .small)
                }
            }
            }
          }
//          .padding(.trailing)
        }
        .padding(.vertical, 4)
        .padding(.horizontal)
      } else {
        // NO WORKOUT DATA AVAILABLE
        HStack {
          if workout.workout_data?.source == "Demoios" {
            Logo(size: 50, logoColor: .white)
          } else {
            Image(systemName: "square.and.arrow.down")
              .foregroundColor(.white)
          }
          
          VStack(alignment: .leading) {
            Text(workout.healthkit_workout_activity_type)
              .DemoiosText(fontSize: .xl)
            
            Text("Workout")
              .DemoiosText(fontSize: .xxs, color: .graySix)
          }
          
          Spacer()
          
          Text("No Workout Data")
            .DemoiosText(fontSize: .small, color: .graySix)
          
        }
//        .padding()
      }
    }
  }
}

struct WorkoutListItem_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutListItem(workout: .defaultWorkout)
      .setupPreview()
      .environmentObject(Datamodel())
    
    WorkoutListItem(workout: .defaultWorkout)
      .setupPreview()
      .environmentObject(Datamodel())
      .previewDisplayName("selected")
  }
}

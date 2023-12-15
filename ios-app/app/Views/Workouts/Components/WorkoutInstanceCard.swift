//
//  WorkoutInstanceCard.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/6/23.
//

import SwiftUI

struct WorkoutInstanceCard: View {
  var workoutData: WorkoutData
    var body: some View {
      SectionItem(fixedBy: .both, cornerRadius: .large, color: .grayTen) {
        VStack {
          HStack {
            
            
            CircleImage(url: workoutData.profile.picture, size: 30)
            Text("\(workoutData.profile.full_name)")
              .DemoiosText(fontSize: .base)
            Spacer()
            
            
          }
          HStack {
            HStack {
              SectionItem(fixedBy: .both, cornerRadius: .medium, color: .red) {
                
                HStack(alignment: .center, spacing: 0) {
                  
                  Image(systemName: "heart")
                  Text(workoutData.formattedHR)
                    .DemoiosText(fontSize: .small)
                }
                .padding(4)
              }
            }
            HStack {
              SectionItem(fixedBy: .both, cornerRadius: .medium, color: .green) {
                
                HStack(alignment: .center, spacing: 0) {
                  
                  Image(systemName: "figure.run")
                  Text(workoutData.miles)
                    .DemoiosText(fontSize: .small)
                }
                .padding(4)
                
              }
            }
            HStack {
              SectionItem(fixedBy: .both, cornerRadius: .medium, color: .blue) {
                
                HStack(alignment: .center, spacing: 0) {
                  
                  Image(systemName: "clock")
                  Text(workoutData.duration.compactTime)
                    .DemoiosText(fontSize: .small)
                }
                .padding(4)
                
                
                
              }
            }
          }
          
          HStack {
            Image(systemName: "clock.badge.checkmark")
            Text(workoutData.healthkit_workout_end_date.intoDate?.timeAgo() ?? "")
            Spacer()
          }
          .DemoiosText(fontSize: .small, color: .graySix)
        }
        .padding(.horizontal)
        .padding(.vertical,8)
        
        
      }
    }
}

struct WorkoutInstanceCard_Previews: PreviewProvider {
    static var previews: some View {
      WorkoutInstanceCard(workoutData: .test)
        .environmentObject(Datamodel())
        .setupPreview()
    }
}

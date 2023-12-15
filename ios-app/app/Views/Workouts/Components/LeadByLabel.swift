//
//  LeadByLabel.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/20/23.
//

import SwiftUI

struct LeadByLabel: View {
  @EnvironmentObject var data: Datamodel
  var workout: Workout
    var body: some View {
      
      HStack {
        CircleImage(url: workout.leader.picture, size: 40)
        VStack(alignment: .leading) {
          Text("\(workout.leader.id == data.profile.id ? "You" : workout.leader.full_name)")
            .multilineTextAlignment(.leading)
            .lineLimit(1)
            .DemoiosText(fontSize: .base)
          
            Text("Lead by")
              .DemoiosText(fontSize: .small, color: .graySix)
        }
      }
      
    }
}

struct LeadByLabel_Previews: PreviewProvider {
    static var previews: some View {
      LeadByLabel(workout: .defaultWorkout)
    }
}

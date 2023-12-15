//
//  ElementTagSet.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/5/23.
//

import SwiftUI

struct ElementTagSet: View {
  var comp: Competition
    var body: some View {
      HStack(spacing: 4) {
        ElementTag(title: "Created by", type: .official, values: [comp.type], scale: .small, stack: .v, color: .primarySeven)
        
        ElementTag(title: "Members", type: .participant_types, values: comp.participant_types, scale: .small, stack: .v, color: .orange)
        
        
        ElementTag(title: "Privacy", type: .privacy, values: [String(comp.public)], scale: .small, stack: .v, color: .blue)
        
        ElementTag(title: "Workouts", type: .workout_types, values: comp.workout_types, scale: .small, stack: .v, color: .green)
      }
    }
}

struct ElementTagSet_Previews: PreviewProvider {
    static var previews: some View {
      ElementTagSet(comp: .test_notstarted_joined)
    }
}

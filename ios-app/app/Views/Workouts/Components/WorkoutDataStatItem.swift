//
//  WorkoutStatItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct WorkoutDataStatItem: View {
    
    var title: String
    var value: String
    
    var body: some View {

      VStack(alignment: .leading) {
        Text(value)
          .DemoiosText(fontSize: .base, weight: .bold)
        Text(title)
          .DemoiosText(fontSize: .small, color: .graySix)
      }
     
    }
}

struct WorkoutStatItem_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutDataStatItem(title: "Title", value: "123")
        .setupPreview()
    }
}

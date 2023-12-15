//
//  CompetitionActivityItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/23/23.
//

import SwiftUI

struct CompetitionActivityItem: View {
    var body: some View {
      
      
      HStack {
        
          
        CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 40)
          .padding(.bottom)
        
        VStack(alignment: .leading) {
          VStack(alignment: .leading) {
          Text("Jacob Maizel just finished a Running Workout")
            .DemoiosText(fontSize: .small)
            
            Text("10 Minutes Ago")
              .DemoiosText(fontSize: .small, color: .graySix)
          }
          
          Divider()
//            .padding(.horizontal)
          
        }
        
      }
      .padding(.horizontal)
      
    }
}

struct CompetitionActivityItem_Previews: PreviewProvider {
    static var previews: some View {
      VStack {
        CompetitionActivityItem()
        CompetitionActivityItem()
        CompetitionActivityItem()
        CompetitionActivityItem()
        CompetitionActivityItem()
      }
        .setupPreview()
    }
}

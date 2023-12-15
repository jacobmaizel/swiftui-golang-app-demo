//
//  SquadActivityItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadActivityItem: View {
    var body: some View {
      SectionItem(fixedBy: .v) {
        HStack {
          
          
          CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 40)
          Text("Jake finished a workout")
          Spacer()
          
        }
        .padding()
      }
    }
}

struct SquadActivityItem_Previews: PreviewProvider {
    static var previews: some View {
        SquadActivityItem()
        .setupPreview()
    }
}

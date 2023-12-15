//
//  SquadListItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadListItem: View {
  var squad: Squad
  var body: some View {
    
    SectionItem(fixedBy: .v) {
      
      
      HStack {
//        
//        HStack {
//          CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
//          CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
//            .offset(x: -15)
//          CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
//            .offset(x: -30)
//        }
        
        VStack(alignment: .leading) {
          Text(squad.title)
            .DemoiosText(fontSize: .base)
          Text("\(squad.members?.count ?? 0) \(squad.members?.count == 1 ? "Member" : "Members")")
            .DemoiosText(fontSize: .small, color: .graySix)
        }
      Spacer()
      }
      .padding()
    }
  }
}

struct SquadListItem_Previews: PreviewProvider {
    static var previews: some View {
      SquadListItem(squad: .not_owned_joined)
        .setupPreview()
    }
}

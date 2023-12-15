//
//  SquadDetailMemberList.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/11/23.
//

import SwiftUI

struct SquadDetailMemberList: View {
  @Binding var squad: Squad
  var body: some View {
    
      ScrollView {
        if let members = squad.members {
          ForEach(members, id: \.id) { member in
            
            NavigationLink(value: member) {
              SquadMemberListItem(prof: member)
                .padding(.horizontal)
            }
          }
        } else {
          Text("No Members yet!")
        }
      }
      .navigationDestination(for: PartialProfile.self) { member in
        NonSelfProfileView(partialProfile: member)
      }
      
  }
}

struct SquadDetailMemberList_Previews: PreviewProvider {
  static var previews: some View {
    SquadDetailMemberList(squad: .constant(.not_owned_joined))
      .setupPreview()
  }
}

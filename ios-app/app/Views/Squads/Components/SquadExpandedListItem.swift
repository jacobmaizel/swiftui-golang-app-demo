//
//  SquadExpandedListItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/13/23.
//

import SwiftUI

struct SquadExpandedListItem: View {
  
//  @EnvironmentObject var data: Datamodel
  var squad: Squad
  
    var body: some View {
      SectionItem(fixedBy: .v) {
        VStack {
          
          HStack(alignment: .center) {
            
            Text(squad.title)
              .DemoiosText(fontSize: .xl, weight: .bold)
            
            Spacer()
            
            Image(systemName: "chevron.right")
              .foregroundColor(.graySix)
          }
          
          SquadDetailHeader(squad: squad)
          
          Divider()
          
          HStack(alignment: .center) {
            
            VStack(alignment: .leading) {
              Text("No Current Competition")
              Text("Current Competition")
                .DemoiosText(fontSize: .small, color: .primarySeven)
            }
            
            Spacer()
            
            
            VStack(alignment: .leading, spacing: 0) {
              
              HStack {
                    if let members = squad.members, members.count > 0 {
                      CircleImage(url: members[0].picture, size: 30)
                    }

                    if let members = squad.members, members.count > 1 {
                      CircleImage(url: members[1].picture, size: 30)
                        .offset(x: -15)
                    }

                    if let members = squad.members, members.count > 2 {
                      CircleImage(url: members[2].picture, size: 30)
                        .offset(x: -30)
                    }

                  }
              Text("\(squad.members?.count ?? 0) \(squad.members?.count == 1 ? "Member" : "Members")")
                .DemoiosText(fontSize: .small, color: .graySix)
            }
          }
        }
        .DemoiosText(fontSize: .base)
        .padding()
      }
      .navigationTitle("Squads")
      .padding(.horizontal)
      .padding(EdgeInsets(top: 4, leading: 0, bottom: 0, trailing: 0))
    }
}

struct SquadExpandedListItem_Previews: PreviewProvider {
    static var previews: some View {
      SquadExpandedListItem(squad: .not_owned_joined)
        .setupPreview()
    }
}

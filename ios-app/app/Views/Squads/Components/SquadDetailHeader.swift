//
//  SquadDetailHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadDetailHeader: View {
  var squad: Squad
  var body: some View {
    
    VStack {
      
      HStack {
        
        Text("\(squad.public ? "Public" : "Private")")
        Image(systemName: squad.public ? "lock.open.fill" : "lock.fill")
        Spacer()
        
      }
      .DemoiosText(fontSize: .base)
      
      HStack {
        
        HStack {
          
          CircleImage(url: squad.owner.picture, size: 40)
          
          VStack(alignment: .leading) {
            Text("Owned By")
              .DemoiosText(fontSize: .small, color: .graySix)
            Text("\(squad.owner.first_name) \(squad.owner.last_name)")
              .DemoiosText(fontSize: .base)
          }
        }
        
        Spacer()
        HStack {
          Image(systemName: "calendar.circle")
            .resizable()
            .frame(width: 40, height: 40)
          
          VStack(alignment: .leading) {
            Text("Created")
              .DemoiosText(fontSize: .small, color: .graySix)
            Text(squad.created_at.intoDate?.timeAgo() ?? "")
              .DemoiosText(fontSize: .base)
          }
        }
      }
    }
  }
}

struct SquadDetailHeader_Previews: PreviewProvider {
  static var previews: some View {
    SquadDetailHeader(squad: .owned_joined)
      .setupPreview()
  }
}

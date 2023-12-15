//
//  CompDetailDataSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/27/23.
//

import SwiftUI

struct CompDetailDataSelector: View {
  
  @Binding var dataChoice: CompDetailDataSelection
  var members: [PartialProfile]
  
  var body: some View {
    
    VStack(spacing: 0) {
      HStack {
        
        Spacer()
        
        Button {
          dataChoice = .leaderboard
        } label: {
          Text("Leaderboard")
        }
        .DemoiosText(fontSize: dataChoice == .leaderboard ? .large : .base, color: dataChoice == .leaderboard ? .grayOne : .graySix, weight: dataChoice == .leaderboard ? .bold : .regular )
        Spacer()
        Button {
          dataChoice = .members
        } label: {
          Text("Members \(members.count != 0 ? String(members.count) : "")")
        }
        .DemoiosText(fontSize: dataChoice == .members ? .large : .base, color: dataChoice == .members ? .grayOne : .graySix, weight: dataChoice == .members ? .bold : .regular )
        Spacer()
      }
      .padding()
      Divider()
        .background(Color.grayTwo)
    }
    
  }
}

struct CompDetailDataSelector_Previews: PreviewProvider {
  static var previews: some View {
    CompDetailDataSelector(dataChoice: .constant(.leaderboard), members: [.default])
      .environmentObject(Datamodel())
      .setupPreview()
    CompDetailDataSelector(dataChoice: .constant(.members), members: [.default])
      .environmentObject(Datamodel())
      .setupPreview()
  }
}


enum CompDetailDataSelection {
  case leaderboard, members
}

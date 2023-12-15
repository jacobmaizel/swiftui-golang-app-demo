//
//  CompNotStartedProfileList.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/25/23.
//

import SwiftUI

struct CompNotStartedProfileList: View {
  
  var comp: Competition
  
  var leaderboardData: [CompetitionLeaderboardItem]
  var graphData: [CompetitionStatGraphDataPoint]
  var profileParticipants: [PartialProfile]
  
  var body: some View {
    ScrollView {
      VStack {

//        Text(comp.start.intoDate!, style: .relative)
//          .DemoiosText(fontSize: .base)
//        Text(timerInterval: comp.start.intoDate!...comp.end.intoDate!, countsDown: true, showsHours: true)
        
        HStack {
          Text("Current Competitors \(profileParticipants.count)")
            .DemoiosText(fontSize: .base, color: .graySix)
          
          Spacer()
        }
        
        ForEach(profileParticipants, id: \.id) { member in
          NavigationLink(value: member) {
            SquadMemberListItem(prof: member)
          }
        }

        Spacer()
      }
//      .padding(.horizontal)
    }

  }
}

struct CompNotStartedProfileList_Previews: PreviewProvider {
  static var previews: some View {
    CompNotStartedProfileList(comp: .test_ended_joined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default])
      .setupPreview()
  }
}

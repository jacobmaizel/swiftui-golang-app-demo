//
//  CompDetailDataSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/27/23.
//

import SwiftUI

struct CompDetailDataSection: View {
  
  var dataChoice: CompDetailDataSelection
  
  var profileParticipants: [PartialProfile]
  var leaderboardData: [CompetitionLeaderboardItem]
  
  var body: some View {
    if dataChoice == .leaderboard {
      CompetitionLeaderboard(leaderboardData: leaderboardData)
    } else if dataChoice == .members {
      ScrollView {
        ForEach(profileParticipants, id: \.id) { member in
          NavigationLink(value: member) {
            SquadMemberListItem(prof: member)
          }
        }
      }
    }
  }
}

struct CompDetailDataSection_Previews: PreviewProvider {
  static var previews: some View {
    
    ScrollView {
      CompDetailDataSection(dataChoice: .leaderboard, profileParticipants: [.default], leaderboardData: CompetitionLeaderboardItem.testData)
        .setupPreview()
    }
    
  }
}

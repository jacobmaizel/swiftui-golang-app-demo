//
//  CompetitionDetail.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/14/23.
//

import SwiftUI

struct CompetitionDetail: View {
  
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) private var dismiss
  
  @State var comp: Competition
  
  @State var leaderboardData: [CompetitionLeaderboardItem] = []
  @State var graphData: [CompetitionStatGraphDataPoint] = []
  @State var profileParticipants: [PartialProfile] = []
  @State var myStats: CompMyStats? = nil
  
  @State var isJoined: Bool
  
  // Data section toggle
  @State var dataChoice: CompDetailDataSelection = .leaderboard
  
  // sheet toggles
  @State var showingAddWorkoutSheet: Bool = false
  @State var showingWorkoutList: Bool = false
  @State var showingInviteMemberSheet: Bool = false
  @State private var showingConfirm: Bool = false
  
  // inviting
  @State var memberListToInvite: [PartialProfile] = []
  @State var inviteSearchString: String = ""
  
  //  @State private var showingConfirm: Bool = false
  @State var loaded: Bool = false
  
  var body: some View {
    
    if comp.state == .active {
      CompActiveView(comp: comp, leaderboardData: $leaderboardData, graphData: $graphData, profileParticipants: $profileParticipants, myStats: $myStats, isJoined: isJoined)
        .task {
          if !loaded {
            let (members, gd, cl, ms) = await data.api.refreshCompData(comp.id)
            profileParticipants = members
            graphData = gd
            leaderboardData = cl
            myStats = ms
            loaded = true
          }
        }
        .navigationDestination(for: PartialProfile.self) { prof in
          NonSelfProfileView(partialProfile: prof)
        }
    } else if comp.state == .notStarted {
      CompNotYetStarted(comp: comp, leaderboardData: $leaderboardData, graphData: $graphData, profileParticipants: $profileParticipants, myStats: $myStats, isJoined: isJoined)
        .task {
          if !loaded {
            let (members, gd, cl, ms) = await data.api.refreshCompData(comp.id)
            profileParticipants = members
            graphData = gd
            leaderboardData = cl
            myStats = ms
            loaded = true
          }
        }
        .navigationDestination(for: PartialProfile.self) { member in
          NonSelfProfileView(partialProfile: member)
        }
    } else if comp.state == .ended {
      CompEndedView(comp: comp, leaderboardData: $leaderboardData, graphData: $graphData, profileParticipants: $profileParticipants, myStats: $myStats, isJoined: isJoined)
        .task {
          if !loaded {
            let (members, gd, cl, ms) = await data.api.refreshCompData(comp.id)
            profileParticipants = members
            graphData = gd
            leaderboardData = cl
            myStats = ms
            loaded = true
          }
        }
        .navigationDestination(for: PartialProfile.self) { prof in
          NonSelfProfileView(partialProfile: prof)
        }
    } else {
      Text("failed to load competition")
    }
  }
}


struct CompetitionDetail_Previews: PreviewProvider {
  
  @State private static var c1 = Competition.test_started_joined
  
  static var previews: some View {
    
    let data = Datamodel()
    
    NavigationStack {
      CompetitionDetail(comp: .test_started_joined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default], isJoined: true)
        .environmentObject(data)
        .setupPreview()
      
    }
    .previewDisplayName("Started, joined")
    
    NavigationStack {
      CompetitionDetail(comp: .test_started_notjoined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default],  isJoined: false)
        .environmentObject(data)
        .setupPreview()
      
    }
    .previewDisplayName("Started, NOT Joined")
    NavigationStack {
      CompetitionDetail(comp: .test_notstarted_joined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default], isJoined: true)
        .environmentObject(data)
        .setupPreview()
      
    }
    .previewDisplayName("Not started")
    
    NavigationStack {
      CompetitionDetail(comp: .test_notstarted_notjoined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default], isJoined: false)
        .environmentObject(data)
        .setupPreview()
      
    }
    .previewDisplayName("Not started, not joined")
    NavigationStack {
      CompetitionDetail(comp: .test_ended_joined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default], isJoined: true)
        .environmentObject(data)
        .setupPreview()
    }
    .previewDisplayName("Ended, joined")
    
    NavigationStack {
      CompetitionDetail(comp: .test_ended_not_joined, leaderboardData: CompetitionLeaderboardItem.testData, graphData: CompetitionStatGraphDataPoint.testData, profileParticipants: [.default], isJoined: false)
        .environmentObject(data)
        .setupPreview()
    }
    .previewDisplayName("Ended, NOT joined")
    
    
  }
}

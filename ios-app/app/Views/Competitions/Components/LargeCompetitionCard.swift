//
//  LargeCompetitionCard.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/4/23.
//

import SwiftUI

struct LargeCompetitionCard: View {
  
  @State var competition: Competition
  
  @EnvironmentObject var data: Datamodel

//  var myPlace: String {
//
//    if let idx = leaderboardData.firstIndex(where: { $0.id == data.profile.id }) {
//
//      return "\(ordinalIndicator(for: idx + 1))"
//
//    }
//
//    return ""
//
//  }
//
//  var myLeaderboardInfo: CompetitionLeaderboardItem? {
//    //        if let leaderboardData = competition.leaderboardData {
//    return leaderboardData.first(where: { $0.id == data.profile.id })
//    //        }
//    //
//    //        return nil
//  }
  
  var body: some View {
    SectionItem(fixedBy: .v) {
      VStack {
        
        HStack(alignment: .center) {
          
          Text(competition.title)
            .DemoiosText(fontSize: .xl, weight: .bold)
          
          Spacer()
          
          Image(systemName: "chevron.right")
            .foregroundColor(.graySix)
          
          
          
          //            Image(systemName: "pin")
        }
        
        CompetitionDetailHeader(comp: competition)
        
        Divider()
        
        HStack(alignment: .center) {
          
          if !competition.place.isEmpty {
            VStack(alignment: .leading) {
              Text("\(ordinalIndicator(for: Int(competition.place) ?? 0))")
                .DemoiosText(fontSize: .base, weight: .bold)
              Text("Place")
                .DemoiosText(fontSize: .small, color: .primarySeven)
            }
          }
            
     
          Spacer()
          
          
          VStack(alignment: .leading, spacing: 0) {
            
            HStack {
              if let members = competition.profile_participants, members.count > 0 {
                CircleImage(url: members[0].picture, size: 30)
              }
              
              if let members = competition.profile_participants, members.count > 1 {
                CircleImage(url: members[1].picture, size: 30)
                  .offset(x: -15)
              }
              
              if let members = competition.profile_participants, members.count > 2 {
                CircleImage(url: members[2].picture, size: 30)
                  .offset(x: -30)
              }
              if competition.profile_participants == nil {
                Circle()
                  .frame(height: 30)
                  .foregroundColor(.clear)
              }
            }
            
            Text("\(competition.profile_participants?.count ?? 0) Participants")
              .DemoiosText(fontSize: .small, color: .graySix)
          }
        }
      }
      .DemoiosText(fontSize: .base)
      .padding()
    }
    .padding(.horizontal)
    .padding(EdgeInsets(top: 4, leading: 0, bottom: 0, trailing: 0))
  }
}

struct LargeCompetitionCard_Previews: PreviewProvider {
  static var previews: some View {
    
    let data = Datamodel()
    
    var personalstats: [CompetitionLeaderboardItem] = CompetitionLeaderboardItem.testData.map { item in
      var newItem = item
      newItem.id = Profile.default.id
      return newItem
    }
    var c1 = Competition.test_started_joined

    
    LargeCompetitionCard(competition: c1)
      .environmentObject(data)
      .setupPreview()
      .onAppear {
        c1.profile_participants = [.default]
        
        
        c1.profile_participants?.append(.default)
      }
  }
}

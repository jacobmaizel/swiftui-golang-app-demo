//
//  CompetitionTabView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/6/23.
//

import SwiftUI

struct CompetitionTabView: View {
  @EnvironmentObject var data: Datamodel
  @EnvironmentObject var path: PathManager
  
  @State var loaded: Bool = false
  
//  @Binding homePath: 
  var body: some View {

    NavigationStack(path: $path.compPath) {
        GeometryReader { geo in
          
          
          ScrollView {
            
            if data.joinedCompetitions.isEmpty {
              
              //            HomeCompCardEmpty()
              VStack {
                
                HStack {
                  Spacer()
                  
                  Text("Search for a competition or create one to begin!")
                    .DemoiosText(fontSize: .large, weight: .bold)
                  Spacer()
                }
                
              }
              
              
            } else {
              
              ScrollView {
                LazyVStack {
                  //
                  //            LazyHStack {
                  
                  ForEach(data.joinedCompetitions) { comp in
                    
                    NavigationLink(value: comp) {
                      
                      LargeCompetitionCard(competition: comp )
                        .frame(width: geo.size.width)
                      
                    }
                  }
                  //            }
                }
              }
              .padding(.bottom)
            }
            
          }
          .navigationDestination(for: Competition.self) { comp in
            CompetitionDetail(comp: comp, isJoined: comp.joined)
          }
          .task {
            if !loaded {
              await data.api.getMyCompetitions()
              loaded = true
            }
          }
          .navigationTitle("Competitions")
          .refreshable {
            await data.api.getMyCompetitions()
          }
          .DemoiosBackgroundColor()
        }
    }
  }
}

struct CompetitionTabView_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    
    var personalstats: [CompetitionLeaderboardItem] = CompetitionLeaderboardItem.testData.map { item in
      var newItem = item
      newItem.id = Profile.default.id
      return newItem
    }
    
    
    
    CompetitionTabView()
      .environmentObject(data)
      .environmentObject(PathManager())
      .setupPreview()
      .onAppear {
        data.profile.id = Profile.default.id
        var c1 = Competition.test_started_joined
        c1.leaderboardData = personalstats
        data.joinedCompetitions = [Competition.test_notstarted_joined, c1]
        
        data.myCompletedWorkouts = [.defaultWorkout]
      }
  }
}

//      VStack {
//        Text("Manage your relationship with your currently joined competition")
//        Text("Easily access YOUR owned competitions you created or manage (admin)")
//        Text("view your currently owned, active competition stats (ie. how many people are currently in there, other stats")
//        Text("view pending competitions and the steps needed in order to start it officially")
//      }

//
//              HStack(alignment: .center) {
//
//                Text("Competition Activity")
//                  .DemoiosText(fontSize: .xl)
//                Spacer()
//
//                Button {
//                  print("refresh")
//                } label: {
//                  Image(systemName: "arrow.clockwise")
//                    .DemoiosText(fontSize: .large, color: .primarySeven)
//                }
//              }
//              .padding(.horizontal)
//
//
//              VStack {
//                CompetitionActivityItem()
//                CompetitionActivityItem()
//                CompetitionActivityItem()
//                CompetitionActivityItem()
//                CompetitionActivityItem()
//              }
//              .padding(.vertical)
//

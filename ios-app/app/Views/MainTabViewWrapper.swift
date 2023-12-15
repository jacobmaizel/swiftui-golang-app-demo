//
//  TabView.swift
//  Demoios
//
//  Created by Jacob Maizel on 2/4/23.
//

import SwiftUI

struct MainTabViewWrapper: View {
  // wrapper for main flow of views after logging in and onboarding

  @StateObject var pathManager = PathManager.shared
  
  @EnvironmentObject var dataModel: Datamodel
  
  //  @State private var tabSelection: Tab = Tab.home
  

  
  var body: some View {
    
    TabView(selection: dataModel.tabSelector) {
      
      HomeTabView()
        .environmentObject(pathManager)
        .tabItem {
          Label("Home", systemImage: "house.fill")
        }
        .tag(Tab.home)
      
      CompetitionTabView()
        .environmentObject(pathManager)
        .tabItem {
          Label("Competitions", systemImage: "person.and.background.dotted")
        }
        .tag(Tab.competitions)
      
      WorkoutsTabView()
        .environmentObject(pathManager)
        .tabItem {
          Label("Workout", systemImage: "figure.run")
        }
        .tag(Tab.workout)
      
      SquadTabView()
        .environmentObject(pathManager)
        .tabItem {
          Label("Squads", systemImage: "person.2.fill")
        }
        .tag(Tab.squads)

      MyProfileTabView(profile: dataModel.profile)
        .environmentObject(pathManager)
        .tabItem {
          Label("Profile", systemImage: "person.crop.circle.fill")
        }
        .tag(Tab.profile)
    }
    .onAppear {
//      let appearance = UITabBarAppearance()
//      appearance.backgroundEffect = UIBlurEffect(style: .systemUltraThinMaterial)
//      appearance.backgroundColor = UIColor(Color.grayTen.opacity(0.4))
//
//      UITabBar.appearance().standardAppearance = appearance
//      UITabBar.appearance().scrollEdgeAppearance = appearance
    }
  }
}

struct TabView_Previews: PreviewProvider {
  static var previews: some View {

    let data = Datamodel()
      
    MainTabViewWrapper()
      .setupPreview()
      .environmentObject(data)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .onAppear {
        data.joinedCompetitions = [Competition.test_started_joined]
        data.joinedCompetitions[0].leaderboardData = CompetitionLeaderboardItem.testData
        data.joinedCompetitions = [Competition.test_notstarted_joined]
        data.joinedCompetitions[0].leaderboardData = CompetitionLeaderboardItem.testData
        data.joinedCompetitions.append(.test_started_joined)
        
        data.joinedCompetitions.append(.test_ended_joined)
        
        data.myCompletedWorkouts.append(.defaultWorkout)
        data.joinedSquads.append(.owned_joined)
      }
      .previewDisplayName("Started, Joined")
    
    MainTabViewWrapper()
      .setupPreview()
      .environmentObject(data)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .onAppear {
        data.joinedCompetitions = [Competition.test_notstarted_joined]
        data.joinedCompetitions[0].leaderboardData = CompetitionLeaderboardItem.testData
        data.joinedCompetitions.append(.test_started_joined)
        
        data.joinedCompetitions.append(.test_ended_joined)
        
        data.myCompletedWorkouts.append(.defaultWorkout)
        data.joinedSquads.append(.owned_joined)
      }
      .previewDisplayName("not started, joined")
//      .previewDisplayName("iPhone 11")
//
//    MainTabViewWrapper()
//
//      .environmentObject(Datamodel())
//      .setupPreview()

  }
}

enum Tab {
  case home
  case workout
  case squads
  case competitions
  //        case goals
  case profile
}

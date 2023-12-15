//
//  HomeView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/12/23.
//

import SwiftUI
import HealthKit



struct HomeTabView: View {
  @EnvironmentObject var data: Datamodel
  @EnvironmentObject var loc: LocationManager
  @EnvironmentObject var pathManager: PathManager
  
  var sectionHeight: CGFloat = 100
  
  var items: [GridItem] = Array(repeating: .init(.flexible(), spacing: 10), count: 2)
  
  
  var body: some View {
    
    NavigationStack(path: $pathManager.homePath) {
      
      GeometryReader { _ in
        ScrollView {
          
          VStack {
            HomeTopHeader()
            
            HStack {
              //
              //          WeatherStack(stack: .h)
              //              .padding(.leading)
              
              Spacer()
              
            }
            
            HStack {
              
              Text("Welcome, \(data.profile.first_name)!")
                .DemoiosText(fontSize: .xxl, color: .grayOne, weight: .bold)
              
              Spacer()
            }
            .padding(.horizontal)
            
            //          Divider()
            NavigationLink(value: HomePaths.search) {
              SectionItem(fixedBy: .v, cornerRadius: .large) {
                HStack(alignment: .center) {
                  Text("Search")
                  Image(systemName: "magnifyingglass")
                }
                .foregroundColor(.graySix)
                .padding()
                //                      .padding(.horizontal)
                //            .padding(.leading)
              }
              .opacity(0.8)
            }
            .padding([.bottom, .leading, .trailing])
            //                  .padding(.trailing)
            
            
            
            LazyVGrid(columns: items, alignment: .center, spacing: 10) {
              
              Button {
                dispatchAsync {
                  data.currentTab = .competitions
                }
              } label: {
                
                HomeActionItem(title: "View your competitions", description: "Check out the latest info on your competitions", icon: .viewComps, background: .primarySeven)
              }
              
              
              NavigationLink(value: HomePaths.workoutcreate) {
                HomeActionItem(title: "Create a workout", description: "Get active and Invite your friends!", icon: .createWorkout, background: .grayTen)
              }
              
              NavigationLink(value: HomePaths.findsquad) {
                
                HomeActionItem(title: "Find a squad", description: "Create or join a Squad!", icon: .findSquad, background: .grayTen)
                
              }
              
              NavigationLink(value: HomePaths.viewstats) {
                
                HomeActionItem(title: "View your stats", description: "Total distance, comp wins, and more", icon: .checkStats, background: .grayTen)
              }
              
              
            }
            .padding()
            
            
            Spacer()
            
            
          }
          
          .navigationDestination(for: HomePaths.self) { pathType in
            if pathType == .findsquad {
              SearchView(searchScope: .squads)
            } else if pathType == .viewstats {
              StatDetailView()
            } else if pathType == .workoutcreate {
              WorkoutConfigurationView()
            } else if pathType == .search {
              SearchView()
            } else if pathType == .compCreate {
              CompetitionCreateForm()
            } else if pathType == .notificationList {
              NotificationList()
            }
            
          }
        }
        .DemoiosBackgroundColor()
        .refreshable {
          //          data.api.loadInitialData(data: data)
          await data.api.asyncLoadInitialData()
        }
        
        //        .navigationDestination(for: Competition.self) { comp in
        //          CompetitionDetail(comp: comp, isJoined: comp.joined)
        //
        //        }
        
      }
      
      
      .task {
        //        print("ON APPEAR TRIGGERED IN HOMEVIEW\n\n")
        loc.lookUpCurrentLocation { _ in
        }
        
        if data.profileIsLoaded {
          if !data.initialDataLoaded {
            //            print("INITAL DATA WAS FALSE IN ONAPPEAR HOMEVIew")
            print("LOADING INITIAL DATA")
            //          data.api.loadInitialData(data: data)
            //          data.api.asy
            await data.api.asyncLoadInitialData()
            
            
            listImportableWorkouts(data: data) { (workouts, error) in
              guard let workouts = workouts, error == nil else {
                print("Could not list importable workouts in on appear")
                return
              }
              
              Task {
                for workout in workouts {
                  do {
                    try await data.api.importWorkout(payload: workout.toDictionary())
                    
                  } catch {
                    triggerLocalNotification(title: "Failed to import the workouts", body: "Please try again", background: .red)
                  }
                } // importable workout loop
                //                  triggerLocalNotification(title: "Imported \(successes)", body: "", background: .green)
                
                await data.api.getMyCompletedWorkouts()
              }
            }
          }
        }
        //      data.api.getMyNotifications(data: data)
      }
    }
  }
}

struct HomeView_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    
    var personalstats: [CompetitionLeaderboardItem] = CompetitionLeaderboardItem.testData.map { item in
      var newItem = item
      newItem.id = Profile.default.id
      return newItem
    }
    
    
    HomeTabView()
      .previewDisplayName("ipad")
      .previewDevice(PreviewDevice(stringLiteral: "iPhone 11"))
      .environmentObject(data)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .environmentObject(PathManager())
      .setupPreview()
      .onAppear {
        data.profile.id = Profile.default.id
        var c1 = Competition.test_started_joined
        c1.leaderboardData = personalstats
        data.joinedCompetitions = [Competition.test_notstarted_joined, c1]
      }
    
    HomeTabView()
      .previewDisplayName("ipad")
      .previewDevice(PreviewDevice(stringLiteral: "iPad Pro (11-inch) (4th generation)"))
      .environmentObject(data)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .environmentObject(PathManager())
      .setupPreview()
  }
}

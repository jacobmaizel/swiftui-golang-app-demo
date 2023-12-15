//
//  ProfileView.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/14/23.
//

import SwiftUI

struct MyProfileTabView: View {
  
  @State private var birthday: Date = Date.now
  @State private var firstName: String = "Jake"
  @State private var lastName: String = "Maizel"
  @State private var dataChoice: ProfileDataOptions = .workouts
  @EnvironmentObject var path: PathManager
  
  @EnvironmentObject var data: Datamodel
  var profile: Profile
  
  @State var profileData: ProfileStatsData = ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0)
  @State var compList: [Competition] = []
  @State var squadList: [Squad] = []
  @State var workoutList: [Workout] = []
  
  @State var initialDataLoaded: Bool = false
  
  @State var showingStatDetailsSheet: Bool = false
  
  var body: some View {
    
    NavigationStack(path: $path.profilePath) {
      ScrollView {
        VStack(spacing: 8) {
          
          ProfileSettingsMenu()
          
          ProfileHeader(picture: profile.picture, first_name: profile.first_name, last_name: profile.last_name, created_at: profile.created_at)
          
          Button {
            showingStatDetailsSheet.toggle()
          } label: {
            
            ProfileStats(profileStats: data.profileStats)
          }
          
          ProfileDataSection(dataChoice: $dataChoice, viewingOwnProfile: true)
          
          switch dataChoice {
          case .squads:
            
            ForEach(data.joinedSquads) { squad in
              NavigationLink(value: squad) {
                SquadSearchResult(squad: squad)
              }
            }
          case .workouts:
            LazyVStack {
              ForEach(data.myCompletedWorkouts.sorted(by: { $0.workout_data?.healthkit_workout_end_date ?? "" > $1.workout_data?.healthkit_workout_end_date ?? "" })) { workout in
                NavigationLink(value: workout) {
                  WorkoutListItem(workout: workout)
                }
              }
            }
            .padding(.horizontal)
          case .competitions:
            
            ForEach(data.joinedCompetitions) { comp in
              NavigationLink(value: comp) {
                CompSearchResult(comp: comp)
              }
            }
          }
        }
      }
      .padding(.bottom)
      .navigationDestination(for: Competition.self) { comp in
        CompetitionDetail(comp: comp, isJoined: comp.joined)
      }
      .navigationDestination(for: Squad.self) { squad in
        SquadDetailView(squad: squad)
      }
      .navigationDestination(for: Workout.self) { workout in
        WorkoutSummaryView(workout: workout)
      }
      .refreshable {
        Task {
          async let profileStats: () = data.api.getMyProfileStats()
          async let squads: () = data.api.listMySquads()
          async let competitions: () = data.api.getMyCompetitions()
          async let completedWorkouts: () = data.api.getMyCompletedWorkouts()
          
          // Await all results before proceeding
          await profileStats
          await squads
          await competitions
          await completedWorkouts
          
        }
      }
      
      .DemoiosBackgroundColor()
    }
    .sheet(isPresented: $showingStatDetailsSheet) {
      StatDetailView()
    }
    .DemoiosBackgroundColor()
    .task {
      if !initialDataLoaded && !data.initialDataLoaded {
        Task {

          async let ps: () = data.api.getMyProfileStats()
          
          async let ls: () = data.api.listMySquads()
          async let mc: () = data.api.getMyCompetitions()
          async let mcw: () = data.api.getMyCompletedWorkouts()
          
          
          await ps
          await ls
          await mc
          await mcw
        }
        initialDataLoaded = true
        data.initialDataLoaded = true
      }
    }
  }
  
}


struct ProfileView_Previews: PreviewProvider {
  static var previews: some View {
    
    let data = Datamodel()
    
    MyProfileTabView(profile: Profile.default)
      .environmentObject(data)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .environmentObject(PathManager())
      .setupPreview()
      .previewDisplayName("viewing mine")
      .onAppear {
        
        data.joinedCompetitions.append(.test_ended_joined)
        
        data.myCompletedWorkouts = [.defaultWorkout, .defaultWorkout2]
        data.joinedSquads.append(.owned_joined)
      }
    
  }
}

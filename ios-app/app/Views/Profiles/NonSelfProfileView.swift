//
//  NonSelfProfileView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/11/23.
//

import SwiftUI

struct NonSelfProfileView: View {
  
  @EnvironmentObject var data: Datamodel
  
  @State private var firstName: String = "Jake"
  @State private var lastName: String = "Maizel"
  @State private var dataChoice: ProfileDataOptions = .workouts
  
  var partialProfile: PartialProfile
  
  @State var profileData: ProfileStatsData = ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0)
  @State var compList: [Competition] = []
  @State var squadList: [Squad] = []
  @State var workoutList: [Workout] = []
  
  var body: some View {
    
    //    NavigationStack {
    
    OverlayBottomToolbar {
      ScrollView {
        
        
        ProfileHeader(picture: partialProfile.picture, first_name: partialProfile.first_name, last_name: partialProfile.last_name)
        
        ProfileStats(profileStats: profileData)
        
        ProfileDataSection(dataChoice: $dataChoice, viewingOwnProfile: true)
        
        switch dataChoice {
        case .squads:
          ForEach(squadList) { squad in
            NavigationLink(value: squad) {
              SquadSearchResult(squad: squad)
            }
          }
          //            .padding(.horizontal)
        case .workouts:
          LazyVStack {
            ForEach(workoutList) { workout in
              NavigationLink(value: workout) {
                WorkoutListItem(workout: workout)
              }
            }
            .padding(.horizontal)
          }
          
        case .competitions:
          ForEach(compList) { comp in
            NavigationLink(value: comp) {
              CompSearchResult(comp: comp)
            }
          }
        }
      }
      .refreshable {
        do {
          let res = try await data.api.loadNonSelfProfileData(data: data, profileId: partialProfile.id)
          
          self.profileData = res.0 ?? .empty
          self.squadList = res.1 ?? []
          self.compList = res.2 ?? []
          self.workoutList = res.3 ?? []
          
        } catch {
          triggerLocalNotification(title: "Failed to load some data for this profile", body: "", background: .red)
        }
      }
        
        .navigationDestination(for: Squad.self) {squad in
          SquadDetailView(squad: squad)
        }
        .navigationDestination(for: Workout.self) { workout  in
          WorkoutSummaryView(workout: workout)
        }
//        .navigationDestination(for: Competition.self) { comp  in
//          //          WorkoutSummaryView(workout: workout)
//          CompetitionDetail(comp: comp, isJoined: comp.joined)
//        }
      
    
      .DemoiosBackgroundColor()
    } toolbarContent: {
//      Button("Callout \(partialProfile.full_name)") {
//
//        print("Callout")
//
//      }
//      .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
//      .padding()
    }
    .DemoiosBackgroundColor()
    .task {
      do {
        let res = try await data.api.loadNonSelfProfileData(data: data, profileId: partialProfile.id)
        
        self.profileData = res.0 ?? .empty
        self.squadList = res.1 ?? []
        self.compList = res.2 ?? []
        self.workoutList = res.3 ?? []
        
      } catch {
        triggerLocalNotification(title: "Failed to load some data for this profile", body: "", background: .red)
      }
    }
  }
}

struct NonSelfProfileView_Previews: PreviewProvider {
  static var previews: some View {
    NonSelfProfileView(partialProfile: .default)
  }
}

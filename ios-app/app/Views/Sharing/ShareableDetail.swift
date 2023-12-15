//
//  ShareableDetail.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/4/23.
//

import SwiftUI

struct ShareableDetail: View {
  @EnvironmentObject var data: Datamodel
  var resourceType: ShareableResource
  var resourceId: UUID
  
  @State var loaded: Bool = false
  
  @State var comp: Competition?
  @State var squad: Squad?
  @State var profile: PartialProfile?
  @State var workout: Workout?
  @State var err: Error?
  

  var body: some View {
    VStack {
      if loaded {
        if let comp = comp, resourceType == .competitions {
          let joined = data.joinedCompetitions.contains(where: { $0.id == resourceId })
          CompetitionDetail(comp: comp, isJoined: joined)
            .navigationTitle(comp.title)
            .navigationBarTitleDisplayMode(.inline)
        }
        if let squad = squad, resourceType == .squads {
          SquadDetailView(squad: squad)
            .navigationTitle(squad.title)
            .navigationBarTitleDisplayMode(.inline)
        }
        if let profile = profile, resourceType == .profiles {
          NonSelfProfileView(partialProfile: profile)
            .navigationTitle(profile.full_name)
            .navigationBarTitleDisplayMode(.inline)
        }
        if let workout = workout, resourceType == .workouts {
          WorkoutSummaryView(workout: workout)
            .navigationTitle(workout.full_title)
            .navigationBarTitleDisplayMode(.inline)
        }
      } else if loaded && err != nil {
       Text("Failed to load")
        if getCurrentEnvType() == .debug {
          Text("\(err.debugDescription)")
        }
      } else {
        ProgressView()
      }
    }
    .task {
      if resourceType == .competitions {
        do {
          comp = try await data.api.getCompById(id: resourceId)
          loaded = true
        } catch {
          err = error
          loaded = true
        }
      } else if resourceType == .squads {
        do {
          squad = try await data.api.getSquadById(squadId: resourceId)
          loaded = true
        } catch {
          err = error
          loaded = true
        }
      } else if resourceType == .workouts {
        do {
          workout = try await data.api.getWorkoutById(id: resourceId)
          loaded = true
        } catch {
          err = error
          loaded = true
        }
      } else if resourceType == .profiles {
        do {
          profile = try await data.api.getProfileById(resourceId)
          loaded = true
        } catch {
          err = error
          loaded = true
        }
      }
      
    }
  }
}

struct ShareableDetail_Previews: PreviewProvider {
  static var previews: some View {
    ShareableDetail(resourceType: .competitions, resourceId: UUID(), loaded: true)
      .environmentObject(Datamodel())
      .setupPreview()
    
  }
}

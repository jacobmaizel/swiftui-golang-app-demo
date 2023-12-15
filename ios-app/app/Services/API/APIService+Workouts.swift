//
//  WorkoutAPI.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/19/23.
//

import Foundation
import HealthKit

extension DemoiosAPIService {
  
  
  @MainActor
  func getMyCompletedWorkouts(data: Datamodel = .shared) async {
    do {
      data.myCompletedWorkouts = try await asyncRequest(method: .get, path: "/workouts", queryItems: [URLQueryItem(name: "for_profile", value: data.profile.id.uuidString)])
    } catch {
      print(error)
    }
  }
  
  func listWorkoutsByProfile(profileId: UUID) async -> [Workout]? {
    let res: [Workout]? = try? await asyncRequest(method: .get, path: "/workouts", queryItems: [URLQueryItem(name: "for_profile", value: profileId.uuidString)])
    return res
  }
  
  func createWorkout(payload: [String: String]) async throws -> String {
    let res: Workout = try await asyncRequest(method: .post, path: "/workouts/", payload: payload)
    return res.id.uuidString
  }
  
  func finishWorkoutPatch(payload: [String: Any]) async throws {
    var payloadToSend = payload
    let workout_id = payloadToSend.removeValue(forKey: "workout_id")
    _ = payloadToSend.removeValue(forKey: "healthkit_workout_activity_type")
    
    guard let workout_id = workout_id as? String else {
      throw DemoiosError.runtimeError("Failed to parse workout id in finishWorkoutPatch")
    }
    /*
     type WorkoutDataWeather struct {
       Temp_with_unit string `json:"temp_with_unit"`
       Symbol         string `json:"symbol"`
       Location_city  string `json:"location_city"`
     }

     */
    if LocationManager.shared.status?.rawValue ?? 0 > 2 {
      
      LocationManager.shared.getLocation()
      
      
      if let loc = LocationManager.shared.location {
        await WeatherManager.shared.getWeather(for: loc)
        
        payloadToSend["weather"] = [
          "temp_with_unit": await WeatherManager.shared.tempF,
          "symbol": await WeatherManager.shared.symbol,
          "location_city": LocationManager.shared.city
        ]
      }
      
      
      
      
      
      
   
    }
      
      print("\n\nADDED weather info to workout: \(payloadToSend["weather"])")
    
    print("\n\nsubmitting workoutdata to api: \(payloadToSend) \n")
    payloadToSend["source"] = "Demoios"
    
      let _: [String: String] = try await asyncRequest(method: .post, path: "/workouts/\(workout_id)/submit_data", payload: payloadToSend)
  }
  
  func inviteProfilesToWorkoutByIds(workoutId: UUID, idsToInvite: [UUID]) async throws -> String {
    let uuidStrings = idsToInvite.map { $0.uuidString }
    let idStringsToInvite: String = uuidStrings.joined(separator: ",")
    do {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/workouts/\(workoutId)/invite", queryItems: [URLQueryItem(name: "send_to", value: idStringsToInvite)])
      return "Invites sent!"
    } catch {
      throw error
    }
  }
  
  func getWorkoutById(id: UUID) async throws -> Workout {
      let res: Workout = try await asyncRequest(method: .get, path: "/workouts/\(id.uuidString)/")
      return res
  }
  
  
  func getMyAcceptedInviteWorkouts() async throws -> [Workout] {
    let res: [Workout] = try await asyncRequest(method: .get, path: "/workouts", queryItems: [URLQueryItem(name: "accepted_invites", value: "true")])
    return res
  }
  
  
  func acceptWorkoutInvite(workout_id: UUID, invite_id_to_accept: UUID) async throws -> String {
    do {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/workouts/\(workout_id.uuidString)/accept_invite/\(invite_id_to_accept.uuidString)")
      return "Successfully Accepted Workout Invite"
    } catch {
      return "Failed to accept Invite, please try again"
    }
  }
  
  func declineWorkoutInvite(workout_id: UUID, invite_id_to_decline: UUID) async throws -> String {
    do {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/workouts/\(workout_id.uuidString)/decline_invite/\(invite_id_to_decline.uuidString)")
      return "Successfully Declined Workout Invite"
    } catch {
      return "Failed to decline Invite, please try again"
    }
  }

  
  func getUserInviteStatusesForWorkout(workoutId: UUID) async -> [PartialProfile] {
    do {
      let res: [PartialProfile] = try await asyncRequest(method: .get, path: "/workouts/\(workoutId)/invite_statuses")
      return res
    } catch {
      print(error)
      return []
    }
  }
  
  func getMyLeadIncompleteWorkouts() async -> [Workout] {
    do {
      let res: [Workout] = try await asyncRequest(method: .get, path: "/workouts/", queryItems: [URLQueryItem(name: "lead_by", value: Datamodel.shared.profile.id.uuidString), URLQueryItem(name: "incomplete", value: "true")])
      return res
    } catch {
      print(error)
      return []
      //      throw error
    }
  }
  
  func importWorkout(payload: [String: Any]) async throws -> Workout {
    var payloadToSend = payload
    _ = payloadToSend.removeValue(forKey: "workout_id")
    
    do {
      let res: Workout = try await asyncRequest(method: .post, path: "/workouts/import", payload: payload)
      return res
    } catch {
      print("ERRORRRRR in import workout \(error)")
      throw error
    }
  }
  
  func getApplicableWorkoutsForCompetition(comp_id: UUID) async throws -> [Workout] {
      let res: [Workout] = try await asyncRequest(method: .get, path: "/workouts", queryItems: [URLQueryItem(name: "applicable_for_comp", value: comp_id.uuidString)])
    return res
  }
  
  func getWorkoutDataByWorkoutId(workoutId: UUID) async -> [WorkoutData] {
    // Returns all NON SELF data for a given workout
    
    do {
      let res: [WorkoutData] = try await asyncRequest(method: .get, path: "/workouts/\(workoutId.uuidString)/data")
      return res
    } catch {
      print(error)
      return []
    }
  }
  
  func deleteWorkout(_ id: UUID) async throws {
    let _: [String: String] = try await asyncRequest(method: .delete, path: "/workouts/\(id.uuidString)")
  }
  func removeWorkoutInvite(workout_id: UUID) async throws {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/workouts/\(workout_id)/remove_invite")
  }
}

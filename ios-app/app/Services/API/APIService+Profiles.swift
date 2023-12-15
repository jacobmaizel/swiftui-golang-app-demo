//
//  ProfileAPI.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/19/23.
//

import Foundation

extension DemoiosAPIService {
  // MARK: - Profile
  
  @MainActor
  func getUserProfileAndConfig(_ dataModel: Datamodel = .shared) async throws {
    do {
      let prof: Profile = try await asyncRequest(method: .get, path: "/profiles/me")
      dataModel.profile = prof
      
      let config = try await getConfig()
      dataModel.config = config
      
      dataModel.profileIsLoaded = true
      
    } catch {
      print(error)
      throw error
    }
  }
  
  @MainActor
  func completeOnboarding(_ dataModel: Datamodel = .shared) async throws {
    // when the user is actually done with onboarding, so we set profile to loaded
    // so that the UI auto updates to the proper screen
    let payload: [String: Any] = ["onboarding_completed": true]
    
    let res: Profile = try await asyncRequest(method: .patch, path: "/profiles/", payload: payload)
    
    dataModel.profile = res
    dataModel.profileIsLoaded = true
  }
  
  @MainActor
  func updateProfile(_ dataModel: Datamodel = .shared, payload: [String: Any]) async throws {
    // used for first step in onboarding where user is just updating name
    do {
      let prof: Profile = try await asyncRequest(method: .patch, path: "/profiles/", payload: payload)
      dataModel.profile = prof
    } catch {
      print("failed to update profile \(error)")
      throw error
    }
  }
  
  @MainActor
  func updateFcmToken(token: String) async {
    let payload: [String: String] = ["token": token]
    do {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/profiles/fcm_token", payload: payload)
      Datamodel.shared.fcmTokenSubmittedThisSession = true
    } catch {
      print("FAILED TO UPDATE FCM TOKEN:: \(error)")
    }
  }
  
  @MainActor
  func getMyProfileStats(data: Datamodel = Datamodel.shared) async {
    do {
      data.profileStats = try await asyncRequest(method: .get, path: "/profiles/\(data.profile.id.uuidString)/stats")
    } catch {
      print(error)
    }
  }
  
  func getProfileStatsById(data: Datamodel = .shared, profileId: UUID) async throws -> ProfileStatsData? {
    
    let data: ProfileStatsData? = try? await asyncRequest(method: .get, path: "/profiles/\(profileId.uuidString)/stats")
    return data
  }
  
  func getProfileById(_ profileId: UUID) async throws -> PartialProfile {
    let data: PartialProfile = try await asyncRequest(method: .get, path: "/profiles/\(profileId.uuidString)/")
    return data
  }
  
  
} // end

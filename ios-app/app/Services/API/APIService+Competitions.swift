//
//  CompetitionAPI.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/19/23.
//

import Foundation


enum DemoiosError: Error {
  case runtimeError(String)
}

extension DemoiosAPIService {
  
  @MainActor
  func getMyCompetitions(data: Datamodel = .shared) async {
    do {
      let compList: [Competition] = try await asyncRequest(method: .get, path: "/competitions/", queryItems: [URLQueryItem(name: "joined", value: "true")])
      data.joinedCompetitions = compList
    } catch {
      print(error)
    }
  }
  
  func listCompetitionsByProfile(_ dataModel: Datamodel = .shared, profileId: UUID) async -> [Competition]? {
    
    let res: [Competition]? = try? await asyncRequest(method: .get, path: "/competitions/", queryItems: [URLQueryItem(name: "for_participant", value: profileId.uuidString)])
    
    return res
  }
  
  // MARK: - JOINING COMPETITION
  func leaveCompetition(compId: String) async throws {
    let endpoint = "/competitions/\(compId)/leave"
    let method: HTTPMethod = .post
    
    do {
      let _: [String: String] = try await asyncRequest(method: method, path: endpoint)
   
      await self.getMyCompetitions()

    } catch {
      print(error)
      throw error
    }
    
  }

  // MARK: - JOINING COMPETITION
  func joinCompetition(compId: String) async throws -> Competition {
    let endpoint = "/competitions/\(compId)/join"
    let method: HTTPMethod = .post
    do {
      let res: Competition = try await asyncRequest(method: method, path: endpoint)
   
      await self.getMyCompetitions()

      return res
    } catch {
      print(error)
      throw error
    }
    
  }
  
  
  func createCompetition(payload: [String: Any]) async throws -> Competition {
    let endpoint = "/competitions/"
    let method: HTTPMethod = .post
    do {
      let res: Competition = try await asyncRequest(method: method, path: endpoint, payload: payload)
      
        await self.getMyCompetitions()
      
      return res
    } catch {
      print(error)
      throw error
    }
    
  }
  
  func getLeaderboardForCompetition(compId: UUID) async -> [CompetitionLeaderboardItem] {
    
    do {
      let res: [CompetitionLeaderboardItem] = try await asyncRequest(method: .get, path: "/competitions/\(compId)/leaderboard")
      return res
    } catch {
      print("comp leaderboard fail\(error)")
      return []
    }
    
  }
  
  func getGraphDataForCompetition(comp_id: UUID) async -> [CompetitionStatGraphDataPoint] {
    do {
      let res: [CompetitionStatGraphDataPoint] = try await asyncRequest(method: .get, path: "/competitions/\(comp_id)/graph")
      return res
    } catch {
      print(error)
      return []
    }
  }
  
  func getMyStatsForComp(comp_id: UUID) async -> CompMyStats? {
    do {
      let res: CompMyStats = try await asyncRequest(method: .get, path: "/competitions/\(comp_id)/stats")
      return res
    } catch {
      print(error)
      return nil
    }
  }
  
  func getParticipantsForCompetition(compId: UUID) async -> [PartialProfile] {
    do {
      let res: [PartialProfile] = try await asyncRequest(method: .get, path: "/competitions/\(compId.uuidString)/participants")
      return res
    } catch {
      print(error)
      return []
    }
  }
  
  func refreshCompData(_ id: UUID) async -> ([PartialProfile], [CompetitionStatGraphDataPoint], [CompetitionLeaderboardItem], CompMyStats?) {
    print("Refreshing comp data")
    async let members = getParticipantsForCompetition(compId: id)
    async let gd = getGraphDataForCompetition(comp_id: id)
    async let cl = getLeaderboardForCompetition(compId: id)
    async let ms = getMyStatsForComp(comp_id: id)
    return await (members, gd, cl, ms)
  }
  
  func competitionSearch(searchTerm: String) async -> CompSearchRes {
    
    let queryItems: [URLQueryItem] = [URLQueryItem(name: "type", value: "competitions"), URLQueryItem(name: "term", value: searchTerm)]
    
    do {
      let res: CompSearchRes = try await asyncRequest(method: .get, path: "/services/search/", queryItems: queryItems)
      return res
    } catch {
      return CompSearchRes(count: 0, results: [])
    }
  }
  
  func applyWorkoutsToComp(data: Datamodel = .shared, compId: UUID, workoutIds: [UUID]) async throws -> [String: String] {
    let payload: [String: Any] = ["workout_data_ids": workoutIds.map { $0.uuidString } ]
    do {
      let res: [String: String] = try await asyncRequest(method: .post, path: "/competitions/\(compId.uuidString)/apply_workout_data", payload: payload)
      return res
    } catch {
      print(error)
      throw error
    }
  }
  
  func inviteMembersToComp(compId: UUID, idsToInvite: [UUID]) async throws {
    let uuidStrings = idsToInvite.map { $0.uuidString }
    let idStringsToInvite: String = uuidStrings.joined(separator: ",")
    
    let _: [String: String] = try await asyncRequest(method: .post, path: "/competitions/\(compId)/invite", queryItems: [URLQueryItem(name: "send_to", value: idStringsToInvite)])
  }
  
  func getCompById(id: UUID) async throws -> Competition {
    let res: Competition = try await asyncRequest(method: .get, path: "/competitions/\(id)")
    return res
  }
  
  func acceptCompInvite(inviteID: UUID, compId: UUID) async throws -> [String: String] {
    
    let endpoint = "/competitions/\(compId)/accept_invite/\(inviteID)"
    let method: HTTPMethod = .post
    
    do {
      let res: [String: String] = try await asyncRequest(method: method, path: endpoint)
      return res
    } catch {
      print(error)
      throw error
    }
  }
  
  func declineCompInvite(inviteID: UUID, compId: UUID) async throws -> [String: String] {
    
    let endpoint = "/competitions/\(compId)/decline_invite/\(inviteID)"
    let method: HTTPMethod = .post
    
    do {
      let res: [String: String] = try await asyncRequest(method: method, path: endpoint)
      return res
    } catch {
      print(error)
      throw error
    }
  }
  
}

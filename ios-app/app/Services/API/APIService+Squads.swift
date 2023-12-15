//
//  SquadAPI.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/19/23.
//

import Foundation

extension DemoiosAPIService {
  
  func createSquad(_ data: Datamodel = .shared, payload: [String: Any]) async throws {
    do {
      let _: Squad = try await asyncRequest(method: .post, path: "/squads", payload: payload)
    } catch {
      throw error
    } }
  
  func leaveSquad(squadID: UUID, data: Datamodel = .shared) async throws {
    let endpoint = "/squads/\(squadID)/leave"
      let _: [String: String] = try await asyncRequest(method: .post, path: endpoint)
  }
  
  func joinSquad(squadID: UUID) async throws {
    let _: [String: String] = try await asyncRequest(method: .post, path: "/squads/\(squadID)/join")
  }
  

  func acceptSquadInvite(inviteID: UUID, squadID: UUID) async throws {
    let endpoint = "/squads/\(squadID.uuidString)/accept_invite/\(inviteID.uuidString)"
    let method: HTTPMethod = .post
    let _: [String: String] = try await asyncRequest(method: method, path: endpoint)
  }
  
  func denySquadInvite(inviteID: UUID, squadID: UUID) async throws {
    let endpoint = "/squads/\(squadID.uuidString)/decline_invite/\(inviteID.uuidString)"
    let method: HTTPMethod = .post
    let _: [String: String] = try await asyncRequest(method: method, path: endpoint)
  }
  
  func inviteProfilesToSquadByIds(squadId: UUID, idsToInvite: [UUID]) async throws -> String {
    let uuidStrings = idsToInvite.map { $0.uuidString }
    let idStringsToInvite: String = uuidStrings.joined(separator: ",")
    do {
      let _: [String: String] = try await asyncRequest(method: .post, path: "/squads/\(squadId)/invite", queryItems: [URLQueryItem(name: "send_to", value: idStringsToInvite)])
      return "Invites Sent!"
    } catch {
      throw error
    }
  }
  
  @MainActor
  func getMyPendingInvites(data: Datamodel = .shared) async throws {
    let res: [Invite] = try await asyncRequest(method: .get, path: "/invites/", queryItems: [.init(name: "pending", value: "true")])
    data.invites = res.sorted { $0.created_at > $1.created_at }
  }
  
  func listSquadsByProfile(profileId: UUID) async -> [Squad]? {
    let res: [Squad]? = try? await asyncRequest(method: .get, path: "/squads/", queryItems: [URLQueryItem(name: "with_member", value: profileId.uuidString)])
    return res
  }
  
  @MainActor
  func listMySquads(data: Datamodel = .shared) async {
    do {
      let squads: [Squad] = try await asyncRequest(method: .get, path: "/squads/", queryItems: [URLQueryItem(name: "joined", value: "true")])
      data.joinedSquads = squads
    } catch {
      print(error)
    }
  }
  
  func getSquadStats(squadId: UUID) async -> SquadStats? {
    let path = "/squads/\(squadId.uuidString)/stats"
    let res: SquadStats? = try? await asyncRequest(method: .get, path: path)
    return res
  }
  
  func getSquadById(squadId: UUID) async throws -> Squad {
    let path = "/squads/\(squadId.uuidString)/"
    let res: Squad = try await asyncRequest(method: .get, path: path)
    return res
  }
  
  
  func deleteSquad(with id: UUID) async throws {
    let _: [String: String] = try await asyncRequest(method: .delete, path: "/squads/\(id.uuidString)/")
  }
  
  func patchSquad(id: UUID, payload: [String: Any]) async throws -> Squad {
    let squadRes: Squad = try await asyncRequest(method: .patch, path: "/squads/\(id.uuidString)/", payload: payload)
    return squadRes
  }
  
} // END

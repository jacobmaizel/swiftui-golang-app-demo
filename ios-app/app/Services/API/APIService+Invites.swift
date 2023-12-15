//
//  APIService+Invites.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/15/23.
//

import Foundation



extension DemoiosAPIService {
  func getInviteById(inviteId: UUID) async throws -> Invite {
    let res: Invite = try await asyncRequest(method: .get, path: "/invites/\(inviteId.uuidString)")
    return res
  }
}

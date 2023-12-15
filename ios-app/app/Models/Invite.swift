//
//  Invite.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/22/23.
//

import Foundation

struct Invite: Codable, Identifiable {
    var id: UUID = UUID()
    var status: String
    var created_at: String
    var updated_at: String
    var sender: PartialProfile
  
  static let test = Invite(status: "pending", created_at: Date.now.adding(weeks: -1).intoString!, updated_at: Date.now.adding(weeks: -1).intoString!, sender: .default)
}

struct Sender: Codable, Identifiable {
    var id: UUID = UUID()
    var picture: String
    var first_name: String
    var last_name: String
}

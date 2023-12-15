//
//  Squad.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/22/23.
//

import Foundation

struct Squad: Codable, Identifiable, Hashable {
  static func == (lhs: Squad, rhs: Squad) -> Bool {
    return lhs.id == rhs.id
  }
  
  var id: UUID = UUID()
  
  var title: String
  //    var member_count: Int
  
  var created_at: String
  var updated_at: String
  
  var owner: PartialProfile
  
  var members: [PartialProfile]?
  
  var joined: Bool
  
  var `public`: Bool = true
  
  var stats: SquadStats?
  
  var member_count: Int
  
//  var owned: Bool = false
  
  static let owned_joined: Squad = Squad(title: "Owned & Joined", created_at: "2023-03-22T18:50:26.692606Z", updated_at: "2023-03-22T18:50:26.692606Z", owner: PartialProfile( first_name: "Jake", last_name: "maizel", picture: "https://asdf.com", public: true), members: [PartialProfile.default], joined: true, public: true, member_count: 5)
  
  static let not_owned_not_joined: Squad = Squad(title: "Not owned OR joined", created_at: "2023-03-22T18:50:26.692606Z", updated_at: "2023-03-22T18:50:26.692606Z", owner: PartialProfile( first_name: "Jake", last_name: "maizel", picture: "https://asdf.com", public: true), members: [PartialProfile.default], joined: false, public: true, member_count: 5)
  
  static let not_owned_joined: Squad = Squad(title: "Not Owned, but joined", created_at: "2023-03-22T18:50:26.692606Z", updated_at: "2023-03-22T18:50:26.692606Z", owner: PartialProfile( first_name: "Jake", last_name: "maizel", picture: "https://asdf.com", public: true), members: [PartialProfile.default], joined: true, public: true, member_count: 5)
  
  
}

//
// type SquadStatsResponse struct {
//  MemberCount int `json:"member_count"`
//  CompetitionWins int `json:"competition_wins"`
//  TotalDistanceFt float64 `json:"total_distance_ft"`
// }

struct SquadStats: Codable, Hashable {
  var member_count: Int
  var competition_wins: Int
  var total_distance_ft: Double
}

#if os(iOS)


#endif

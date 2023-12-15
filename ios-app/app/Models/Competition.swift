//
//  Competition.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/23/23.
//

import Foundation

struct Competition: Codable, Identifiable, Hashable {

  
  static func == (lhs: Competition, rhs: Competition) -> Bool {
    return lhs.id == rhs.id
  }
  
  
  var id: UUID = UUID()
  var title: String
  var description: String
  
  var start: String
  var end: String
  
  var scheduled: Bool = false
  
  var workout_types: [String] = []
  var participant_types: [String]
  
  var type: String
  
  var joined: Bool = false
  
  var owner: PartialProfile
  
  var `public`: Bool = true
  
  var status: String
  
  var created_at: String
  var updated_at: String
  
  var leaderboardData: [CompetitionLeaderboardItem]?
  var place: String
  var distance_total: Double
  
  var profile_participants: [PartialProfile]?
  
  var squad_participants: [Squad]?
  
  var participant_count: Int
  
  
}



struct CompetitionLeaderboardItem: Identifiable, Codable, Hashable, Comparable {
  static func < (lhs: CompetitionLeaderboardItem, rhs: CompetitionLeaderboardItem) -> Bool {
    return lhs.distance_total < rhs.distance_total
  }
  
  var id: UUID // profile id
  
  var first_name: String
  var last_name: String
  var picture: String
  
  var distance_total: Double
  
}

struct CompMyStats: Decodable {
  var place: String
  var distance_total: Double
  
  static let test = CompMyStats(place: "hi", distance_total: 123.0)
}

struct CompetitionStatGraphDataPoint: Identifiable, Decodable {
  enum CodingKeys: CodingKey {
    case healthkit_workout_end_date
    case distance
  }
  
  var id = UUID()
  
  var healthkit_workout_end_date: String
  var distance: Double

}


// MARK: - Extensions

#if os(iOS)
extension Competition {
//  func calcIsJoined(data: Datamodel = .shared) -> Bool {
//    print("calced is comp joined")
//    return data.joinedCompetitions.contains(where: { $0.id == self.id})
//  }
}
#endif




//
//  Extensions+Competition.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/26/23.
//

import Foundation

enum CompetitionTimeState {
  case notStarted
  case active
  case ended
  case invalid
}




extension Competition {
  
  func hasEnded() -> Bool {
    
    guard let compStart = self.start.intoDate, let compEnd = self.end.intoDate else {
      print("failed to convert dates in hasEnded")
      return false
    }
    
    let now = Date.now
    
    if compStart < now && compEnd < now {
      return true
    }
    
    return false
  }
  
  func isActive() -> Bool {
    
    guard let d1 = self.start.intoDate, let d2 = self.end.intoDate else {
      return false
    }
    
    let now = Date()
    
    if d1 < now && d2 >= now {
      return true
    }
    
    return false
    
  }
  
  func notStarted() -> Bool {
    
    guard let start = self.start.intoDate, let end = self.end.intoDate else {
      return false
    }
    
    let currentDate = Date.now
    
    return start > currentDate && end > currentDate
  }
  var timeLeft: String {
    
    guard let date = self.end.intoDate else {
      return ""
    }
    
    return Date.now.timeUntil(date: date)
  }
  
  var timeToStart: String {
    
    guard let date = self.end.intoDate else {
      return ""
    }
    return Date.now.timeUntil(date: date)
  }
  
  var timeSinceEnd: String {
    guard let date = self.end.intoDate else {
      return ""
    }
    
    return date.timeAgo() ?? ""
  }
  
  var state: CompetitionTimeState {
    
    guard self.start.intoDate != nil, self.end.intoDate != nil else {
      return .invalid
    }
    
//    print("CALLED STATE")
    
    
    // not started - currentDate < Start Date
    if notStarted() {
      return .notStarted
    } else if isActive() {
      return .active
    } else if hasEnded() {
      return .ended
    } else {
      return .invalid
    }
  }
  
  
}


#if os(iOS)
extension Competition {
  
  static let test_started_notjoined = Competition(title: "test_started_notjoined", description: "A great competition:)", start: Date.now.adding(days: 5).intoString ?? "", end: Date.now.adding(weeks: 2).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: false, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ],place: "1", distance_total: 20, participant_count: 4)
  
  static let test_started_joined = Competition(title: "Started, Joined", description: "A great competition:)", start: Date.now.adding(days: -2).intoString ?? "", end: Date.now.adding(weeks: 2).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: true, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ], place: "1", distance_total: 20, profile_participants: [.default], participant_count: 4)
  // not started
  
  static let test_notstarted_notjoined = Competition(title: "NOT started, NOT joined", description: "A great competition:)", start: Date.now.adding(days: 5).intoString ?? "", end: Date.now.adding(weeks: 2).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: false, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ], place: "1", distance_total: 20, participant_count: 4)
  
  static let test_notstarted_joined = Competition(title: "NOT started, joined", description: "A great competition:)", start: Date.now.adding(days: 10).intoString ?? "", end: Date.now.adding(weeks: 2).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: true, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ],place: "1", distance_total: 20, participant_count: 4)
  
  static let test_ended_joined = Competition(title: "Ended, joined", description: "A great competition:)", start: Date.now.adding(days: -20).intoString ?? "", end: Date.now.adding(weeks: -1).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: true, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ],place: "1", distance_total: 20, participant_count: 4)
  static let test_ended_not_joined = Competition(title: "Ended, NOT joined", description: "A great competition:)", start: Date.now.adding(days: -20).intoString ?? "", end: Date.now.adding(weeks: -1).intoString ?? "", workout_types: ["running", "cycling"], participant_types: ["squad", "solo"], type: "official", joined: false, owner: .default, status: "pending", created_at: "2023-02-21T19:01:44.428143Z", updated_at: "2023-11-18T19:01:44.428143Z", leaderboardData: [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0)
  ],place: "1", distance_total: 20, participant_count: 4)
  
}
#endif
extension CompetitionLeaderboardItem {
  
  static let testData = [
    CompetitionLeaderboardItem(id: UUID(), first_name: "John", last_name: "Doe", picture: "john_doe.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Jane", last_name: "Smith", picture: "jane_smith.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Michael", last_name: "Johnson", picture: "michael_johnson.png", distance_total: 75.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emily", last_name: "Brown", picture: "emily_brown.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "David", last_name: "Wilson", picture: "david_wilson.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sarah", last_name: "Davis", picture: "sarah_davis.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Robert", last_name: "Anderson", picture: "robert_anderson.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Olivia", last_name: "Taylor", picture: "olivia_taylor.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "William", last_name: "Martin", picture: "william_martin.png", distance_total: 140.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Clark", picture: "sophia_clark.png", distance_total: 160.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Daniel", last_name: "Walker", picture: "daniel_walker.png", distance_total: 130.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Emma", last_name: "Thomas", picture: "emma_thomas.png", distance_total: 170.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "James", last_name: "Anderson", picture: "james_anderson.png", distance_total: 110.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Sophia", last_name: "Lee", picture: "sophia_lee.png", distance_total: 120.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Matthew", last_name: "Harris", picture: "matthew_harris.png", distance_total: 150.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Ava", last_name: "Wilson", picture: "ava_wilson.png", distance_total: 90.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Benjamin", last_name: "Johnson", picture: "benjamin_johnson.png", distance_total: 200.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Mia", last_name: "Roberts", picture: "mia_roberts.png", distance_total: 100.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Elijah", last_name: "Turner", picture: "elijah_turner.png", distance_total: 180.0),
    CompetitionLeaderboardItem(id: UUID(), first_name: "Charlotte", last_name: "Scott", picture: "charlotte_scott.png", distance_total: 130.0)
  ]
}


extension CompetitionStatGraphDataPoint {
  
  static let testData: [CompetitionStatGraphDataPoint] = [
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-13T12:00:00.000000Z", distance: 100.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-14T12:00:00.000000Z", distance: 150.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-15T12:00:00.000000Z", distance: 75.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-16T12:00:00.000000Z", distance: 200.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-17T12:00:00.000000Z", distance: 120.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-18T12:00:00.000000Z", distance: 90.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-19T12:00:00.000000Z", distance: 180.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-20T12:00:00.000000Z", distance: 110.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-21T12:00:00.000000Z", distance: 140.0),
    CompetitionStatGraphDataPoint(healthkit_workout_end_date: "2023-06-22T12:00:00.000000Z", distance: 160.0)
  ]
}

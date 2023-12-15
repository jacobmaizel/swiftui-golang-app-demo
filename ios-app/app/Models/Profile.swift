//
//  Profile.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/12/23.
//

import Foundation

struct Profile: Codable, Identifiable {
    var id: UUID = UUID()
    var first_name: String
    var last_name: String
    var picture: String
    var birthday: String?
    var onboarding_completed: Bool
    
    var created_at: String
    var updated_at: String
  
  var `public`: Bool = true
  
  var full_name: String {
    return "\(first_name) \(last_name)"
  }
    
    
    static let `default` = Profile(first_name: "Jacob", last_name: "Maizel", picture: "https://lh3.googleusercontent.com/a/AGNmyxbgRZWxwYyzzpADX4VGuKMkdA0Xu2la2nSz5FxmDdQ=s96-c", birthday: "2023-03-04T19:01:44.428143Z", onboarding_completed: true, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z")
  
  
  
  
  static let test_profiles: [Profile] = {
      let profile1 = Profile(first_name: "Emily", last_name: "Smuch", picture: "https://lh3.googleusercontent.com/a/AGNmyxbgRZWxwYyzzpAmDdQ=s96-c", birthday: "2023-03-04T19:01:44.428143Z", onboarding_completed: true, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z")
      
      let profile2 = Profile(first_name: "Bobby", last_name: "Smith", picture: "https://lh3.googleusercontent.com/a/AGNmyxbgRZWxwYyzzpkdA0Xu2la2nSz5FxmDdQ=s96-c", birthday: "2023-03-04T19:01:44.428143Z", onboarding_completed: true, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z" )
      
      let profile3 = Profile(first_name: "James", last_name: "Meezels", picture: "https://example.com/profile.jpg", birthday: "2023-03-04T19:01:44.428143Z", onboarding_completed: true, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z")
      
      let profile4 = Profile(first_name: "Joba", last_name: "Menlo", picture: "https://lh3.googleusercontent.com/a/AGNmyxbgRZWxwYyzzFxmDdQ=s96-c", birthday: "1990-10-15T00:00:00Z", onboarding_completed: true, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z" )
      
      let profile5 = Profile(first_name: "Jack", last_name: "Meezol", picture: "https://lh3.googleusercontent.com/ADX4VGuKMkdA0Xu2la2nSz5FxmDdQ=s96-c", birthday: "2023-03-04T19:01:44.428143Z", onboarding_completed: false, created_at: "2023-03-04T19:01:44.428143Z", updated_at: "2023-03-04T19:01:44.428143Z")
      
      return [Profile.default, profile1, profile2, profile3, profile4, profile5]
  }()
}



/*
 type ProfileStatsResponse struct {
   Workout_count     int `json:"workout_count"`
   Competition_wins  int `json:"competition_wins"`
   Total_distance_ft int `json:"total_Distance_ft"`
 }
 */

struct ProfileStatsData: Decodable {
  var workout_count: Int
  var competition_wins: Int
  var total_distance_ft: Double
}

extension ProfileStatsData {
  
  static let empty = ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0)
  
  var miles: String {
    return "\(feetToMiles(feet: total_distance_ft))"
  }
}
 
 
struct PartialProfile: Codable, Identifiable, Hashable {
  var id: UUID = UUID()
  var first_name: String
  var last_name: String
  var picture: String
  
  var invite_status: String?
  
  var `public`: Bool
  
  var full_name: String {
    return first_name + " " + last_name
  }
  
  static let `default`: PartialProfile = PartialProfile( first_name: "Jake", last_name: "maizel", picture: "https://asdf.com", public: true)
  
  
  
}

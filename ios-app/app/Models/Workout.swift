//
//  Workout.swift
//  Demoios
//
//  Created by Jacob Maizel on 4/7/23.
//

import Foundation

struct Workout: Codable, Identifiable, Hashable, Comparable {
  static func < (lhs: Workout, rhs: Workout) -> Bool {
    guard let d1 = lhs.created_at.intoDate, let d2 = rhs.created_at.intoDate else {
      return false
    }
    return d1 < d2
  }
  
  static func == (lhs: Workout, rhs: Workout) -> Bool {
    lhs.id == rhs.id
  }
  
  
  var id: UUID = UUID()
  
  var created_at: String
  var updated_at: String
  
  var healthkit_workout_activity_type: String
  
  var leader: PartialProfile
  var workout_data: WorkoutData?
  
  var route_data: String?
  
  var competition: Competition?
  
  var full_title: String {
    var location = ""
    if let wd = workout_data, Consts.cardioActivityList.contains(where: { $0 == healthkit_workout_activity_type.lowercased() }) {
      location = wd.location_type.capitalized
    }
    
    return "\(location + " ")\(healthkit_workout_activity_type.capitalized)"
  }
  
  
  static let defaultWorkout2 = Workout(
    created_at: "2023-04-20T10:00:00Z",
    updated_at: "2023-04-20T11:00:00Z",
    healthkit_workout_activity_type: "Walking",
    leader: .default,
    workout_data: .test2,
    //        healthkit_workout_id: "123456",
    //        healthkit_workout_start_date: "2023-04-20T10:00:00Z",
    //        healthkit_workout_end_date: "2023-04-20T10:00:00Z",
    //        average_heart_rate: "120",
    //        distance: 5.0,
    //        duration: "30:00",
    //        energy_burned: "300",
    //        status: "completed",
    //        location_type: "indoor",
    //        source: "Demoios",
    competition: Competition.test_notstarted_joined
  )
  
  static let testGraphData: [WorkoutGraphDataPoint] = [
    WorkoutGraphDataPoint(value: 120, date: Date.now.adding(days: -12)),
    WorkoutGraphDataPoint(value: 40, date: Date.now.adding(days: -10)),
    WorkoutGraphDataPoint(value: 20, date: Date.now.adding(days: -8)),
    WorkoutGraphDataPoint(value: 120, date: Date.now.adding(days: -6)),
    WorkoutGraphDataPoint(value: 70, date: Date.now.adding(days: -2))
  ]
  
  static let defaultWorkout = Workout(
    created_at: "2023-04-20T10:00:00Z",
    updated_at: "2023-04-20T11:00:00Z",
    healthkit_workout_activity_type: "Running",
    leader: .default,
    workout_data: .test,
    //        healthkit_workout_id: "123456",
    //        healthkit_workout_start_date: "2023-04-20T10:00:00Z",
    //        healthkit_workout_end_date: "2023-04-20T10:00:00Z",
    //        average_heart_rate: "120",
    //        distance: 5.0,
    //        duration: "30:00",
    //        energy_burned: "300",
    //        status: "completed",
    //        location_type: "indoor",
    //        source: "Demoios",
    competition: Competition.test_notstarted_joined
    )
    
      
  
}

struct ImportableWorkout: Hashable, Encodable {
  
  var healthkit_workout_activity_type: String
  var healthkit_workout_id: String
  var healthkit_workout_start_date: String
  var healthkit_workout_end_date: String
  
  var average_heart_rate: String
  var distance: Double
  var duration: String
  var energy_burned: String
//  var status: String
  
  var location_type: String
  
  var source: String = "import"
  
  var miles: String {
    return "\(feetToMiles(feet: distance)) Miles"
  }
  
  var calories: String {
    let x = roundToWholeNumber(from: energy_burned)
    return "\(x ?? "")"
  }
  
  var formattedHR: String {
    let x = roundToWholeNumber(from: average_heart_rate)
    return "\(x ?? "")"
  }
  
  static let defaultWorkout = ImportableWorkout(
    healthkit_workout_activity_type: "Running",
    healthkit_workout_id: "123456",
    healthkit_workout_start_date: "2023-04-20T10:00:00Z",
    healthkit_workout_end_date: "2023-04-20T10:00:00Z",
    average_heart_rate: "120",
    distance: 5.0,
    duration: "30:00",
    energy_burned: "300",
//    status: "completed",
    location_type: "indoor",
    source: "Demoios"
  )
  
  func toDictionary() -> [String: Any] {
    let encoder = JSONEncoder()
//    print(self)
    if let data = try? encoder.encode(self) {
      if let jsonObject = try? JSONSerialization.jsonObject(with: data, options: .allowFragments) as? [String: Any] {
        return jsonObject
      }
    }
    print("returning empty dictionary uh oh ")
    return [:] // Return an empty dictionary in case of failure
  }
}



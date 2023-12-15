//
//  WorkoutData.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/8/23.
//

import Foundation


struct WorkoutData: Identifiable, Codable, Hashable {
  var id = UUID()
  
    var healthkit_workout_id: String
    var healthkit_workout_start_date: String
    var healthkit_workout_end_date: String
  
  var distance: Double
  var duration: String
  var energy_burned: String
  var average_heart_rate: String
  
  var source: String = "Demoios"
  var location_type: String = "indoor"
  
  var competition: Competition?
  
  var profile: PartialProfile
  
  var weather: WorkoutDataWeather?
  
}

struct WorkoutDataWeather: Codable, Hashable {
  var temp_with_unit: String
  var symbol: String
  var location_city: String
}
/*
 type WorkoutDataWeather struct {
   Temp_with_unit string `json:"temp_with_unit"`
   Symbol         string `json:"symbol"`
   Location_city  string `json:"location_city"`
 }

 */


extension WorkoutData {
  
  // MARK: - - WorkoutDataMethods
  
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
  
  static let test = WorkoutData(
    id: UUID(),
    healthkit_workout_id: UUID().uuidString,
    healthkit_workout_start_date: "2023-04-20T10:00:00Z",
    healthkit_workout_end_date: "2023-04-20T12:00:00Z",
    distance: 5.0,
    duration: "30",
    energy_burned: "391",
    //        status: "completed",
    average_heart_rate: "117",
    source: "Demoios",
    competition: .test_started_joined,
    profile: .default,
    weather: WorkoutDataWeather(temp_with_unit: "86 F", symbol: "cloud", location_city: "Test York")
    
  )
  static let test2 = WorkoutData(
    id: UUID(),
    healthkit_workout_id: UUID().uuidString,
    healthkit_workout_start_date: "2023-04-20T10:00:00Z",
    healthkit_workout_end_date: "2023-04-20T12:00:00Z",
    distance: 1235.0,
    duration: "1:20",
    energy_burned: "200",
    //        status: "completed",
    average_heart_rate: "97",
    source: "Demoios",
    profile: .default
  )
}

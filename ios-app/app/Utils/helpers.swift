//
//  helpers.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/26/23.
//

import Foundation
import HealthKit
import SwiftUI

extension Date {
  func days(from date: Date) -> Int {
    Calendar.current.dateComponents([.day], from: date, to: self).day!
  }
  func adding(days: Int) -> Date {
    Calendar.current.date(byAdding: .day, value: days, to: self)!
  }
  func adding(minutes: Int) -> Date {
    Calendar.current.date(byAdding: .minute, value: minutes, to: self)!
  }
  func adding(weeks: Int) -> Date {
    Calendar.current.date(byAdding: .weekOfYear, value: weeks, to: self)!
  }
  func adding(months: Int) -> Date {
    Calendar.current.date(byAdding: .month, value: months, to: self)!
  }
  func adding(years: Int) -> Date {
    Calendar.current.date(byAdding: .year, value: years, to: self)!
  }
}

extension Date {
  
  func timeAgo() -> String? {
    guard self < Date.now else {
      return nil
    }
    
    let currentDate = Date()
    let calendar = Calendar.current
    let components = calendar.dateComponents([.year, .month, .day, .hour, .minute, .second], from: self, to: currentDate)
    
    if let year = components.year, year > 0 {
      return "\(year) year\(year == 1 ? "" : "s") ago"
    } else if let month = components.month, month > 0 {
      return "\(month) month\(month == 1 ? "" : "s") ago"
    } else if let day = components.day, day > 0 {
      return "\(day) day\(day == 1 ? "" : "s") ago"
    } else if let hour = components.hour, hour > 0 {
      return "\(hour) hour\(hour == 1 ? "" : "s") ago"
    } else if let minute = components.minute, minute > 0 {
      return "\(minute) minute\(minute == 1 ? "" : "s") ago"
    } else if let second = components.second, second > 0 {
      return "\(second) second\(second == 1 ? "" : "s") ago"
    } else {
      return nil
    }
  }
  
  func timeUntil(date: Date) -> String {
    guard date >= Date.now else {
      print("TIME UNTIL FAILED: DATE WAS >= DATE.now")
      return ""
    }
    
    let calendar = Calendar.current
    let components = calendar.dateComponents([.day, .hour, .minute, .second], from: self, to: date)
    if let day = components.day, day > 0 {
      return "\(day) days"
    } else if let hour = components.hour, hour > 0 {
      return "\(hour) hours"
    } else if let minute = components.minute, minute > 0 {
      return "\(minute) minutes"
    } else if let second = components.second, second > 0 {
      return "\(second) seconds"
    } else {
      print("FAILED in TIMEUNTIL FUNCTION, RETURNING EMPTY STRING")
      return ""
    }
  }
}

extension Formatter {
  static let formatDate: DateFormatter = {
    
    let formatter = DateFormatter()
    formatter.timeZone = TimeZone(abbreviation: "UTC")
    formatter.locale = .init(identifier: "en_US_POSIX")
    formatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSSSSSS'Z'"
    
    return formatter
  }()
}

func betweenTwoDates(start: Date?, end: Date?) -> Bool {
  
  guard let d1 = start, let d2 = end else {
    return false
  }
  
  let now = Date()
  
  if d1 < now && d2 > now {
    return true
  }
  
  return false
}

extension Date {
  var intoString: String? {
    Formatter.formatDate.string(from: self)
  }
}

extension String {
  
  var compactTime: String {
    let inputFormatter = DateFormatter()
    inputFormatter.locale = .init(identifier: "en_US_POSIX")
    inputFormatter.dateFormat = "HH:mm:ss"
    
    if let date = inputFormatter.date(from: self) {
      let outputFormatter = DateFormatter()
      outputFormatter.dateFormat = "H:mm"
      return outputFormatter.string(from: date)
    }
    
    return ""
  }
  
  var stringToFullDateTime: String {
    let dateFormatter = DateFormatter()
    dateFormatter.timeZone = TimeZone.autoupdatingCurrent
    dateFormatter.locale = .init(identifier: "en_US_POSIX")
    
    if let date = self.intoDate {
      dateFormatter.dateFormat = "MMM dd, yyyy - hh:mm:ss a"
      return dateFormatter.string(from: date)
    }
    return ""
  }
  var stringToShortDate: String {
    let dateFormatter = DateFormatter()
    dateFormatter.timeZone = TimeZone.autoupdatingCurrent
    dateFormatter.locale = .init(identifier: "en_US_POSIX")
    
    if let date = self.intoDate {
      dateFormatter.dateFormat = "MMM dd, yyyy"
      return dateFormatter.string(from: date)
    }
    return ""
  }
  
  var intoDate: Date? {
    let formatter = DateFormatter()
    formatter.timeZone = TimeZone.autoupdatingCurrent
    formatter.timeZone = TimeZone(abbreviation: "UTC")
    formatter.locale = .init(identifier: "en_US_POSIX")
    
    let possibleDateFormats = ["yyyy-MM-dd'T'HH:mm:ss.SSSSSSSSZ", "yyyy-MM-dd'T'HH:mm:ss.SSSZ", "yyyy-MM-dd'T'HH:mm:ssZ"]
    
    for format in possibleDateFormats {
      
      formatter.dateFormat = format
      
      let date = formatter.date(from: self)
      
      if date != nil {
        return date
      }
    }
    return nil
  }
}

// func compById(comp_id: UUID, dataModel: Datamodel) -> Competition? {
//    return dataModel.competitions.first(where: { $0.id == comp_id })
// }


func dispatchAsync(_ functionToRun: @escaping () -> Void) {
  DispatchQueue.main.async {
    functionToRun()
  }
}


//
// func calcCompTimeLeft(comp: Competition) -> String {
//
//    let timeLeft = Date.now.timeUntil(date: comp.end.intoDate!)
//
//    return timeLeft ?? ""
// }
//
// func calcCompTimeUntil(comp: Competition) -> String {
//
//    let timeUntil = Date.now.timeUntil(date: comp.start.intoDate!)
//
//    return timeUntil ?? ""
// }

func getWorkoutsCompletedInDemoios(from healthStore: HKHealthStore, completion: @escaping ([HKWorkout]?, Error?) -> Void) {
  let workoutType = HKObjectType.workoutType()
  
  // Define the workout source for your custom app
  let workoutSource = HKSource.default()
  
  //    print("Trying to list workouts for Source bundleid::: \(workoutSource.bundleIdentifier) ::: source name:: \(workoutSource.name) ::: source string description: \(workoutSource.description)")
  
  // Define the workout predicate to retrieve workouts saved by your app
  let workoutPredicate = HKQuery.predicateForObjects(from: workoutSource)
  
  
  // Combine the workout type and workout source predicate
  let compoundPredicate = NSCompoundPredicate(andPredicateWithSubpredicates: [workoutPredicate])
  
  // Define the workout query to retrieve the workouts
  let workoutQuery = HKSampleQuery(sampleType: workoutType,
                                   predicate: compoundPredicate,
                                   limit: HKObjectQueryNoLimit,
                                   sortDescriptors: nil) { (_, results, error) in
    guard let workouts = results as? [HKWorkout], error == nil else {
      completion(nil, error)
      return
    }
    
    // Filter out workouts that were not completed in your app
    let filteredWorkouts = workouts.filter { $0.sourceRevision.source == workoutSource }
    
    // Return the filtered workouts to the completion handler
    completion(filteredWorkouts, nil)
  }
  
  healthStore.execute(workoutQuery)
}

func deleteWorkoutsForCurrentSource(store: HKHealthStore) {
  
  let predicate = HKQuery.predicateForObjects(from: .default())
  
  let workoutQuery = HKSampleQuery(sampleType: HKObjectType.workoutType(),
                                   predicate: predicate,
                                   limit: HKObjectQueryNoLimit,
                                   sortDescriptors: nil) { (_, samples, error) in
    if let workouts = samples as? [HKWorkout], !workouts.isEmpty {
      store.delete(workouts) { (success, error) in
        if success {
          print("Deleted \(workouts.count) workouts for current source.")
        } else {
          if let deleteError = error {
            print("Failed to delete workouts: \(deleteError.localizedDescription)")
          } else {
            print("Failed to delete workouts: unknown error")
          }
        }
      }
    } else {
      print("No workouts found for current source.")
    }
  }
  
  store.execute(workoutQuery)
}
func printWorkoutData(_ workouts: [HKWorkout]) {
  //    let formatter = DateFormatter()
  //    formatter.dateStyle = .medium
  //    formatter.timeStyle = .medium
  
  for workout in workouts {
    print("______________")
    
    print(workout.allStatistics)
    
    print("WORKOUT META DATA", workout.metadata ?? "NO META DATA")
    
    print("______________")
    //        let duration = workout.duration
    //        let distance = workout.totalDistance?.doubleValue(for: HKUnit.mile()) ?? 0.0
    //        let workoutType = workout.workoutActivityType
    //        let avgHeartRate = workout.averageHeartRate?.doubleValue(for: HKUnit.count().unitDivided(by: HKUnit.minute())) ?? 0.0
    //        let endDate = workout.endDate
    //
    //        print("Duration: \(duration.formatted())")
    //        print("Distance: \(String(format: "%.2f", distance)) miles")
    //        print("Workout type: \(workoutType)")
    //        print("Average heart rate: \(String(format: "%.0f", avgHeartRate)) bpm")
    //        print("End date: \(formatter.string(from: endDate))\n")
  }
}
func feetToMiles(feet: Double) -> String {
  
  let milesValue = feet / 5280.0
  let formatter = NumberFormatter()
  formatter.maximumFractionDigits = 2
  formatter.minimumFractionDigits = 1
  return formatter.string(from: NSNumber(value: milesValue)) ?? "999"
}

extension String {
  var formatNumber: String {
    let formatter = NumberFormatter()
    formatter.maximumFractionDigits = 2
    
    guard let doubleValue = Double(self) else {
      return "0"
    }
    
    return formatter.string(from: NSNumber(value: doubleValue)) ?? "FAILED TO CONVERT"
  }
}

func roundToWholeNumber(from string: String) -> String? {
  if let doubleValue = Double(string) {
    let roundedValue = Int(doubleValue.rounded())
    return String(roundedValue)
  }
  return nil
}

extension HKWorkout {
  
  func extractAPIPayload() -> [String: Any] {
    let allstats = self.allStatistics
    let heartRateUnit = HKUnit.count().unitDivided(by: HKUnit.minute())
    let hrType = HKQuantityType(.heartRate)
    let energyType = HKQuantityType(.activeEnergyBurned)
    // need to make this dynamic based on workout eventually
    let distanceType = HKQuantityType(.distanceWalkingRunning)
    
    let avgHr = allstats[hrType]?.averageQuantity()?.doubleValue(for: heartRateUnit) ?? 0
    let totalDistance = allstats[distanceType]?.sumQuantity()?.doubleValue(for: HKUnit.foot()) ?? 0
    let totalEnergy = allstats[energyType]?.sumQuantity()?.doubleValue(for: HKUnit.kilocalorie()) ?? 0
    let healthkitWorkoutId = self.uuid
    let duration = formatTimeInterval(self.duration )
    
    let activityType: HKWorkoutActivityType = self.workoutActivityType
    
    var locationType: String = ""
    
    if let fromMeta = self.metadata?[HKMetadataKeyIndoorWorkout] as? Bool {
      if fromMeta {
        locationType = "indoor"
      } else {
        locationType = "outdoor"
      }
    } else {
      locationType = "unknown"
    }
    
    let finishedWorkoutMessage: [String: Any] = [
      
      "workout_id": self.metadata?[HKMetadataKeyExternalUUID] as? String ?? "",
      "healthkit_workout_id": healthkitWorkoutId.uuidString,
      "healthkit_workout_end_date": self.endDate.intoString ?? "failed to get end_date",
      "healthkit_workout_start_date": self.startDate.intoString ?? "start date not found",
      "average_heart_rate": String(avgHr),
      "distance": totalDistance,
      "duration": String(duration),
      "energy_burned": String(totalEnergy),
      //            "status": "finished",
      //            "source": "Demoios",
      "location_type": locationType,
      "healthkit_workout_activity_type": activityType.name
    ]
    
    //        print("Extracting payload::::\(finishedWorkoutMessage)")
    
    return finishedWorkoutMessage
  }
}
func formatTimeInterval(_ timeInterval: TimeInterval) -> String {
  let formatter = DateComponentsFormatter()
  formatter.unitsStyle = .positional
  formatter.allowedUnits = [.hour, .minute, .second]
  formatter.zeroFormattingBehavior = .pad
  
  return formatter.string(from: timeInterval) ?? ""
}


enum EnvironmentType {
  case debug
  case release
  case preview
  case invalid
}

func getCurrentEnvType() -> EnvironmentType {
  
  if ProcessInfo.processInfo.environment["XCODE_RUNNING_FOR_PREVIEWS"] == "1" {
    return .preview
  }
  
  switch Configuration.environment.absoluteString {
  case "DEBUG":
    return .debug
  case "RELEASE":
    return .release
  default:
    return .invalid
  }
}



#if os(iOS)

func triggerLocalNotification(title: String, body: String, background: Color, icon: NotificationBannerIcons = .info ) {
  dispatchAsync {
    NotificationHandler.shared.new(title: title, body: body, background: background, icon: icon)
  }
}





// func syncDemoiosWorkoutsWithAPI(data: Datamodel) {
//
//    // If any of the workouts from the api response are active, do a query in the healthstore
//    // for a workout matching the healthkitID in the API Workout.
//
//    // if we find one, check if it is completed.
//    // if it IS, then we want to grab the payload and patch it on the api.
//    // if not, it is still active and we want to leave it alone. :)
//
//    getWorkoutsCompletedInDemoios(from: data.health.store) { workouts, error in
//        guard let workouts = workouts, error == nil else {
//            print("error from getting workouts", error ?? "no error")
//            return
//        }
//
//        for apiWorkout in data.myWorkouts.filter({ $0.status == "active" }) {
//
//            print("WE FOUND AN ACTIVE ONE with id \(apiWorkout.id) and HEALTHKID ID \(apiWorkout.healthkit_workout_id)")
//
//            if let matchingHKWorkout = workouts.first(where: { $0.metadata?[HKMetadataKeyExternalUUID] as? String == apiWorkout.id.uuidString }) {
//
//                print("AND NOW WE GOT A MATCHHHH")
//                // we got a match, this means that the active workout in the DB is out of sync with a finished workout in healthkit
//
//                var payloadToPatch = matchingHKWorkout.extractAPIPayload()
//
//                payloadToPatch["source"] = "Demoios"
//
//                data.api.finishWorkoutPatch(payload: payloadToPatch) { _ in
//                    data.api.listMyWorkouts(dataModel: data)
//                }
//            }
//        }
//    }
// }

func listImportableWorkouts(data: Datamodel, completion: @escaping ([ImportableWorkout]?, Error?) -> Void) {
  // get all workouts not sourced from Demoios, and their ID is NOT one of the existing workouts' healthkit id.
  //    print("called list importable workouts")
  
  let workoutType = HKObjectType.workoutType()
  
  // Define the workout source for your custom app
  let workoutSource = HKSource.default()
  
  //    print("Trying to list workouts for Source bundleid::: \(workoutSource.bundleIdentifier) ::: source name:: \(workoutSource.name) ::: source string description: \(workoutSource.description)")
  
  // Define the workout predicate to retrieve workouts saved by your app
  //    let workoutPredicate = HKQuery.predicateForObjects(from: workoutSource)
  
  // Combine the workout type and workout source predicate
  
  //    let sortDescriptors = [NSSortDescriptor(key: HKSampleSortIdentifierStartDate, ascending: true)]
  
  let startDate = Calendar.current.date(byAdding: .day, value: -30, to: Date())
  let endDate = Date()
  let datePredicate = HKQuery.predicateForSamples(withStart: startDate, end: endDate, options: [])
  
  let workoutSourcePredicate = HKQuery.predicateForObjects(from: workoutSource)
  
  let excludedIds: [UUID] = data.myCompletedWorkouts.map { $0.id }
  
  let idPredicate = NSCompoundPredicate(notPredicateWithSubpredicate: NSPredicate(format: "UUID IN %@", argumentArray: [excludedIds]))
  
  let notDemoiosAppPredicate = NSCompoundPredicate(notPredicateWithSubpredicate: workoutSourcePredicate)
  
  let predicate = NSCompoundPredicate(andPredicateWithSubpredicates: [datePredicate, notDemoiosAppPredicate, idPredicate])
  
  // Define the workout query to retrieve the workouts
  let workoutQuery = HKSampleQuery(sampleType: workoutType,
                                   predicate: predicate,
                                   limit: HKObjectQueryNoLimit,
                                   sortDescriptors: nil) { (_, results, error) in
    guard let workouts = results as? [HKWorkout], error == nil else {
      print("Failed to query for importable workouts")
      completion(nil, error)
      return
    }
    
    let resImportableWorkouts = workouts.map { workout in
      let payload = workout.extractAPIPayload()
      
      return ImportableWorkout(
        healthkit_workout_activity_type: payload["healthkit_workout_activity_type"] as? String ?? "",
        healthkit_workout_id: payload["healthkit_workout_id"] as? String ?? "",
        healthkit_workout_start_date: payload["healthkit_workout_start_date"] as? String ?? "",
        healthkit_workout_end_date: payload["healthkit_workout_end_date"] as? String ?? "",
        average_heart_rate: payload["average_heart_rate"] as? String ?? "",
        distance: payload["distance"] as? Double ?? 0.0,
        duration: payload["duration"] as? String ?? "",
        energy_burned: payload["energy_burned"] as? String ?? "",
        //                status: payload["status"] as? String ?? "",
        location_type: payload["location_type"] as? String ?? "",
        source: payload["source"] as? String ?? "import")
    }
    // success!
    
    let workoutListExternalIds: [String] = data.myCompletedWorkouts.map({ $0.workout_data?.healthkit_workout_id ?? "" })
    
    var filtered = resImportableWorkouts.filter({ !workoutListExternalIds.contains($0.healthkit_workout_id) })
    
    filtered.sort(by: { $0.healthkit_workout_start_date > $1.healthkit_workout_start_date })
    
    completion(filtered, nil)
  }
  
  //    print("Executing query")
  data.health.store.execute(workoutQuery)
}
#endif

func ordinalIndicator(for number: Int) -> String {
  let tens = number % 100
  let units = number % 10
  
  if tens >= 11 && tens <= 13 {
    return "\(number)th"
  }
  
  switch units {
  case 1:
    return "\(number)st"
  case 2:
    return "\(number)nd"
  case 3:
    return "\(number)rd"
  default:
    return "\(number)th"
  }
}




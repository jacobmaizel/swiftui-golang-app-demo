//
//  WorkoutController.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/1/23.
//

import Foundation
import HealthKit
import UIKit
import WatchConnectivity
import UserNotifications
import CoreLocation


class WorkoutController: NSObject, ObservableObject, CLLocationManagerDelegate {
  static let shared = WorkoutController()
  
  let locManager = CLLocationManager()
  let healthStore = HKHealthStore()
  
  @Published var locationAuthorized: String?
  @Published var notificationsAuthorized: String?
  @Published var healthStoreAuthorized: String?
  
  var session: HKWorkoutSession?
  
  var builder: HKLiveWorkoutBuilder?
  var routeBuilder: HKWorkoutRouteBuilder?
  
  @Published var errStrings: [String] = []
//  @Published var notiTestStr: String = ""
  
  @Published var externalWorkoutId: String?
  @Published var acceptedWorkoutLeaderPicture: String?
  @Published var acceptedWorkoutLeaderName: String?
  
  // MARK: - Session State Control
  @Published var running = false
  
  // MARK: - Workout Metrics
  @Published var averageHeartRate: Double = 0
  @Published var heartRate: Double = 0
  @Published var activeEnergy: Double = 0
  @Published var distance: Double = 0
  @Published var workout: HKWorkout?
//  @Published var workout_debug: String = "No Error"
  
  @Published var competitions: [[String: String]] = [[:]]
  
  @Published var preConfiguredWorkoutData: [String: String]? {
    didSet {
      self.selectedWorkout = preConfiguredWorkoutData?["workout"]?.workoutActivityType
      
      if let locType = preConfiguredWorkoutData?["location"] {
        if locType == "indoor" {
          self.selectedWorkoutLocationType = .indoor
        } else {
          self.selectedWorkoutLocationType = .outdoor
        }
      }
    }
  }
  
  
  @Published var acceptedInviteWorkoutData: [String: String]? {
    
    /*
     data.send(message: ["start_workout_from_accepted_invite":
     [
     "activityType": workout.healthkit_workout_activity_type.capitalized,
     "dbWorkoutId": workout.id.uuidString,
     "leaderPicture": workout.leader.picture,
     "leaderName": workout.leader.first_name
     ]
     ])
     */
    didSet {
      self.selectedWorkout = acceptedInviteWorkoutData?["activityType"]?.workoutActivityType
      self.externalWorkoutId = acceptedInviteWorkoutData?["dbWorkoutId"]
      self.acceptedWorkoutLeaderName = acceptedInviteWorkoutData?["leaderName"]
      self.acceptedWorkoutLeaderPicture = acceptedInviteWorkoutData?["leaderPicture"]
    }
  }
  
  
  @Published var selectedWorkout: HKWorkoutActivityType? {
    didSet {
      print("Selected workout \(selectedWorkout?.name ?? "")")
      guard selectedWorkout != nil else {
        return
      }
    }
  }

  @Published var selectedCompetition: [String: String]? {
    didSet {
      print("Selected COMP \(String(describing: selectedCompetition))")
      guard selectedCompetition != nil else {
        return
      }
    }
  }
  
  @Published var selectedWorkoutLocationType: HKWorkoutSessionLocationType = .indoor
  
  @Published var showingSummaryView: Bool = false {
    didSet {
      if showingSummaryView == false {
        resetWorkout()
      }
    }
  }
  
  
  override init() {
    super.init()
    locManager.delegate = self
  }
  

  // Trigger a workout session
  func startConfirmedWorkout(externalId: String) {
    if let workout = selectedWorkout {
      startWorkout(workoutType: workout, externalId: externalId)
    } else {
      return
    }
  }
  
  func stopWorkoutHandshake() {
    //
  }
  
  func requestCompetitionList() {
    
    DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
      sendWCMessage(message: ["requesting_competitions": "true"])
    }
  }
  
  func startWorkoutFromAcceptedInviteState() {
    
    guard let wID = externalWorkoutId, let actType = selectedWorkout else {
      
      self.errStrings.append("FAILED TO GET EXISTING WORKOUT ID /activity type IN startWorkoutFromACceptedInviteState function")
      return
    }
    
    startWorkout(workoutType: actType, externalId: wID)
    
  }
  
  func initiateWorkoutStartHandshake() {
    // send msg to iphone and let it know to hit api and get a fresh workout ID
    //    let compId = selectedCompetition?["id"]
    if let workout = selectedWorkout {
      
      // start workout start handshake
      sendWCMessage(message:
                      ["workout_start": ["payload": ["healthkit_workout_activity_type": workout.name] ]])
      
      //      startWorkout(workoutType: workout, externalId: wID)
    }
  }

  
  // Start the workout.
  private func startWorkout(workoutType: HKWorkoutActivityType, externalId: String) {
    let configuration = HKWorkoutConfiguration()
    configuration.activityType = workoutType
    configuration.locationType = selectedWorkoutLocationType
    
    // if it is an outdoor workout, start tracking user location IF we have perm
    
    if locManager.authorizationStatus.rawValue > 2 && selectedWorkoutLocationType == .outdoor {
      
      locManager.desiredAccuracy = kCLLocationAccuracyBest
      locManager.startUpdatingLocation()
      routeBuilder = HKWorkoutRouteBuilder(healthStore: healthStore, device: nil)
      
    }
    
    
    // Create the session and obtain the workout builder.
    do {
      session = try HKWorkoutSession(healthStore: healthStore, configuration: configuration)
      builder = session?.associatedWorkoutBuilder()
    } catch {
      // Handle any exceptions.
      
      self.errStrings.append("Error in startWorkout function: \(error)")
      return
    }
    
    // Setup session and builder.
    session?.delegate = self
    builder?.delegate = self
    
    // Set the workout builder's data source.
    builder?.dataSource = HKLiveWorkoutDataSource(healthStore: healthStore,
                                                  workoutConfiguration: configuration)
    
    // Start the workout session and begin data collection.
    let startDate = Date()
    
    session?.startActivity(with: startDate)
    
    // Tell iphone app that we have started a workout BEFORE BEGINNING COLLECTION
    
    // wait for the specific
    
    // Add UUID for workout  from api call here to the HK External key ID.
    builder?.addMetadata(
      [
        HKMetadataKeyExternalUUID: externalId,
        "location_type": selectedWorkoutLocationType.getString()
      ]
    ) { (success, err) in
      guard err == nil, success else {
        print("something went wrong adding external ID to workout \(success) \(String(describing: err))")
        return
      }
    }
    
    builder?.beginCollection(withStart: startDate) { (success, error) in
      // The workout has started.
      
      guard success else {
        print("failed to begin workout data collection")
        self.errStrings.append("\("failed to begin workout data collection" + "\(error)" )")
        return
      }
    }
  }
  
  // Request authorization to access HealthKit.
  func requestAuthorizationHealthkit(completion: @escaping (Bool) -> Void) {
    // The quantity type to write to the health store.
    let typesToShare: Set = [
      HKQuantityType.workoutType(),
      HKSeriesType.workoutRoute()
    ]
    
    WorkoutController.shared.errStrings.append("Starting req for health kit auth")
    
    // The quantity types to read from the health store.
    let typesToRead: Set = [
      HKObjectType.workoutType(),
      HKSeriesType.workoutRoute(),
      HKQuantityType.quantityType(forIdentifier: .heartRate)!,
      HKQuantityType.quantityType(forIdentifier: .activeEnergyBurned)!,
      HKQuantityType.quantityType(forIdentifier: .appleExerciseTime)!,
      HKQuantityType.quantityType(forIdentifier: .distanceWalkingRunning)!,
      HKQuantityType.quantityType(forIdentifier: .stepCount)!,
      HKQuantityType.quantityType(forIdentifier: .flightsClimbed)!,
      HKQuantityType.quantityType(forIdentifier: .basalEnergyBurned)!,
      HKQuantityType.quantityType(forIdentifier: .appleStandTime)!,
      HKQuantityType.quantityType(forIdentifier: .vo2Max)!,
      HKQuantityType.quantityType(forIdentifier: .distanceCycling)!,
      HKQuantityType.quantityType(forIdentifier: .distanceSwimming)!,
      HKObjectType.activitySummaryType()
      //            HKQuantityType.quantityType(forIdentifier: .heartRate)!,
      //            HKQuantityType.quantityType(forIdentifier: .activeEnergyBurned)!,
      //            HKQuantityType.quantityType(forIdentifier: .distanceWalkingRunning)!,
      //            HKQuantityType.quantityType(forIdentifier: .appleExerciseTime)!,
      //            HKQuantityType.quantityType(forIdentifier: .distanceCycling)!,
      //            HKQuantityType.quantityType(forIdentifier: .distanceSwimming)!,
      //            HKObjectType.activitySummaryType()
    ]
    
    // Request authorization for those quantity types.
    healthStore.requestAuthorization(toShare: typesToShare, read: typesToRead) { (granted, error) in
      // Handle error.
      
      
      guard error == nil else {
        self.errStrings.append("Failed to request healthstore auth \(error)")
        completion(false)
        return
      }

      completion(granted)
      
    }
  }
  func togglePause() {
    if running == true {
      self.pause()
    } else {
      resume()
    }
  }
  
  func pause() {
    session?.pause()
  }
  
  func resume() {
    session?.resume()
  }
  
  func endWorkout() {
    session?.end()
  }

  func resetWorkout() {
    selectedWorkout = nil
    selectedCompetition = nil
    builder = nil
    routeBuilder = nil
    workout = nil
    session = nil
    activeEnergy = 0
    averageHeartRate = 0
    heartRate = 0
    distance = 0
    acceptedInviteWorkoutData = nil
    acceptedWorkoutLeaderName = nil
    acceptedWorkoutLeaderPicture = nil
  }
  
  func deleteWorkout(with uuid: UUID) {
    let predicate = HKQuery.predicateForObject(with: uuid)
    let workoutQuery = HKSampleQuery(sampleType: HKObjectType.workoutType(),
                                     predicate: predicate,
                                     limit: HKObjectQueryNoLimit,
                                     sortDescriptors: nil) { (_, samples, error) in
      if let workout = samples?.first as? HKWorkout {
        self.healthStore.delete(workout) { (success, error) in
          if success {
            print("Workout deleted successfully.")
            //            self.workout_debug = "workout successfully deleted"
          } else {
            if let deleteError = error {
              print("Failed to delete workout: \(deleteError.localizedDescription)")
            } else {
              print("Failed to delete workout: unknown error")
            }
          }
        }
      } else {
        print("Workout not found.")
      }
    }
    healthStore.execute(workoutQuery)
  }
  
  func deleteWorkoutsForCurrentSource() {
    
    let predicate = HKQuery.predicateForObjects(from: HKSource.default())
    let workoutQuery = HKSampleQuery(sampleType: HKObjectType.workoutType(),
                                     predicate: predicate,
                                     limit: HKObjectQueryNoLimit,
                                     sortDescriptors: nil) { (_, samples, error) in
      if let workouts = samples as? [HKWorkout], !workouts.isEmpty {
        self.healthStore.delete(workouts) { (success, error) in
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
    
    healthStore.execute(workoutQuery)
  }
}

// MARK: - HKWorkoutSessionDelegate
extension WorkoutController: HKWorkoutSessionDelegate {
  func workoutSession(_ workoutSession: HKWorkoutSession,
                      didChangeTo toState: HKWorkoutSessionState,
                      from fromState: HKWorkoutSessionState,
                      date: Date) {
    
    DispatchQueue.main.async {
      self.running = toState == .running
//      self.errStrings.append("debu")
    }
    
    // Wait for the session to transition states before ending the builder.
    if toState == .ended {
      builder?.endCollection(withEnd: date) { (_, _) in
        
        self.builder?.finishWorkout { (workout, err) in
          
          if err != nil {
            self.errStrings.append(err?.localizedDescription ?? "ERROR in end collection / finish workout\(String(describing: err))")
            return
          }
          
          let payload = workout?.extractAPIPayload()
          
          sendWCMessage(message: ["workout_finished": payload as Any] )
          
          
          // Finish workout route if there was one set up
          if self.routeBuilder != nil {
            
            self.locManager.stopUpdatingLocation()
            
            let extId = workout?.metadata?[HKMetadataKeyExternalUUID]
            let routeMeta: [String: Any] = [HKMetadataKeyExternalUUID: extId as Any]
            self.routeBuilder?.finishRoute(with: workout!, metadata: routeMeta) { route, err in
              
              guard route != nil && err == nil else {
                self.errStrings.append("Error in route builder completion in end workout: \(String(describing: err))")
                
                sendWCMessage(message: [
                  "route_processed": [
                  
                    "success": false,
                    "message": "Something went wrong when processing your workout route.\(String(describing: Configuration.environment.absoluteString == "DEBUG" ? err?.localizedDescription : ""))"
                  ] as [String : Any]
                ])
                return
              }
              if let route = route {
                sendWCMessage(message: [
                  "route_processed": [
                    "success": true,
                    "message": "Your workout route has finished processing and is ready to view."
                  ] as [String : Any]
                ])
                
              }
              
            }
          }
          
          
          
          DispatchQueue.main.async {
            self.workout = workout
            if Configuration.environment.absoluteString == "DEBUG" {
              //                            self.deleteWorkout(with: workout?.uuid ?? UUID())
            }
          }
        } // Builder.Finishworkout
      }
    }
  }
  
  func workoutSession(_ workoutSession: HKWorkoutSession, didFailWithError error: Error) {
    DispatchQueue.main.async {
      self.errStrings.append(error.localizedDescription)
    }
  }
}

// MARK: - HKLiveWorkoutBuilderDelegate
extension WorkoutController: HKLiveWorkoutBuilderDelegate {
  
  func updateForStatistics(_ statistics: HKStatistics?) {
    guard let statistics = statistics else {
      return
    }
    
    DispatchQueue.main.async {
      switch statistics.quantityType {
      case HKQuantityType.quantityType(forIdentifier: .heartRate):
        let heartRateUnit = HKUnit.count().unitDivided(by: HKUnit.minute())
        self.heartRate = statistics.mostRecentQuantity()?.doubleValue(for: heartRateUnit) ?? 0
        self.averageHeartRate = statistics.averageQuantity()?.doubleValue(for: heartRateUnit) ?? 0
      case HKQuantityType.quantityType(forIdentifier: .activeEnergyBurned):
        let energyUnit = HKUnit.kilocalorie()
        self.activeEnergy = statistics.sumQuantity()?.doubleValue(for: energyUnit) ?? 0
      case HKQuantityType.quantityType(forIdentifier: .distanceWalkingRunning), HKQuantityType.quantityType(forIdentifier: .distanceCycling):
        let meterUnit = HKUnit.meter()
        self.distance = statistics.sumQuantity()?.doubleValue(for: meterUnit) ?? 0
      default:
        return
      }
    }
  }
  
  func workoutBuilderDidCollectEvent(_ workoutBuilder: HKLiveWorkoutBuilder) {
  }
  
  func workoutBuilder(_ workoutBuilder: HKLiveWorkoutBuilder, didCollectDataOf collectedTypes: Set<HKSampleType>) {
    for type in collectedTypes {
      guard let quantityType = type as? HKQuantityType else {
        return // Nothing to do.
      }
      
      let statistics = workoutBuilder.statistics(for: quantityType)
      updateForStatistics(statistics)
    }
  }
}

//
//  Utils+HK.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/7/23.
//

import Foundation

import HealthKit
import CoreLocation


//
/*
 case HKQuantityType.quantityType(forIdentifier: .heartRate):
 let heartRateUnit = HKUnit.count().unitDivided(by: HKUnit.minute())
 self.heartRate = statistics.mostRecentQuantity()?.doubleValue(for: heartRateUnit) ?? 0
 self.averageHeartRate = statistics.averageQuantity()?.doubleValue(for: heartRateUnit) ?? 0
 */


func workoutRouteQuery(healthStore: HKHealthStore = HealthService.shared.store, for hkIdentifier: HKQuantityTypeIdentifier, fromWorkout hkWorkoutId: UUID?, completion: @escaping ([WorkoutGraphDataPoint]) -> Void) {
  
  guard let hkWorkoutId = hkWorkoutId else {
    print("WORKOUT ID REQUIRED")
    return
  }
  
  //  print("Starting health store query with id: \(hkWorkoutId.uuidString)")
  
  let workouts = HKObjectType.workoutType()
  
  let workoutQuery = HKSampleQuery(sampleType: workouts,
                                   predicate: HKQuery.predicateForObject(with: hkWorkoutId),
                                   limit: 1,
                                   sortDescriptors: nil) { (query, samples, error) in
    
    if let error = error {
      print("Error fetching workout: \(error)")
      return
    }
    
    guard let workout = samples?.first as? HKWorkout else {
      print("No workout found")
      return
    }
    
    let predicate = HKQuery.predicateForSamples(withStart: workout.startDate,
                                                end: workout.endDate,
                                                options: .strictStartDate)
    
    let heartRateType = HKQuantityType.quantityType(forIdentifier: hkIdentifier)!
    
    
    let heartRateQuantityUnit: HKUnit = HKUnit.count().unitDivided(by: HKUnit.minute())
    
    let query = HKSampleQuery(sampleType: heartRateType,
                              predicate: predicate,
                              limit: HKObjectQueryNoLimit,
                              sortDescriptors: nil) { (_, samples, error) in
      if let error = error {
        print("Error fetching heart rates: \(error)")
        return
      }
      
      guard let heartRates = samples as? [HKQuantitySample] else {
        print("No heart rates found")
        return
      }
      
      
      var res: [WorkoutGraphDataPoint] = []
      
      for heartRate in heartRates {
        //        print("QUANTITY::: \(heartRate.quantity.doubleValue(for: heartRateQuantityUnit)) COUNT:::: \(heartRate.count), END DATE::: \(heartRate.endDate)")
        
        let new: WorkoutGraphDataPoint = WorkoutGraphDataPoint(value: heartRate.quantity.doubleValue(for: heartRateQuantityUnit), date: heartRate.endDate)
        
        res.append(new)
      }
      
      completion(res)
    }
    
    healthStore.execute(query)
  }
  
  healthStore.execute(workoutQuery)
}

struct WorkoutGraphDataPoint: Identifiable {
  var id = UUID()
  let value: Double
  let date: Date
}

func healthStoreWorkoutRouteQuery(healthStore: HKHealthStore = HealthService.shared.store, fromWorkout hkWorkoutId: UUID?, completion: @escaping ([CLLocation]) -> Void, finished: @escaping (Bool) -> Void) {
  
  guard let hkWorkoutId = hkWorkoutId else {
    print("WORKOUT ID REQUIRED")
    return
  }
  
  //  print("Starting health store query with id: \(hkWorkoutId.uuidString)")
  
  //MARK: - START WorkoutQuery
  let workouts = HKObjectType.workoutType()
  
  let workoutQuery = HKSampleQuery(sampleType: workouts,
                                   predicate: HKQuery.predicateForObject(with: hkWorkoutId),
                                   limit: 1,
                                   sortDescriptors: nil) { (query, samples, error) in
    
    if let error = error {
      print("Error fetching workout: \(error)")
      return
    }
    
    guard let workout = samples?.first as? HKWorkout else {
      print("No workout found")
      return
    }
    
    //MARK: - START RouteQuery
    let workoutPredicate = HKQuery.predicateForObjects(from: workout)
    let routeQuery = HKSampleQuery(sampleType: HKSeriesType.workoutRoute(), predicate: workoutPredicate, limit: HKObjectQueryNoLimit, sortDescriptors: nil) { query, samples, errorOrNil in
      
      if let error = error {
        print("Error fetching workout route: \(error)")
        return
      }
      
      guard let routeData = samples?.first as? HKWorkoutRoute else {
        print("No route sample found found")
        return
      }
      
      
      //MARK: - START RouteLocationDataQuery
      
      let routeLocationDataQuery = HKWorkoutRouteQuery(route: routeData) { (query, locationsOrNil, done, errorOrNil) in
        
        // This block may be called multiple times.
        
        if let error = errorOrNil {
          // Handle any errors here.
          return
        }
        
        guard let locations = locationsOrNil else {
          fatalError("*** Invalid State: This can only fail if there was an error. ***")
        }
        
        // Do something with this batch of location data.
        completion(locations)
        
        if done {
          
          finished(true)
          // The query returned all the location data associated with the route.
          // Do something with the complete data set.
          
          
          // You can stop the query by calling:
        } else {
          finished(false)
        }
        //        healthStore.stop(routeLocationDataQuery)
        
        
      } // route data query
      healthStore.execute(routeLocationDataQuery)
      // MARK: - END RouteLocationDataQuery
      
    } // route query
    healthStore.execute(routeQuery)
    // MARK: - END RouteQuery
    
    
  } // workout query
  healthStore.execute(workoutQuery)
  // MARK: - END WORKOUTQUERY
  
}

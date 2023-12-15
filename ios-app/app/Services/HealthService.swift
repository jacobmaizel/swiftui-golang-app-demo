//
//  HealthService.swift
//  Demoios
//
//  Created by Jacob Maizel on 4/1/23.
//

import Foundation
import HealthKit

class HealthService {
    
    static let shared = HealthService()
    
    let store: HKHealthStore = HKHealthStore()
    
    let typesToShare: Set = [HKQuantityType.workoutType(), HKSeriesType.workoutRoute(),]
    
    let typesToRead: Set<HKObjectType> = [
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
    ]
    
    private init() {}

 
    func requestHealthKitAuthorization(completion: @escaping (Bool) -> Void) {
        // Check if HealthKit is available on the device
        guard HKHealthStore.isHealthDataAvailable() else {
            print("HealthKit is not available on this device.")
            completion(false)
            return
        }
 
        // Request authorization from the user
        store.requestAuthorization(toShare: typesToShare, read: typesToRead) { success, error in
            if let error = error {
                print("Error requesting HealthKit authorization: \(error.localizedDescription)")
                completion(false)
                return
            }
            
            if success {
                print("HealthKit authorization request succeeded.")
                completion(true)
            } else {
                print("HealthKit authorization request failed.")
                completion(false)
            }
        }
    }
}

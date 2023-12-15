//
//  WorkoutController+Location.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 8/5/23.
//

import Foundation
import CoreLocation

extension WorkoutController {
  func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
    //      print("Loc man did change auth!!! now set to:::  \(manager.authorizationStatus)")
    
    self.errStrings.append("locationManagerDidChangeAuthorization:::\(manager.authorizationStatus)")
    
    self.locationAuthorized = manager.authorizationStatus.rawValue > 2 ? "Got Auth lvl \(manager.authorizationStatus.rawValue)" : "Auth denied, status \(manager.authorizationStatus.rawValue)"
  }
  
  func locationManager(_ manager: CLLocationManager, didFailWithError error: Error) {
    print("Location requets error:::\(error)")
    self.errStrings.append("Location requets error:::\(error)")
    //    triggerLocalNotification(title: "Failed to get location info", body: error.localizedDescription, background: .red, icon: .warning)
  }
  
  func locationManager(_ manager: CLLocationManager, didUpdateLocations locations: [CLLocation]) {
    
    // Filter the raw data.
    let filteredLocations = locations.filter { (location: CLLocation) -> Bool in
      location.horizontalAccuracy <= 50.0
    }
    
    guard !filteredLocations.isEmpty else { return }
    
    // Add the filtered data to the route.
    routeBuilder?.insertRouteData(filteredLocations) { (success, error) in
      if !success {
        // Handle any errors here.
      }
    }
  }
  
  
}

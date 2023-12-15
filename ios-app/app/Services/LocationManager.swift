//
//  LocationManager.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/5/23.
//

import Foundation
import CoreLocation

class LocationManager: NSObject, ObservableObject, CLLocationManagerDelegate {
  static let shared = LocationManager()
  let manager = CLLocationManager()
  
  @Published var status: CLAuthorizationStatus?
  
  @Published var location: CLLocationCoordinate2D?
  
  @Published var placemark: CLPlacemark?
  @Published var city: String?
  
  override init() {
    super.init()
    manager.delegate = self
  }
  
  func askForWhenInUsePermission() {
    manager.requestWhenInUseAuthorization()
  }
  
  func getLocation() {
    manager.requestLocation()
    location = manager.location?.coordinate
    lookUpCurrentLocation { _ in
      
    }
   
  }
  
  func hasPerms() -> Bool {
    
    guard let status = status else {
      return false
    }
    
    return status.rawValue > 2
  }
  
  func locationManagerDidChangeAuthorization(_ manager: CLLocationManager) {
    print("Loc man did change auth!!! now set to:::  \(manager.authorizationStatus)")
    
    status = manager.authorizationStatus
  }
  
  func locationManager(_ manager: CLLocationManager, didFailWithError error: Error) {
    print("Location requets error:::\(error)")
//    triggerLocalNotification(title: "Failed to get location info", body: error.localizedDescription, background: .red, icon: .warning)
  }
  
  func locationManager(_ manager: CLLocationManager, didUpdateLocations locations: [CLLocation]) {
    location = locations.first?.coordinate
//    print("Got location update:::\(locations.first?.coordinate)")
  }
  
  
  func lookUpCurrentLocation(completionHandler: @escaping (CLPlacemark?)
                             -> Void ) {
    // Use the last reported location.
    if let lastLocation = self.manager.location {
      let geocoder = CLGeocoder()
      
      // Look up the location and pass it to the completion handler
      geocoder.reverseGeocodeLocation(lastLocation,
                                      completionHandler: { (placemarks, error) in
        if error == nil {
          let firstLocation = placemarks?[0]
          completionHandler(firstLocation)
          self.placemark = firstLocation
          self.city = firstLocation?.locality
        } else {
          // An error occurred during geocoding.
          print("Failed to geocode!!")
          completionHandler(nil)
        }
      })
    } else {
      // No location was available.
      print("NO LOCATION WAS AVAILABLE")
      completionHandler(nil)
    }
  }
  
  
  
  func getCoordinate( addressString: String,
                      completionHandler: @escaping(CLLocationCoordinate2D, NSError?) -> Void ) {
    let geocoder = CLGeocoder()
    geocoder.geocodeAddressString(addressString) { (placemarks, error) in
      if error == nil {
        if let placemark = placemarks?[0] {
          let location = placemark.location!
          
          completionHandler(location.coordinate, nil)
          return
        }
      }
      
      completionHandler(kCLLocationCoordinate2DInvalid, error as NSError?)
    }
  }
}

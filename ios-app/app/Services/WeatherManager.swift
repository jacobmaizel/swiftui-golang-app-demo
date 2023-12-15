//
//  WeatherService.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/5/23.
//

import Foundation
import WeatherKit
import CoreLocation


@MainActor class WeatherManager: ObservableObject {
  
  static let shared = WeatherManager()
  
  @Published var weather: Weather?
  
  @Published var lastRequested: Date?
  
  func getWeather(for coords: CLLocationCoordinate2D) async {
    
    do {
      weather = try await Task.detached(priority: .userInitiated) {
        return try await WeatherService.shared.weather(for: .init(latitude: coords.latitude, longitude: coords.longitude))
      }.value
    } catch {
      print("\(error)")
    }
  }
  
  var symbol: String {
    weather?.currentWeather.symbolName ?? ""
  }
  
  var rawTempF: String {
    let temp = weather?.currentWeather.temperature
    let convert = temp?.converted(to: .fahrenheit).description
    return convert?.trimmingCharacters(in: .whitespacesAndNewlines) ?? ""
    
  }
  
  var tempF: String {
    let temp = weather?.currentWeather.temperature
    
    let convert: String = temp?.converted(to: .fahrenheit).description.replacingOccurrences(of: "°F", with: "").trimmingCharacters(in: .whitespacesAndNewlines) ?? ""
    
    if let tmp = Double(convert) {
      let re = Int(tmp.rounded())
      
//      print("returning tmp \(tmp)")
      return "\(re) °F"
    }
    
//    print("Returning convert \(convert)")
    
    return convert.isEmpty ? "" : convert + "°F"
  }
}

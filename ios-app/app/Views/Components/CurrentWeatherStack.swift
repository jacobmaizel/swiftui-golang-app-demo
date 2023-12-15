//
//  WeatherStack.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/5/23.
//

import SwiftUI



struct CurrentWeatherStack: View {
  
  @EnvironmentObject var weather: WeatherManager
  @EnvironmentObject var loc: LocationManager
  
  var stack: StackType
  
  @ViewBuilder
  var body: some View {
    Group {
      

        if stack == .v {
          
          VStack(spacing: 2) {
            Image(systemName: weather.symbol)
            Text("\(weather.tempF)")
          }
        } else {
          HStack {
            Image(systemName: weather.symbol)
            Text("\(weather.tempF)\(loc.city != nil ? " in " : "")\(loc.city ?? "")")
          }
        }
      
      
    }
    .DemoiosText(fontSize: .base)
    .task {
      if getCurrentEnvType() != .preview && loc.status?.rawValue ?? 0 > 2 {
        loc.getLocation()
        if let loc = loc.location {
          
          if weather.weather == nil {
            await weather.getWeather(for: loc)
            
            dispatchAsync {
              weather.lastRequested = Date.now
            }
          } else if let lastRequested = weather.lastRequested {
            
            // if there WAS a last requested date, is it > 60*5 ago?
            
            // time since last requested
            let timeSince = Date().timeIntervalSince(lastRequested)
            
            if timeSince >= 60 * 5 {
              //            print("REFRESHING WEATHER:: GREATER THAN 5 sec \n\n\n")
              await weather.getWeather(for: loc)
              
              dispatchAsync {
                weather.lastRequested = Date.now
              }
            }
            
          } else {
            
            //          print("GETTING WEATHER FOR THE FIRST TIME:::::::::")
            
            if getCurrentEnvType() != .preview {
              await weather.getWeather(for: loc)
              
              dispatchAsync {
                weather.lastRequested = Date.now
              }
            }
          }
        }
      }
    }
  }
}

struct WeatherStack_Previews: PreviewProvider {
  static var previews: some View {
    CurrentWeatherStack(stack: .h)
      .environmentObject(WeatherManager())
      .environmentObject(LocationManager())
      .setupPreview()
  }
}



//
//  Config.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/21/23.
//

import Foundation

struct Config: Identifiable, Codable {
  var id = UUID()
  
  var auto_sync_workouts: Bool
}

extension Config {
  static let `default` = Config(auto_sync_workouts: true)
}

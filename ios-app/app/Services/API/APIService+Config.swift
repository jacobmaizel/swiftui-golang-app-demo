//
//  APIService+Config.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/21/23.
//

import Foundation

extension DemoiosAPIService {
  
  func getConfig() async throws -> Config {
    let config: Config = try await asyncRequest(method: .get, path: "/config")
    return config
  }
  
  func updateConfig(payload: [String: Any]) async throws -> Config {
    let config: Config = try await asyncRequest(method: .patch, path: "/config/", payload: payload)
    return config
  }
}

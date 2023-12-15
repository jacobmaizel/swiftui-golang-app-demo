//
//  APIService+Search.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/21/23.
//

import Foundation

enum SearchableResourceType: String {
  case squads
  case profiles
  case competitions
}


extension DemoiosAPIService {
  
  func search<T: Decodable & SearchResultStructure>(for resourceType: SearchableResourceType, searchTerm: String) async throws -> T where T.ResultType: Decodable {
    
    let queryItems: [URLQueryItem] = [URLQueryItem(name: "type", value: resourceType.rawValue), URLQueryItem(name: "term", value: searchTerm)]
    
    do {
      let res: T = try await asyncRequest(method: .get, path: "/services/search", queryItems: queryItems)
      return res
    } catch {
      print(error)
      throw error
    }
  }
}

protocol SearchResultStructure {
  associatedtype ResultType
  
  var count: Int { get set }
  var results: ResultType { get set }
}


struct ProfileSearchRes: Decodable, SearchResultStructure {
  var count: Int
  var results: [PartialProfile]
}

struct SquadSearchRes: Decodable, SearchResultStructure {
  var count: Int
  var results: [Squad]
}

struct CompSearchRes: Decodable, SearchResultStructure {
  var count: Int
  var results: [Competition]
}

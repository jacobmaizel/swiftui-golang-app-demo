//
//  APIService.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/30/23.
//

import Foundation
import Auth0

// enum that conforms to Error?
enum APIError: Error {
  case invalidURL
  case invalidResponse
  case unauthorized
  case noData
  case decodingError
  case error(NSError)
}

enum HTTPMethod: String {
  case get = "GET"
  case post = "POST"
  case patch = "PATCH"
  case delete = "DELETE"
}

protocol APIService {
}
class DemoiosAPIService: APIService {
  
  static let shared = DemoiosAPIService()
  
  static let audience = Configuration.authAudience.absoluteString
  static let clientId = Configuration.authClientId.absoluteString
  static let domain = Configuration.authDomain.absoluteString
  
  let auth0Manager = CredentialsManager(authentication: Auth0.authentication(clientId: clientId, domain: domain))
  
  let webAuthClient = Auth0.webAuth(clientId: clientId, domain: domain).audience(audience).scope("openid profile offline_access")
  
  let baseURL = Configuration.apiBaseURL.absoluteString
  
  private let urlSession = URLSession.shared
  private let decoder = JSONDecoder()
  
  private init() {}
  
  func createURL(path: String, with queryItems: [URLQueryItem]? ) -> URL? {
    guard var url = URLComponents(string: baseURL) else {
      print("INSIDE CREATEURL FUNCTION::: Unable to guard, failed to create url from baseurl")
      return nil
    }
    
    if queryItems != nil {
      url.queryItems = queryItems
    }
    
    url.path = path
    //        print("API CALL:::", url)
    return url.url
  }
  
  func genRequest(path: String, method: HTTPMethod = .get, payload: [String: Any]? = nil, queryItems: [URLQueryItem]? = nil ) -> URLRequest? {
    
//            print("\n\ngenerating request for \(path)\n\n")
    
    var accessTokenCred: String?
    
    //        let methodNeedsPayload = ["PATCH", "POST", "PUT"]
    
    guard let genURL = createURL(path: path, with: queryItems) else {
      print("createURL Failed for \(path) in genRequest")
      return nil
    }
    
    var request = URLRequest(url: genURL)
    
    let semaphore = DispatchSemaphore(value: 0)
    
    auth0Manager.credentials { result in
      
      switch result {
      case(.success(let creds)):
        accessTokenCred = creds.accessToken
        //                                print("Creds found in genrequest for \(path)")
        //                                print(":::",creds.accessToken,":::")
        semaphore.signal()
        
      case(.failure(let err)):
        print("Failed in genRequest::::\(err):::while trying to hit \(path)")
        semaphore.signal()
      }
    }
    
    semaphore.wait()
    
    guard let token = accessTokenCred else {
      print("Unable to guard against nil access token in genRequest")
      return nil
    }
    
//    print("got token in rq builder")
    
    // Create request
    request.addValue("Bearer \(token)", forHTTPHeaderField: "Auth")
    request.httpMethod = method.rawValue
    
    if let nonNilPayload = payload {
      do {
        
        request.httpBody = try JSONSerialization.data(withJSONObject: nonNilPayload)
      } catch {
        print("Unable to serialize and attach the payload for the \(method.rawValue) request to path \(path)")
      }
    }
    
    request.addValue("application/json", forHTTPHeaderField: "Content-Type")
    
//            print("Returning created request", request)
    
    return request
  }
  
  func sendRequest<T: Decodable>(method: HTTPMethod, path: String, payload: [String: Any]? = nil, queryItems: [URLQueryItem]? = nil, completion: @escaping (Result<T, APIError>) -> Void) {
    guard let req = genRequest(path: path, method: method, payload: payload, queryItems: queryItems) else {
      return
    }
    print("API CALL::: \(req.httpMethod ?? "NO METHOD") \(req.description)")
    loadAndDecode(with: req) { (result: Result<T, APIError>) in
      switch result {
      case .success(let res):
        completion(.success(res))
      case .failure(let error):
        completion(.failure(error))
      }
    }
  }
  
  
  func loadAndDecode<D: Decodable>(with req: URLRequest, completion: @escaping (Result<D, APIError>) -> Void) {
    urlSession.dataTask(with: req) { (data, response, error) in
      if let error = error {
        completion(.failure(.error(error as NSError)))
        return
      }
      
      guard let httpRes = response as? HTTPURLResponse else {
        completion(.failure(.invalidResponse))
        return
      }
      
      switch httpRes.statusCode {
      case 200...299:
        guard let data = data else {
          completion(.failure(.noData))
          return
        }
        do {
          let decodedData = try self.decoder.decode(D.self, from: data)
          completion(.success(decodedData))
        } catch let DecodingError.dataCorrupted(context) {
          print("\(req) Data corrupted: \(context)")
        } catch let DecodingError.keyNotFound(key, context) {
          print("Key '\(key)' not found: \(context.debugDescription)")
          print("codingPath: \(context.codingPath)")
        } catch let DecodingError.valueNotFound(value, context) {
          print("Value '\(value)' not found: \(context.debugDescription)")
          print("codingPath: \(context.codingPath)")
        } catch let DecodingError.typeMismatch(type, context) {
          print("Type '\(type)' mismatch: \(context.debugDescription)")
          print("codingPath: \(context.codingPath)")
        } catch {
          print("Other error: \(error)")
        }
      case 401:
        completion(.failure(.unauthorized))
      default:
        completion(.failure(.invalidResponse))
      }
    }.resume()
  }
  
  
  func asyncLoadAndDecode<D: Decodable>(with req: URLRequest, for type: D.Type) async throws -> D {
    let (data, _) = try await urlSession.data(for: req)
    
    do {
      let decodedData = try self.decoder.decode(D.self, from: data)
      return decodedData
    } catch let DecodingError.dataCorrupted(context) {
      print("\(req) Data corrupted: \(context)")
    } catch let DecodingError.keyNotFound(key, context) {
      print("Key '\(key)' not found: \(context.debugDescription)")
      print("codingPath: \(context.codingPath)")
    } catch let DecodingError.valueNotFound(value, context) {
      print("Value '\(value)' not found: \(context.debugDescription)")
      print("codingPath: \(context.codingPath)")
    } catch let DecodingError.typeMismatch(type, context) {
      print("Type '\(type)' mismatch: \(context.debugDescription)")
      print("codingPath: \(context.codingPath)")
    } catch {
      print("Other error: \(error)")
    }
    
    throw DemoiosError.runtimeError("Failed to decode for req: \(String(describing: req.url))")
  }
  
  func asyncRequest<T: Decodable>(method: HTTPMethod, path: String, payload: [String: Any]? = nil, queryItems: [URLQueryItem]? = nil) async throws -> T {
    
    guard let request = genRequest(path: path, method: method, payload: payload, queryItems: queryItems) else {
      print("FAILED TO CREATE REQUEST")
      throw DemoiosError.runtimeError("Failed to gen request")
    }
    
    print("[ Async: \(request.httpMethod ?? "NO METHOD") \(request.url?.absoluteString ?? "") ]")
    
    let data = try await asyncLoadAndDecode(with: request, for: T.self)
    
    return data
    
  }
  
  @MainActor
  func asyncLoadInitialData(data: Datamodel = Datamodel.shared) async {
    async let notis: () = getMyNotifications()
    async let squads: () = listMySquads()
    async let profileStats: () = getMyProfileStats()
    async let compList: () = getMyCompetitions()
    async let workoutList: () = getMyCompletedWorkouts()

    await notis
    await workoutList
    await profileStats
    await compList
    await squads
    
    data.initialDataLoaded = true
    
  }
  
  func loadNonSelfProfileData(data: Datamodel, profileId: UUID) async throws -> (ProfileStatsData?, [Squad]?, [Competition]?, [Workout]?) {
    
    async let profStats: ProfileStatsData? = getProfileStatsById(profileId: profileId)
    
    // Competition list
    async let compList: [Competition]? = listCompetitionsByProfile(profileId: profileId)
    
    // Squad list
    async let squadList = listSquadsByProfile(profileId: profileId)
    
    // Workout List
    
    async let workoutList = listWorkoutsByProfile(profileId: profileId)
    
    
    return try await (profStats, squadList, compList, workoutList)
    
  }
} // APIService

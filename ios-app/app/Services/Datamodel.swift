//
//  Datamodel.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/25/23.
//

import Foundation
import HealthKit
import Auth0
import WatchConnectivity
import UserNotifications
import SwiftUI
import FirebaseMessaging

class Datamodel: NSObject, ObservableObject, WCSessionDelegate {
  
  static let shared = Datamodel()
  
  //  @Published var path: NavigationPath = .init()
  
  @Published var fcmTokenSubmittedThisSession = false
  @Published var isAuthenticated = false
  @Published var error: APIError?
  @Published var profileStats: ProfileStatsData = ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0)
  @Published var profileIsLoaded = false
  @Published var profile: Profile = Profile.default {
    didSet {
      editingProfile = profile
    }
  }
  
  @Published var config: Config = Config.default
  
  var editingProfile: Profile = Profile.default
  
  @Published var joinedCompetitions: [Competition] = [] {
    didSet {
      //            print("we got some new joined comps!")
      if self.connectivitySession.isReachable {
        let compDictList = joinedCompetitions.filter({ $0.isActive() }).map { ["title": $0.title, "id": $0.id.uuidString] }
        send(message: ["competition_list": compDictList])
      }
    }
  }
  
  @Published var competitions: [Competition] = []
  
  var initialDataLoaded: Bool = false
  
  @Published var currentTab: Tab = .home
  
  var tabSelector: Binding<Tab> {
    Binding {
      self.currentTab
    } set: { newTab in
      if newTab == self.currentTab {
        if newTab == .home {
//          print("re pressed:: \(newTab)")
          PathManager.shared.homePath.removeLast(PathManager.shared.homePath.count)
        } else if newTab == .profile {
//          print("re pressed:: \(newTab)")
          PathManager.shared.profilePath.removeLast(PathManager.shared.profilePath.count)
        } else if newTab == .competitions {
//          print("re pressed:: \(newTab)")
          PathManager.shared.compPath.removeLast(PathManager.shared.compPath.count)
        } else if newTab == .squads {
//          print("re pressed:: \(newTab) \(PathManager.shared.squadPath.count)")
          PathManager.shared.squadPath.removeLast(PathManager.shared.squadPath.count)
        } else if newTab == .workout {
//          print("re pressed:: \(newTab)")
          PathManager.shared.workoutPath.removeLast(PathManager.shared.workoutPath.count)
        }
      }
      self.currentTab = newTab
    }
  }

  @Published var invites: [Invite] = []
  
  @Published var notifications: [Notification] = []
  
  @Published var squads: [Squad] = []
  @Published var joinedSquads: [Squad] = []
  
  @Published var myCompletedWorkouts: [Workout] = []
  
  @Published var tabSelection: Tab = .home
  
  let api: DemoiosAPIService
  let health: HealthService
  
  private var connectivitySession: WCSession
  
  init(apiService: DemoiosAPIService = DemoiosAPIService.shared, healthService: HealthService = HealthService.shared) {
    self.api = apiService
    self.health = healthService
    
    self.connectivitySession = WCSession.default
    
    super.init()
    
    self.connectivitySession.delegate = self
    self.connect()
  }
  
  // MARK: - WCSession related
  
  func connect() {
    guard WCSession.isSupported() else {
      print("WCSession is not supported")
      return
    }
    
    connectivitySession.activate()
    
  }
  
  func send(message: [String: Any]) {
    
    if connectivitySession.isWatchAppInstalled && connectivitySession.isPaired && connectivitySession.isReachable {
      
      print("Session live, sending message directly")
      
      connectivitySession.sendMessage(message, replyHandler: nil) { (error) in
        print(error.localizedDescription)
      }
      
    } else if connectivitySession.isWatchAppInstalled && !connectivitySession.isReachable {
      print("sending in background")
      
      connectivitySession.transferUserInfo(message)
      //    print("info sending \n is transferring: \(u.isTransferring) \n data: \(u.userInfo)")
    } else {
      
      print("Did not send message\n\n")
    }
    
  }
  
  
  func session(_ session: WCSession, didFinish userInfoTransfer: WCSessionUserInfoTransfer, error: Error?) {
    print("done transfering::: err: \(error?.localizedDescription ?? "") and sesh: \(userInfoTransfer.description)")
  }
  
  func session(_ session: WCSession, didReceiveMessage message: [String: Any]) {
    
    if let startWorkoutPayload = message["workout_start"] as? [String: Any], let payload = startWorkoutPayload["payload"] as? [String: String] {
      print("GOT THE MESSAGE START with payload::: \(startWorkoutPayload)" )
      
      // Call API to initiate the workout and send back the externalId to start the workout
      Task {
        do {
          let createWorkoutRes: String = try await api.createWorkout(payload: payload)
          print("Sending external id to watch \(createWorkoutRes) to start workout!!!")
          self.send(message: ["externalId": createWorkoutRes])
        } catch {
          self.send(message: ["initiate_workout_error": "Failed to create workout, try again"])
          print(error)
        }
      }
      
      
      
      
    }
    
    
    if let route_processed = message["route_processed"] as? [String: Any], let success = route_processed["success"] as? Bool, let msg = route_processed["message"] as? String {
      if success {
        triggerLocalNotification(title: msg, body: "", background: .blue)
      } else {
        triggerLocalNotification(title: msg, body: "", background: .red)
      }
      
    }
    
    if message["requesting_competitions"] != nil {
      print("requested to send the comps over", competitions.count)
      
      // Create an array of dictionaries representing competitions and send it back to the watch
      let compDictList = joinedCompetitions.map { ["title": $0.title, "id": $0.id.uuidString] }
      send(message: ["competition_list": compDictList])
    }
    
    if let payload = message["workout_finished"] as? [String: Any] {
      
      Task {
        do {
          // Call API to finish the workout
          try await self.api.finishWorkoutPatch(payload: payload)
        } catch {
          triggerLocalNotification(title: "Failed to save workout", body: "", background: .red)
        }
      }
    }
    
    
  }
  
  func sessionDidBecomeInactive(_ session: WCSession) {
    print("session did become inactive")
  }
  
  func sessionDidDeactivate(_ session: WCSession) {
    print("Session did deactivate ")
  }
  
  func session(_ session: WCSession, activationDidCompleteWith activationState: WCSessionActivationState, error: Error?) {
    //
  }
  
  func clear() {
    
    print("\n\nCLEARING ALL DATA\n\n")
    self.isAuthenticated = false
    
    self.profileIsLoaded = false
    self.profileStats = ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0)
    self.profile = Profile.default
    
    self.competitions = []
    self.joinedCompetitions = []
    
    self.notifications = []
    
    self.invites = []
    
    self.squads = []
    self.joinedSquads = []
    
    self.myCompletedWorkouts = []
    
    self.joinedSquads = []
    
    self.tabSelection = .home
    self.currentTab = .home
    
    self.initialDataLoaded = false
    
    self.fcmTokenSubmittedThisSession = false
    
    self.config = Config.default
    
    PathManager.shared.clear()
  }
}

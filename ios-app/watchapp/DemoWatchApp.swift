//
//  DemoiosWatchApp.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 3/25/23.
//

import SwiftUI
import UIKit
import UserNotifications
import WatchConnectivity

class MyWatchAppDelegate: NSObject, WKApplicationDelegate, UNUserNotificationCenterDelegate, WCSessionDelegate {
  
  
  
  func applicationDidFinishLaunching() {

    
    UNUserNotificationCenter.current().delegate = self
    
    let category = UNNotificationCategory(identifier: "Demoiosalerts", actions: [], intentIdentifiers: [], options: [])
    
    UNUserNotificationCenter.current().setNotificationCategories([category])
  }

  private var connectivitySession: WCSession
//
//  override init() {
//    self.connectivitySession = WCSession.default
//    super.init()
//    connectivitySession.delegate = self
//    self.connect()

  override init() {
    self.connectivitySession = WCSession.default
    super.init()
    connectivitySession.delegate = self
    self.connect()
    
    UNUserNotificationCenter.current().delegate = self
    
    let category = UNNotificationCategory(identifier: "Demoiosalerts", actions: [], intentIdentifiers: [], options: [])
    
    UNUserNotificationCenter.current().setNotificationCategories([category])
  }

  func userNotificationCenter(_ center: UNUserNotificationCenter, didReceive response: UNNotificationResponse, withCompletionHandler completionHandler: @escaping () -> Void) {
    
//    triggerLocalNotification(body: "didReceiv!")
    
    completionHandler()
  }

  
  func userNotificationCenter(_ center: UNUserNotificationCenter, willPresent notification: UNNotification, withCompletionHandler completionHandler: @escaping (UNNotificationPresentationOptions) -> Void) {
    
//    WorkoutController.shared.notiTestStr = "Got a noti while in foreground, \(notification.request.content.body)"
//
////    triggerLocalNotification(body: "From foreground!")
    
    completionHandler([.badge, .sound, .banner])
    
  }
  

}

@main
struct Demoios_Watch_AppApp: App {
  
  @WKApplicationDelegateAdaptor var appDelegate: MyWatchAppDelegate

  @StateObject var toplevelworkoutController = WorkoutController.shared
    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(toplevelworkoutController)
//                .task {
//                    toplevelworkoutController.requestCompetitionList()
//                }
        }
    }
}




func triggerLocalNotification(title: String, body: String, interruption: UNNotificationInterruptionLevel) {
  
  // Create the request
  let content = UNMutableNotificationContent()
  content.title = title
  content.body = body
  content.sound = .default
  content.categoryIdentifier = "Demoiosalerts"
  content.interruptionLevel = interruption
  
  let uuidString = UUID().uuidString
  
  //    let trigger = UNTimeIntervalNotificationTrigger(timeInterval: 5, repeats: false)
  
  let request = UNNotificationRequest(
    identifier: uuidString, content: content, trigger: .none
  )
  
  // Schedule the request with the system.
  let notificationCenter = UNUserNotificationCenter.current()
  
  //    notificationCenter.removeAllPendingNotificationRequests()
  notificationCenter.add(request) { (error) in
    
    
//    WorkoutController.shared.notiTestStr = "ADDED NOTIFICATION"
    
    if error != nil {
      // Handle any errors.
      
//      WorkoutController.shared.errString = "\(error?.localizedDescription ?? "no error")"
      
    }
  }
}





extension MyWatchAppDelegate {
  
  private enum MessageActions {
    case startWorkout(String)
    case initiateWorkoutError(String)
    case competitionList([[String: String]])
    case handlePreConfiguredWorkoutData([String: String])
    case handleStartingWorkoutFromAcceptedInvite([String: String])
  }
  
//  data.send(message: ["start_workout_from_accepted_invite":
//                        [
//                          "activityType": workout.healthkit_workout_activity_type.capitalized,
//                          "dbWorkoutId": workout.id.uuidString,
//                        ]
//                     ])
//
  
  func connect() {
    guard WCSession.isSupported() else {
      print("WCSession is not supported")
      return
    }
    
    connectivitySession.activate()
  }
  
  
  
  func session(_ session: WCSession, activationDidCompleteWith activationState: WCSessionActivationState, error: Error?) {
    
  }
  
  

  
  private func messageActions(for message: [String: Any]) -> MessageActions? {
    if let idFromApi = message["externalId"] as? String {
      return .startWorkout(idFromApi)
    } else if let err = message["initiate_workout_error"] as? String {
      return .initiateWorkoutError(err)
    } else if let comps = message["competition_list"] as? [[String: String]] {
      return .competitionList(comps)
    } else if let data = message["pre_configured_workout"] as? [String: String] {
      return .handlePreConfiguredWorkoutData(data)
    } else if let acceptedInviteWorkoutData = message["start_workout_from_accepted_invite"] as? [String: String] {
      return .handleStartingWorkoutFromAcceptedInvite(acceptedInviteWorkoutData)
    }
    return nil
  }
  
  func session(_ session: WCSession, didReceiveUserInfo userInfo: [String: Any] = [:]) {
    // BACKGROUND HANDLER
    
    if let workoutData = userInfo["start_workout_from_accepted_invite"] as? [String: String] {
      triggerLocalNotification(title: "Workout Ready", body: "Start your workout!", interruption: .timeSensitive)
      WorkoutController.shared.acceptedInviteWorkoutData = workoutData
    }

    if let workoutData = userInfo["pre_configured_workout"] as? [String: String] {
      WorkoutController.shared.preConfiguredWorkoutData = workoutData
    }
    }
    
  
  func session(_ session: WCSession, didReceiveMessage message: [String: Any]) {
    
    // FOREGROUND HANDLER
    guard let action = messageActions(for: message) else {
      print("Unknown message received")
      return
    }
    
//    triggerLocalNotification(body: "got message in foreground!! ")
    
    switch action {
      
    case .handleStartingWorkoutFromAcceptedInvite(let acceptedInviteWorkoutData):
      
      WorkoutController.shared.acceptedInviteWorkoutData = acceptedInviteWorkoutData
      
    case .handlePreConfiguredWorkoutData(let data):
      //      DispatchQueue.main.async {
      WorkoutController.shared.preConfiguredWorkoutData = data
      //      }
      
    case .startWorkout(let idFromApi):
      WorkoutController.shared.startConfirmedWorkout(externalId: idFromApi)
      
    case .initiateWorkoutError(let err):
      print("WE GOT AN ERROR, CANNOT START WORKOUT \(err)")
      
    case .competitionList(let comps):
      DispatchQueue.main.async {
        WorkoutController.shared.competitions = comps
      }
      
    }
  }
  
}

//
//  WorkoutController+Notifications.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 8/29/23.
//

import Foundation
import UserNotifications


extension WorkoutController {
  func requestNotificationPermission(completion: @escaping (Bool) -> Void) {
    let authOptions: UNAuthorizationOptions = [.alert, .badge, .sound]
    UNUserNotificationCenter.current().requestAuthorization(
      options: authOptions,
      completionHandler: { (granted, err) in
        guard err == nil else {
          self.errStrings.append(err?.localizedDescription ?? "error in requestNotificationPermission")
          completion(false)
          return
        }
        if granted {
          self.notificationsAuthorized = "Got it!"
        } else {
          self.notificationsAuthorized = "Granted was false"
        }
        completion(granted)
      }
    )
  }
}

//
//  Utils+Notifications.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/5/23.
//

import Foundation
import UserNotifications
import UIKit

func registerForPushNotifications(completion: @escaping (Bool) -> Void) {
  
  let authOptions: UNAuthorizationOptions = [.alert, .badge, .sound]
  UNUserNotificationCenter.current().requestAuthorization(
    options: authOptions,
    completionHandler: { (granted, error) in
      print("Permission granted: \(granted)")
      // 1. Check if permission granted
      guard granted else {
        print("\n\nUnable to grant notification permission\(error?.localizedDescription ?? "ERROR msg failed")")
        completion(false)
        return
      }
      
      // 2. Attempt registration for remote notifications on the main thread
      DispatchQueue.main.async {
        UIApplication.shared.registerForRemoteNotifications()
      }
      completion(true)
    }
  )
  
  
}

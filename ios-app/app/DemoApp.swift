//
//  DemoiosApp.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/7/23.
//

import SwiftUI
import Auth0
import Sentry
import SentrySwiftUI
import FirebaseCore
import UserNotifications
import FirebaseMessaging
import NewRelic

class AppDelegate: NSObject, UIApplicationDelegate, UNUserNotificationCenterDelegate, MessagingDelegate {
  
  func messaging(_ messaging: Messaging, didReceiveRegistrationToken fcmToken: String?) {
    if let tok = fcmToken {
      DispatchQueue.main.asyncAfter(deadline: .now() + 1) {
        if !Datamodel.shared.fcmTokenSubmittedThisSession && Datamodel.shared.profileIsLoaded {
          
//          print("Token recieved in app delegate messaging, sending to backend ")
          Task {
            await DemoiosAPIService.shared.updateFcmToken(token: tok)
          }
        }
      }
    }
  }
  
  func application(_ application: UIApplication, didRegisterForRemoteNotificationsWithDeviceToken deviceToken: Data) {
    // 1. Convert device token to string
    let tokenParts = deviceToken.map { data -> String in
      return String(format: "%02.2hhx", data)
    }
    let token = tokenParts.joined()
    // 2. Print device token to use for PNs payloads
    print("Device Token: \(token)")
    
    Messaging.messaging().token { token, error in
      if let error = error {
        print("\n\nError fetching FCM registration token: \(error)")
      } else if let token = token {
        print("\n\nFCM registration token: \(token)")
        //            self.fcmRegTokenMessage.text  = "Remote FCM registration token: \(token)"
      }
    }
  }
  
  func application(_ application: UIApplication, didFailToRegisterForRemoteNotificationsWithError error: Error) {
    // 1. Print out error if PNs registration not successful
    print("Failed to register for remote notifications with error: \(error)")
  }
  
  func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]? = nil) -> Bool {
    
    if getCurrentEnvType() == .release {
      NewRelic.start(withApplicationToken: "")
      
      NewRelic.enableFeatures(.NRFeatureFlag_DistributedTracing)
    }
    //////////////////////////////////
    /// RESPONDS TO NOTIFICATIONS WHILE APP IS NOT RUNNING
    ///
    /// The launchOptions dictionary contains the key UIApplication.LaunchOptionsKey.remoteNotification if the app was launched from a remote notification.
    /*
     {
     "aps" : {
     "alert" : {
     "body" : "great match!",
     "title" : "Portugal vs. Denmark",
     },
     "badge" : 1,
     },
     "customKey" : "customValue"
     }
     */
    ///
    ///
    ////////////////////////////////
    FirebaseApp.configure()
    UNUserNotificationCenter.current().delegate = self
    Messaging.messaging().delegate = self
    
    if let resToNoti = launchOptions?[.remoteNotification] as? [String: Any], let aps = resToNoti["aps"] as? [String: Any] {
      // the app was launched from pressing on a notification while app is in the background
      //      let aps = remoteNotif!["aps"] as? [String:AnyObject]
      print("app was NOT OPENED -> Responding to notification:: \(resToNoti), APS INFO ::: \(aps)")
    }
    return true
  }
  
  func userNotificationCenter(_ center: UNUserNotificationCenter, willPresent notification: UNNotification, withCompletionHandler completionHandler: @escaping (UNNotificationPresentationOptions) -> Void) {
    /// FOREGROUND / BACKGROUND HANDLING -> presents the actual notification banner in app
    
    let userinfo = notification.request.content.userInfo
    let bod = notification.request.content.body
    let title = notification.request.content.title
    
    if notification.request.content.categoryIdentifier != "Demoiosalerts", let receiverId = notification.request.content.userInfo[AnyHashable("receiver_id")] as? String, receiverId.lowercased() == Datamodel.shared.profile.id.uuidString.lowercased() {
      dispatchAsync {
        NotificationHandler.shared.new(title: title, body: bod, background: .green, icon: .info)
      }
      
      Task {
        await DemoiosAPIService.shared.getMyNotifications()
      }
      
      completionHandler([.badge, .sound])
    }
    
  }
  
  func userNotificationCenter(_ center: UNUserNotificationCenter, didReceive response: UNNotificationResponse, withCompletionHandler completionHandler: @escaping () -> Void) {
    /// HANDLER FOR WHEN APP IS IN FOREGROUND OR BACKGROUND -> when we got the noti, just info
    
    let userinfo = response.notification.request.content.userInfo
    let bod = response.notification.request.content.body
    let title = response.notification.request.content.title
    let ci = response.notification.request.content.categoryIdentifier
    
    print("\n\n\nRESPONDING TO PushNoti press \n:::TITLE:: \(title) \nBODY:: \(bod) \nUSER INFO \(userinfo) \ncategory identifier:: \(ci)\n\n\n")
  }
  
}



@main
struct DemoiosApp: App {
  @UIApplicationDelegateAdaptor(AppDelegate.self) var delegate
  
  @StateObject var dataModel = Datamodel.shared
  @StateObject var notificationHandler = NotificationHandler.shared
  @StateObject var weather = WeatherManager.shared
  @StateObject var location = LocationManager.shared
  
  var body: some Scene {
    WindowGroup {
      ContentView()
        .environmentObject(dataModel)
        .environmentObject(notificationHandler)
        .environmentObject(weather)
        .environmentObject(location)
    }
  }
  
  init () {
    
    let curEnv = getCurrentEnvType()
    
    if curEnv == .preview {
      print("Running in Preview")
    }
    
    if curEnv == .release {
      SentrySDK.start { options in
        
        options.dsn = ""
        
        options.environment = Configuration.environment.absoluteString
        
        options.sessionTrackingIntervalMillis = 5_000
        
        options.enableFileIOTracing = true
        
        options.profilesSampleRate = 0.5
        
        options.enableUserInteractionTracing = true
        
        options.enableTracing = true
        options.tracesSampleRate = 0.5
        
        // options.debug = Configuration.environment.absoluteString == "DEBUG"
        
        options.enableAutoPerformanceTracing = true
        
        options.enablePreWarmedAppStartTracing = true
        options.attachScreenshot = true
        options.attachViewHierarchy = true
        options.enableMetricKit = true
      }
    } // END SENTRY INIT
    
  }
}



extension AppDelegate {
  
  // MARK: - UNIVERSAL LINK HANDLER
  
  func application(_ application: UIApplication,
                   continue userActivity: NSUserActivity,
                   restorationHandler: @escaping ([UIUserActivityRestoring]?) -> Void) -> Bool {
    // Get URL components from the incoming user activity.
    guard userActivity.activityType == NSUserActivityTypeBrowsingWeb,
          let incomingURL = userActivity.webpageURL,
          let components = NSURLComponents(url: incomingURL, resolvingAgainstBaseURL: true) else {
      return false
    }
    
    
    // Check for specific URL components that you need.
    guard let path = components.path,
          let params = components.queryItems else {
      return false
    }
    print("path = \(path)")
    
    
    if let albumName = params.first(where: { $0.name == "albumname" })?.value,
       let photoIndex = params.first(where: { $0.name == "index" })?.value {
      
      
      print("album = \(albumName)")
      print("photoIndex = \(photoIndex)")
      return true
      
      
    } else {
      print("Either album name or photo index missing")
      return false
    }
  }
}

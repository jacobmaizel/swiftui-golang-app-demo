//
//  NotificationSoftRequest.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/24/23.
//

import SwiftUI
import FirebaseMessaging

struct NotificationSoftRequest: View {
  
  @State private var userRespondedToNotiRequest = false
  @EnvironmentObject var notihandler: NotificationHandler
  @EnvironmentObject var data: Datamodel
  
    var body: some View {
      
      NavigationStack {
        VStack {
          
          Spacer()
          HStack {
            
            Spacer()
            Image(systemName: "bell")
              .resizable()
              .frame(width: 100, height: 100)
            
            Image(systemName: "plus")
              .resizable()
              .frame(width: 50, height: 50)
              .padding(.horizontal)
            
            Logo(size: 125, logoColor: .white)
            
            Spacer()
            
          }
          
          Spacer()
          
          Text("Enable Notifications to stay up to date with invites, competition updates, and squad activity!")
            .DemoiosText(fontSize: .large, weight: .bold)
            .multilineTextAlignment(.center)
            .padding(.horizontal)
          
          Spacer()
          Spacer()
          
          
          VStack(spacing: 0) {
            
            Button {
              registerForPushNotifications { res in
                  dispatchAsync {
                    userRespondedToNotiRequest = true
                  }
                
                if res {
                  print("got notification permission, sending FCM to backend")
                  DispatchQueue.main.asyncAfter(deadline: .now() + 0.5) {
                    
                    Messaging.messaging().token { token, error in
                      if let error = error {
                        print("\n\nCould not get fcm token in noti soft request::: \(error)")
                      } else if let token = token {
                        print("\n\nGOT token in fcm soft request, updating backend: \(token)")
                        //            self.fcmRegTokenMessage.text  = "Remote FCM registration token: \(token)"
                        if !Datamodel.shared.fcmTokenSubmittedThisSession && Datamodel.shared.profileIsLoaded {
                        DispatchQueue.main.asyncAfter(deadline: .now() + 1) {
                          Task {
                            await DemoiosAPIService.shared.updateFcmToken(token: token)
                          }
                        }
                        }
                      }
                    }
                  }
                }
                
                if !res {
                  // user answered but answered with DO NOT ALLOW, let them know
                  // they can always go to settings to enable notifications
                  notihandler.new(title: "Maybe next time!", body: "You can always head to settings and enable Notifications if you change your mind", background: .blue, icon: .ok)
                }
                Task {
                  do {
                    try await data.api.completeOnboarding(data)
                  } catch {
                    triggerLocalNotification(title: "Failed to complete onboarding", body: "Try again", background: .red)
                  }
                }
              }
              
            } label: {
              Text("Enable Notifications")
              
            }
            .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
            .padding(.horizontal)
//            .simultaneousGesture(TapGesture().onEnded {
//
//              registerForPushNotifications()
//
//            })
//
            NavigationLink {
              // send profile update to onboarding_completed!
              Text("Done with onboarding!")
            } label: {
              Text("No Thanks")
            }
            .buttonStyle(CustomizableResponsiveButton(fontSize: .base, fontWeight: .light, bgColor: .bg, fgColor: .graySeven))
            .padding(.top)
            .simultaneousGesture(TapGesture().onEnded {
              print("denied, update prefs")
      //          DispatchQueue.main.async {
      //              self.workoutController.selectedCompetition = comp_choice
      //          }
            })
          }
          
          
        }
        .DemoiosBackgroundColor()
      }
    }
}

struct NotificationSoftRequest_Previews: PreviewProvider {
    static var previews: some View {
        NotificationSoftRequest()
        .setupPreview()
    }
}

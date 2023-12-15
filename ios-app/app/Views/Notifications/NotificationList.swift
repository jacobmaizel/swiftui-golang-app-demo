//
//  NotificationList.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/24/23.
//

import SwiftUI
import FirebaseMessaging

struct NotificationList: View {
  @EnvironmentObject var data: Datamodel
  
  @State var hasNotiPerms: Bool?
  @State var selectedNoti: Notification?
  @State var showNotiDetailSheet: Bool = false
  
  var body: some View {
    
    Group {
      
      if hasNotiPerms == true {
        ScrollView {
          ForEach($data.notifications) { $noti in
            NotificationItem(notification: $noti)
            //              .overlay(
            //                Rectangle()
            //                  .fill(noti.opened != nil ? Color.graySix.opacity(0.8) : Color.clear)
            //                  .cornerRadius(8)
            //              )
              .padding(.horizontal)
              .onTapGesture {
                if noti.opened == nil {
                  Task {
                   await data.api.openNotification(data: data, id: noti.id)
                  }
                }
                
                dispatchAsync {
                  selectedNoti = noti
                }
                showNotiDetailSheet.toggle()
              }
            
          }
        }
        .navigationTitle("Notifications")
        .padding(.vertical)
        .task {
          await data.api.getMyNotifications()
        }
        .refreshable {
          await data.api.getMyNotifications()
        }
        .sheet(isPresented: $showNotiDetailSheet) {
          NotificationDetailSheet(noti: selectedNoti)
            .presentationDetents([.fraction(0.4), .fraction(1.0)])
          //            .presentationDetents([.fraction(0.4)])
        }
        .onChange(of: showNotiDetailSheet) { newVal in
          if !newVal {
            selectedNoti = nil
          }
        }
        
      } else if hasNotiPerms == false {
        Text("Notification Permission Denied, please head to your settings app to enable notifications for Demoios :)")
      } else {
        VStack {
          ProgressView()
        }
        .DemoiosBackgroundColor()
      }
    } // group
    .DemoiosBackgroundColor()
    .onAppear {
      registerForPushNotifications { got_perms in
        dispatchAsync {
          hasNotiPerms = got_perms
        }
        if got_perms {
          DispatchQueue.main.asyncAfter(deadline: .now() + 0.2) {
            Messaging.messaging().token { token, error in
              if let error = error {
                print("\n\nCould not get fcm token in noti soft request::: \(error)")
              } else if let token = token {
                //            self.fcmRegTokenMessage.text  = "Remote FCM registration token: \(token)"
                if !Datamodel.shared.fcmTokenSubmittedThisSession && Datamodel.shared.profileIsLoaded {
                  
                  print("\n\nGOT token in fcm soft request, AND the token had not been loaded this session yet: \(token)")
                  DispatchQueue.main.asyncAfter(deadline: .now() + 3) {
                    Task {
                      await DemoiosAPIService.shared.updateFcmToken(token: token)
                    }
                  }
                }
              }
            }
          }
          
        }
      }
      Task {
       await data.api.getMyNotifications()
      }
    } // onappear
  }
}

struct NotificationList_Previews: PreviewProvider {
  static var previews: some View {
    
    var data = Datamodel()
    
    NotificationList(hasNotiPerms: true)
      .environmentObject(data)
      .onAppear {
        data.notifications = [.test_opened, .test_opened, .test_opened, .test_unopened_comp]
      }
      .setupPreview()
  }
}


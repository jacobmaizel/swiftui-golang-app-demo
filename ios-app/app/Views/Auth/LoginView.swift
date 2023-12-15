//
//  Login.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/7/23.
//

import SwiftUI
import Auth0
import FirebaseMessaging


struct LoginView: View {
  @State private var username: String = ""
  @State private var password: String = ""
  @State private var showPassword: Bool = false
  
  @EnvironmentObject var dataModel: Datamodel
  
  var body: some View {
    
    VStack(alignment: .center) {
      
      if Configuration.environment.absoluteString == "DEBUG" {
        Text("DEBUG MODE ENABLED")
          .DemoiosText(fontSize: .xxl, weight: .bold)
      }
      Spacer()
      
      Image("white-logo")
        .resizable()
        .frame(width: 200, height: 200)
      
      Spacer()
      
      if Configuration.environment.absoluteString == "DEBUG" {
        Button {
          dataModel.api.auth0Manager.credentials { result in
            switch result {
            case .success(let credentials):
              print("Obtained credentials: \(credentials.accessToken)")
              
              print("Obtained refresh token: \(credentials.refreshToken ?? "")")
            case .failure(let error):
              print("Failed with: \(error)")
            }
          }
        } label: {
          Text("Get logged in info from cred manager")
        }
        
        Button {
          dataModel.api.webAuthClient
            .clearSession { result in
              switch result {
              case .success:
                
                let cleared = dataModel.api.auth0Manager.clear()
                
                if cleared {
                  print("Auth session cleared")
                } else {
                  print("auth failed to clear keychain")
                }
                
                dataModel.clear()
              case .failure(let error):
                print("Failed with: \(error)")
              }
            }
        } label: {
          Text("logout")
        }
      }
      
      Button {
        
        dataModel.api.webAuthClient
          .start { result in
            switch result {
            case .success(let credentials):
              //                            print("Obtained credentials: \(credentials)")
              //                            print("Here ya go access:::\(credentials.accessToken):::")
              //                            print("refresh token::: \(credentials.refreshToken ?? "woopsei")")
              
              _ = dataModel.api.auth0Manager.store(credentials: credentials)
              
              //                            if stored {
              //                                print("successfully stored creds")
              //                            }
              
              dataModel.isAuthenticated = true
              
              DispatchQueue.main.asyncAfter(deadline: .now() + 5) {
                UNUserNotificationCenter.current().getNotificationSettings { settings in
                  
                  if settings.authorizationStatus == .authorized {
                    print("got notification permission, sending FCM to backend")
                    
                    Messaging.messaging().token { token, error in
                      if let error = error {
                        print("\n\nLOGIN VIEW TOKEN:::: \(error)")
                      } else if let token = token {
                        //            self.fcmRegTokenMessage.text  = "Remote FCM registration token: \(token)"
                        if !Datamodel.shared.fcmTokenSubmittedThisSession && Datamodel.shared.profileIsLoaded {
                          
                          print("LOGIN VIEW TOKEN SUCCESS, posting to backend")
                          Task {
                            await DemoiosAPIService.shared.updateFcmToken(token: token)
                          }
                        }
                      }
                    }
                  }
                }
                
              }
              
            case .failure(let error):
              print("Auth Failed with: \(error)")
            }
          }
      } label: {
        Text("Sign in")
      }
      .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
      .padding()
      
      //            .padding()
      
    } // Top Vstack
    .background(Color.bg)
    .DemoiosText(fontSize: .xxl)
  }
}

struct Login_Previews: PreviewProvider {
  static var previews: some View {
    LoginView()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}

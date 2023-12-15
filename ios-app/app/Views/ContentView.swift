//
//  ContentView.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/7/23.
//

import SwiftUI
import Auth0

struct ContentView: View {
    
  @EnvironmentObject var data: Datamodel
  @EnvironmentObject var notiHandler: NotificationHandler
  @State var showing: Bool = false
  @State var url: URL?
  @State var sharedResourceType: ShareableResource?
  @State var sharedResourceId: UUID?
//  @State var path = NavigationPath()
  
  
  @ViewBuilder
  var body: some View {
      Group {
        
        if !data.api.auth0Manager.canRenew() && !data.isAuthenticated {
          LoginView()
        } else if !data.profileIsLoaded {
          VStack {
            ProgressView()
            Text("Loading Profile")
              .DemoiosText(fontSize: .base, weight: .bold)
              .task {
                do {
                  try await data.api.getUserProfileAndConfig(data)
                } catch {
                  print(error)
                }
              }
            Button {
              data.api.webAuthClient
                .clearSession { result in
                  switch result {
                  case .success:
                    
                    let cleared = data.api.auth0Manager.clear()
                    
                    if cleared {
                      print("Auth session cleared")
                    } else {
                      print("auth failed to clear keychain")
                    }
                    
                    data.clear()
                    
                  case .failure(let error):
                    print("Failed with: \(error)")
                  }
                }
            } label: {
              Text("Force logout")
            }
          }
        } else if !data.profile.onboarding_completed {
          OnboardingView()
        } else if data.profile.onboarding_completed {
          MainTabViewWrapper()

            .sheet(isPresented: $showing) {
             if let sharedResourceType, let sharedResourceId = sharedResourceId {
               NavigationStack {
                 ShareableDetail(resourceType: sharedResourceType, resourceId: sharedResourceId)
                   .onDisappear {
                     self.url = nil
                     self.sharedResourceId = nil
                     self.sharedResourceType = nil
                   }
               }
               
             }
            }
            .modifier(BannerModifier(notiToDisplay: $notiHandler.currentNotification))
            .onOpenURL { url in
              let resourceInfo = processOpenedShareableUrl(with: url)
              guard let resourceInfo = resourceInfo else {
                print("ignoring url, invalid resource info \(url)")
                return
              }
              dispatchAsync {
                self.url = url
                self.sharedResourceType = resourceInfo.0
                self.sharedResourceId = resourceInfo.1
              }
              showing.toggle()
            }
           

        }
      }
      .task {
        Datamodel.shared.health.requestHealthKitAuthorization { success in
          print("Got auth for healthkit? \(success)")
          LocationManager.shared.manager.requestWhenInUseAuthorization()
        }
        registerForPushNotifications { _ in
        }
        
        if LocationManager.shared.hasPerms(), let loc = LocationManager.shared.location {
          LocationManager.shared.getLocation()
          await WeatherManager.shared.getWeather(for: loc)
        }
      }
//      .onChange(of: path) { pi in
//        print("Path changed in content view: \(pi)")
//      }
//    }
//      .task {
//        data.health.requestHealthKitAuthorization { success in
//          print("Got auth for healthkit? \(success)")
//        }
//      }
    
  }
}

struct ContentView_Previews: PreviewProvider {
  static var previews: some View {
    ContentView()
      .environmentObject(Datamodel())
  }
}

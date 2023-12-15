//
//  HomeView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 6/26/23.
//

import SwiftUI
import UserNotifications
struct HomeTabView: View {
  
  @EnvironmentObject var wc: WorkoutController
  
  var body: some View {
    NavigationStack {
      ScrollView {
        
        VStack {
//
//          Text("Noti debug:::\(wc.notiTestStr)")
//          Text("err debug:::\(wc.errString)")
//
          
          if wc.acceptedInviteWorkoutData != nil {
            NavigationLink {
              WorkoutLocationChoiceList()
            } label: {
              Text("Start \(wc.acceptedWorkoutLeaderName ?? "")'s Workout")
            }
          }
          
          if wc.preConfiguredWorkoutData != nil {
            NavigationLink("Confirm Workout") {
              StartWorkoutConfirmation()
            }
            .buttonStyle(.borderedProminent)
          }
          
          NavigationLink("Configure") {
            StartWorkoutView()
          }
          
          if getCurrentEnvType() == .debug {
            ForEach(wc.errStrings, id: \.self) { str in
              Text(str)
            }
            
            Text("Location auth: \(wc.locationAuthorized ?? " no location auth")")
            Text("Noti auth: \(wc.notificationsAuthorized ?? "no noti auth")")
            Text("Healthkit auth: \(wc.healthStoreAuthorized ?? "no healthkit auth")")
            
          }
        }
    
      }
      .onAppear {
        wc.requestAuthorizationHealthkit { granted in
          
          if granted {
            wc.healthStoreAuthorized = "Got auth"
          } else {
            
            wc.healthStoreAuthorized = "FAILED to get auth"
          }
        }
        wc.requestNotificationPermission { granted in
                  }
        
        
        
      
      }

    }
    .navigationBarBackButtonHidden()
  }
}

struct HomeView_Previews: PreviewProvider {
  static var previews: some View {
    HomeTabView()
      .environmentObject(WorkoutController())
  }
}

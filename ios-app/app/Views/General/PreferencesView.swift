//
//  PreferencesView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/24/23.
//

import SwiftUI

struct PreferencesView: View {
  
  @EnvironmentObject var data: Datamodel
  
  @State var loaded: Bool = false
  
  @State var auto_sync_workouts: Bool = true
  
  
  
  var body: some View {
    
    Group {
      
   
        Form {
          if loaded {
          Section {
            Toggle(isOn: Binding(get: {
              auto_sync_workouts
            }, set: { val in
              Task {
                let payload = [
                  "auto_sync_workouts": val
                ]
                do {
                  let config = try await data.api.updateConfig(payload: payload)
                  data.config = config
                } catch {
                  triggerLocalNotification(title: "Failed to update", body: "", background: .red, icon: .warning)
                }
              }
              auto_sync_workouts = val
            })
            ) {
              Text("Auto Sync Workouts")
            }
          } footer: {
            Text("Demoios can automatically sync your workout data when you open the app.")
          }
          
        }
       
        
      }
        .scrollContentBackground(.hidden)
        .DemoiosBackgroundColor()
      
    }
    .DemoiosBackgroundColor()
    .task {
      
      auto_sync_workouts = data.config.auto_sync_workouts
      loaded = true
    }
    
  }
}

struct PreferencesView_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    PreferencesView()
      .setupPreview()
      .environmentObject(data)
      .task {
        
        data.config = Config.default
      }
  }
}

/*
 Section(header: Text("Workout Types")) {
 ForEach(CompetitionAcceptedActivityTypes.allCases, id: \.self) { option in
 Toggle(isOn: Binding(
 get: { self.allowedActivityType.contains(option) },
 set: { newValue in
 if newValue {
 self.allowedActivityType.insert(option)
 } else {
 self.allowedActivityType.remove(option)
 }
 }
 )) {
 Text(option.capped)
 }
 }
 }
 
 */




enum SquadNotificationTypes: String, ListableEnum {
  case squad_invites, squad_competition_state
}

enum WorkoutNotificationTypes: String, ListableEnum {
  case workout_invites, post_workout_summary
}

enum CompetitionNotificationTypes: String, ListableEnum {
  case competiton_state, competition_invites
}

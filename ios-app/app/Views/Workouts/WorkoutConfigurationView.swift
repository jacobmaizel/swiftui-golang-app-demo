//
//  WorkoutConfigurationView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/20/23.
//

import SwiftUI

struct WorkoutConfigurationView: View {
  @State private var location: WorkoutConfigLocationOptions = .outdoor
  @State private var activityType: WorkoutConfigActivityTypeOptions = .running
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) var dismiss
  
  
  // TODO: need to have a way for the user to manage current workout aka re invite people if they fail to join the first time
  @State var usersToInviteArray: [PartialProfile] = []
  @State private var showingInviteSheet: Bool = false
  @State private var showModal: Bool = false
  @State var inviteSearchString: String = ""
  
  var body: some View {
    VStack {
      Form {
        Picker("Location", selection: $location) {
          ForEach(WorkoutConfigLocationOptions.allCases) { choice in
            Text(choice.capped)
          }
        }
        .pickerStyle(.inline)
        
        Picker("Workout Type", selection: $activityType) {
          ForEach(WorkoutConfigActivityTypeOptions.allCases) { choice in
            Text(choice.capped)
          }
        }
        .pickerStyle(.inline)
        Section("Want to invite a friend to this workout?") {
          Button {
            showingInviteSheet.toggle()
          } label: {
            HStack {
              Image(systemName: "magnifyingglass")
                .padding(.trailing)
              Text("\(usersToInviteArray.count > 0 ? "Search for others" : "Search")")
            }
          }
        }
        
        
        if !usersToInviteArray.isEmpty {
          Section("Members to Invite") {
            List {
              ForEach(usersToInviteArray) { array in
                HStack {
                  CircleImage(url: array.picture, size: 30)
                  Text(array.full_name)
                }
              }
              
            }
          }
        }
        
        Section {
          
          Button("Quick Start (Watch)") {
            print("Quick start pressed")
          }
          
          Button("Configure workout\(!usersToInviteArray.isEmpty ? " + Send Invites" : "")") {
            
            Task {
              
              
              do {
                let newWorkoutId: String = try await data.api.createWorkout(payload: ["healthkit_workout_activity_type": activityType.rawValue])
                
                let profilesToInvite = usersToInviteArray.map({ $0.id })

                guard let workoutUUID = UUID(uuidString: newWorkoutId) else {
                  print("Failed to decode uuid from res in invite calback")
               
                    triggerLocalNotification(title: "Failed to start workout and send invites", body: "", background: .red, icon: .warning)
              
                  return
                }
                
                print("CREATED WORKOUT ::: \(workoutUUID.uuidString)")
                
                if !profilesToInvite.isEmpty {
                  do {
                    let res: String = try await data.api.inviteProfilesToWorkoutByIds(workoutId: workoutUUID, idsToInvite: profilesToInvite)
                    triggerLocalNotification(title: res, body: "", background: .green, icon: .ok)
                  } catch {
                    triggerLocalNotification(title: "Failed to send invites", body: "\(getCurrentEnvType() == .debug ? error.localizedDescription : "")", background: .red, icon: .warning)
                  }
                }
                
                data.send(message: ["start_workout_from_accepted_invite":
                                      [
                                        "activityType": activityType.rawValue.capitalized,
                                        "dbWorkoutId": workoutUUID.uuidString,
                                        "leaderPicture": data.profile.picture,
                                        "leaderName": data.profile.first_name
                                      ]
                                   ])

                dismiss()

            } catch {
              
            }
              
            } // bottom task

          
          
        }
        } footer: {
          Text("Sending invites creates the workout and alerts all users currently in the Invite List")
        }
        
      }
      .formStyle(.grouped)
    } // vstack here !!!!!!!!!!!!!!!
    .navigationTitle("Configure a Workout")
    .scrollContentBackground(.hidden)
    .DemoiosBackgroundColor()
    .onChange(of: showingInviteSheet) { newValue in
      if !newValue {
        inviteSearchString = ""
      }
    }
    .sheet(isPresented: $showingInviteSheet) {
      NavigationStack {
        //        WorkoutInviteSearch(searchText: $inviteSearchString, selectedResults: $usersToInviteArray)
        MemberInviteSearch(resourceType: .profiles, workout: .defaultWorkout, searchText: $inviteSearchString, selectedResults: $usersToInviteArray)
      }
    }
    
    
  }
}

struct WorkoutConfigurationView_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutConfigurationView()
      .setupPreview()
      .environmentObject(Datamodel())
  }
}



enum WorkoutConfigLocationOptions: String, ListableEnum {
  case indoor, outdoor
}

enum WorkoutConfigActivityTypeOptions: String, ListableEnum {
  case running, walking, swimming, cycling
}

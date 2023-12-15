//
//  NotificationWorkoutDetail.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/3/23.
//

import SwiftUI


struct NotificationWorkoutDetail: View {
  
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) var dismiss
  
  var noti: Notification?
  var workoutId: UUID
  @State var loadedWorkout: Workout?
  @State var err: APIError?
  @State var loaded: Bool = false
  
  @State var relatedInvite: Invite?
  @State var inviteLoaded: Bool = false
  @State var inviteLoadingErr: APIError?
  
  var body: some View {
    
    VStack {
      if noti == nil {
        Text("No Notification Found.")
      } else if !loaded {
        ProgressView()
      } else if err != nil {
        Text("Failed to load workouts.")
      } else if let loadedWorkout = loadedWorkout {
        
        VStack {
          
          HStack {
            
            VStack(alignment: .leading) {
              
              Text(loadedWorkout.healthkit_workout_activity_type.capitalized)
                .DemoiosText(fontSize: .xl, weight: .bold)
              
              Text("Created \(loadedWorkout.created_at.stringToFullDateTime)")
                .DemoiosText(fontSize: .small, color: .grayFour)
            }
            
            Spacer()
            
            LeadByLabel(workout: loadedWorkout)
            
          }
          
          Divider()
          Spacer()
          
          if let inv = relatedInvite, inv.status == "pending" {
            HStack {
              Button("Accept Invite") {
                
                if let wId = noti?.data.workout_id, let iID = noti?.data.invite_id {
                  
                  Task {
                    do {
                      
                      let _: String = try await data.api.acceptWorkoutInvite(workout_id: wId, invite_id_to_accept: iID)
                      
                      
                      
                      
                      await data.api.getMyCompletedWorkouts()
                      
                      triggerLocalNotification(title: "Joined the workout!", body: "Go to the workout tab to start!", background: .green, icon: .info)
                    } catch {
                      
                      NotificationHandler.shared.new(title: "Failed to accept invite", body: "Please try again, or ask the leader to re-invite you.", background: .primarySeven, icon: .warning)
                    }
                  }
                  
                  dismiss()
                  
                }
                
              }
              .buttonStyle(CustomizableResponsiveButton(fontSize: .large, fontWeight: .bold, bgColor: .primarySeven, fgColor: Color.grayOne))
              
              Button("Decline Invite") {
                
                if let wId = noti?.data.workout_id, let iID = noti?.data.invite_id {
                  
                  
                  Task {
                    do {
                      
                      let _: String = try await data.api.declineWorkoutInvite(workout_id: wId, invite_id_to_decline: iID)
                      
                      await data.api.getMyCompletedWorkouts()
                      
                      triggerLocalNotification(title: "Declined the invite!", body: "", background: .green, icon: .info)
                    } catch {
                      
                      NotificationHandler.shared.new(title: "Failed to decline invite", body: "", background: .primarySeven, icon: .warning)
                    }
                  }
                  
                  dismiss()
                  
                }
                
                
              }
              .buttonStyle(CustomizableResponsiveButton(fontSize: .base, fontWeight: .regular, bgColor: .clear, fgColor: Color.grayThree))
              
            }
          } else if let inv = relatedInvite {
            if inv.status == "accepted" {
              Image(systemName: "checkmark")
                .foregroundColor(.green)
            } else if inv.status == "declined" {
              Image(systemName: "xmark")
                .foregroundColor(.red)
            }
          } else if !inviteLoaded {
            ProgressView()
          } else {
            EmptyView()
          }
          
          
        }
        
        
      }
      
    }
    .DemoiosText(fontSize: .base)
    .task {
      
      if !inviteLoaded {
        if let inviteId = noti?.data.invite_id {
          do {
            self.relatedInvite = try await data.api.getInviteById(inviteId: inviteId)
          } catch {
            inviteLoaded = true
            inviteLoadingErr = APIError.invalidResponse
          }

        }
      }
      
      if !loaded {
        Task {
          do {
            let workout: Workout = try await data.api.getWorkoutById(id: workoutId)
            self.loadedWorkout = workout
            self.loaded = true
          } catch {
            self.loaded = true
            self.err = .invalidResponse
          }
        }
      }
      
    } // here
    
  }
}

struct NotificationWorkoutDetail_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    NotificationWorkoutDetail(noti: .test_workout_unopened, workoutId: UUID(), loadedWorkout: .defaultWorkout, err: nil, loaded: true, relatedInvite: .test, inviteLoaded: true, inviteLoadingErr: nil)
      .environmentObject(data)
      .setupPreview()
      .frame(height: 300)
    
    
  }
}

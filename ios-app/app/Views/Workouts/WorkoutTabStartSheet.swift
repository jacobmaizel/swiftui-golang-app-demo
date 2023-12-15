//
//  AcceptedWorkoutStartSheet.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/19/23.
//

import SwiftUI

enum WorkoutSheetType {
  case accepted
  case pending
}
struct WorkoutTabStartSheet: View {
  
  @Binding var workout: Workout?
  @Environment(\.dismiss) var dismiss
  @EnvironmentObject var data: Datamodel
  
  @State var inviteStatuses: [PartialProfile]?
  
  var sheetType: WorkoutSheetType
  
  @Binding var acceptedWorkouts: [Workout]
  @Binding var pendingWorkouts: [Workout]
  
  
  var body: some View {
    if let workout = workout {
      VStack {
        VStack {
          HStack {
            LeadByLabel(workout: workout)
            
            Spacer()
            
            if let created = workout.created_at.intoDate?.timeAgo() {
              VStack {
                Text(created)
                  .DemoiosText(fontSize: .base)
                Text("Created")
                  .DemoiosText(fontSize: .small, color: .graySix)
              }
              
            }
          }
          
          if let invStats = inviteStatuses, !invStats.isEmpty {
            Divider()
            
            Text("Invite Statuses")
              .DemoiosText(fontSize: .small, color: .graySix)
            
            ForEach(invStats) { user in
              HStack {
                CircleImage(url: user.picture, size: 30)
                
                Text(user.full_name)
                  .DemoiosText(fontSize: .base)
                
                
                Spacer()
                
                getInviteStatusIcon(status: user.invite_status ?? "")
              }
              //              .padding(.horizontal)
              //              .padding([ .bottom], 8)
              
              
            } // foreach
          }
          
        }
        .navigationTitle(workout.full_title)
        Divider()
        Spacer()
        
        VStack {
          
          VStack {
            
            Text("If you are ready to start this workout, send it to your watch!")
              .DemoiosText(fontSize: .xs, color: .graySix, weight: .bold)
              .lineLimit(1)
            
            Text("Once you send it, open the Demoios Watch App or look for a notification on your watch to get started!")
              .DemoiosText(fontSize: .xs, color: .graySix)
            
            Spacer()
            
          }
          
          Button("Send workout to Apple Watch") {
            
            //              triggerLocalNotification(body: "Start your workout!")
            
            data.send(message: ["start_workout_from_accepted_invite":
                                  [
                                    "activityType": workout.healthkit_workout_activity_type.capitalized,
                                    "dbWorkoutId": workout.id.uuidString,
                                    "leaderPicture": workout.leader.picture,
                                    "leaderName": workout.leader.first_name
                                  ]
                               ])
            
            dismiss()
          }
          .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
        }
      }
      .task {
        if sheetType == .pending {
          inviteStatuses = await data.api.getUserInviteStatusesForWorkout(workoutId: workout.id)
        }
      }
      //      .navigationTitle(workout.full_title)
      .navigationBarTitleDisplayMode(.inline)
      .toolbar {
        ToolbarItem(placement: .primaryAction) {
          Button("Delete") {
            Task {
              if sheetType == .accepted {
                do {
                  try await data.api.removeWorkoutInvite(workout_id: workout.id)
                  dispatchAsync {
                    self.acceptedWorkouts.removeAll(where: { $0.id == workout.id })
                  }
                  dismiss()
                } catch {
                  triggerLocalNotification(title: "Failed to delete workout", body: "", background: .red, icon: .warning)
                }
              } else {
                do {
                  try await data.api.deleteWorkout(workout.id)
                  dispatchAsync {
                    self.pendingWorkouts.removeAll(where: { $0.id == workout.id} )
                  }
                  dismiss()
                } catch {
                  triggerLocalNotification(title: "Failed to delete workout", body: "", background: .red, icon: .warning)
                }
              }
            }
          }
          
        }
      }
      .environment(\.colorScheme, .dark)
      .padding()
      
      
    } else {
      ProgressView()
    }
    
  }
}

struct AcceptedWorkoutStartSheet_Previews: PreviewProvider {
  static var previews: some View {
    var data = Datamodel()
    NavigationStack {
      
      WorkoutTabStartSheet(workout: .constant(.defaultWorkout), inviteStatuses: [.default], sheetType: .accepted, acceptedWorkouts: .constant([]), pendingWorkouts: .constant([]))
        .environmentObject(data)
        .setupPreview()
    }
  }
}

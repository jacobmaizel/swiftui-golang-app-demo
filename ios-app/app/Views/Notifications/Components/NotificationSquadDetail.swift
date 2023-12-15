//
//  NotificationSquadDetail.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/4/23.
//

import SwiftUI

struct NotificationSquadDetail: View {
  
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) var dismiss
  
  var noti: Notification?
  var squadId: UUID
  @State var loadedSquad: Squad?
  @State var err: APIError?
  
  @State var loaded: Bool = false
  
  @State var relatedInvite: Invite?
  @State var inviteLoaded: Bool = false
  @State var inviteLoadingErr: APIError?
  
  var body: some View {
    
    VStack {
      if noti == nil {
        Text("Failed to load Squad Info.")
      } else if !loaded {
        ProgressView()
      } else if err != nil {
        Text("Failed to load workout.")
      } else if let loadedSquad = loadedSquad {
        VStack {
          HStack {
            VStack(alignment: .leading) {
              Text(loadedSquad.title)
                .DemoiosText(fontSize: .xl, weight: .bold)
              Text("\(loadedSquad.member_count) Member\(loadedSquad.member_count == 1 ? "" : "s")")
              
                .DemoiosText(fontSize: .small, weight: .regular)
              Text("Created \(loadedSquad.created_at.stringToFullDateTime)")
              
                .DemoiosText(fontSize: .small, color: .grayFour)
                .padding(.top, 4)
            }
            Spacer()
              HStack {
                CircleImage(url: loadedSquad.owner.picture, size: 30)
                VStack(alignment: .leading) {
                  Text("\(loadedSquad.owner.full_name)")
                    .DemoiosText(fontSize: .base, weight: .bold)
                Text("Created by")
                  .DemoiosText(fontSize: .small, color: .grayFour)
                }
              }
          }
          .padding()
          
          Divider()
          Spacer()
          
          if let inv = relatedInvite, inv.status == "pending" {
            HStack {
              Button("Accept Invite") {
                
                if let sId = noti?.data.squad_id, let iID = noti?.data.invite_id {
                  
                  Task {
                    do {
                      try await data.api.acceptSquadInvite(inviteID: iID, squadID: sId)
                      
                      triggerLocalNotification(title: "Joined the squad!", body: "", background: .green, icon: .info)
                      
                      dismiss()
                    } catch {
                      NotificationHandler.shared.new(title: "Failed to accept invite", body: "Please try again, or ask the owner to re-invite you.", background: .primarySeven, icon: .warning)
                      dismiss()
                    }
                  }
                  
                } else {
                  print("COULD NOT GET WORKOUT ID OR INVITE ID FROM NOTI TO ACCEPT WORKOUT")
                  
                  
                  dispatchAsync {
                    NotificationHandler.shared.new(title: "Failed to accept invite", body: "Please try again, or ask the leader to re-invite you.", background: .primarySeven, icon: .warning)
                  }
                  
                }
                
                
              }
              .buttonStyle(CustomizableResponsiveButton(fontSize: .large, fontWeight: .bold, bgColor: .primarySeven, fgColor: Color.grayOne))
              
              Button("Decline") {
                
                if let sId = noti?.data.squad_id, let iID = noti?.data.invite_id {
                  
                  
                  
                  Task {
                    do {
                      try await data.api.denySquadInvite(inviteID: iID, squadID: sId)
                      triggerLocalNotification(title: "Invitation declined", body: "", background: .green, icon: .info)
                      dismiss()
                      
                    } catch {
                      triggerLocalNotification(title: "Failed to decline invite", body: "Please try again, or contact support", background: .primarySeven, icon: .warning)
                      dismiss()
                    }
                  }
                  
                } else {
                  print("COULD NOT GET WORKOUT ID OR INVITE ID FROM NOTI TO ACCEPT WORKOUT")
                  
                  
                  triggerLocalNotification(title: "Failed to accept invite", body: "Please try again, or ask the leader to re-invite you.", background: .primarySeven, icon: .warning)
               
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
            let success: Squad = try await data.api.getSquadById(squadId: squadId)
            
            self.loadedSquad = success
            self.loaded = true
          } catch {
            
            triggerLocalNotification(title: "Failed to get Squad info", body: "", background: .red)
            self.loaded = true
          }
        }
        
      }
    }
  }
} // here

struct NotificationSquadDetail_Previews: PreviewProvider {
    static var previews: some View {
      let data = Datamodel()
      NotificationSquadDetail(noti: .test_workout_unopened, squadId: UUID(), loadedSquad: .owned_joined, err: nil, loaded: true, relatedInvite: .test, inviteLoaded: true, inviteLoadingErr: nil)
        .environmentObject(data)
        .setupPreview()
        .frame(height: 300)
    }
}

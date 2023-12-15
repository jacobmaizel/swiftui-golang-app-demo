//
//  NotificationCompetitionDetail.swift
//  Demoios
//
//  Created by Jacob Maizel on 9/4/23.
//

import SwiftUI


struct NotificationCompetitionDetail: View {
  
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) var dismiss
  
  var noti: Notification?
  var compId: UUID
  @State var loadedComp: Competition?
  @State var err: APIError?
  
  @State var loaded: Bool = false
  
  @State var relatedInvite: Invite?
  @State var inviteLoaded: Bool = false
  @State var inviteLoadingErr: APIError?
  
  var body: some View {
    
    VStack {
      if noti == nil {
        Text("Failed to load comp.")
      } else if !loaded {
        ProgressView()
      } else if err != nil {
        Text("Failed to load Comp.")
      } else if let loadedComp = loadedComp {
        
        VStack {
          
          HStack {
            
            VStack(alignment: .leading) {
              
              Text(loadedComp.title)
                .DemoiosText(fontSize: .xl, weight: .bold)
              Text("\(loadedComp.participant_count) Participant\(loadedComp.participant_count == 1 ? "" : "s")")
                .DemoiosText(fontSize: .small, weight: .regular)
              
              Text("Created \(loadedComp.created_at.stringToFullDateTime)")
                .padding(.top, 4)
                .DemoiosText(fontSize: .small, color: .grayFour)
            }
            
            Spacer()
            
            
            HStack {
              CircleImage(url: loadedComp.owner.picture, size: 30)
              VStack(alignment: .leading) {
                Text("\(loadedComp.owner.full_name)")
                  .DemoiosText(fontSize: .base, weight: .bold)
                Text("Created by")
                  .DemoiosText(fontSize: .small, color: .grayFour)
              }
              
            }
            
          }
          
          Divider()
          Spacer()
          
          if let inv = relatedInvite, inv.status == "pending" {
            HStack {
              Button("Accept Invite") {
                
                if let cId = noti?.data.competition_id, let iID = noti?.data.invite_id {
                  
                  
                  Task {
                    do {
                      try await data.api.acceptCompInvite(inviteID: iID, compId: cId)
                      
                      triggerLocalNotification(title: "Joined the comp!", body: "Next, Start a workout!", background: .green, icon: .info)
                      
                      dismiss()
                    } catch {
                      NotificationHandler.shared.new(title: "Failed to accept invite", body: "Please try again, or ask the leader to re-invite you.", background: .primarySeven, icon: .warning)
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
              
              Button("Decline Invite") {
                
                if let cId = noti?.data.competition_id, let iID = noti?.data.invite_id {
                  
                  
                  
                  Task {
                    do {
                      try await data.api.declineCompInvite(inviteID: iID, compId: cId)
                      triggerLocalNotification(title: "Invitation declined", body: "", background: .green, icon: .info)
                      dismiss()
                      
                    } catch {
                      
                      triggerLocalNotification(title: "Failed to decline invite", body: "Please try again, or contact support", background: .primarySeven, icon: .warning)
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
            let success: Competition = try await data.api.getCompById(id: compId)
            
            self.loadedComp = success
            self.loaded = true
          } catch {
            
            triggerLocalNotification(title: "Failed to get Competition info", body: "", background: .red)
            self.loaded = true
          }
        }
        
      }
    }
  }
} // here





struct NotificationCompetitionDetail_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    NotificationCompetitionDetail(noti: .test_comp_unopened, compId: UUID(), loadedComp: .test_started_notjoined, err: nil, loaded: true, relatedInvite: .test, inviteLoaded: true, inviteLoadingErr: nil)
      .environmentObject(data)
      .setupPreview()
      .frame(height: 300)
  }
}

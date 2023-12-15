//
//  SquadInvitesView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/23/23.
//

import SwiftUI

struct SquadInvitesView: View {
  
  @EnvironmentObject var data: Datamodel
  
  var body: some View {
    
    ScrollView {
      
      VStack {
        
        ForEach(data.invites, id: \.id) { invite in
          
          ZStack {
            RoundedRectangle(cornerRadius: 16)
              .foregroundColor(.grayTen)
            
            VStack {
              HStack {
                
                //                Text("\(invite.sender.first_name) invited you to \(invite.squad_title)")
                //                  .DemoiosText(fontSize: .base)
                
                Spacer()
                VStack {
                  
                  Text("Invited")
                    .foregroundColor(.graySix)
                    .DemoiosText(fontSize: .xxs)
                  
                  Text("\(invite.created_at.intoDate?.timeAgo() ?? "Just now")")
                    .DemoiosText(fontSize: .base)
                }
              }
              .padding(.horizontal)
              
              HStack {
                Spacer()
                Button {
                  print("Deprecated?")
//                  Task {
//                    do {
//                      try await data.api.acceptSquadInvite(inviteID: invite.id)
//
//                      do {
//                        try await data.api.getMyPendingInvites()
//                      } catch {
//                        triggerLocalNotification(title: "Failed to load invites", body: "Please try again", background: .red)
//                      }
//                    } catch {
//                      triggerLocalNotification(title: "Failed to accept squad invite", body: "" , background: .red)
//                    }
//                  }
                } label: {
                  Text("Accept")
                    .foregroundColor(.green)
                    .DemoiosText(fontSize: .base)
                    .padding(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
                    .background(Color.grayTen)
                    .cornerRadius(8)
                    .overlay {
                      RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                        .stroke(.green)
                    }
                }
                
                Spacer()
                Button {
//                  Task {
//                    do {
//                      try await data.api.acceptSquadInvite(inviteID: invite.id, squadID: squ)
//                      
//                      do {
//                        try await data.api.getMyPendingInvites()
//                      } catch {
//                        triggerLocalNotification(title: "Failed to load invites", body: "Please try again", background: .red)
//                      }
//                    } catch {
//                      triggerLocalNotification(title: "Failed to accept squad invite", body: "" , background: .red)
//                    }
//                    
//                    
//                  }
                  
                } label: {
                  Text("Deny")
                    .foregroundColor(.primarySeven)
                    .DemoiosText(fontSize: .base)
                    .padding(EdgeInsets(top: 4, leading: 16, bottom: 4, trailing: 16))
                    .background(Color.grayTen)
                    .cornerRadius(8)
                    .overlay {
                      RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                        .stroke(Color.primarySeven)
                    }
                }
                Spacer()
              }
            }
            .padding()
          }
          .fixedSize(horizontal: false, vertical: true)
          //                        .frame(height: 50)
          .padding(.horizontal)
        }
      } // Top VSTack
      .padding()
    }
    .task {
      //      data.api.getMyPendingInvites(data)
      do {
        try await data.api.getMyPendingInvites()
      } catch {
        triggerLocalNotification(title: "Failed to load invites", body: "Please try again", background: .red)
      }
      
    }
  }
}

struct SquadInvitesView_Previews: PreviewProvider {
  static var previews: some View {
    SquadInvitesView()
  }
}

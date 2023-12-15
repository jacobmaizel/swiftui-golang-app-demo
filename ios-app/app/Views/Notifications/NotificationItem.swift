//
//  NotificationItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/24/23.
//

import SwiftUI

struct NotificationItem: View {
  
  @Binding var notification: Notification
  
  var body: some View {
    SectionItem(fixedBy: .v) {
      HStack(alignment: .center) {
        
        VStack(alignment: .leading) {
          
          VStack(alignment: .leading) {
            
            Text(notification.title)
              .DemoiosText(fontSize: .base)
            Text(notification.body)
              .DemoiosText(fontSize: .base, color: .grayFour)
          }
          //          }
//          .padding(.bottom, 2)
          
          VStack(alignment: .leading) {
            
            
            if notification.opened != nil {
              
              Text("Opened \(notification.opened?.stringToFullDateTime ?? "")")
                .DemoiosText(fontSize: .small, color: .graySix)
            } else {
            Text("Received \(notification.sent.stringToFullDateTime)")
              .DemoiosText(fontSize: .small, color: .graySix)
            }
            
          }
        }
        Spacer()
        VStack(alignment: .trailing) {
          if notification.opened == nil {
            Image(systemName: "envelope.fill")
              .foregroundColor(.primarySeven)
            
          } else {
            Image(systemName: "envelope.open")
              .foregroundColor(.primarySeven)
          }
          
          Spacer()
          
          HStack(spacing: 0) {
            Group {
              if notification.data.squad_id != nil {
                Text("Squad Invite")
              }
              if notification.data.workout_id != nil {
                Text("Workout Invite")
              }
              if notification.data.competition_id != nil {
                Text("Competition Invite")
              }
            }
            .DemoiosText(fontSize: .base, color: .primarySeven)
          }
          
        }
      }
      .padding()
    }
    //    .sheet(isPresented: $showNotiInviteSheetDetail) {
    //
    //    }
  }
}

struct NotificationItem_Previews: PreviewProvider {
  static var previews: some View {
    NotificationItem(notification: .constant(.test_opened))
      .setupPreview()
    NotificationItem(notification: .constant(.test_unopened))
      .setupPreview()
  }
}

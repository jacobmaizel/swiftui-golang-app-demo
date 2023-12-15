//
//  NotifcationDetailSheet.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/17/23.
//

import SwiftUI

struct NotificationDetailSheet: View {
  var noti: Notification?
  var body: some View {
    
    VStack {
      
      if let wid = noti?.data.workout_id {
        NotificationWorkoutDetail(noti: noti, workoutId: wid)
          .padding()
      }
      if let cid = noti?.data.competition_id {
        NotificationCompetitionDetail(noti: noti, compId: cid)
          .padding()
      }
      if let sid = noti?.data.squad_id {
        NotificationSquadDetail(noti: noti, squadId: sid)
      }
      
      
      // MARK: - Default view
      
      
      if noti?.data.workout_id == nil && noti?.data.competition_id == nil && noti?.data.squad_id == nil {
        
        VStack {
          
          HStack {
            Text(noti?.title ?? "")
              .DemoiosText(fontSize: .large, weight: .bold)
            Text(noti?.body ?? "")
              .DemoiosText(fontSize: .base, color: .graySix)
            Spacer()
          }
          Divider()
          Spacer()
        }
        .padding()
      }
    }
  }
}

struct NotifcationDetailSheet_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    
    NotificationDetailSheet(noti: .test_workout_unopened)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("Workout")
    
    NotificationDetailSheet(noti: .test_comp_unopened)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("Competition")
    
    
    NotificationDetailSheet(noti: .test_squad_unopened)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("Squad")
    
  }
}



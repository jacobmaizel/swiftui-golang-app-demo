//
//  CompetitionDetailHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/26/23.
//

import SwiftUI

struct CompetitionDetailHeader: View {
  var comp: Competition
    var body: some View {
      VStack {
        
//        if comp.state == .ended {
//
//          Text("Ended \(comp.end.intoDate?.timeAgo() ?? "")")
//            .DemoiosText(fontSize: .base, color: .primarySeven, weight: .bold)
//          Divider()
//        }
        
        HStack {
          
          Text("\(comp.public ? "Public" : "Private")")
          Image(systemName: comp.public ? "lock.open.fill" : "lock.fill")
          Spacer()
          
          
//          ElementTagSet(comp: comp)
          
          ElementTag(title: "Workouts", type: .workout_types, values: comp.workout_types, scale: .small, stack: .v, color: .green)
            
        }
        .DemoiosText(fontSize: .base)
        
      HStack {
        
        HStack {
          
          CircleImage(url: comp.owner.picture, size: 40)
          
          VStack(alignment: .leading) {
            Text("Owned By")
              .DemoiosText(fontSize: .small, color: .graySix)
            Text("\(comp.owner.first_name) \(comp.owner.last_name)")
              .DemoiosText(fontSize: .base)
          }
          
        }
        
        Spacer()
        
        HStack {
          
          
          if comp.state == .ended {
            Image(systemName: "timer")
              .resizable()
              .foregroundColor(.primarySeven)
            .frame(width: 40, height: 40)
          } else {

            Image(systemName: "timer")
              .resizable()
              .foregroundColor(.white)
            .frame(width: 40, height: 40)
          }
          
          
          VStack(alignment: .leading) {
            switch comp.state {
            case .active:
            Text("Ends in")
              .DemoiosText(fontSize: .small, color: .graySix)
              Text(comp.end.intoDate!, style: .relative)
              .DemoiosText(fontSize: .base)
            case .ended:
              Text("Ended")
                .DemoiosText(fontSize: .small, color: .primarySeven)
              Text("\(comp.end.intoDate?.timeAgo() ?? "")")
                .DemoiosText(fontSize: .base, color: .primarySeven, weight: .bold)
            case .invalid:
              Text("Fail")
            case .notStarted:
              Text("Starts in")
                .DemoiosText(fontSize: .small, color: .graySix)
              Text(comp.start.intoDate!, style: .relative)
                .DemoiosText(fontSize: .base)
              
            }
          }
          
        }
        
        
      }
    }
    }
}

struct CompetitionDetailHeader_Previews: PreviewProvider {
    static var previews: some View {
      
      let data = Datamodel()
      VStack {
        
      CompetitionDetailHeader(comp: .test_ended_joined)
      CompetitionDetailHeader(comp: .test_started_joined)
      CompetitionDetailHeader(comp: .test_notstarted_notjoined)
      CompetitionDetailHeader(comp: .test_started_notjoined)
      }
        .environmentObject(data)
        .setupPreview()
    }
}

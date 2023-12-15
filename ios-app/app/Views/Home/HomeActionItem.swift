//
//  HomeActionItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/25/23.
//

import SwiftUI

struct HomeActionItem: View {
  
  var title: String
  var description: String
  var icon: ActionItemIcons
  var background: Color = .graySix
  
    var body: some View {
      
      SectionItem(fixedBy: .none, cornerRadius: .large, color: background) {
        HStack {
          
        VStack(alignment: .leading) {
          
          Text(icon.rawValue)
            .DemoiosText(fontSize: .xxl)
          
          Spacer()
          
          Text(title)
            .DemoiosText(fontSize: .base, color: background == .primarySeven ? .grayTen : .grayOne, weight: .bold)
            .multilineTextAlignment(.leading)
          Text(description)
            .DemoiosText(fontSize: .xxs, color: background == .primarySeven ? .bg : .grayTwo)
          
            .multilineTextAlignment(.leading)
        }
        .padding()
          Spacer()
        }
      }
      .frame(minHeight: 180)
      .background(
        Color.black
          .cornerRadius(DemoiosCornerRadius.large.rawValue)
          .shadow(color: .black, radius: 5, x: 0, y: 5)
      )
    }
}

struct HomeActionItem_Previews: PreviewProvider {
    static var previews: some View {
      VStack(alignment: .leading) {
        HStack(alignment: .center) {
          
          HomeActionItem(title: "View your competitions", description: "Check out the latest info on your competitions!", icon: .viewComps, background: .primarySeven)
            .setupPreview()
            .frame(width: 200, height: 200)
          
          HomeActionItem(title: "Create a workout", description: "Get active and Invite your friends!", icon: .createWorkout, background: .grayTen)
            .setupPreview()
            .frame(width: 200, height: 200)
        }
        HStack {
          
          HomeActionItem(title: "Find a squad", description: "Create or join a Squad!", icon: .findSquad, background: .grayTen)
            .setupPreview()
            .frame(width: 200, height: 200)
          
          HomeActionItem(title: "View your stats", description: "Total distance, comp wins, and more", icon: .checkStats, background: .grayTen)
            .setupPreview()
            .frame(width: 200, height: 200)
        }
      }
      .setupPreview()
    }
}


enum ActionItemIcons: String {
  // Create a workout
  // View your competitions
  // Find a squad
  // Check out your stats

  case viewComps = "👀"
  case createWorkout = "🏋🏽‍♂️"
  case findSquad = "🔎"
  case checkStats = "📈"
}

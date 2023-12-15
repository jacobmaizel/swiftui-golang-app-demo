//
//  ProfileDataSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/17/23.
//

import SwiftUI

struct ProfileDataSection: View {
  
  
  @Binding var dataChoice: ProfileDataOptions
  var viewingOwnProfile: Bool
  
  
    var body: some View {
      
      
      VStack(spacing: 0) {
        HStack {
          
          Spacer()
          if viewingOwnProfile {
            Button {
              dataChoice = .squads
            } label: {
            Text("Squads")
            }
            .DemoiosText(fontSize: dataChoice == .squads ? .large : .base, color: dataChoice == .squads ? .grayOne : .graySix, weight: dataChoice == .squads ? .bold : .regular )
            
            Spacer()
          }
          Button {
            dataChoice = .workouts
          } label: {
          Text("Workouts")
          }
          .DemoiosText(fontSize: dataChoice == .workouts ? .large : .base, color: dataChoice == .workouts ? .grayOne : .graySix, weight: dataChoice == .workouts ? .bold : .regular )
          
          Spacer()
          Button {
            dataChoice = .competitions
          } label: {
          Text("Comps")
          }
          .DemoiosText(fontSize: dataChoice == .competitions ? .large : .base, color: dataChoice == .competitions ? .grayOne : .graySix, weight: dataChoice == .competitions ? .bold : .regular )
          Spacer()
        }
        .padding()
        
        Divider()
          .background(Color.grayTwo)
      }


    }
}

struct ProfileDataSection_Previews: PreviewProvider {
    static var previews: some View {
      
      ProfileDataSection(dataChoice: .constant(.competitions), viewingOwnProfile: false)
        .setupPreview()
        .previewDisplayName("Viewing other")
      ProfileDataSection(dataChoice: .constant(.competitions), viewingOwnProfile: true)
        .setupPreview()
        .previewDisplayName("Viewing own")
    }
}

enum ProfileDataOptions: String, ListableEnum {
  case squads, workouts, competitions
}

//
//  ProfileDataSectionSelector.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadDataSelector: View {
  
  @Binding var dataChoice: SquadDataChoices
  
    var body: some View {
      VStack(spacing: 0) {
        HStack {
          
          Spacer()
          Button {
            dataChoice = .activity
          } label: {
            Text("Activity")
          }
          .DemoiosText(fontSize: dataChoice == .activity ? .large : .base, color: dataChoice == .activity ? .grayOne : .graySix, weight: dataChoice == .activity ? .bold : .regular )
          
          Spacer()
          Button {
            dataChoice = .members
          } label: {
            Text("Members")
          }
          .DemoiosText(fontSize: dataChoice == .members ? .large : .base, color: dataChoice == .members ? .grayOne : .graySix, weight: dataChoice == .members ? .bold : .regular )
          
          Spacer()
        }
        .padding()
        
          Divider()
            .background(Color.grayTwo)
      }
    }
}

struct ProfileDataSectionSelector_Previews: PreviewProvider {
    static var previews: some View {
      SquadDataSelector(dataChoice: .constant(.activity))
        .setupPreview()
    }
}

enum SquadDataChoices: String, ListableEnum {
  case members, activity
}

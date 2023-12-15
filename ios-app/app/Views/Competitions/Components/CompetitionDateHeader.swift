//
//  CompetitionDateHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/10/23.
//

import SwiftUI

struct CompetitionDateHeader: View {
  
  @EnvironmentObject var data: Datamodel
  var competition: Competition
  var body: some View {
    //    //Date
    VStack(spacing: 0) {
      
      HStack(alignment: .center, spacing: 0) {
        Image(systemName: "calendar")
          .font(.system(size: DemoiosFontSizes.large.size))
          .foregroundColor(.grayFour)
        
        Text("\(competition.start.stringToShortDate) - \(competition.end.stringToShortDate)")
          .foregroundColor(.grayFour)
          .DemoiosText(fontSize: .xxs)
        
        if competition.isActive() {
          Text("Active")
            .foregroundColor(.primarySeven)
            .fontWeight(.bold)
            .DemoiosText(fontSize: .xs)
            .padding(.horizontal)
        }
      } // date hstack
      HStack {
        CircleImage(url: competition.owner.picture, size: 30)
        Text(competition.owner.id == data.profile.id ? "Created By You" : "Created by \(competition.owner.first_name) \(competition.owner.last_name)")
          .DemoiosText(fontSize: .xxs)
      }
      .padding()
    }
    
    .DemoiosText(fontSize: .xxs)
  }
}

struct CompetitionDateHeader_Previews: PreviewProvider {
  static var previews: some View {
    CompetitionDateHeader(competition: Competition.test_notstarted_notjoined)
      .environmentObject(Datamodel())
      .setupPreview()
  }
}

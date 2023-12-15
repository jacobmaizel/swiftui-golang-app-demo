//
//  HomeCompetitionHeaderCard.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/13/23.
//

import SwiftUI

struct HomeCompetitionCard: View {
  
  @State var comp: Competition
  
  var body: some View {
//    GeometryReader { geo in
      
//    NavigationStack {
LargeCompetitionCard(competition: comp )
        
//      }
//    }
  }
}

struct HomeCompetitionHeaderCard_Previews: PreviewProvider {
  static var previews: some View {
    HomeCompetitionCard(comp: .test_notstarted_joined)
      .previewLayout(.sizeThatFits)
      .environment(\.colorScheme, .dark)
      .setupPreview()
      .environmentObject(Datamodel())
    //      .setupPreview
  }
}

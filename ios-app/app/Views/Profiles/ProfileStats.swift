//
//  ProfileStats.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/3/23.
//

import SwiftUI

struct ProfileStats: View {
  
  var profileStats: ProfileStatsData
  var sectionHeight: CGFloat = 100
  
  
//  @EnvironmentObject var data: Datamodel
  
  var body: some View {
    
    SectionItem(fixedBy: .v) {
      
      HStack(alignment: .center) {
        
        GeometryReader { geometry in
          
          HStack(alignment: .center, spacing: 0) {
            // Lifetime points section
            VStack {
              Text("\(profileStats.workout_count)")
                .DemoiosText(fontSize: .xxl)
                .fontWeight(.bold)
              
              Text("Workouts")
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .xs)
            }
            .frame(width: geometry.size.width / 3)
            
            Divider()
              .foregroundColor(.graySix)
              .frame(height: geometry.size.height / 2)
            
            // comps completed section
            VStack {
              
              Text("\(profileStats.competition_wins)")
                .DemoiosText(fontSize: .xxl)
                .fontWeight(.bold)
              
              Text("Comp Wins")
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .xs)
            }
            .frame(width: geometry.size.width / 3)
            
            Divider()
              .foregroundColor(.graySix)
              .frame(height: geometry.size.height / 2)
            VStack {
              
              Text("\(profileStats.miles)")
                .DemoiosText(fontSize: .xxl)
                .fontWeight(.bold)
              
              Text("Total Miles")
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .xs)
            }
            .frame(width: geometry.size.width / 3)
          }// points/comps section hstack
          .frame(height: sectionHeight)
          
        } // geometry reader
      }
      .frame(height: sectionHeight)
    }
    .padding(.horizontal)
    
    
  } // END BODY
}

struct ProfileStats_Previews: PreviewProvider {
  static var previews: some View {
    ProfileStats(profileStats: ProfileStatsData(workout_count: 0, competition_wins: 0, total_distance_ft: 0))
      .environmentObject(Datamodel())
      .previewLayout(.sizeThatFits)
      .setupPreview()
  }
}

//
//  SquadDetailStats.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadDetailStats: View {
  
  @Binding var squad: Squad
  var stats: SquadStats?
  var sectionHeight: Double = 100
  

  
    var body: some View {
      SectionItem(fixedBy: .v) {
        
        HStack(alignment: .center) {
          
          GeometryReader { geometry in
            
            HStack(alignment: .center, spacing: 0) {
              // Lifetime points section
              VStack {
                Text("\(squad.stats?.member_count ?? stats?.member_count ?? 0)")
                  .DemoiosText(fontSize: .xxl)
                  .fontWeight(.bold)
                
                Text("Members")
                  .foregroundColor(.grayEight)
                  .DemoiosText(fontSize: .xs)
              }
              .frame(width: geometry.size.width / 3)
              
              Divider()
                .foregroundColor(.graySix)
                .frame(height: geometry.size.height / 2)
              
              // comps completed section
              VStack {
                
                Text("\(squad.stats?.competition_wins ?? stats?.competition_wins ?? 0)")
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
                
                Text("\(feetToMiles(feet: squad.stats?.total_distance_ft ?? stats?.total_distance_ft ?? 0))")
                  .DemoiosText(fontSize: .xxl)
                  .fontWeight(.bold)
                
                Text("Total Mileage")
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
      
    }
}

struct SquadDetailStats_Previews: PreviewProvider {
    static var previews: some View {
      SquadDetailStats(squad: .constant(.not_owned_joined))
        .environmentObject(Datamodel())
        .setupPreview()
    }
}

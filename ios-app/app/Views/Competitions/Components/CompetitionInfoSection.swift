//
//  CompetitionInfoSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/10/23.
//

import SwiftUI

struct CompetitionInfoSection: View {
  
  var graphData: [CompetitionStatGraphDataPoint]
  var myStats: CompMyStats?
  
  var height: CGFloat = 180
  var body: some View {
    
    VStack {
      VStack(alignment: .center) {
        HStack(spacing: 0) {
          Text("Stats")
            .foregroundColor(.grayEight)
            .DemoiosText(fontSize: .xs)
            .offset(x: 5, y: 10)
          Spacer()
        }
        
        if let myStats = myStats, !myStats.place.isEmpty {
          HStack(alignment: .center, spacing: 0) {
            
            // MARK: - Total Distance
            Spacer()
            
            VStack {
              Text("\(feetToMiles(feet: myStats.distance_total)) Miles")
                .fontWeight(.bold)
                .DemoiosText(fontSize: .xl)
              
              Text("Total Distance")
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .base)
            }
            Spacer()
            // MARK: - Place
            VStack {
              Text("\(ordinalIndicator(for: Int(myStats.place) ?? 0 + 1))")
                .DemoiosText(fontSize: .xl)
                .fontWeight(.bold)
              
              Text("Place")
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .base)
            }
            Spacer()
            
            
          }
        } else {
          
          Spacer()
          VStack {
            Spacer()
            Text("No Stats")
              .DemoiosText(fontSize: .xl)
              .fontWeight(.bold)
            Spacer()
          }
          Spacer()
        }
        
      }
      .padding(.bottom)
      
      VStack(alignment: .center, spacing: 0) {
        HStack {
          Text("Workout Graph")
            .foregroundColor(.grayEight)
            .DemoiosText(fontSize: .xs)
            .offset(x: 5, y: 0)
          
          Spacer()
        }
        
        CompetitionStatGraph(graphData: graphData)
      }
    }
  } // top body
} // top view

struct CompetitionInfoSection_Previews: PreviewProvider {
  static var previews: some View {
    
    //    let data = Datamodel()
    
    CompetitionInfoSection(graphData: CompetitionStatGraphDataPoint.testData)
    //      .environmentObject(data)
      .setupPreview()
    
  }
}

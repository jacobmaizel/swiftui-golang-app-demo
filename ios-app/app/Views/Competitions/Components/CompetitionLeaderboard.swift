//
//  CompetitionLeaderboard.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/8/23.
//

import SwiftUI

struct CompetitionLeaderboard: View {
  
  var leaderboardData: [CompetitionLeaderboardItem]
  
  var body: some View {
    if !leaderboardData.isEmpty {
      VStack(spacing: 0) {
        VStack(alignment: .leading, spacing: 0) {
          HStack {
            Spacer()
            Text("Mi")
          }
          .foregroundColor(.graySix)
          .DemoiosText(fontSize: .base)
          
          ForEach(Array(zip(leaderboardData.indices, leaderboardData)), id: \.0) { index, item in
            
            HStack(alignment: .center) {
              
              ZStack(alignment: .center) {
                
                if index + 1 == 1 {
                  
                  
                Image(systemName: "circle.fill")
                    .foregroundColor(.gold)
                  .scaleEffect(CGSize(width: 2, height: 2))
                 
                  
              Text("\(index + 1)")
                    .DemoiosText(fontSize: .base, color: .grayTen)
                  
                } else if index + 1 == 2 {
                  
                Image(systemName: "circle.fill")
                    .foregroundColor(.silver)
                  .scaleEffect(CGSize(width: 2, height: 2))
                  
              Text("\(index + 1)")
                    .DemoiosText(fontSize: .base, color: .grayTen)
                  
                } else if index + 1 == 3 {
                  
                Image(systemName: "circle.fill")
                    .foregroundColor(.bronze)
                  .scaleEffect(CGSize(width: 2, height: 2))
                  
              Text("\(index + 1)")
                    .DemoiosText(fontSize: .base, color: .grayTen)
                  
                } else {
                  
                Image(systemName: "circle.fill")
                    .foregroundColor(.clear)
                  .scaleEffect(CGSize(width: 2, height: 2))
                  
              Text("\(index + 1)")
                .DemoiosText(fontSize: .base)
                .foregroundColor(.grayFour)
                }
                
              }
              .padding(.trailing)
              
              CircleImage(url: item.picture, size: 50)
              
              Text("\(item.first_name) \(item.last_name)")
                .padding(.trailing)
              
              Spacer()
              
              HStack(spacing: 12) {
                Text("\(feetToMiles(feet: item.distance_total))")
              }
//              .padding(.trailing)
            }
            .DemoiosText(fontSize: .base)
            .padding(EdgeInsets(top: 4, leading: 8, bottom: 4, trailing: 8))
            
            Divider()
          }
        }
//        .padding()
      }
    } else {
      EmptyView()
    }
  }
}

struct CompetitionLeaderboard_Previews: PreviewProvider {
  static var previews: some View {

    ScrollView {
      
      CompetitionLeaderboard(leaderboardData: CompetitionLeaderboardItem.testData.sorted(by: >))
        .setupPreview()
        .environmentObject(Datamodel())
    }
  }
}

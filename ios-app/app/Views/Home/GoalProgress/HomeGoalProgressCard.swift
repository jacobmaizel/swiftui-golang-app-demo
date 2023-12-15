//
//  HomeGoalProgressCard.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import SwiftUI

struct HomeGoalProgressCard: View {
  var body: some View {
    
    GeometryReader { geometry in
      
      HStack(spacing: 0) {
        HomeGoalProgressLeft {
          Image(systemName: "figure.run")
            .resizable()
            .scaledToFit()
            .frame(width: 80, height: 80)
        }
        .frame(width: geometry.size.width * 0.2)
        
        HomeProgressGoalRight {
          
          VStack(spacing: 0) {
            HStack {
              Text("Run 10 Miles")
                .DemoiosText(fontSize: .xl, weight: .bold)
              Spacer()
              Image(systemName: "gearshape.fill")
                .resizable()
                .frame(width: 30, height: 30)
                .foregroundColor(.primarySeven)
              
            }
            .padding(.bottom)
            
            HStack {
              
              ProgressBar(progress: .constant(0.5))
              Text("3 Miles Left")
                .DemoiosText(fontSize: .base, color: .graySix, weight: .bold)
            }
            
            
            
            VStack(spacing: 0) {
              
              HStack {
                
                VStack(alignment: .leading, spacing: 0) {
                  
                  Text("Running")
                    .DemoiosText(fontSize: .base, color: .graySix)
                  
                  Text("Today")
                    .DemoiosText(fontSize: .small, color: .grayEight)
                }
                
                Spacer()
                Text("+0.5 Miles")
                  .DemoiosText(fontSize: .base, color: .green)
                
                Spacer()
                
                HStack {
                  
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -15)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -30)
                }
                
              }
              
              Divider()
              
              HStack {
                
                VStack(alignment: .leading, spacing: 0) {
                  
                  Text("Running")
                    .DemoiosText(fontSize: .base, color: .graySix)
                  
                  Text("Today")
                    .DemoiosText(fontSize: .small, color: .grayEight)
                }
                
                Spacer()
                
                Text("+0.5 Miles")
                  .DemoiosText(fontSize: .base, color: .green)
                
                Spacer()
                
                HStack {
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -15)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -30)
                }
                
              }
              
              Divider()
              HStack {
                
                VStack(alignment: .leading, spacing: 0) {
                  
                  Text("Running")
                    .DemoiosText(fontSize: .base, color: .graySix)
                  
                  Text("Today")
                    .DemoiosText(fontSize: .small, color: .grayEight)
                }
                
                Spacer()
                
                Text("+0.5 Miles")
                  .DemoiosText(fontSize: .base, color: .green)
                
                Spacer()
                
                HStack {
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -15)
                  CircleImage(url: Consts.JAKE_M_PROFILE_IMAGE, size: 30)
                    .offset(x: -30)
                }
              }
              Divider()
            } // MAIN VSTACK
          }
          .padding(EdgeInsets(top: 4, leading: 8, bottom: 4, trailing: 8))
        }
      }
      .padding()
    }
  }
}

struct HomeGoalProgressCard_Previews: PreviewProvider {
  static var previews: some View {
    HomeGoalProgressCard()
      .setupPreview()
  }
}

//
//  HomeGoalCardEmpty.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import SwiftUI

struct HomeGoalCardEmpty: View {
    var body: some View {
        ZStack(alignment: .top) {
            
            Rectangle()
                .fill(Color.grayTen)
                .frame(height: 230)
            
            VStack {
                
                VStack {
                    
                    Text("No Active Goal!")
                        .DemoiosText(fontSize: .xl)
                        .fontWeight(.bold)
                        .padding(.top)
                    
                    Divider()
                }
                .padding(.bottom)
                .padding(.bottom)
                
                HStack {
                    Button {
                        print("create")
                    } label: {
                        Text("Create a Goal")
                            .DemoiosText(fontSize: .base)
                    }.buttonStyle(HomeEmptyCardButtonStyle())
                    
          
                }
            }
        }
        .cornerRadius(16)
        .padding()
    }
}

struct HomeGoalCardEmpty_Previews: PreviewProvider {
    static var previews: some View {
      Group {
        
        HomeGoalCardEmpty()
     
        
        HomeGoalCardEmpty()
//        .environment(\.colorScheme, .dark)
        
      }
      .setupPreview()
    }
}

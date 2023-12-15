//
//  CompCardStatView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/6/23.
//

import SwiftUI

struct CompCardStatView: View {
  
  var sectionHeight: CGFloat = 100
  
    var body: some View {
      SectionItem(fixedBy: .v) {
        
        HStack(alignment: .center) {
          
          GeometryReader { geometry in
            
            HStack(alignment: .center, spacing: 0) {
              // Lifetime points section
              VStack {
                Text("10")
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
                
                Text("2.8M")
                  .DemoiosText(fontSize: .xxl)
                  .fontWeight(.bold)
                
                Text("Mi / Workout")
                  .foregroundColor(.grayEight)
                  .DemoiosText(fontSize: .xs)
              }
              .frame(width: geometry.size.width / 3)
              
              Divider()
                .foregroundColor(.graySix)
                .frame(height: geometry.size.height / 2)
              VStack {
                
                Text("26.1")
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
//      .padding(.horizontal)
    }
}

struct CompCardStatView_Previews: PreviewProvider {
    static var previews: some View {
        CompCardStatView()
    }
}

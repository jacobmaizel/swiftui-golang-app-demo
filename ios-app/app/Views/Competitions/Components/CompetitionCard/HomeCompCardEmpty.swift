//
//  HomeCompCardEmpty.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import SwiftUI

struct HomeCompCardEmpty: View {
  var body: some View {
    
//    NavigationStack {
    
    ZStack(alignment: .center) {
      
      Rectangle()
        .fill(Color.grayTen)
        .frame(height: 230)
      
//
//        VStack {
//
//          Text("No Competition Joined!")
//            .DemoiosText(fontSize: .xl)
//            .fontWeight(.bold)
//            .padding(.top)
//
//          Divider()
//        }
//        .padding(.bottom)
//        .padding(.bottom)
        
        HStack {
          NavigationLink {
            CompetitionCreateForm()
          } label: {
            Text("Create A Competition")
              .DemoiosText(fontSize: .base)
          }
          .buttonStyle(CustomizableResponsiveButton())
          
          
          NavigationLink {
            SearchView()
          } label: {
            
            HStack {
              Text("Search for one")
                .DemoiosText(fontSize: .base)
              Image(systemName: "magnifyingglass")
            }
            
          }
            .buttonStyle(CustomizableResponsiveButton())
        }
     
    }
    .cornerRadius(16)
    .padding()
//  }
        
    }
}

struct HomeCompCardEmpty_Previews: PreviewProvider {
    static var previews: some View {
        HomeCompCardEmpty()
    }
}

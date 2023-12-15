//
//  ProfilePastComps.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/4/23.
//

import SwiftUI

struct ProfilePastComps: View {
    
    var body: some View {
        
        NavigationStack {
            
            NavigationLink {
                CompetitionListView()
                    .navigationTitle("Past Competitions")
            } label: {
                SectionItem(fixedBy: .v) {
                    
                    HStack {
                        
                        Text("View Past Competitions")
                            .DemoiosText(fontSize: .base)
                        
                        Spacer()
                        
                        Image(systemName: "arrow.right")
                            .foregroundColor(.grayFour)
                    }
                    .padding()
                }
                .padding(.horizontal)
            }
        }
    }// body
}// Top View

struct ProfilePastComps_Previews: PreviewProvider {
    static var previews: some View {
        ProfilePastComps()
    }
}

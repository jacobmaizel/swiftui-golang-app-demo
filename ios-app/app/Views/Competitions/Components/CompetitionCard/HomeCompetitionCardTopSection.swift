//
//  HomeCompetitionCardTopSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/13/23.
//

import SwiftUI

struct HomeCompetitionCardTopSection<Content: View>: View {
    
    let content: Content
    
    init(@ViewBuilder _ content: () -> Content) {
        self.content = content()
    }
    
    var body: some View {
        ZStack(alignment: .top) {
            RoundedCorner(radius: DemoiosCornerRadius.large.rawValue, corners: [.topLeft, .topRight])
                .fill(Color.grayTen)
//                .frame(height: 250)
            content
                .padding()
        }
        .fixedSize(horizontal: false, vertical: true)
    }
}

struct HomeCompetitionCardTopSection_Previews: PreviewProvider {
    static var previews: some View {
        HomeCompetitionCardTopSection {
            
        }
    }
}

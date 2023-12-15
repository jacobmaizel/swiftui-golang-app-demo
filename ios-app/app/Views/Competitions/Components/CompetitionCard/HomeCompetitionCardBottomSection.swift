//
//  HomeCompetitionCardBottomSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/13/23.
//

import SwiftUI

struct HomeCompetitionCardBottomSection<Content: View>: View {
    
    let content: Content
    
    init(@ViewBuilder _ content: () -> Content) {
        self.content = content()
    }
    
    var body: some View {
        ZStack {
            RoundedCorner(radius: DemoiosCornerRadius.large.rawValue, corners: [.bottomLeft, .bottomRight])
                .fill(Color.graySix)
                .frame(height: 50)
            
            content
        }
    }
}

struct HomeCompetitionCardBottomSection_Previews: PreviewProvider {
    static var previews: some View {
        HomeCompetitionCardBottomSection {
            
        }
    }
}

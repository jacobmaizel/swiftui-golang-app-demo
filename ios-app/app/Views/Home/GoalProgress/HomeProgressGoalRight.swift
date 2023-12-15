//
//  HomeProgressGoalRight.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import SwiftUI

struct HomeProgressGoalRight<Content: View>: View {
    
    let content: Content
    
    init(@ViewBuilder _ content: () -> Content) {
        self.content = content()
    }
    
    var body: some View {
        ZStack {
            RoundedCorner(radius: DemoiosCornerRadius.large.rawValue, corners: [.bottomRight, .topRight])
                .fill(Color.grayTen)
                .frame(height: 230)
            
            content
        }
    }
}

struct HomeProgressGoalRight_Previews: PreviewProvider {
    static var previews: some View {
        HomeProgressGoalRight {
            
        }
    }
}

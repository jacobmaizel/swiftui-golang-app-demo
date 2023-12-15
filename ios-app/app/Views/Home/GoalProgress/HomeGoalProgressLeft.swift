//
//  HomeGoalProgressLeft.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import SwiftUI

struct HomeGoalProgressLeft<Content: View>: View {
    
    let content: Content
    
    init(@ViewBuilder _ content: () -> Content) {
        self.content = content()
    }
    
    var body: some View {
        ZStack {
            RoundedCorner(radius: DemoiosCornerRadius.large.rawValue, corners: [.bottomLeft, .topLeft])
                .fill(Color.primarySeven)
                .frame(height: 230)
            
            content
        }
    }
}

struct HomeGoalProgressLeft_Previews: PreviewProvider {
    static var previews: some View {
        HomeGoalProgressLeft {
            
        }
    }
}

//
//  ActiveWorkoutView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/2/23.
//

import SwiftUI

struct ActiveWorkoutView: View {
    
    enum ActiveWorkoutTab {
        case stateControls
        case statView
    }
    
    @State private var selectedTab: ActiveWorkoutTab = .statView
    
    var body: some View {
        TabView(selection: $selectedTab) {
            ActiveWorkoutStateControlView()
                .tabItem {
                    Image(systemName: "circle.fill")
                }
                .tag(ActiveWorkoutTab.stateControls)
            ActiveWorkoutStatsView()
                .tabItem {
                    Image(systemName: "circle.fill")
                }
                .tag(ActiveWorkoutTab.statView)
        }
        .navigationBarBackButtonHidden()
    }
}

struct ActiveWorkoutView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveWorkoutView()
    }
}

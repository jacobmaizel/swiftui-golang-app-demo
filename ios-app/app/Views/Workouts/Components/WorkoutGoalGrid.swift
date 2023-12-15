//
//  WorkoutGoalGrid.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct WorkoutGoalGrid: View {
    let columns = Array(repeating: GridItem(.flexible(), spacing: 8), count: 2)
    
    var state: WorkoutState
    
    var body: some View {
        
        if state == .active {
            
            LazyVGrid(columns: columns) {
                WorkoutGoalItem(title: "Total Miles")
                WorkoutGoalItem(title: "Fastest Mile")
                WorkoutGoalItem(title: "WIN EVERYTHING" )
                WorkoutGoalItem(title: "Destroy" )
            }
        } else {
            
            LazyVGrid(columns: columns) {
                WorkoutGoalItem(title: "Total Miles", progress: 1)
                WorkoutGoalItem(title: "Fastest Mile", progress: 1)
            }
        }
    }
}

struct WorkoutGoalGrid_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutGoalGrid(state: .active)
        .setupPreview()
    }
}

//
//  StartWorkoutView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/2/23.
//

import SwiftUI
import HealthKit

struct StartWorkoutView: View {
    
    @EnvironmentObject var workoutController: WorkoutController
    
    var body: some View {
            CompetitionChoiceList()
//                .navigationTitle("Demoios")
//                .navigationBarTitleDisplayMode(.inline)

                .navigationTitle("Choose A Competition")
    }
}

struct StartWorkoutView_Previews: PreviewProvider {
    static var previews: some View {
        StartWorkoutView()
    }
}

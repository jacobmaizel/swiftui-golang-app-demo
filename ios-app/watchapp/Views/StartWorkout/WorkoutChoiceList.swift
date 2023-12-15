//
//  WorkoutChoiceList.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/3/23.
//

import SwiftUI
import HealthKit

struct WorkoutChoiceList: View {
    @EnvironmentObject var wc: WorkoutController
    
    var workoutChoices: [HKWorkoutActivityType] = [.cycling, .running, .walking, .swimming]
    
    var body: some View {
        
      
            VStack {
                Text("Choose a Workout")
                    .font(.caption)
                
                ScrollView {
                    ForEach(workoutChoices, id: \.self) { choice in
                        
                        NavigationLink {
                            WorkoutLocationChoiceList()
                        } label: {
                          Text(choice.name)
                        }
                        .simultaneousGesture(TapGesture().onEnded {
                            DispatchQueue.main.async {
                                self.wc.selectedWorkout = choice
                            }
                        })
                    }
                }
            }
       
    }
}

struct WorkoutChoiceList_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutChoiceList()
        .environmentObject(WorkoutController())
        .setupPreview()
    }
}

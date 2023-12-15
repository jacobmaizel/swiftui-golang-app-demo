//
//  WorkoutCompetitionHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct WorkoutCompetitionHeader: View {
    
    var state: WorkoutState
    var workout: Workout
    
    var body: some View {
        
        SectionItem(fixedBy: .v) {
            
            HStack {
                VStack(alignment: .leading, spacing: 0) {
                    HStack(alignment: .center) {
                        
                        Text(workout.competition?.title ?? "No Competition")
                            .DemoiosText(fontSize: .large)
                        
                        Spacer()
                    }
                    
//                    Spacer()
                    
                    if workout.competition != nil {
//                        HStack {
//                            ElementTag(color: .blue, type: "Endure", size: .medium)
//                            ElementTag(color: .green, type: "Run", size: .medium)
//                            ElementTag(color: .orange, type: "Fast", size: .medium)
//                        }
                    }
                    Spacer()
                    
                    Text("Competition")
                        .foregroundColor(.graySix)
                        .DemoiosText(fontSize: .xxs)
                }
                .padding(.horizontal)
                
                Spacer()
                
                if state == .active {
                    VStack(alignment: .leading, spacing: 0) {
                        
                        HStack(alignment: .bottom) {
                            
                            Text("999th")
                                .fontWeight(.bold)
                                .DemoiosText(fontSize: .xl)
                            Text("/ 100")
                                .foregroundColor(.grayFour)
                                .DemoiosText(fontSize: .xxs)
                        }
                        Text("Place")
                            .foregroundColor(.graySix)
                            .DemoiosText(fontSize: .xxs)
                    }
                    .padding(.horizontal)
                } else {
                    // STATE IS SUMMARY
                    
                    VStack(alignment: .leading, spacing: 0) {
                        
                        HStack(alignment: .bottom) {
                            
                          Text(workout.workout_data?.healthkit_workout_end_date.intoDate?.timeAgo() ?? "")
                                .fontWeight(.bold)
                                .DemoiosText(fontSize: .large)
                        }
                        Text("Completed")
                            .foregroundColor(.graySix)
                            .DemoiosText(fontSize: .xxs)
                    }
                    .padding(.horizontal)
                }
            }
            //                    .padding()
        }
    }
}

struct WorkoutCompetitionHeader_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutCompetitionHeader(state: .active, workout: .defaultWorkout)
    }
}

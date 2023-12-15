//
//  ActiveWorkoutStateControlView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/2/23.
//

import SwiftUI

struct ActiveWorkoutStateControlView: View {
    
    @EnvironmentObject var wc: WorkoutController
    
    var body: some View {
        NavigationStack {
            ScrollView {
                VStack {
                    Text(wc.running ? "Active" : "Paused")
                        .font(.title3)
                        .fontWeight(.bold)
                        .foregroundColor(wc.running ? .green : .yellow)
                  
                  
//                  if wc.errString != "" {
//                    Text("ERROR::: \(wc.errString)")
//                  }
                    
                    VStack(alignment: .leading) {
                        if let comp = wc.selectedCompetition?["title"]{
                          VStack(alignment: .leading) {
                              Text("Competition")
                                  .font(.caption)
                                  .foregroundColor(.gray)
                              
                              Text(comp)
                                  .font(.body)
                                  .fontWeight(.bold)
                          }
                        }
                        Divider()
                      if let workoutType = wc.selectedWorkout?.name {
                          VStack(alignment: .leading) {
                              Text("Workout Type")
                                  .font(.caption)
                                  .foregroundColor(.gray)
                              
                            Text("\(wc.selectedWorkoutLocationType.getString()) \(workoutType)")
                                  .font(.body)
                                  .fontWeight(.bold)
                          }
                        }
                        
                        HStack {
                            
                            Button {
                                wc.togglePause()
                            } label: {
                                Text(wc.running ? "Pause" : "Resume")
                            }
                            .tint(wc.running ? .yellow : .green)
                            
                            NavigationLink {
                                WorkoutSummaryView()
                            } label: {
                                Text("End")
                            }
                            .tint(.red)
                            .simultaneousGesture(TapGesture().onEnded {
                                self.wc.endWorkout()
                            })
                        }
                        .padding(.horizontal)
                    }
                } // TOP VSTACK
            }
            .navigationTitle("Demoios")
            .navigationBarBackButtonHidden()
        } // Nav stack
    }
}

struct WorkoutStateControlView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveWorkoutStateControlView()
        .environmentObject(WorkoutController())
    }
}

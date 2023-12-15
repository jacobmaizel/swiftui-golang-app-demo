//
//  StartWorkoutConfirmation.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/3/23.
//

import SwiftUI

struct StartWorkoutConfirmation: View {
    @EnvironmentObject var wc: WorkoutController
    
    var body: some View {
        
        NavigationStack {
            VStack {
                
                VStack {
                  if let compTitle = wc.selectedCompetition?["title"] {
                    VStack {
                      HStack {
                        
                        Text("Competition")
                          .font(.footnote)
                          .foregroundColor(.gray)
                        Spacer()
                      }
                      HStack {
                        Spacer()
                        
                        Text(compTitle)
                          .fontWeight(.bold)
                          .font(.body)
                      }
                    }
                    .padding(.horizontal)
                  }
                  
                  
                  if wc.acceptedInviteWorkoutData != nil {
                    
                    HStack {
                      CircleImage(url: wc.acceptedWorkoutLeaderPicture ?? "", size: 30)
                      VStack(alignment: .leading) {
                        Text("Lead By")
                          .foregroundColor(.gray)
                        Text(wc.acceptedWorkoutLeaderName ?? "User")
                        
                      }
                    }
                    
                    Divider()
                    
                  }
                  
                  
                    VStack {
                        HStack {
                            
                            Text("Workout Type")
                                .font(.footnote)
                                .foregroundColor(.gray)
                            Spacer()
                        }
                        
                        HStack {
                            Spacer()
                            
                            Text(wc.selectedWorkout?.name ?? "No Workout Selected")
                                .fontWeight(.bold)
                                .font(.body)
                        }
                    }
                    .padding(.horizontal)
                    
                    VStack {
                        HStack {
                            
                            Text("Location")
                                .font(.footnote)
                                .foregroundColor(.gray)
                            Spacer()
                        }
                        HStack {
                            Spacer()
                            
                            Text(wc.selectedWorkoutLocationType.getString())
                                .fontWeight(.bold)
                                .font(.body)
                        }
                    }
                    .padding(.horizontal)
                }
                .padding(.vertical)
                
                //                Text(wc.workout_debug)
                
                Spacer()
                
                NavigationLink {
                    ActiveWorkoutView()
                } label: {
                    Text("Start Workout")
                }
                .simultaneousGesture(TapGesture().onEnded {
                  if wc.acceptedInviteWorkoutData != nil {
                    
                    DispatchQueue.main.async {
                      self.wc.startWorkoutFromAcceptedInviteState()
                    }
                    
                  } else {
                    
                    DispatchQueue.main.async {
                        self.wc.initiateWorkoutStartHandshake()
                    }
                  }
                })
                .buttonStyle(.borderedProminent)
                .tint(.green)
            }
        }
    }
}

struct StartWorkoutConfirmation_Previews: PreviewProvider {
    static var previews: some View {
        StartWorkoutConfirmation()
        .environmentObject(WorkoutController())
    }
}
struct GreenButtonStyle: ButtonStyle {
    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
            .frame(maxWidth: .infinity)
            .padding()
            .padding(.vertical)
            .background(Color.green.saturation(1.4))
            .foregroundColor(.white)
            .cornerRadius(16)
    }
}

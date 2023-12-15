//
//  CompetitionChoiceList.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/3/23.
//

import SwiftUI

struct CompetitionChoiceList: View {
  @EnvironmentObject var workoutController: WorkoutController
  
  var body: some View {
    
    Group {
      
      if workoutController.competitions.isEmpty {
        VStack {
//          
//          Text(workoutController.errString)
//          Text(workoutController.notiTestStr)
          
          Button {
            workoutController.requestCompetitionList()
          } label: {
            Text("Refresh Competitions")
          }
          Spacer()
        }
      } else {
        VStack {
          //                Text(workoutController.workout_debug)
          Text("Choose competition")
            .font(.caption)
          ScrollView {
            
            NavigationLink {
              WorkoutChoiceList()
            } label: {
              Text("No Comp")
            }
            .simultaneousGesture(TapGesture().onEnded {
              DispatchQueue.main.async {
                self.workoutController.selectedCompetition = nil
              }
            })
            ForEach(workoutController.competitions, id: \.self) { comp_choice in
              
              NavigationLink {
                WorkoutChoiceList()
              } label: {
                Text(comp_choice["title"] ?? "")
              }
              .simultaneousGesture(TapGesture().onEnded {
                DispatchQueue.main.async {
                  self.workoutController.selectedCompetition = comp_choice
                }
              })
            }
          }
        }
      }
    }

    
    //        .navigationBarBackButtonHidden()
    .navigationTitle("Demoios")
    .onAppear {
      workoutController.requestCompetitionList()
    }
  }
}

struct CompetitionPickerList_Previews: PreviewProvider {
  static var previews: some View {
    CompetitionChoiceList()
      .environmentObject(WorkoutController())
      .setupPreview()
  }
}

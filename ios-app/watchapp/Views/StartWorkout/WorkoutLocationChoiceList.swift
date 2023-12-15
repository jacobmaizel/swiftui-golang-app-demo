//
//  WorkoutLocationChoiceList.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 5/3/23.
//

import SwiftUI
import HealthKit

struct WorkoutLocationChoiceList: View {
  @EnvironmentObject var wc: WorkoutController
  
  var body: some View {
    
    NavigationStack {
      VStack {
        Text("Location")
          .font(.caption)
        
        ScrollView {
          ForEach(LocationChoiceEnum.allCases) { location in
            
            NavigationLink {
              StartWorkoutConfirmation()
            } label: {
              Text(location.rawValue.capitalized)
            }
            .simultaneousGesture(TapGesture().onEnded {
              DispatchQueue.main.async {
                self.wc.selectedWorkoutLocationType = location.getHKSessionLocationType()
                if self.wc.selectedWorkoutLocationType == .outdoor {
                  wc.locManager.requestWhenInUseAuthorization()
                }
                
              }
            })
          }
        }
      }
    }
  }
}

struct WorkoutLocationChoiceList_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutLocationChoiceList()
  }
}

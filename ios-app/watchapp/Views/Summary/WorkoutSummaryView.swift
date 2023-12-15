//
//  WorkoutSummaryView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/2/23.
//

import SwiftUI
import Foundation
import WatchKit
import HealthKit

struct WorkoutSummaryView: View {
  @EnvironmentObject var wc: WorkoutController
  @State private var durationFormatter: DateComponentsFormatter = {
    let formatter = DateComponentsFormatter()
    formatter.allowedUnits = [.hour, .minute, .second]
    formatter.zeroFormattingBehavior = .pad
    return formatter
  }()
  var body: some View {
    
    NavigationStack {
      
      ScrollView {
        
        VStack {
          VStack {
            
            Text("Workout Completed")
              .font(.title3)
              .fontWeight(.bold)
            //
            //                      if wc.errString != "" {
            //                        Text("ERROR::: \(wc.errString)")
            //                      }
            
            if getCurrentEnvType() == .debug {
              ForEach(wc.errStrings, id: \.self) { str in
                Text(str)
              }
              Text("Location auth: \(wc.locationAuthorized ?? " no location auth")")
              Text("Noti auth: \(wc.notificationsAuthorized ?? "no noti auth")")
              Text("Healthkit auth: \(wc.healthStoreAuthorized ?? "no healthkit auth")")
              
            }
            
            HStack {
              if let comp = wc.selectedCompetition {
                VStack(alignment: .center) {
                  
                  Text("Competiton")
                    .font(.caption2)
                    .foregroundColor(.gray)
                  
                  Text(comp["title"] ?? "")
                    .font(.body)
                }
              }
              
              if let name = wc.workout?.workoutActivityType.name {
                VStack(alignment: .center) {
                  
                  Text("Workout Type")
                    .font(.caption2)
                    .foregroundColor(.gray)
                  
                  Text(name)
                    .font(.body)
                }
              }
            }
          } // HEADER VSTACK
          
          Divider()
          
          VStack {
            WorkoutSummaryStatItem(title: "Total Time", value: durationFormatter.string(from: wc.workout?.duration ?? 0.0) ?? "")
              .foregroundStyle(.yellow)
            
            WorkoutSummaryStatItem(title: "Avg. Heart Rate", value: wc.averageHeartRate.formatted(.number.precision(.fractionLength(0))) + " bpm")
              .foregroundStyle(.red)
            
            WorkoutSummaryStatItem(title: "Total Distance",
                                   value: Measurement(value: wc.workout?.totalDistance?.doubleValue(for: .meter()) ?? 0,
                                                      unit: UnitLength.meters)
                                    .formatted(
                                      .measurement(width: .abbreviated,
                                                   usage: .road,
                                                   numberFormatStyle:
                                          .number
                                        .precision(.fractionLength(2)))))
          }
          
          NavigationLink {
            HomeTabView()
          } label: {
            Text("Done")
          }
          .simultaneousGesture(TapGesture().onEnded {
            DispatchQueue.main.async {
              self.wc.resetWorkout()
            }
          })
          .buttonStyle(.borderedProminent)
          .padding(.horizontal)
          .tint(.blue)
          .navigationBarBackButtonHidden()
        } // TOP VSTACK
      }
    }
    .navigationBarBackButtonHidden()
  }
}

struct WorkoutSummaryView_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutSummaryView()
  }
}

//
//  WorkoutStatGraph.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/7/23.
//

import SwiftUI
import Charts

enum WorkoutGraphType {
  case heartrate, distance, pace
}

/*
 
 HeartRate
 - Average BPM, 3 min intervals, 
 
 */


struct WorkoutStatGraph: View {
  
  var type: WorkoutGraphType = .pace
  var workout: Workout = .defaultWorkout
  
  var graphData: [WorkoutGraphDataPoint] = []
  
  
  var height: CGFloat = 200
  
  
  
    var body: some View {
      
      VStack {
        
        if !graphData.isEmpty {

          Chart(graphData) {
//            ForEach(graphData) { item in
              
              
              LineMark(
                x: .value("Date", $0.date),
                y: .value("HR", $0.value)
              )

          }//        .padding()

           .chartPlotStyle { content in
               content.frame(height: 100)
           }
//
//          .chartYAxisLabel("Ft", position: .bottomTrailing, spacing: 8)
        } else {
          //
          Text("No Stats for this Workout! ")
        }
      }
    }
}

struct WorkoutStatGraph_Previews: PreviewProvider {
    static var previews: some View {
      WorkoutStatGraph(graphData: Workout.testGraphData)
        .setupPreview()
      WorkoutStatGraph(graphData: [])
        .setupPreview()
    }
}
//              .foregroundStyle(.blue)
//              .annotation(position: .bottom, alignment: .bottom, spacing: 4) {
//
//                Text(item.healthkit_workout_end_date.stringToShortDate ?? "")
//              }
//              .annotation(position: .automatic, alignment: .center, spacing: 5) {
//                Text("\(Int(item.value))")
//              }
              
              //                        PointMark(
              //                          x: .value("Date", item.healthkit_workout_end_date.intoDate ?? Date.now),
              //                            y: .value("Distance", item.distance)
              //                        )
//            }
//          .chartXAxis {
//            AxisMarks(values: .automatic) { value in
//                  if let date = value.as(Date.self) {
//                      let hour = Calendar.current.component(.hour, from: date)
//                      AxisValueLabel {
//                          VStack(alignment: .leading) {
//                              switch hour {
//                              case 0, 12:
//                                  Text(date, format: .dateTime.hour())
//                              default:
//                                  Text(date, format: .dateTime.hour(.defaultDigits(amPM: .omitted)))
//                              }
//                              if value.index == 0 || hour == 0 {
//                                  Text(date, format: .dateTime.month().day())
//                              }
//                          }
//                      }
//
//
//                      if hour == 0 {
//                          AxisGridLine(stroke: StrokeStyle(lineWidth: 0.5))
//                          AxisTick(stroke: StrokeStyle(lineWidth: 0.5))
//                      } else {
//                          AxisGridLine()
//                          AxisTick()
//                      }
//                  }
//              }
//          }
//          .frame(height: height)

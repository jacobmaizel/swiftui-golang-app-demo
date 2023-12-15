//
//  CompetitionStatGraph.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/8/23.
//

import SwiftUI
import Charts

struct CompetitionStatGraph: View {
  @EnvironmentObject var data: Datamodel
  
  var graphData: [CompetitionStatGraphDataPoint] = []
  
  @State var highlightedPoint: CompetitionStatGraphDataPoint?
  @State var highlightedLocation: CGPoint?
  @State var graphGeometryOrigin: CGPoint?
  
  @State var hoveringGraphAnnotationSizeValue: CGSize = .zero
  
  @State private var topContainerSize: CGSize = .zero
  
  struct HoveringGraphAnnotationSize: PreferenceKey {
    typealias Value = CGSize
    
    static var defaultValue: CGSize = .zero
    
    static func reduce(value: inout CGSize, nextValue: () -> CGSize) {
      value = nextValue()
    }
  }
  
  var body: some View {
    VStack {
      if !graphData.isEmpty {
        
        Chart {
          ForEach(graphData) { item in
            BarMark(
              x: .value("Date", item.healthkit_workout_end_date.intoDate!),
              y: .value("Distance", item.distance)
            )
            //            .foregroundStyle(highlightedPoint?.id == item.id ? Color.primarySeven : .graySix)
            .foregroundStyle(by: .value("Day", .now, unit: .day))
            //            .annotation(position: .top, alignment: .center) {
            //              if Int(item.distance) > 100 && item.id == highlightedPoint?.id {
            //
            //                Text("\(Int(item.distance)) ft")
            //                  .font(.caption2)
            //              }
            //            }
          }
        }
        .chartOverlay { proxy in
          GeometryReader { geometry in
            ZStack(alignment: .top) {
              
              Rectangle().fill(.clear).contentShape(Rectangle())
                .gesture(
                  DragGesture()
                    .onEnded { _ in
                      highlightedPoint = nil
                    }
                    .onChanged { value in
                      // Convert the gesture location to the coordinate space of the plot area.
                      let origin = geometry[proxy.plotAreaFrame].origin
                      let location = CGPoint(
                        x: value.location.x - origin.x,
                        y: value.location.y - origin.y
                      )
                      graphGeometryOrigin = origin
                      if   let (date, _) = proxy.value(at: location, as: (Date, Double).self) {
                        let firstGreater = graphData.lastIndex(where: {
                          
                          let targetMonth = Calendar.current.component(.month, from: $0.healthkit_workout_end_date.intoDate!)
                          let targetDay = Calendar.current.component(.day, from: $0.healthkit_workout_end_date.intoDate!)
                          let targetHour = Calendar.current.component(.hour, from: $0.healthkit_workout_end_date.intoDate!)
                          
                          
                          let dataMonth = Calendar.current.component(.month, from: date)
                          let dataDay = Calendar.current.component(.day, from: date)
                          
                          // return targetMonth == dataMonth && targetDay == dataDay && (dataHour - 3 <= targetHour && targetHour <= dataHour + 3)
                          return targetMonth == dataMonth && targetDay == dataDay
                          
                          
                        })
                        
                        if let index = firstGreater {
                          
                          
                          dispatchAsync {
                            self.highlightedPoint = graphData[index]
                            self.highlightedLocation = location
                          }
                          
                        }
                        
                      }
                    }
                )
            }
          }
        }
        .overlay(alignment: .topLeading) {
          if let loc = highlightedLocation, let point = highlightedPoint, let ori = graphGeometryOrigin {
            
            //      Image(systemName: "circle")
            //        .foregroundColor(.yellow)
            
            //                            GeometryReader { geo in
            SectionItem(fixedBy: .both, cornerRadius: .medium, color: .graySix) {
              
              
              //                              Color.clear
              //                                .preference(key: HoveringGraphAnnotationSize.self, value: geo.size)
              
              VStack {
                
                Text("\(point.healthkit_workout_end_date.stringToShortDate)")
                Text("\(Int(point.distance))")
              }
              
              
            }
            .offset(x: loc.x - hoveringGraphAnnotationSizeValue.width / 2 + ori.x, y: -hoveringGraphAnnotationSizeValue.height)
            .background(
              GeometryReader { proxy in
                Color.clear.preference(
                  key: HoveringGraphAnnotationSize.self,
                  value: proxy.size
                )
              }
            )
            .onPreferenceChange(HoveringGraphAnnotationSize.self) { value in
              print("Got new value\(value)")
              self.hoveringGraphAnnotationSizeValue = value
            }
          }
        }
        .chartYAxisLabel("Ft", position: .bottomTrailing, spacing: 8)
        .chartPlotStyle { content in
          content
            .frame(height: 150)
            .padding(.horizontal, 8)
        }
        .chartXAxis {
          AxisMarks(preset: .automatic, position: .automatic, values: .automatic) { value in
            if let date = value.as(Date.self) {
              let day = Calendar.current.component(.day, from: date)
              AxisValueLabel {
                VStack(alignment: .leading) {
                  Text("\(day)")
                  Text(date, format: .dateTime.month())
                }
              }
              AxisTick()
            }
          }
          
        }
        
      } else {
        Text("No Stats for this competition yet! Apply a workout to it.")
      }
    }
  }
}

struct CompetitionStatGraph_Previews: PreviewProvider {
  static var previews: some View {
    CompetitionStatGraph( graphData: CompetitionStatGraphDataPoint.testData)
      .environmentObject(Datamodel())
      .scaledToFit()
      .setupPreview()
  }
}

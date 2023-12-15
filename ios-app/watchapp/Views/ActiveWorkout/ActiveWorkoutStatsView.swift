//
//  ActiveWorkoutStatsView.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/4/23.
//

import SwiftUI

struct ActiveWorkoutStatsView: View {
    let date: Date = Date()
    @EnvironmentObject var wc: WorkoutController
    var body: some View {
            TimelineView(MetricsTimelineSchedule(from: wc.builder?.startDate ?? Date(),
                                                 isPaused: wc.session?.state == .paused)) { context in
                VStack(alignment: .leading) {
                    ActiveWorkoutElapsedTime(elapsedTime: wc.builder?.elapsedTime(at: context.date) ?? 0, showSubseconds: context.cadence == .live)
                        .foregroundStyle(.yellow)
                    Text(Measurement(value: wc.activeEnergy, unit: UnitEnergy.kilocalories)
                            .formatted(.measurement(width: .abbreviated, usage: .workout, numberFormatStyle: .number.precision(.fractionLength(0)))))
                    Text(wc.heartRate.formatted(.number.precision(.fractionLength(0))) + " bpm")
                    Text(Measurement(value: wc.distance, unit: UnitLength.meters).formatted(.measurement(width: .abbreviated, usage: .road)))
                }
                .font(.system(.title, design: .rounded).monospacedDigit().lowercaseSmallCaps())
                .frame(maxWidth: .infinity, alignment: .leading)
                .ignoresSafeArea(edges: .bottom)
                .scenePadding()
            }
//        NavigationStack {
//
//            ScrollView {
//
//                VStack(alignment: .leading) {
//
//                    Text(wc.workout_debug)
//
//                    ActiveWorkoutStatItem(title: "Duration", value: "\(wc.builder?.elapsedTime(at: date) ?? 0)")
//
//                    Divider()
//
//                    ActiveWorkoutStatItem(title: "Total Distance", value: "\(Measurement(value: wc.distance, unit: UnitLength.meters).formatted(.measurement(width: .abbreviated, usage: .road)))")
//
//                    Divider()
//
//                    ActiveWorkoutStatItem(title: "HR", value: "\(wc.heartRate)")
//
//                    Divider()
//
//                    ActiveWorkoutStatItem(title: "Calories", value: "\(wc.activeEnergy)")
//                }
//            }
//        }
//        .navigationTitle("Demoios")
//        .navigationBarBackButtonHidden()
    }
}

struct ActiveWorkoutStatsView_Previews: PreviewProvider {
    static var previews: some View {
        ActiveWorkoutStatsView()
    }
}

private struct MetricsTimelineSchedule: TimelineSchedule {
    var startDate: Date
    var isPaused: Bool
    
    init(from startDate: Date, isPaused: Bool) {
        self.startDate = startDate
        self.isPaused = isPaused
    }
    
    func entries(from startDate: Date, mode: TimelineScheduleMode) -> AnyIterator<Date> {
        var baseSchedule = PeriodicTimelineSchedule(from: self.startDate,
                                                    by: (mode == .lowFrequency ? 1.0 : 1.0 / 30.0))
            .entries(from: startDate, mode: mode)
        
        return AnyIterator<Date> {
            guard !isPaused else {
                return nil
            }
            return baseSchedule.next()
        }
    }
}

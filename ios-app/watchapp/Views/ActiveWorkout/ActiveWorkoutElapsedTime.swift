//
//  ActiveWorkoutElapsedTime.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/14/23.
//

import SwiftUI

struct ActiveWorkoutElapsedTime: View {
    var elapsedTime: TimeInterval = 0
    var showSubseconds: Bool = true
    @State private var timeFormatter = ElapsedTimeFormatter()

    var body: some View {
        Text(NSNumber(value: elapsedTime), formatter: timeFormatter)
            .fontWeight(.semibold)
            .onChange(of: showSubseconds) {
                timeFormatter.showSubseconds = $0
            }
    }
}

struct ActiveWorkoutElapsedTime_Previews: PreviewProvider {
    static var previews: some View {
        ActiveWorkoutElapsedTime()
    }
}

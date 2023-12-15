//
//  WorkoutStartView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/26/23.
//

import SwiftUI

struct WorkoutStartView: View {
    @EnvironmentObject var data: Datamodel
    
//    let store = HealthService.store
    
    @State var selectedComp: String = "Comp 1"
    var compsToChoose: [String] = ["Comp 1", "Comp 2", "Comp 3"]
    
    @State var selectedWorkoutType: String = "Running"
    var workoutTypesToChoose: [String] = ["Running", "Walking", "Cycling", "Swimming"]
    @State private var isAuthorized = false

    var body: some View {
        
        VStack {
            
//            Text("Start a Workout")
//                .DemoiosText(fontSize: .xl)
            
            Spacer()

            FormMenuBox(title: "Competition", choices: compsToChoose, selection: $selectedComp)
            FormMenuBox(title: "Workout Type", choices: workoutTypesToChoose, selection: $selectedWorkoutType)
            
            Spacer()
            Spacer()
            
            Button {
                print("hi")
            } label: {
                Text("Start")
                    .DemoiosText(fontSize: .large)
                    .frame(maxWidth: .infinity)
            }
            .DemoiosFullWidthButtonLabelStyling(cornerRadius: .medium, color: .green)
        }
        .navigationTitle("Start a Workout")
        .DemoiosBackgroundColor()
    }
}

struct WorkoutStartView_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutStartView()
    }
}

//
//  WorkoutGoalItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct WorkoutGoalItem: View {
    var title: String
    @State var progress: Float = 0.5
    
    var body: some View {
        
        SectionItem(fixedBy: .v) {
            
            VStack {
                HStack {
                    Text(title)
                        .foregroundColor(.graySix)
                        .lineLimit(1)
                        .DemoiosText(fontSize: .xxs)
                    Spacer()
                }
                
                Spacer()
                
                ProgressBar(progress: $progress)
                    .frame(height: 30)
                
                Spacer()
            }
            .padding(EdgeInsets(top: 0, leading: 8, bottom: 0, trailing: 8))
        }
    }
}

struct WorkoutGoalItem_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutGoalItem(title: "", progress: 10)
    }
}

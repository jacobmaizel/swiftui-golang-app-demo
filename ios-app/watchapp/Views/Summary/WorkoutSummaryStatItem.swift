//
//  WorkoutStatItem.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/12/23.
//

import SwiftUI

struct WorkoutSummaryStatItem: View {
    
    var title: String
    var value: String
    
    var body: some View {
        HStack {
            Text(title)
                .font(.caption2)
                .foregroundColor(.gray)
            Spacer()
            
            Text(value)
                .font(.body)
        }
    }
}

struct WorkoutSummaryStatItem_Previews: PreviewProvider {
    static var previews: some View {
        WorkoutSummaryStatItem(title: "", value: "")
    }
}

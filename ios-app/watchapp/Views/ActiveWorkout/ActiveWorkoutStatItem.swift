//
//  ActiveWorkoutStatItem.swift
//  Demoios Watch App
//
//  Created by Jacob Maizel on 4/12/23.
//

import SwiftUI

struct ActiveWorkoutStatItem: View {
    var title: String
    var value: String
    
    var body: some View {
        VStack(alignment: .leading) {
            
            Text(title)
                .font(.caption)
                .foregroundColor(.gray)
            
            Text(value)
                .font(.headline)
                .fontWeight(.bold)
        }
    }
}

struct ActiveWorkoutStatItem_Previews: PreviewProvider {
    static var previews: some View {
        ActiveWorkoutStatItem(title: "", value: "")
    }
}

//
//  ProgressBar.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/5/23.
//

import SwiftUI

struct ProgressBar: View {
    
    var backgroundColor: Color = .graySix
    var height: Double = 30
    @Binding var progress: Float
    
    
    
    var progressColor: Color {
        progress >= 1 ? .green : .primarySeven
    }
    
    var corners: UIRectCorner {
      progress >= 1 ? [.allCorners] : [.bottomLeft, .topLeft]
    }
    
    var body: some View {
        
        GeometryReader { geometry in
            ZStack(alignment: .leading) {
                
                RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                    .frame(width: geometry.size.width, height: geometry.size.height)
                //                            .opacity(0.3)
                    .foregroundColor(backgroundColor)
                
                RoundedCorner(radius: DemoiosCornerRadius.medium.rawValue, corners: corners)
                    .frame(width: min(CGFloat(self.progress) * geometry.size.width, geometry.size.width), height: geometry.size.height)
                    .foregroundColor(progressColor)
                    .animation(.linear, value: progress)
            }
        }
        .frame(height: height)
    }
}

struct ProgressBar_Previews: PreviewProvider {
    static var previews: some View {
        ProgressBar(progress: .constant(10))
    }
}

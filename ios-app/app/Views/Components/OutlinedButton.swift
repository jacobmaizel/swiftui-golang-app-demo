//
//  OutlinedButton.swift
//  Demoios
//
//  Created by Jacob Maizel on 4/24/23.
//

import SwiftUI

struct OutlinedButton<Content: View>: View {
    
    var fontSize: DemoiosFontSizes = .base
    var mainColor: Color = .primarySeven
    var background: Color = .grayTen
    var cornerRadius: DemoiosCornerRadius = .medium
    var action: () -> Void
    var content: () -> Content
    
    var body: some View {
        
        Button {
            action()
        } label: {
            content()
                .font(.system(size: fontSize.size))
                .foregroundColor(mainColor)
                .padding()
                .background(background)
                .cornerRadius(cornerRadius.rawValue)
                .overlay {
                    RoundedRectangle(cornerRadius: cornerRadius.rawValue)
                        .stroke(mainColor, lineWidth: 1)
                }
        }
    }
}

struct OutlinedButton_Previews: PreviewProvider {
    static var previews: some View {
        OutlinedButton(fontSize: .base) {
            print("hi")
        } content: {
            Text("hey")
        }
    }
}

//
//  Styles.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/15/23.
//

import Foundation
import SwiftUI


struct HomeEmptyCardButtonStyle: ButtonStyle {
    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
            .padding()
            .background(configuration.isPressed ? Color.primarySeven.opacity(1.1) : Color.primarySeven)
            .DemoiosText(fontSize: .large, weight: .bold)
            .cornerRadius(16)
            .shadow(radius: configuration.isPressed ? 5 : 0)
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
        
    }
}

struct MainCTAButtonStyle: ButtonStyle {
  
  var color: Color

    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
            .frame(maxWidth: .infinity)
            .padding()
            .background(configuration.isPressed ? color.opacity(1.1) : color)
            .DemoiosText(fontSize: .large, weight: .bold)
            .cornerRadius(8)
            .shadow(radius: configuration.isPressed ? 5 : 0)
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
        
    }
}

struct CustomizableResponsiveButton: ButtonStyle {
  
  var fontSize: DemoiosFontSizes = .base
  var fontWeight: Font.Weight = .regular
  var bgColor: Color = .primarySeven
  var fgColor: Color = .grayOne
  
    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
//            .frame(maxWidth: .infinity)
            .padding()
            .background(configuration.isPressed ? bgColor.opacity(0.9) : bgColor)
            .DemoiosText(fontSize: fontSize, weight: fontWeight)
            .cornerRadius(8)
            .shadow(radius: configuration.isPressed ? 5 : 0)
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
        
    }
}


struct ResponsiveButtonBase: ButtonStyle {

    func makeBody(configuration: Self.Configuration) -> some View {
        configuration.label
            .shadow(radius: configuration.isPressed ? 5 : 0)
            .scaleEffect(configuration.isPressed ? 0.95 : 1.0)
        
    }
}

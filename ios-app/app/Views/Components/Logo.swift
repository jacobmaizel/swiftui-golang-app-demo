//
//  Logo.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct Logo: View {
    
    enum LogoColor {
        case red
        case white
    }
    
    var size: CGFloat
    var logoColor: LogoColor
    
    var body: some View {
        
        Image(logoColor == .red ? "red-logo" : "white-logo")
            .resizable()
            .frame(width: size, height: size)
    }
}

struct Logo_Previews: PreviewProvider {
    static var previews: some View {
        Logo(size: 50, logoColor: .white)
    }
}

//
//  AboutDemoiosView.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/20/23.
//

import SwiftUI

struct AboutDemoiosView: View {
    var body: some View {
        
        VStack(alignment: .center) {
            Spacer()
            
            Text("Thanks for using Demoios!")
            
            Text("Created by [Jacob Maizel](https://www.linkedin.com/in/jacobmaizel/)")
            
            Spacer()
        }
        .frame(maxWidth: .infinity, maxHeight: .infinity)
        .DemoiosText(fontSize: .xxl)
        .DemoiosBackgroundColor()
    }
}

struct AboutDemoiosView_Previews: PreviewProvider {
    static var previews: some View {
        AboutDemoiosView()
        .setupPreview()
    }
}

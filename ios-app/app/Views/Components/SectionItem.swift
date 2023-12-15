//
//  Section.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct SectionItem<Content: View>: View {
    
    var fixedBy: FixedBy
    var cornerRadius: DemoiosCornerRadius
    var color: Color
  
    var content: Content
  
  init(fixedBy: FixedBy, cornerRadius: DemoiosCornerRadius = .medium, color: Color = .grayTen, @ViewBuilder _ content: () -> Content) {
    self.fixedBy = fixedBy
    self.cornerRadius = cornerRadius
    self.color = color
    self.content = content()
    
  }
    
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: cornerRadius.rawValue)
                .foregroundColor(color)
          
            content
          
        }
        .fixedSize(horizontal: fixedBy == .both || fixedBy == .h ? true : false, vertical: fixedBy == .both || fixedBy == .v ? true : false)
    }
}

struct Section_Previews: PreviewProvider {
    static var previews: some View {
        
        SectionItem(fixedBy: .both) {
            Text("hi")
            .padding()
        }
        .setupPreview()
        SectionItem(fixedBy: .h) {
            Text("hi")
            .padding()
        }
        .setupPreview()
        SectionItem(fixedBy: .v) {
            Text("hi")
            .padding()
        }
        .setupPreview()
      
    }
}

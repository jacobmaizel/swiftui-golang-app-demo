//
//  RedToppedBackground.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/13/23.
//

import SwiftUI

struct RedToppedBackground<Content: View>: View {
  
  let content: Content
  
  init(@ViewBuilder _ content: () -> Content) {
    self.content = content()
  }
  
  @State var offset: CGFloat = 0
  
  
  var body: some View {
    
    GeometryReader { geo in
      
      ZStack(alignment: .top) {
        
        Color.bg.frame(height: offset >= 0 ? geo.safeAreaInsets.top + geo.size.height * 0.6 : geo.safeAreaInsets.top, alignment: .top)
          .ignoresSafeArea()
        
        ZStack(alignment: .top) {
          
          ReadableScrollview(scrollOffset: $offset) {
            
            Rectangle()
              .foregroundColor(.bg)
              .frame(height: geo.size.height * 0.6)
            
            content
          }

        } // INNER Z

      } // TOP LEVEL Z
      
    }
    .DemoiosBackgroundColor()
  }
}

struct RedToppedBackground_Previews: PreviewProvider {
  static var previews: some View {
    RedToppedBackground {

      VStack(spacing: 0) {
          
          HomeTopHeader()
          
        HomeCompetitionCard(comp: .test_notstarted_joined)

          HStack(alignment: .center) {
            
          Text("Competition Activity")
              .DemoiosText(fontSize: .xl)
            Spacer()
            
            Button {
              print("refresh")
            } label: {
              Image(systemName: "arrow.clockwise")
              .DemoiosText(fontSize: .large, color: .primarySeven)
            }
          }
          .padding(.horizontal)
        }
    }
    .setupPreview()
  
  }
}

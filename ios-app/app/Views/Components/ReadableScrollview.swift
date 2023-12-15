//
//  ReadableScrollview.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/22/23.
//

import SwiftUI


struct ScrollViewOffsetPreferenceKey: PreferenceKey {

  static var defaultValue = CGFloat.zero
  
  static func reduce(value: inout CGFloat, nextValue: () -> CGFloat) {
    
//    print("got new value \(nextValue())")
      value = nextValue()
  }
}


struct ReadableScrollview<Content>: View where Content: View {
  @Namespace var scrollSpace
  
  @Binding var scrollOffset: CGFloat
  
  let content: () -> Content
  
  init(scrollOffset: Binding<CGFloat>,
       @ViewBuilder content: @escaping () -> Content) {
    _scrollOffset = scrollOffset
    self.content = content
    
    UIScrollView.appearance().alwaysBounceVertical = false
  }
  
  var body: some View {
    ScrollView {
      ZStack(alignment: .top) {
        
        content()

        GeometryReader { proxy in
          let offset = proxy.frame(in: .named("scroll")).minY
          Color.clear.preference(key: ScrollViewOffsetPreferenceKey.self, value: offset)
        }
      }
    }
    .coordinateSpace(name: "scroll")
    .onPreferenceChange(ScrollViewOffsetPreferenceKey.self) { value in
      
      dispatchAsync {
        scrollOffset = value
      }
      //              print(value)
      //              print(scrollOffset)
    }
  }
}
struct ReadableScrollview_Previews: PreviewProvider {
  static var previews: some View {
    @State var ugh: CGFloat = CGFloat.zero
    ReadableScrollview(scrollOffset: $ugh) {
      Text("\(ugh)")
        .background(ugh <= 0 && ugh > -200 ? Color.green : Color.blue)
    }
  }
}

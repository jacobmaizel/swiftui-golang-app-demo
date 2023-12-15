//
//  OverlayBottomToolbar.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct OverlayBottomToolbar<Content: View, Content2: View>: View {
  
  let viewContent: Content
  let menuContent: Content2
  
  init(@ViewBuilder viewContent: () -> Content, @ViewBuilder toolbarContent: () -> Content2) {
    self.viewContent = viewContent()
    self.menuContent = toolbarContent()
  }
  
  var body: some View {
    ZStack {
      viewContent
      VStack {
        Spacer()
        menuContent
      }
    }
  }
}

struct OverlayedMenu_Previews: PreviewProvider {
  static var previews: some View {
    OverlayBottomToolbar {
      Text("Main Content")
    } toolbarContent: {
      SectionItem(fixedBy: .v) {
        Text("Menu Content")
          .padding()
      }
      .padding()
    }
    .setupPreview()
  }
}

//
//  SectionTitle.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/27/23.
//

import SwiftUI

struct SectionTitle: View {
  var title: String
  var size: DemoiosFontSizes = .base
  var body: some View {
    HStack {
      Text(title)
        .DemoiosText(fontSize: size)
      Spacer()
    }
//    .padding(.horizontal)
//    .padding(.top)
  }
}

struct SectionTitle_Previews: PreviewProvider {
  static var previews: some View {
    SectionTitle(title: "hi")
  }
}

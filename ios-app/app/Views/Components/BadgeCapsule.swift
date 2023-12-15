//
//  BadgeCapsule.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/3/23.
//

import SwiftUI

struct BadgeCapsule: View {
    var title: String
    var color: Color
    var textSize: DemoiosFontSizes
    var body: some View {
        Text(title)
//            .foregroundColor(Color.graySix)
            .DemoiosText(fontSize: textSize, color: color )
            .padding(EdgeInsets(top: 4, leading: 8, bottom: 4, trailing: 8))
            .overlay {
                Capsule()
                    .stroke(color, lineWidth: 2)
                    .frame(minWidth: 75)
            }
    }
}

struct BadgeCapsule_Previews: PreviewProvider {
    static var previews: some View {
        BadgeCapsule(title: "Hero 402", color: Color.primarySeven, textSize: .small)
        .environment(\.colorScheme, .dark)
//        .previewLayout(.sizeThatFits)
    }
}

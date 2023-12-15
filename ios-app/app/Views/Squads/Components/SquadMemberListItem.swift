//
//  SquadMemberListItem.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/11/23.
//

import SwiftUI

struct SquadMemberListItem: View {
  var prof: PartialProfile
  var body: some View {
    SectionItem(fixedBy: .v, cornerRadius: .large) {
      HStack {
        CircleImage(url: prof.picture, size: 50)
        Text(prof.full_name)
        Spacer()
        Image(systemName: "chevron.right")
          .foregroundColor(.graySix)
      }
      .DemoiosText(fontSize: .base)
      .padding(.horizontal)
      .padding(.vertical, 6)
    }
  }
}

struct SquadMemberListItem_Previews: PreviewProvider {
  static var previews: some View {
    SquadMemberListItem(prof: .default)
      .setupPreview()
  }
}

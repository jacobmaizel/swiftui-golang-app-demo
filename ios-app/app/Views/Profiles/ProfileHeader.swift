//
//  ProfileHeaderSection.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/1/23.
//

import SwiftUI

struct ProfileHeader: View {
  var picture: String
  var first_name: String
  var last_name: String
  var created_at: String?
  
  @EnvironmentObject var data: Datamodel
  
  var body: some View {
    
    VStack(spacing: 0) {
    
      
      CircleImage(url: picture, size: 120)
      
      Text("\(first_name) \(last_name)")
        .DemoiosText(fontSize: .xxl)
      
//      Text("Joined \(created_at?.intoDate?.timeAgo() ?? " a billion years ago")")
//        .DemoiosText(fontSize: .small, color: .graySix)
    }
  }
}

struct ProfileHeaderSection_Previews: PreviewProvider {
  static var previews: some View {
    ProfileHeader(picture: "", first_name: "", last_name: "", created_at: "")
      .environmentObject(Datamodel())
      .setupPreview()
  }
}

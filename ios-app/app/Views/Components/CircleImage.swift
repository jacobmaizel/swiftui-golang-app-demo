//
//  CircleImage.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/3/23.
//

import SwiftUI

struct CircleImage: View {
  
  var url: String
  var size: CGFloat
  
  @State var timeOut: Bool = false
  @State var loadedImageSuccess: Bool = false
  
  var body: some View {
    AsyncImage(url: URL(string: url)) { phase in
      if let image = phase.image {
        
        image
          .resizable()
          .aspectRatio(contentMode: .fit)
          .onAppear {
            
//            DispatchQueue.main.async {
//              self.loadedImageSuccess = true
//            }
            
          }
        
      } else if phase.error != nil {
        Image(systemName: "person.fill")
      } else if timeOut {
        Image(systemName: "person.fill")
      } else {
        ProgressView()
      }
    }
    .frame(width: size, height: size)
    .clipShape(Circle())
    .onAppear {
      let timer = Timer.scheduledTimer(withTimeInterval: 5.0, repeats: false) { _ in
        DispatchQueue.main.async {
          self.timeOut = true
        }
        
      }
    }
  }
}

struct CircleImage_Previews: PreviewProvider {
  static var previews: some View {
    CircleImage(url: "", size: 80)
      .setupPreview()
    
    CircleImage(url: "https://lh3.googleusercontent.com/a/AGNmyxbgRZWxwYyzzpADX4VGuKMkdA0Xu2la2nSz5FxmDdQ=s96-c", size: 30)
      .setupPreview()
      
  }
}

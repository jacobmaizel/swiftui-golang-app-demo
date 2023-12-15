//
//  ExploreFilterControlBar.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/5/23.
//

import SwiftUI

struct ExploreFilterControlBar: View {
    
    @State private var selectedButton: Int = 1
    
    var body: some View {
        
        ZStack(alignment: .center) {
            
            RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                .fill(Color.grayTen)
//                .frame(maxWidth: .infinity, idealHeight: 30)
            
            HStack {
                
                Button {
                    self.selectedButton = 1
                } label: {
                    Text("All")
                        .foregroundColor(selectedButton == 1 ? .grayOne : .graySix)
                        .DemoiosText(fontSize: .base)
                        .padding()
                        .overlay {
                            selectedButton == 1 ? RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                                .stroke(Color.primarySeven, lineWidth: 1) : nil
                        }
                } // BUTTON END LABEL
                
                Spacer()
                
                Button {
                    self.selectedButton = 2
                } label: {
                    Text("Recommended")
                        .foregroundColor(selectedButton == 2 ? .grayOne : .graySix)
                        .DemoiosText(fontSize: .base)
                        .padding()
                        .overlay {
                            selectedButton == 2 ? RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                                .stroke(Color.primarySeven, lineWidth: 1) : nil
                        }
                } // BUTTON END LABEL
                
                Spacer()
                
                Button {
                    self.selectedButton = 3
                } label: {
                    Text("Outside Your Zone")
                        .foregroundColor(selectedButton == 3 ? .grayOne : .graySix)
                        .DemoiosText(fontSize: .base)
                        .padding()
                        .overlay {
                            selectedButton == 3 ? RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
                                .stroke(Color.primarySeven, lineWidth: 1) : nil
                        }
                } // BUTTON END LABEL
            }
        }
        .fixedSize(horizontal: true, vertical: true)
        .padding(.horizontal)
        .padding(.bottom)
    }
}

struct ExploreFilterControlBar_Previews: PreviewProvider {
    static var previews: some View {
        ExploreFilterControlBar()
    }
}

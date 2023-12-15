//
//  CompetitionCompetitorList.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/10/23.
//

import SwiftUI

struct CompetitionCompetitorList: View {
    
    var competition: Competition
    var body: some View {
        
        VStack {
            
            if competition.joined == true {
                HStack(alignment: .firstTextBaseline) {
                    
                    Text("Competitor Highlight")
                        .DemoiosText(fontSize: .xl)
                        .padding(.horizontal)
                    
                    Spacer()
                }
            }
            
            VStack(alignment: .leading) {
                
                if competition.joined == true {
                    
                    ForEach(1..<6, id: \.self) { _ in
                        
                        HStack {
                            
                            CircleImage(url: "https://asdf.com", size: 30)
                            
//                            BadgeCapsule(title: "Big Boi", color: .green, textSize: .xxs)
                            
                            Spacer()
                            
                            HStack(spacing: 12) {
                                Text("123,456")
                            }
                            .DemoiosText(fontSize: .xxs)
                            .padding(.trailing)
                        }
                        .padding(.horizontal)
                    }
                } else {
                    
                    VStack {
                        
                        Text("Join Now to see your competitors!")
                            .DemoiosText(fontSize: .large)
                        ForEach(1..<6, id: \.self) { _ in
                            
                            HStack {
                                Circle()
                                    .fill(Color.graySix)
                                
                                Capsule()
                                    .stroke( Color.graySix, lineWidth: 1)
                            }
                            .padding(.horizontal)
                        }
                    }
                    .frame(height: 200)
                }
            }
        }
        
        // end
    }
}

struct CompetitionCompetitorList_Previews: PreviewProvider {
    static var previews: some View {
        CompetitionCompetitorList(competition: Competition.test_notstarted_notjoined)
        .setupPreview()
    }
}

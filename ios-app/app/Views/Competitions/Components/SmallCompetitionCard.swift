//
//  SmallCompetitionCard.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/5/23.
//

import SwiftUI

struct SmallCompetitionCard: View {
    
    var competition: Competition
    
    var body: some View {
        ZStack {
            RoundedRectangle(cornerRadius: 16)
                .fill(Color.grayTen)
                .frame(maxWidth: .infinity, minHeight: 120, maxHeight: 120)
            
            VStack(alignment: .leading, spacing: 0) {
                HStack(alignment: .center, spacing: 0) {
                    Image(systemName: "calendar")
                        .font(.system(size: DemoiosFontSizes.large.size))
                    
                    Text("\(competition.start.stringToShortDate) - \(competition.end.stringToShortDate)")
                } // date hstack
                .foregroundColor(.grayEight)
                .DemoiosText(fontSize: .xxs)
                
                if competition.isActive() {
                    Text("Active")
                        .foregroundColor(.primarySeven)
                        .fontWeight(.bold)
                        .DemoiosText(fontSize: .xs)
//                        .padding()
                }
                
                Spacer()
                
                Text(competition.title)
                    .DemoiosText(fontSize: .base)
                    .padding(.bottom)
                
//                    HStack {
//                        ElementTag(color: .primarySeven, type: competition.category, size: .small)
//                        ElementTag(color: .blue, type: competition.level, size: .small)
//                        ElementTag(color: .green, type: competition.group_style, size: .small)
//                    }
            } // TOP VSTACK !!
            .padding(EdgeInsets(top: 8, leading: 8, bottom: 8, trailing: 8))
        } // TOP LEVEL ZSTACK
        .fixedSize(horizontal: true, vertical: true)
        
//        .frame(width:190, height:120)
            
    }
}

struct SmallCompetitionCard_Previews: PreviewProvider {
    static var previews: some View {
        SmallCompetitionCard(competition: Competition.test_notstarted_joined)
    }
}

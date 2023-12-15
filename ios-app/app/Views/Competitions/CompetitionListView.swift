//
//  CompetitionListView.swift
//  Demoios
//
//  Created by Jacob Maizel on 4/15/23.
//

import SwiftUI

struct CompetitionListView: View {
    
    @EnvironmentObject var data: Datamodel
    
    var body: some View {
      Text("Deprecated")
//        NavigationStack {
//
//            ScrollView {
//              ForEach(data.competitions) { comp in
//                    NavigationLink {
//                        CompetitionDetail()
//                    } label: {
//                      CompSearchResult(comp: comp)
//                    }
//                }
//            }
//            .padding()
//        }
    }
}

struct CompetitionListView_Previews: PreviewProvider {
    static var previews: some View {
      CompetitionListView()
        .environmentObject(Datamodel())
        .setupPreview()
    }
}

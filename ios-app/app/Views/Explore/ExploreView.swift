//
//  ExploreView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/4/23.
//

import SwiftUI
import SentrySwiftUI

struct ExploreView: View {
  
  
  var body: some View {
//    NavigationStack {
    Text("Deprecated")
//
//      VStack {
//
//        // MARK: - HEADER SECTION
//
//        HStack(alignment: .center) {
//
//          Image("white-logo")
//            .resizable()
//            .frame(width: 90, height: 90)
//
//          Spacer()
//
//          VStack(alignment: .center, spacing: 4) {
//
//            NavigationLink {
//              MyProfileView(profile: dataModel.profile)
//            } label: {
//              CircleImage(url: dataModel.profile.picture, size: 60)
//            }
//          }
//        } // HEADER HSTACK END
//        .padding()
//
//        ScrollView {
//
//          if !dataModel.joinedCompetitions.isEmpty {
//            VStack(alignment: .leading, spacing: 0) {
//              Text("My Competitions")
//                .DemoiosText(fontSize: .xl)
//                .padding(.horizontal)
//
//
//
//              ScrollView(.horizontal) {
//                HStack {
//
//                  ForEach(dataModel.competitions, id: \.id) { comp in
//                    if comp.joined && comp.isActive() {
//                      NavigationLink {
//                        //                                                CompetitionDetail(competition: $comp, isJoined)
//                        CompetitionDetail()
//                      } label: {
//                        LargeCompetitionCard(competition: comp)
//                      }
//                    }
//                  }
//                }
//                .padding(.horizontal)
//              }
//              //                            .sentryTrace("Explore View - Joined Comps List")
//            }
//          } else {
//            Text("No Joined Competitions!")
//          }
//
//          // MARK: - EXPLORE SECTION
//          VStack(alignment: .center) {
//
//            HStack {
//              Text("Explore")
//                .DemoiosText(fontSize: .xl)
//
//              Spacer()
//
//              Group {
//                NavigationLink {
//                  CompetitionCreateForm()
//                } label: {
//                  Image(systemName: "plus")
//                    .font(.system(size: DemoiosFontSizes.xl.size))
//                    .foregroundColor(.primarySeven)
//                    .padding()
//                    .background(Color.grayTen)
//                    .cornerRadius(8)
//                    .overlay {
//                      RoundedRectangle(cornerRadius: DemoiosCornerRadius.medium.rawValue)
//                        .stroke(Color.primarySeven, lineWidth: 1)
//                    }
//                } // LINK END
//              }
//            }
//            .padding(.horizontal)
//            .padding(.top)
//
//            ExploreFilterControlBar()
//
//            LazyVGrid(columns: columns, alignment: .center, spacing: 16) {
//
//              ForEach(dataModel.competitions, id: \.id) { comp in
//                if !comp.joined && comp.isActive() {
//                  NavigationLink {
//                    CompetitionDetail(competition: comp, isJoined: comp.joined)
//                  } label: {
//                    SmallCompetitionCard(competition: comp)
//                  }
//                }
//              }
//            }
//            .padding(EdgeInsets(top: 0, leading: 8, bottom: 16, trailing: 8))
//            //                        .sentryTrace("Explore View - Competition explore list")
//          } // EXPLORE VSTACK END
//
//        }// SCROLL VIEW END
//
//      } // TOP VSTACK
//      .background(Color.bg.ignoresSafeArea())
//      //            .sentryTrace("Explore View - Top Level")
////    }
//    .task {
//      if dataModel.competitions.isEmpty {
//        await dataModel.api.getMyCompetitions()
//      }
//    }
//    .refreshable {
//      await dataModel.api.getMyCompetitions()
//    }
  }
}

struct ExploreView_Previews: PreviewProvider {
  static var previews: some View {
    ExploreView()
   
      .setupPreview()
  }
}

//
//  CompEndedView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/26/23.
//

import SwiftUI

struct CompEndedView: View {
  
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) private var dismiss
  
  @State var comp: Competition
  
  @Binding var leaderboardData: [CompetitionLeaderboardItem]
  @Binding var graphData: [CompetitionStatGraphDataPoint]
  @Binding var profileParticipants: [PartialProfile]
  @Binding var myStats: CompMyStats?
  
  @State var isJoined: Bool
  
  // Data section toggle
  @State var dataChoice: CompDetailDataSelection = .leaderboard
  
  // sheet toggles
  @State var showingAddWorkoutSheet: Bool = false
  @State var showingWorkoutList: Bool = false
  @State var showingInviteMemberSheet: Bool = false
  @State private var showingConfirm: Bool = false
  
  // inviting
  @State var memberListToInvite: [PartialProfile] = []
  @State var inviteSearchString: String = ""
  
  //  @State private var showingConfirm: Bool = false
  @State var loaded: Bool = false
  var body: some View {
    //    NavigationStack {

      OverlayBottomToolbar {

        if isJoined == true {

          ScrollView {
            VStack {
              
              CompetitionDetailHeader(comp: comp)
              Divider()

              CompetitionInfoSection(graphData: graphData, myStats: myStats)

              CompDetailDataSelector(dataChoice: $dataChoice, members: profileParticipants)

              CompDetailDataSection(dataChoice: dataChoice, profileParticipants: profileParticipants, leaderboardData: leaderboardData)
                 .padding(.bottom, 75)

              Spacer()
            }
            .padding(.horizontal)
          }

          .sheet(isPresented: $showingWorkoutList) {
            NavigationStack {
              CompetitionWorkoutListView(forComp: comp.id)
            }
          }

          .DemoiosBackgroundColor()
          .navigationTitle(comp.title)
          .toolbar {

          }
          .sheet(isPresented: $showingInviteMemberSheet) {
            //            NavigationStack {
            MemberInviteSearch(resourceType: .profiles, comp: comp, searchText: $inviteSearchString, selectedResults: $memberListToInvite)
            //            }
          }

        } else {

          // MARK: - Competition NOT Joined

          ScrollView {

            VStack {
              CompetitionDetailHeader(comp: comp)

              //            Divider()

              CompDetailDataSelector(dataChoice: $dataChoice, members: profileParticipants)
              //                .padding(.top)

              CompDetailDataSection(dataChoice: dataChoice, profileParticipants: profileParticipants, leaderboardData: leaderboardData)
                .padding(.bottom, 75)
              //                    CompetitionInfoSection(competition: $competition)
              //            CompetitionCompetitorList(competition: competition)

              //              Spacer()

            }
            .padding()

          }
          //          .padding()
          .navigationTitle(comp.title)
          .DemoiosBackgroundColor()
        }
      } toolbarContent: {
        if isJoined {
          SectionItem(fixedBy: .both, cornerRadius: .medium, color: Color.graySeven) {
            HStack {
              ShareLink(item: createShareableLink(type: .competitions, id: comp.id), subject: Text("\(comp.title)"), message: Text("Check out \(comp.title) competition!")) {
                Label("Send Link", systemImage: "square.and.arrow.up")
                  .labelStyle(.iconOnly)
                  .padding()
                  .padding(.horizontal)
                  .contentShape(Rectangle())
              }
              Divider()
                .background(.black)

              //              Spacer()

              Button {

                //                WorkoutListView(forComp: competition.id)
                showingWorkoutList.toggle()


              } label: {
                Image(systemName: "list.bullet")
                  .padding(.horizontal)
                  .padding(.horizontal)
                  .padding(.vertical)

                  .contentShape(Rectangle())
              }
              .buttonStyle(ResponsiveButtonBase())


              //            Spacer()

              //              Divider()
              //                .background(.black)

              //            Spacer()
              //
              //              Button {
              //                showingAddWorkoutSheet.toggle()
              //              } label: {
              //                HStack {
              //                  Image(systemName: "plus")
              //                    .padding(.horizontal)
              //                    .padding(.leading)
              //                    .contentShape(Rectangle())
              //                }
              //              }
              //              .buttonStyle(ResponsiveButtonBase())
            }
            .DemoiosText(fontSize: .base)
            .padding(.horizontal)
          }
          .padding()
          //          .border(.red)

          //          VStack {
          //            Text("There is still time to compete!")
          //              .DemoiosText(fontSize: .large, weight: .bold)
          //            Button {
          //              data.api.joinCompetition(compId: competition.id.uuidString, dataModel: data)
          //              isJoined = true
          //            } label: {
          //              Text("Join")
          //                .DemoiosText(fontSize: .xl, weight: .bold)
          //            }
          //            .buttonStyle(MainCTAButtonStyle(color: .green))
          //          }
          //          .padding()
        }
      }
      //    }

  }
}

struct CompEndedView_Previews: PreviewProvider {
  static var previews: some View {
    
    let data = Datamodel()
    
    var personalstats: [CompetitionLeaderboardItem] = CompetitionLeaderboardItem.testData.map { item in
      var newItem = item
      newItem.id = Profile.default.id
      return newItem
    }
    var c1 = Competition.test_ended_joined

    CompEndedView(comp: .test_ended_joined, leaderboardData: .constant(CompetitionLeaderboardItem.testData), graphData: .constant(CompetitionStatGraphDataPoint.testData), profileParticipants: .constant([.default]), myStats: .constant(nil),  isJoined: true)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("not joined")
      .onAppear {
        
        data.profile.id = Profile.default.id
        c1.leaderboardData = personalstats
        
        data.joinedCompetitions = [Competition.test_notstarted_joined, c1]
        
      }
  }
}

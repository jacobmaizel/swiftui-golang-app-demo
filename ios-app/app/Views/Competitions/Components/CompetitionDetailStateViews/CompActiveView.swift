//
//  CompActiveView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/26/23.
//

import SwiftUI

struct CompActiveView: View {
  
  @EnvironmentObject var data: Datamodel
  
//  @Environment(\.dismiss) private var dismiss
  
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
            .refreshable {
                let (members, gd, cl, ms) = await data.api.refreshCompData(comp.id)
                profileParticipants = members
                graphData = gd
                leaderboardData = cl
                myStats = ms
            }
            .sheet(isPresented: $showingAddWorkoutSheet) {
              NavigationStack {
                AddWorkoutSheet(comp: $comp, leaderboardData: $leaderboardData, graphData: $graphData, profileParticipants: $profileParticipants, myStats: $myStats)
                  .presentationDragIndicator(.visible)
              }
            }
            .sheet(isPresented: $showingWorkoutList) {
              NavigationStack {
                CompetitionWorkoutListView(forComp: comp.id)
                  .presentationDragIndicator(.visible)
              }
            }
            .DemoiosBackgroundColor()
            .navigationTitle(comp.title)
            .toolbar {
              if isJoined {
                ToolbarItemGroup(placement: .primaryAction) {
                  Button("Leave") {
                    showingConfirm.toggle()
                  }
                  .confirmationDialog("Are you sure you want to leave?", isPresented: $showingConfirm) {
                    Button("Leave Competition", role: .destructive) {
                      Task {
                        do {
                          try await data.api.leaveCompetition(compId: comp.id.uuidString)
                          isJoined = false
                        } catch {
                          triggerLocalNotification(title: "Failed to leave the competition", body: "Please try again", background: .red)
//                            isJoined = true
                        }
                      }
                    }
                  } message: {
                    Text("Are you sure you want to leave this competition?")
                  }
                }
              }
            }
            .sheet(isPresented: $showingInviteMemberSheet) {
              NavigationStack {
                MemberInviteSearch(resourceType: .profiles, comp: comp, searchText: $inviteSearchString, selectedResults: $memberListToInvite)
                  .presentationDragIndicator(.visible)
              }

              }

  //          }

          } else {

            // MARK: - Competition NOT Joined

            VStack {
              CompetitionDetailHeader(comp: comp)
              //                    CompetitionInfoSection(competition: $competition)
              //            CompetitionCompetitorList(competition: competition)

              Spacer()

            }
            .padding()
            .navigationTitle(comp.title)
            .DemoiosBackgroundColor()
            .onAppear {
              let joined: Bool = data.joinedCompetitions.contains(where: { $0.id == comp.id })
              print(joined)
            }
          }
        } toolbarContent: {
          if isJoined {


            SectionItem(fixedBy: .both, cornerRadius: .medium, color: Color.graySix) {

              HStack {

                //            Spacer()

                Button {
                  showingInviteMemberSheet.toggle()
                } label: {
                  HStack {
                    Image(systemName: "envelope")
                      .padding(.horizontal)
                      .padding(.trailing)
                      .padding(.vertical)
                    //                  .padding(.horizontal)

                      .contentShape(Rectangle())

                    //                  .border(.red)
                  }

                }
                .buttonStyle(ResponsiveButtonBase())

                ShareLink(item: createShareableLink(type: .competitions, id: comp.id), subject: Text("\(comp.title)"), message: Text("Check out \(comp.title) competition!")) {
                  Label("Send Link", systemImage: "square.and.arrow.up")
                    .labelStyle(.iconOnly)
                    .padding()
                    .padding(.horizontal)
                    .contentShape(Rectangle())
                  //                  .padding(.horizontal)
                  //                  .padding(.trailing)
                  //                  .padding(.vertical)
                  //                  .padding(.leading)
                }


                //              Spacer()

  //              Divider()
  //                .background(.black)

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

                Button {
                  showingAddWorkoutSheet.toggle()
                } label: {
                  HStack {
                    Image(systemName: "plus")
                      .padding(.horizontal)
                      .padding(.leading)
                      .contentShape(Rectangle())
                  }
                }
                .buttonStyle(ResponsiveButtonBase())
              }
              .DemoiosText(fontSize: .base)
              .padding(.horizontal)
            }
            .padding()
          } else {
            VStack {
              Text("There is still time to compete!")
                .DemoiosText(fontSize: .large, weight: .bold)
              Button {
                Task {
                  do {
                    try await data.api.joinCompetition(compId: comp.id.uuidString)
                    isJoined = true
                  } catch {
                    triggerLocalNotification(title: "Failed to join the competition", body: "Please try again", background: .red)
                    isJoined = false
                  }
                }
              } label: {
                Text("Join")
                  .DemoiosText(fontSize: .xl, weight: .bold)
              }
              .buttonStyle(MainCTAButtonStyle(color: .green))
            }
            .padding()
          }
        }

  }
}

struct CompActiveView_Previews: PreviewProvider {
  static var previews: some View {
    
    let data = Datamodel()
    
    var personalstats: [CompetitionLeaderboardItem] = CompetitionLeaderboardItem.testData.map { item in
      var newItem = item
      newItem.id = Profile.default.id
      return newItem
    }
    var c1 = Competition.test_started_joined
    
   
    CompActiveView(comp: .test_ended_joined, leaderboardData: .constant(CompetitionLeaderboardItem.testData), graphData: .constant(CompetitionStatGraphDataPoint.testData), profileParticipants: .constant([.default]), myStats: .constant(nil),  isJoined: true)
      .environmentObject(data)
      .setupPreview()
      .onAppear {

        data.profile.id = Profile.default.id
        c1.leaderboardData = personalstats
        data.myCompletedWorkouts = [.defaultWorkout]

        data.joinedCompetitions = [Competition.test_notstarted_joined, c1]

      }
  }
}

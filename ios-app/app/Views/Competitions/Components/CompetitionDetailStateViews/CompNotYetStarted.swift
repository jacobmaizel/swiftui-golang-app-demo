//
//  CompNotYetStarted.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/26/23.
//

import SwiftUI


struct CompNotYetStarted: View {
  
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
            CompNotStartedProfileList(comp: comp, leaderboardData: leaderboardData, graphData: graphData, profileParticipants: profileParticipants)
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
        .sheet(isPresented: $showingWorkoutList) {
          NavigationStack {
            CompetitionWorkoutListView(forComp: comp.id)
          }
        }
        .padding(.bottom, 75)
        
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
          }
        }
        
      } else {
        
        // MARK: - Competition NOT Joined
        
        VStack {
          CompetitionDetailHeader(comp: comp)
          //                    CompetitionInfoSection(competition: $competition)
          //            CompetitionCompetitorList(competition: competition)
          
          Divider()
          
          
          Text("Current Competitors \(profileParticipants.count)")
            .DemoiosText(fontSize: .base, color: .graySix)
          
          
          Text("Join now to see the competitors!")
            .DemoiosText(fontSize: .base)
          
          
          Spacer()
          
        }
        .padding()
        .navigationTitle(comp.title)
        .DemoiosBackgroundColor()
      }
    } toolbarContent: {
      if isJoined {
        
        
        SectionItem(fixedBy: .both, cornerRadius: .large, color: Color.graySeven) {
          
          HStack {

            Button {
              showingInviteMemberSheet.toggle()
            } label: {
              HStack {
                Label("envelope", systemImage: "envelope")
                  .labelStyle(.iconOnly)
                  .padding()
                  .padding(.horizontal)
                  .contentShape(Rectangle())
              }
            }
            .buttonStyle(ResponsiveButtonBase())

            ShareLink(item: createShareableLink(type: .competitions, id: comp.id), subject: Text("\(comp.title)"), message: Text("Check out \(comp.title) competition!")) {
              Label("Send Link", systemImage: "square.and.arrow.up")
                .labelStyle(.iconOnly)
                .padding()
                .padding(.horizontal)
                .contentShape(Rectangle())
            }
          }
          .DemoiosText(fontSize: .base)
          .padding(.horizontal)
        }
        .padding()
        //          .border(.red)
      } else {
        
        VStack {
          //            Text("There is still time to compete!")
          //              .DemoiosText(fontSize: .large, weight: .bold)
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
    //    }
    
  }
}

struct CompNotYetStarted_Previews: PreviewProvider {
  static var previews: some View {
    
    let data = Datamodel()
    
    CompNotYetStarted(comp: .test_notstarted_joined, leaderboardData: .constant(CompetitionLeaderboardItem.testData), graphData: .constant(CompetitionStatGraphDataPoint.testData), profileParticipants: .constant([.default]), myStats: .constant(nil),  isJoined: true)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("joined")
    
    CompNotYetStarted(comp: .test_notstarted_notjoined, leaderboardData: .constant(CompetitionLeaderboardItem.testData), graphData: .constant(CompetitionStatGraphDataPoint.testData), profileParticipants: .constant([.default]), myStats: .constant(nil),  isJoined: false)
      .environmentObject(data)
      .setupPreview()
      .previewDisplayName("Not joined")
    
  }
}

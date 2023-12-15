//
//  SquadDetailView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/19/23.
//

import SwiftUI

struct SquadDetailView: View {
  @EnvironmentObject var data: Datamodel
  
  @State var squad: Squad
  
  @State var dataChoice: SquadDataChoices = .activity
  
  @State var showingInviteSheet: Bool = false
  @State var showingEditSquadSheet: Bool = false
  @State var inviteSearchString: String = ""
  @State var usersToInviteArray: [PartialProfile] = []
  @Environment(\.dismiss) var dismiss
  
  @State var showingConfirm: Bool = false
  
  @State var squadStats: SquadStats?
  
  @State var joined: Bool = false
  
  var body: some View {
    
    OverlayBottomToolbar {
      VStack {
        SquadDetailHeader(squad: squad)
          .padding(.horizontal)
        SquadDetailStats(squad: $squad, stats: squadStats)
        SquadDataSelector(dataChoice: $dataChoice)
        
        switch dataChoice {
        case .members:
          SquadDetailMemberList(squad: $squad)
        case .activity:
          Text("activity list")
          Spacer()
        }
      }
      .task {
        if squad.stats == nil {
          //          print("Squad stats was nil for \(squad.title)")
          
          let res: SquadStats? = await data.api.getSquadStats(squadId: squad.id)
          
          self.squad.stats = res
          self.squadStats = res
          //          triggerLocalNotification(title: "Failed to get squad stats", body: "Pull down to refresh", background: .red, icon: .warning)
        }
      }
      
      .DemoiosBackgroundColor()
      .navigationTitle(squad.title)
      .toolbar {
        if data.joinedSquads.contains(where: { $0.id == squad.id }) && data.profile.id != squad.owner.id {
          ToolbarItem(placement: .primaryAction) {
            Button("Leave", role: .destructive) {
              showingConfirm.toggle()
            }
            .confirmationDialog("Are you sure you want to leave?", isPresented: $showingConfirm, titleVisibility: .visible) {
              Button("Leave Squad", role: .destructive) {
                
                Task {
                  do {
                    try await data.api.leaveSquad(squadID: squad.id)
                    self.joined = false
                    await data.api.listMySquads()
                  } catch {
                    triggerLocalNotification(title: "Failed to leave squad", body: "Please try again", background: .red)
                  }
                }
                
                dismiss()
              }
            }
          }
        }
      }
      
    } toolbarContent: {
      
      if data.joinedSquads.contains(where: { $0.id == squad.id }) || joined {
        // HANDLE JOINED
        
        SectionItem(fixedBy: .both, cornerRadius: .large, color: .graySix) {
          
          HStack {
            
            Button {
              showingInviteSheet.toggle()
            } label: {
              Label("Invite", systemImage: "envelope")
                .labelStyle(.iconOnly)
              //                .padding(.leading)
              //                    .padding(.trailing)
                .padding()
              
                .contentShape(Rectangle())
              //                                  .border(.red)
              
            }
            .buttonStyle(ResponsiveButtonBase())
            
            ShareLink(item: createShareableLink(type: .squads, id: squad.id), subject: Text("Check out \(squad.title)"), message: Text("Join the squad now!")) {
              Label("Send Link", systemImage: "square.and.arrow.up")
                .labelStyle(.iconOnly)
                .padding(.leading)
              //                    .padding(.trailing)
                .padding()
              
                .contentShape(Rectangle())
              //                                  .border(.red)
              
            }
            
            if data.profile.id == squad.owner.id {
              Button {
                showingEditSquadSheet.toggle()
              } label: {
                
                Label("Edit", systemImage: "gearshape.fill")
                  .labelStyle(.iconOnly)
                //                    .padding(.trailing)
                  .padding()
                  .padding(.leading)
//                  .padding(.leading)
                
                  .contentShape(Rectangle())
                //                    .border(.red)
                
              }
              .buttonStyle(ResponsiveButtonBase())
              
            }
            
          } // main hstack
          .DemoiosText(fontSize: .base)
          .padding(.horizontal)
          //          .padding(.horizontal, 4)
          //          .padding(.vertical, 4)
          //          .border(.blue)
          
        } // SECTION ITEM
        .padding()
        
      } else {
        // HANDLE NOT JOINED
        
        Button {
          Task {
            do {
              try await data.api.joinSquad(squadID: squad.id)
              
              self.squad.joined = true
              self.joined = true
              
              let res: SquadStats? = await data.api.getSquadStats(squadId: squad.id)
              
              self.squad.stats = res
              self.squadStats = res
              
              let updatedSquad: Squad? = try await data.api.getSquadById(squadId: squad.id)
              
              if let updatedSquad = updatedSquad {
                self.squad = updatedSquad
              }
              
              await data.api.listMySquads()
            } catch {
              triggerLocalNotification(title: "Failed to Join Squad", body: error.localizedDescription, background: .red)
            }
            
          }
        } label: {
          //          Text("Accept Invite")
          Text("JOIN")
        }
        .buttonStyle(MainCTAButtonStyle(color: .green))
        .padding()
      }
      
    }
    .sheet(isPresented: $showingInviteSheet) {
      NavigationStack {
        MemberInviteSearch(resourceType: .profiles, squad: squad, searchText: $inviteSearchString, selectedResults: $usersToInviteArray)
          .onDisappear {
            dispatchAsync {
              self.inviteSearchString = ""
              self.usersToInviteArray = []
            }
          }
      }
    }
    .sheet(isPresented: $showingEditSquadSheet) {
      NavigationStack {
        EditSquadView(squad: squad)
          .onDisappear {
            if !data.joinedSquads.contains(where: { $0.id == squad.id }) {
              dismiss()
            }
          }
      }
    }
  }
}

struct SquadDetailView_Previews: PreviewProvider {
  static var previews: some View {
    Group {
      SquadDetailView(squad: Squad.owned_joined)
        .setupPreview()
        .previewDisplayName("Owned")
      SquadDetailView(squad: Squad.not_owned_joined, joined: true)
        .setupPreview()
        .previewDisplayName("Not owned joined")
      SquadDetailView(squad: Squad.not_owned_not_joined, joined: true)
        .setupPreview()
        .previewDisplayName("Not owned not joined")
    }
    .environmentObject(Datamodel())
  }
}

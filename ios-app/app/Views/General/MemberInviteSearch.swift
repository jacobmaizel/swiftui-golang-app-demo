//
//  CompetitionInviteSearch.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/12/23.
//

import SwiftUI

struct MemberInviteSearch: View {
  @State var searchResults: ProfileSearchRes = ProfileSearchRes(count: 0, results: [])
  @EnvironmentObject var data: Datamodel
  @Environment(\.dismiss) var dismiss
  
  var resourceType: SearchableResourceType
  
  var squad: Squad?
  var comp: Competition?
  var workout: Workout?
  
  @Binding var searchText: String
  //  @Binding var selectedResults: Set<PartialProfile>
  @Binding var selectedResults: [PartialProfile]
  
  var body: some View {
    VStack {
      ScrollView {
        ForEach(searchResults.results.filter({ $0.id != data.profile.id })) { prof in
          Button {
            if selectedResults.map({ $0.id }).contains(prof.id) {
              selectedResults.removeAll(where: { $0.id == prof.id })
            } else {
              selectedResults.append(prof)
            }
          } label: {
            ProfileSearchResult(prof: prof)
              .overlay {
                if selectedResults.map({ $0.id }).contains(prof.id) {
                  HStack {
                    Spacer()
                    Image(systemName: "circle.fill")
                  }
                  .DemoiosText(fontSize: .base, color: .orange)
                  .padding(.trailing)
                  .padding(.trailing)
                }
              }
          }
        }
      }
      
      
      if let comp = comp {
        Button("Done") {
          
          if getCurrentEnvType() != .preview {
            dismiss()
            
            Task {
              do {
                try await data.api.inviteMembersToComp(compId: comp.id, idsToInvite: selectedResults.map( {$0.id} ))
                triggerLocalNotification(title: "Sent Invites", body: "", background: .blue)
              } catch {
                triggerLocalNotification(title: "Failed to send invites", body: "Please try again", background: .red, icon: .warning)
              }
              
              
            }
            
          }
          
        }
        .buttonStyle(CustomizableResponsiveButton())
      } else if let squad = squad {
        Button("Send Invites") {
          
          if getCurrentEnvType() != .preview {
            dismiss()
            Task {
              do {
                let _: String = try await data.api.inviteProfilesToSquadByIds(squadId: squad.id, idsToInvite: selectedResults.map( {$0.id} ))
                
                triggerLocalNotification(title: "Sent Invites", body: "", background: .blue)
              } catch {
                triggerLocalNotification(title: "Failed to send invites", body: "Please try again", background: .red, icon: .warning)
              }
            }
            
            
            
          }
          
        }
        .buttonStyle(CustomizableResponsiveButton())
      } else {
        Button("Done") {
          
          if getCurrentEnvType() != .preview {
            dismiss()
          }
          
        }
        .buttonStyle(CustomizableResponsiveButton())
      }
    }
    .padding(.bottom)
    .searchable(text: $searchText, placement: .navigationBarDrawer(displayMode: .always), prompt: "Search")
    .autocorrectionDisabled(true)
    .onSubmit(of: .search) {
      Task {
        do {
          self.searchResults = try await data.api.search(for: resourceType, searchTerm: searchText)
        } catch {
          print(error)
        }
      }
    }
    .onChange(of: searchText) { texty in
      Task {
        do {
          self.searchResults = try await data.api.search(for: resourceType, searchTerm: texty)
        } catch {
          print(error)
        }
      }
    }
  }
}

struct CompetitionInviteSearch_Previews: PreviewProvider {
  static var previews: some View {
    NavigationStack {
      MemberInviteSearch(resourceType: .squads, squad: .not_owned_joined, searchText: .constant(""), selectedResults: .constant([]))
    }
    .setupPreview()
  }
}

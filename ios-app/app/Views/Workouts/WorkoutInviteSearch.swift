//
//  WorkoutInviteSearch.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/16/23.
//

import SwiftUI

struct WorkoutInviteSearch: View {
  @State var searchResults: ProfileSearchRes = ProfileSearchRes(count: 0, results: [])
  @EnvironmentObject var data: Datamodel
  
  @Binding var searchText: String
  //  @Binding var selectedResults: Set<PartialProfile>
  @Binding var selectedResults: [PartialProfile]
  
  var body: some View {
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
    .padding(.bottom)
    .searchable(text: $searchText, placement: .navigationBarDrawer(displayMode: .always), prompt: "Search")
    .autocorrectionDisabled(true)
    .onSubmit(of: .search) {
      print("Search")
    }
    .onChange(of: searchText) { texty in
      Task {
        do {
          searchResults = try await data.api.search(for: .profiles, searchTerm: texty)
        } catch {
          print(error)
        }
      }
 
    }
    
  }
}

struct WorkoutInviteSearch_Previews: PreviewProvider {
  static var previews: some View {
    WorkoutInviteSearch(searchText: .constant(""), selectedResults: .constant([PartialProfile.default]))
  }
}

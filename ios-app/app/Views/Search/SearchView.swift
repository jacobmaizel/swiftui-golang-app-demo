//
//  SearchView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/16/23.
//

import SwiftUI

struct SearchView: View {
  
  @EnvironmentObject var data: Datamodel
  
  @State private var searchText: String = ""
  @State var searchScope: SearchScope = .competitions
  
  @State private var profSearchResults: ProfileSearchRes = ProfileSearchRes(count: 0, results: [])
  @State private var squadSearchResults: SquadSearchRes = SquadSearchRes(count: 0, results: [])
  @State private var compSearchResults: CompSearchRes = CompSearchRes(count: 0, results: [])
  
  @State private var showEndedComps: Bool = false
  
  var body: some View {
    
    ScrollView {
      
      Picker("", selection: $searchScope) {
        ForEach(SearchScope.allCases) { scope in
          Text(scope.capped)
        }
      }
      .pickerStyle(.segmented)
      .padding(.horizontal)
      .padding(.bottom)
      .DemoiosBackgroundColor()
      
      if searchScope == .competitions {
        VStack {
          
          if !compSearchResults.results.isEmpty {
            Button(showEndedComps ? "Hide Ended Comps" : "Show Ended Comps") {
              showEndedComps.toggle()
            }
            .buttonStyle(.borderedProminent)
          }
          if showEndedComps {
            ForEach(compSearchResults.results.sorted(by: { $0.end > $1.end })) { comp in
              NavigationLink(value: comp) {
                CompSearchResult(comp: comp)
              }
            }
          } else {
            ForEach(compSearchResults.results.sorted(by: { $0.end > $1.end }).filter( { $0.state == .active } )) { comp in
              NavigationLink(value: comp) {
                CompSearchResult(comp: comp)
              }
            }
          }
        }
        .padding(.bottom)
        .navigationDestination(for: Competition.self) { comp in
          let joined: Bool = data.joinedCompetitions.contains(where: { $0.id == comp.id }) || comp.joined
          CompetitionDetail(comp: comp, isJoined: joined)
//          Text("\(joined.description)")
        }
      }
      if searchScope == .squads {
        VStack {
          ForEach(squadSearchResults.results) { squad in
            
            NavigationLink(value: squad) {
              SquadSearchResult(squad: squad)
            }
          }
        }
        .padding(.bottom)
        .navigationDestination(for: Squad.self) { squad in
          SquadDetailView(squad: squad)
        }
      }
      if searchScope == .people {
        VStack {
          ForEach(profSearchResults.results) { prof in
            NavigationLink(value: prof) {
              ProfileSearchResult(prof: prof)
            }
          }
        }
        .padding(.bottom)
        .navigationDestination(for: PartialProfile.self) { prof in
          NonSelfProfileView(partialProfile: prof)
        }
      }
    }
    .DemoiosBackgroundColor()
    .searchable(text: $searchText, placement: .navigationBarDrawer(displayMode: .always), prompt: "Search")
    .autocorrectionDisabled(true)
    .onSubmit(of: .search) {
      switch searchScope {
      case .competitions:
        Task {
          self.compSearchResults = try await data.api.search(for: .competitions, searchTerm: searchText)
        }
      case .squads:
        Task {
          squadSearchResults = try await data.api.search(for: .squads, searchTerm: searchText)
        }
      case .people:
        Task {
          profSearchResults = try await data.api.search(for: .profiles, searchTerm: searchText)
        }
        
      }
    }
    .onChange(of: searchText) { newVal in
      switch searchScope {
      case .competitions:
        Task {
          self.compSearchResults = try await data.api.search(for: .competitions, searchTerm: newVal)
        }
      case .squads:
        Task {
          squadSearchResults = try await data.api.search(for: .squads, searchTerm: newVal)
        }
      case .people:
        Task {
          profSearchResults = try await data.api.search(for: .profiles, searchTerm: newVal)
        }
      }
    }
  }
}


struct SearchView_Previews: PreviewProvider {
  static var previews: some View {
    
    Group {
      NavigationStack {
        SearchView()
      }
    }
    .setupPreview()
    
  }
}


enum SearchScope: String, ListableEnum {
  case competitions, squads, people
}

enum SearchResultType: String, ListableEnum {
  case competitions, squads, people
}


struct SearchResultItem: View {
  
  let resType: SearchResultType
  
  var body: some View {
    
    
    
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      
      HStack {
        
        Spacer()
        
        Image(systemName: "chevron.right")
          .foregroundColor(.graySix)
      }
      .padding()
    }
    .padding(.horizontal)
  }
}

struct CompSearchResult: View {
  
  var comp: Competition
  

  var body: some View {
    
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      
      HStack {
        VStack(alignment: .leading, spacing: 6) {
          
          
          HStack(spacing: 0) {
            Text(comp.title)
              .DemoiosText(fontSize: .base, weight: .bold)
            //              .padding(.bottom)
            
            Spacer()
//            if comp.state == .active {
//
//              Text("Active ")
//                .DemoiosText(fontSize: .base, color: .primarySeven, weight: .bold)
//
//              HStack(spacing: 0) {
//                Text("for ")
//                  .DemoiosText(fontSize: .xs, color: .primarySeven, weight: .regular)
//
//                Text(comp.end.intoDate!, style: .relative)
//                  .DemoiosText(fontSize: .xs, color: .primarySeven, weight: .regular)
//              }
//            } else if comp.state == .ended {
//
//              Text("Ended ")
//                .DemoiosText(fontSize: .base, color: .primarySeven, weight: .bold)
//
//              Text("\(comp.timeSinceEnd)")
//                .DemoiosText(fontSize: .xs, color: .primarySeven, weight: .regular)
//            }
          }
          
          Divider().padding(.bottom)
          
          ElementTagSet(comp: comp)
          
          DateHeader(from: comp.start.stringToShortDate, to: comp.end.stringToShortDate)
            .DemoiosText(fontSize: .xxs, color: .graySix)
            .lineLimit(1)
        }
        Spacer()
        
        Image(systemName: "chevron.right")
          .foregroundColor(.graySix)
      }
      .padding()
      
      
    }
    .padding(.horizontal)
    
    
  }
}

struct SquadSearchResult: View {
  var squad: Squad
  var body: some View {
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      HStack {
        HStack {
          
          HStack {
            if let members = squad.members, members.count > 0 {
              CircleImage(url: members[0].picture, size: 40)
            }
            
            if let members = squad.members, members.count > 1 {
              CircleImage(url: members[1].picture, size: 40)
                .offset(x: -20)
            }
            
            if let members = squad.members, members.count > 2 {
              CircleImage(url: members[2].picture, size: 40)
                .offset(x: -40)
            }
            
          }
          VStack(alignment: .leading) {
            Text(squad.title)
              .DemoiosText(fontSize: .base)
            //            Text("350 Members")
            //              .DemoiosText(fontSize: .small, color: .graySix)
          }
          //          .offset(x: -30)
        }
        Spacer()
        
        Image(systemName: "chevron.right")
          .foregroundColor(.graySix)
      }
      .padding()
    }
    .padding(.horizontal)
    
  }
}

struct ProfileSearchResult: View {
  
  var prof: PartialProfile
  
  var body: some View {
    
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      HStack {
        HStack {
          CircleImage(url: prof.picture, size: 40)
          VStack(alignment: .leading) {
            Text(prof.full_name)
              .DemoiosText(fontSize: .base)
            
          }
          //          .offset(x: -30)
        }
        Spacer()
        
        Image(systemName: "chevron.right")
          .foregroundColor(.graySix)
      }
      .padding()
    }
    .padding(.horizontal)
    //    SectionItem(fixedBy: .v, cornerRadius: .medium) {
    //
    //      HStack {
    //        HStack {
    //
    //          CircleImage(url: prof.picture, size: 40)
    //
    //          VStack(alignment: .leading) {
    //            Text(prof.full_name)
    //              .DemoiosText(fontSize: .base)
    //            //            Text(prof.created_at?.intoDate?.timeAgo() ?? "")
    //            //              .DemoiosText(fontSize: .small, color: .graySix)
    //          }
    //        }
    //        Spacer()
    //
    //        //        Image(systemName: "chevron.right")
    //        //          .foregroundColor(.graySix)
    //      }
    //      .padding()
    //    }
    //    .padding(.horizontal)
    
  }
}

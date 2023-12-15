//
//  SquadListView.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/11/23.
//

import SwiftUI
import SentrySwiftUI

struct SquadTabView: View {
  
  @EnvironmentObject var data: Datamodel
  @EnvironmentObject var path: PathManager
  
  @State var invitesListIsPresented: Bool = false
  
  var body: some View {
    
    NavigationStack(path: $path.squadPath) {
      OverlayBottomToolbar(viewContent: {
        
        ScrollView {
          if !data.joinedSquads.isEmpty {
            
            ForEach(data.joinedSquads) { squad in
              NavigationLink(value: squad) {
                SquadExpandedListItem(squad: squad)
              }
            }
          } else {
            VStack {
              HStack {
                Spacer()
                Text("Create a Squad or search for one to join!")
                  .DemoiosText(fontSize: .base, color: .graySix)
                Spacer()
              }
            }
          }
        }
        .refreshable {
          await data.api.listMySquads()
        }
        .navigationDestination(for: Squad.self) { squad in
          SquadDetailView(squad: squad)
        }
        
        // in here
        
      }, toolbarContent: {
        // toolbar stuff
        
        SectionItem(fixedBy: .both, cornerRadius: .large, color: .graySix) {
          
          HStack {
            
            NavigationLink(value: SquadPaths.create) {
              
              Label("Create", systemImage: "plus")
                .padding()
                .contentShape(Rectangle())
              //                .labelStyle(.iconOnly)
            }
            .buttonStyle(ResponsiveButtonBase())

          } // top level boolbar hstack

          .navigationDestination(for: SquadPaths.self) { val in
            
            if val == .create {
              SquadCreateForm()
            }
          }
          
        }
        .DemoiosText(fontSize: .base)
        .padding()
      })
      .navigationTitle("Squads")
      
      
      .DemoiosBackgroundColor()
      
      .sheet(isPresented: $invitesListIsPresented) {
        NavigationStack {
          SquadInvitesView()
        }
      }
    }
  }
}

struct SquadListView_Previews: PreviewProvider {
  static var previews: some View {
    
    var data = Datamodel()
    
    
    
    SquadTabView()
      .onAppear {
        
        data.joinedSquads.append(.not_owned_joined)
      }
      .environmentObject(data)
      .environmentObject(PathManager())
      .setupPreview()
  }
}

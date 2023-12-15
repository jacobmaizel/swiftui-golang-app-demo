//
//  EditSquadView.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/27/23.
//

import SwiftUI

struct EditSquadView: View {
  @EnvironmentObject var data: Datamodel
  @Environment(\.dismiss) var dismiss
  var squad: Squad
  
  
  @State var title: String = ""
  @State var publicSetting: Bool = false
  
  @State var showingDeleteConfirmation: Bool = false
  
    var body: some View {
      
      Form {
        Section("Squad Info") {
          TextField("Title", text: $title)
            .autocorrectionDisabled()
          
          Toggle(isOn: $publicSetting) {
            Text("Public")
          }
          
          Button("Update") {
            Task {
              let payload: [String: Any] = [
                "title": title,
                "public": publicSetting
              ]
              
              print("payload::", payload)
              
              do {
                let _: Squad = try await data.api.patchSquad(id: squad.id, payload: payload)
                
                triggerLocalNotification(title: "Updated Squad", body: "", background: .blue, icon: .ok)
                
                data.joinedSquads = []
                
                await data.api.listMySquads()
                
                dismiss()
              } catch {
                triggerLocalNotification(title: "Failed to update Squad", body: "Please try again", background: .red, icon: .warning)
                
              }
            }
            
            
          }
          .foregroundColor(.blue)
        }
        
        Section {
          Button("Delete", role: .destructive) {
//            print("delete")
            showingDeleteConfirmation.toggle()
            
          }
          .confirmationDialog("Are you sure you want to delete this squad? You cannot reverse this action.", isPresented: $showingDeleteConfirmation) {
            Button("Yes, I am sure.", role: .destructive) {
              Task {
              do {
                try await data.api.deleteSquad(with: squad.id)
                triggerLocalNotification(title: "Squad deleted", body: "", background: .blue)
                data.joinedSquads.removeAll(where: { $0.id == squad.id })
              } catch {
                triggerLocalNotification(title: "Failed to delete squad", body: "Try again", background: .red)
              }
              }
              dismiss()
              
            }
            
          } message: {
            Text("Are you sure you want to delete this squad? You cannot reverse this action.")
          }
          
        }
          
        
        
        
      }
      .navigationTitle("Edit Squad")
      .onAppear {
        self.title = squad.title
        self.publicSetting = squad.public
      }
      
      
        
    }
}

struct EditSquadView_Previews: PreviewProvider {
    static var previews: some View {
      NavigationStack {
        EditSquadView(squad: .owned_joined)
          .environmentObject(Datamodel())
      }
    }
}


func updateSquad(data: Datamodel, id: UUID, title: String, publicSetting: Bool) async throws -> Squad {
  
  let payload: [String: Any] = [
    "title": title,
    "public": publicSetting
  ]
  
  let res: Squad = try await data.api.patchSquad(id: id, payload: payload)
  
  return res
  
}

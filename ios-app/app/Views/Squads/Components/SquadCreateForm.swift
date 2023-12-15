//
//  SquadCreateForm.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/22/23.
//

import SwiftUI

struct SquadCreateForm: View {
    @State var title: String = ""
    
    @EnvironmentObject var data: Datamodel
    @Environment(\.dismiss) private var dismiss
    
  var body: some View {

    Form {
      TextField("Title", text: $title, prompt: Text("Title for your new Squad"))
        Button {
          Task {
            let payload: [String: Any] = ["title": title]
            
            do {
              try await data.api.createSquad(payload: payload)
              await data.api.listMySquads()
            } catch {
              triggerLocalNotification(title: "Failed to create squad", body: error.localizedDescription, background: .red)
            }
          }
          dismiss()
        } label: {
          Text("Create Squad")
        }
    }
    .DemoiosBackgroundColor()
    .scrollContentBackground(.hidden)
    .navigationTitle("Create a Squad")

    }
}

struct SquadCreateForm_Previews: PreviewProvider {
    static var previews: some View {
        SquadCreateForm()
        .setupPreview()
        .environmentObject(Datamodel())
    }
}

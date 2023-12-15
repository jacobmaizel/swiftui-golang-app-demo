//
//  ProfileView.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/13/23.
//

import SwiftUI

struct ProfileEditView: View {
  @Environment(\.dismiss) private var dismiss
  @EnvironmentObject var dataModel: Datamodel
  @State private var birthday: Date = Date.now
  @State private var firstName: String = ""
  @State private var lastName: String = ""
  @State private var publicSetting: Bool = false
  
  var body: some View {
    
    
    Form {
      
      TextField("First name", text: $firstName)
      TextField("Last name", text: $lastName)
      
    
      
      
      DatePicker("Birthday", selection: $birthday, in: ...Date.now, displayedComponents: .date)
      Toggle("Public", isOn: $publicSetting)
      Section {
        
        Button("Update Profile") {
          let payload: [String: Any] = ["first_name": firstName, "last_name": lastName, "birthday": birthday.intoString!, "public": publicSetting]
          
          Task {
          do {
            try await dataModel.api.updateProfile(dataModel, payload: payload)
            try await dataModel.api.getUserProfileAndConfig(dataModel)
            
            triggerLocalNotification(title: "Updated profile", body: "", background: .blue, icon: .ok)
          } catch {
            triggerLocalNotification(title: "Could not update profile", body: "Try again", background: .red, icon: .warning)
          }
            dismiss()
          }
        }
      }
    }
    .scrollContentBackground(.hidden)
    .DemoiosBackgroundColor()
    .onAppear {
      self.firstName = dataModel.editingProfile.first_name
      self.lastName = dataModel.editingProfile.last_name
      self.birthday = dataModel.editingProfile.birthday?.intoDate! ?? Date.now
      self.publicSetting = dataModel.editingProfile.public
    }
  }
}



struct ProfileEditView_Previews: PreviewProvider {
  static var previews: some View {
    ProfileEditView()
      .setupPreview()
      .environmentObject(Datamodel())
  }
}

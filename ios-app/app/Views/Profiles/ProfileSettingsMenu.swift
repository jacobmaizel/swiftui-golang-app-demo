//
//  ProfileSettingsMenu.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/17/23.
//

import SwiftUI

struct ProfileSettingsMenu: View {
  
  @EnvironmentObject var dataModel: Datamodel
  
  var body: some View {
    
    HStack {
      
      Spacer()
      
      
      NavigationLink(value: ProfilePaths.edit) {
        Text("Edit Profile")
          .padding(EdgeInsets(top: 4, leading: 4, bottom: 4, trailing: 4))
          .overlay {
            
            RoundedRectangle(cornerRadius: 8)
              .stroke(Color.primarySeven, lineWidth: 1)
          }
          .DemoiosText(fontSize: .base, color: .primarySeven)
      }
      .buttonStyle(ResponsiveButtonBase())
      
      Menu {
        
        if Configuration.environment.absoluteString == "DEBUG" {
          
          Button {
            dataModel.api.auth0Manager.credentials { result in
              switch result {
              case .success(let credentials):
                print("Obtained credentials: \(credentials.accessToken)")
                
                print("Obtained refresh token: \(credentials.refreshToken ?? "")")
              case .failure(let error):
                print("Failed with: \(error)")
              }
            }
          } label: {
            Text("Print Token")
          }
        }
        
        // MARK: - ABOUT
        NavigationLink(value: ProfilePaths.about) {
          Text("About Demoios")
        }
        
        NavigationLink(value: ProfilePaths.preferences) {
          Text("Preferences")
        }
        
        Button {
          dataModel.api.webAuthClient
            .clearSession { result in
              switch result {
              case .success:
                
                let cleared = dataModel.api.auth0Manager.clear()
                
                if cleared {
                  print("Auth session cleared")
                } else {
                  print("auth failed to clear keychain")
                }
                
                dataModel.clear()
                
              case .failure(let error):
                print("Failed with: \(error)")
              }
            }
        } label: {
          Text("Logout")
        }
      } label: {
        Image(systemName: "gearshape.fill")
          .padding()
          .foregroundColor(Color.primarySeven)
          .font(.system(size: DemoiosFontSizes.large.size + 4))
      }
    }
    .navigationDestination(for: ProfilePaths.self) { profpath in
      //      let path = profpath as ProfilePaths
      
      if profpath == .edit {
        ProfileEditView()
      }
      if profpath == .preferences {
        
        PreferencesView()
      }
      if profpath == .about {
        AboutDemoiosView()
      }
    }
  }
}


struct ProfileSettingsMenu_Previews: PreviewProvider {
  static var previews: some View {
    ProfileSettingsMenu()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}

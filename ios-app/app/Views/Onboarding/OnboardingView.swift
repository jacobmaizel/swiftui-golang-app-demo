//
//  OnboardingView.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/12/23.
//

import SwiftUI

struct OnboardingView: View {
  @EnvironmentObject var dataModel: Datamodel
  @State private var firstName: String = ""
  @State private var lastName: String = ""
  @State private var birthday: Date = Date.now
  
  var body: some View {
    NavigationStack {
      VStack {
        
        Text("Lets get to know you!")
          .DemoiosText(fontSize: .xxl)
          .padding()
        
        Form {
          
          Section("Name") {
            TextField("First Name", text: $firstName)
            TextField("Last Name", text: $lastName)
          }
          
          DatePicker(
            "Birthday",
            selection: $birthday,
            in: Date.now.adding(years: -100)...Date.now,
            displayedComponents: .date
          )
        }
        .scrollContentBackground(.hidden)
        
        NavigationLink {
          
          LocationSoftRequest()
          
        } label: {
          Text("Next")
        }
        .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
        .padding()
        .padding(.bottom)
        .padding(.bottom)
        .padding(.bottom)
        .simultaneousGesture(TapGesture().onEnded {
          print("going to location soft request")
          
          let payload: [String: Any] = [
            "first_name": self.firstName,
            "last_name": self.lastName,
            "birthday": self.birthday.intoString
          ]
          
          Task {
            do {
              try await dataModel.api.updateProfile(dataModel, payload: payload)
            } catch {
              triggerLocalNotification(title: "Failed to onboard profile", body: "Please try again", background: .red, icon: .warning)
            }
          }
    
        })
        
      }
      .DemoiosBackgroundColor()
      .onAppear {
        dispatchAsync {
        self.firstName = dataModel.profile.first_name
        self.lastName = dataModel.profile.last_name
        self.birthday = dataModel.profile.birthday?.intoDate ?? Date.now
        }
      }
      
    }
    
  }
}

struct OnboardingView_Previews: PreviewProvider {
  static var previews: some View {
    OnboardingView()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}



/*
 
 VStack(alignment: .leading) {
 
 
 Spacer()
 
 TextField("First name", text: $dataModel.profile.first_name, prompt: Text("First Name"))
 .DemoiosFormInput(title: "First Name")
 
 TextField("Last name", text: $dataModel.profile.last_name, prompt: Text("Last Name"))
 .DemoiosFormInput(title: "Last Name")
 
 DatePicker("", selection: $birthday, in: ...Date.now, displayedComponents: .date)
 .padding(.trailing, 10)
 .DemoiosFormInput(title: "Birthday")
 
 Spacer()
 Spacer()
 
 Button {
 let payload: [String: Any] = ["first_name": dataModel.profile.first_name, "last_name": dataModel.profile.last_name, "birthday": birthday.intoString!, "onboarding_completed": true]
 
 dataModel.api.updateProfile(dataModel, payload: payload) {
 }
 } label: {
 Text("Complete Onboarding")
 .DemoiosText(fontSize: .xl)
 .frame(maxWidth: .infinity)
 }
 .DemoiosFullWidthButtonLabelStyling(cornerRadius: .large, color: .primarySeven)
 } // end main VStack
 .DemoiosText(fontSize: .base)
 */

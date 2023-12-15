//
//  CompetitionCreateForm.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/17/23.
//

import SwiftUI

struct CompetitionCreateForm: View {
  
  @State var title: String = ""
  
  @State var start: Date = Date.now
  @State var end: Date = Date().adding(weeks: 1)
  
  @State var allowedMembers: CompetitionAllowedMembers = .all
  
  @State private var allowedActivityType = Set<CompetitionAcceptedActivityTypes>()
  @State private var privacySelection: CompetitionPrivacyOptions = .anyone
  @State private var scheduled: Bool = false
  
  @EnvironmentObject var dataModel: Datamodel
//  @Environment(\.presentationMode) var presentationMode
  @Environment(\.editMode) private var editMode
  @Environment(\.dismiss) private var dismiss
  
  @State var usersToInviteArray: [PartialProfile] = []
  @State var inviteSearchString: String = ""
  
  @State private var usersToInvite = Set<PartialProfile>() {
    didSet {
      print("added more users:::", usersToInvite)
    }
  }
    
  
  
  @State var showingInviteSheet: Bool = false {
    didSet {
    }
  }
  
  
  var body: some View {
    
    
//    NavigationStack {
      Form {
        
        Section("Competition Info") {
          TextField("Title", text: $title)
          
          Picker("Privacy", selection: $privacySelection) {
            ForEach(CompetitionPrivacyOptions.allCases) { opt in
              Text(opt.capped)
            }
          }
          
          Picker("Who Can Join", selection: $allowedMembers) {
            ForEach(CompetitionAllowedMembers.allCases) { opt in
              
              Text(opt.capped)
              
            }
          }
          
          
          Toggle("Scheduled", isOn: $scheduled)
          
          
          
          if scheduled {
            DatePicker("Start", selection: $start, displayedComponents: [.date, .hourAndMinute])
//              .datePickerStyle(.graphical)
          } else {
            LabeledContent("Start", value: "Now")
          }
          
          
          
          DatePicker("End", selection: $end, displayedComponents: [.date, .hourAndMinute])
//            .datePickerStyle(.graphical)
        }
        
        
        
        Section(header: Text("Workout Types")) {
          ForEach(CompetitionAcceptedActivityTypes.allCases, id: \.self) { option in
            Toggle(isOn: Binding(
              get: { self.allowedActivityType.contains(option) },
              set: { newValue in
                if newValue {
                  self.allowedActivityType.insert(option)
                } else {
                  self.allowedActivityType.remove(option)
                }
              }
            )) {
              Text(option.capped)
            }
          }
        }
        
        Section("Members to Invite") {
          List {
            ForEach(usersToInviteArray) { array in
              HStack {
                CircleImage(url: array.picture, size: 30)
                Text(array.full_name)
              }
            }
            
            
//            if usersToInviteArray.count > 0 {
//              HStack {
//                CircleImage(url: usersToInviteArray[0].picture, size: 30)
//                Text("\(usersToInviteArray[0].full_name)")
//              }
//            }
//            if usersToInviteArray.count > 1 {
//              HStack {
//                CircleImage(url: usersToInviteArray[1].picture, size: 30)
//                Text("\(usersToInviteArray[1].full_name)")
//              }
//            }
//            if usersToInviteArray.count > 2 {
//              HStack {
//                CircleImage(url: usersToInviteArray[2].picture, size: 30)
//                Text("\(usersToInviteArray[2].full_name)")
//              }
//            }
//
//            if usersToInviteArray.count > 3 {
//              let left = usersToInviteArray.count - 3
//              Text("\(left) other \(left == 1 ? "Invite" : "Invites") hidden")
//                .foregroundColor(.gray)
//            }
//
          }
          
          Button {
            showingInviteSheet.toggle()
          } label: {
            Text("\(usersToInviteArray.count > 0 ? "Invite Others" : "Invite")")
          }
        }
        
        
        Button("Create Competition") {
          let workoutTypes: [String] = allowedActivityType.map { $0.rawValue }
          
          let payload = [
            "title": title,
            "description": "Description",
            "start": start.intoString ?? "",
            "end": end.intoString ?? "",
            "scheduled": scheduled,
            "participant_types": allowedMembers == .all ? ["squad", "solo"] : allowedMembers == .solo ? ["solo"] : allowedMembers == .squad ? ["squad"] : [""],
            "workout_types": workoutTypes,
            "type": "community",
            "status": "pending",
            "public": privacySelection == .invite_only ? false : true
          ] as [String: Any]
          
          print("Creating competition \n\n", payload, "\n\n")
          
          Task {
            do {
              try await dataModel.api.createCompetition(payload: payload)
              dismiss()
            } catch {
              triggerLocalNotification(title: "Failed to create the competition", body: "Please try again", background: .red)
            }
          }
        }
      }
      .scrollContentBackground(.hidden)
      .DemoiosBackgroundColor()
      .onChange(of: scheduled) { newValue in
        if !newValue {
          start = Date.now
          print("resetting date")
        }
      }
      .onChange(of: showingInviteSheet) { newValue in
        
        if !newValue {
          
          inviteSearchString = ""
          
        }
      }
        
        
      
//    }
    .onAppear {
//        print("setting user invite array ")
//        usersToInviteArray = Array(usersToInvite)
    }
    .navigationTitle("Create Competition")
    .sheet(isPresented: $showingInviteSheet) {
      NavigationStack {
        
        MemberInviteSearch(resourceType: .profiles, searchText: $inviteSearchString, selectedResults: $usersToInviteArray)
      }
    }
  }
}

struct CompetitionCreateForm_Previews: PreviewProvider {
  static var previews: some View {
    CompetitionCreateForm()
      .environmentObject(Datamodel())
      .setupPreview()
  }
}


enum CompetitionAcceptedActivityTypes: String, ListableEnum {
  case running, walking, swimming, cycling
}


enum CompetitionPrivacyOptions: String, ListableEnum {
  case invite_only, anyone
}

enum CompetitionAllowedMembers: String, ListableEnum {
  case all, squad, solo
}

//
//  WorkoutsView.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/17/23.
//

import SwiftUI

struct WorkoutsTabView: View {
  @EnvironmentObject var data: Datamodel
  
  @Environment(\.dismiss) var dismiss
  
  @State var acceptedWorkouts: [Workout] = []
  
  @State var loadedAcceptedWorkouts: Bool = false
  @State var selectedAcceptedWorkout: Workout?
  @State var showingAcceptedWorkoutStartSheet: Bool = false
  
  
  @State var leadPendingWorkouts: [Workout] = []
  
  @State var memberInviteStatuses: [PartialProfile] = []
  
  @State var loadedLeadPendingWorkouts: Bool = false
  @State var selectedLeadPendingWorkout: Workout?
  @State var showingLeadPendingWorkoutSheet: Bool = false
  
  @State var showingImportWorkoutSheet: Bool = false
  @State var showingCreateWorkoutSheet: Bool = false
  
//  @State var invitesByPendingWorkout: [UUID: [PartialProfile]] = [:]
  
  @State var err: APIError?
  @EnvironmentObject var path: PathManager
  
  var body: some View {
    NavigationStack(path: $path.workoutPath) {
//      OverlayBottomToolbar {
        ScrollView {
          
          SectionTitle(title: "Your Workouts")
            .DemoiosText(fontSize: .large, weight: .bold)
            .padding(.leading)
          
          if err != nil && loadedLeadPendingWorkouts {
            Text("Failed to load pending workouts")
          } else if loadedLeadPendingWorkouts && leadPendingWorkouts.isEmpty {
            Text("Create and invite users to a workout!")
              .foregroundColor(.graySix)
          } else if loadedLeadPendingWorkouts && !leadPendingWorkouts.isEmpty {
            
            ForEach(leadPendingWorkouts.sorted(by: >)) { workout in
              Button {
                dispatchAsync {
                  self.selectedLeadPendingWorkout = workout
                  self.showingLeadPendingWorkoutSheet.toggle()
                }
              } label: {
                WorkoutTabListItem(workout: workout)
                //                  .padding(.top, 8)
                  .padding(.bottom, 4)
              }
            }
          } else {
            VStack {
              ProgressView()
            }
            
          }
          
          SectionTitle(title: "Accepted Workout Invites")
            .DemoiosText(fontSize: .large, weight: .bold)
            .padding(.leading)
          
          // when you accepted an invite to a workout, but have not done the workout for it yet
          
          if err != nil && loadedAcceptedWorkouts {
            Text("Failed to load pending workouts")
          } else if loadedAcceptedWorkouts && acceptedWorkouts.isEmpty {
            Text("You will see your pending workouts here after you accept an invite!")
              .foregroundColor(.graySix)
          } else if loadedAcceptedWorkouts && !acceptedWorkouts.isEmpty {
            
            ForEach(acceptedWorkouts.sorted(by: >)) { workout in
              Button {
                dispatchAsync {
                  self.selectedAcceptedWorkout = workout
                  self.showingAcceptedWorkoutStartSheet.toggle()
                }
              } label: {
                //                AcceptedWorkoutItem(workout: workout)
                WorkoutTabListItem(workout: workout)
                  .padding(.bottom, 4)
              }
            }
          } else {
            VStack {
              ProgressView()
            }
          }
          Rectangle()
            .frame(height: 100)
            .foregroundColor(.clear)
        }
        .DemoiosBackgroundColor()
        .toolbar {
          ToolbarItemGroup {
            Button {
              showingImportWorkoutSheet.toggle()
            } label: {
              Image(systemName: "arrow.down.square")
              
            }
              Button {
                showingCreateWorkoutSheet.toggle()
              } label: {
                Image(systemName: "plus.square")
              }
          }
        }
        .refreshable {
          let pendingWorkoutRes: [Workout] = await data.api.getMyLeadIncompleteWorkouts()
          
          leadPendingWorkouts = pendingWorkoutRes
          loadedLeadPendingWorkouts = true
          
          print("got pending workouts \(leadPendingWorkouts.count)")
          
          do {
            acceptedWorkouts = try await data.api.getMyAcceptedInviteWorkouts()
            loadedAcceptedWorkouts = true
          } catch {
//            triggerLocalNotification(title: "Failed to load your workouts", body: "Please try again.", background: .red)
            print(error)
          }
          
        }
        
//      } toolbarContent: {
//        SectionItem(fixedBy: .both, cornerRadius: .large, color: .graySix) {
//          HStack {
//            
//            Button {
//              showingImportWorkoutSheet.toggle()
//            } label: {
//              Label("Import", systemImage: "square.and.arrow.down")
//                .DemoiosText(fontSize: .base)
//            }
//            .padding(.trailing)
//            NavigationLink(value: WorkoutPaths.create) {
//              Label("Create", systemImage: "plus")
//                .DemoiosText(fontSize: .base)
//            }
//            .padding(.leading)
//            //          .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
//          }
//          .padding()
//        }
//        .navigationDestination(for: WorkoutPaths.self) { workoutPath in
//          WorkoutConfigurationView()
//        }
//        .padding()
//      }
//      .navigationTitle("Workouts")
//      .DemoiosBackgroundColor()
    }
    .sheet(isPresented: $showingImportWorkoutSheet, onDismiss: {
      //      print("Dismissed Import sheet")
    }, content: {
      NavigationStack {
        WorkoutImportView()
      }
    })
    .sheet(isPresented: $showingCreateWorkoutSheet, onDismiss: {
      //      print("Dismissed Import sheet")
    }, content: {
      NavigationStack {
        WorkoutConfigurationView()
      }
    })
    .sheet(isPresented: $showingAcceptedWorkoutStartSheet) {
      NavigationStack {
        WorkoutTabStartSheet(workout: $selectedAcceptedWorkout, sheetType: .accepted, acceptedWorkouts: $acceptedWorkouts, pendingWorkouts: $leadPendingWorkouts)
        
      }
      .presentationDetents([.medium, .large])
      
    }
    .sheet(isPresented: $showingLeadPendingWorkoutSheet) {
      NavigationStack {
        WorkoutTabStartSheet(workout: $selectedLeadPendingWorkout, sheetType: .pending, acceptedWorkouts: $acceptedWorkouts, pendingWorkouts: $leadPendingWorkouts)
          .navigationBarTitleDisplayMode(.inline)
        
      }
      .presentationDetents([.medium, .large])
          .navigationBarTitleDisplayMode(.inline)
    }
    .onChange(of: showingAcceptedWorkoutStartSheet) { newVal in
      if !newVal {
        self.selectedAcceptedWorkout = nil
      }
    }.onChange(of: showingLeadPendingWorkoutSheet) { newVal in
      if !newVal {
        self.selectedLeadPendingWorkout = nil
      }
    }
    .task {
      if !loadedLeadPendingWorkouts {
        //        do {
        let pendingWorkoutRes: [Workout] = await data.api.getMyLeadIncompleteWorkouts()
        
        leadPendingWorkouts = pendingWorkoutRes
        loadedLeadPendingWorkouts = true

      }
      if !loadedAcceptedWorkouts {
        do {
          acceptedWorkouts = try await data.api.getMyAcceptedInviteWorkouts()
          loadedAcceptedWorkouts = true
        } catch {
          triggerLocalNotification(title: "Failed to load your workouts", body: "Please try again.", background: .red)
        }
      }
    }
  }
}

struct WorkoutsView_Previews: PreviewProvider {
  static var previews: some View {
    let data = Datamodel()
    
    WorkoutsTabView(acceptedWorkouts: [.defaultWorkout, .defaultWorkout2], loadedAcceptedWorkouts: true, leadPendingWorkouts: [.defaultWorkout, .defaultWorkout2], loadedLeadPendingWorkouts: true)
    
      .environmentObject(data)
      .environmentObject(PathManager())
      .environment(\.colorScheme, .dark)
      .onAppear {
        
      }
      .setupPreview()
  }
}



struct AcceptedWorkoutItem: View {
  
  var workout: Workout
  @EnvironmentObject var data: Datamodel
  var body: some View {
    
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      
      HStack {
        
        VStack(spacing: 4) {
          HStack {
            
            Text(workout.healthkit_workout_activity_type.capitalized)
              .DemoiosText(fontSize: .large, weight: .bold)
            
            Spacer()
          }
          
          HStack {
            
            HStack {
              CircleImage(url: workout.leader.picture, size: 40)
              HStack {
                
                
                Text("Lead by")
                  .DemoiosText(fontSize: .small, color: .graySix)
                
                
                Text("\(workout.leader.id == data.profile.id ? "You" : workout.leader.full_name)")
                  .multilineTextAlignment(.leading)
                  .lineLimit(1)
                  .DemoiosText(fontSize: .base)
              }
            }
            Spacer()
            
            if let created = workout.created_at.intoDate?.timeAgo() {
              VStack {
                Text(created)
                  .DemoiosText(fontSize: .base)
                Text("Created")
                  .DemoiosText(fontSize: .small, color: .graySix)
              }
              
            }
          }
        }
        Spacer()
        
        Image(systemName: "circle.fill")
          .foregroundColor(.orange)
        
      }
      .padding()
      
    }
    .padding(.horizontal)
    
    
  }
}


struct WorkoutTabListItem: View {
  @EnvironmentObject var data: Datamodel
  var workout: Workout
//  var invitedUsers: [PartialProfile]
  @State private var dragOffset: CGSize = CGSize.zero
  
  var body: some View {
    
    
    SectionItem(fixedBy: .v, cornerRadius: .large) {
      VStack {
        HStack {
          
          VStack(spacing: 4) {
            HStack {
              Text(workout.workout_data?.location_type.capitalized ?? "")
                .DemoiosText(fontSize: .xl, weight: .bold)
              Text(workout.healthkit_workout_activity_type.capitalized)
                .DemoiosText(fontSize: .xl, weight: .bold)
              
              Spacer()
            }
            HStack(spacing: 2) {
              
              VStack(alignment: .leading) {
                HStack(alignment: .top) {
                  CircleImage(url: workout.leader.picture, size: 40)
                  VStack(alignment: .leading) {
                    
                    Text("\(workout.leader.id == data.profile.id ? "You" : workout.leader.full_name)")
                      .multilineTextAlignment(.leading)
                      .lineLimit(1)
                      .DemoiosText(fontSize: .base)
                    
                    Text("Leader")
                      .DemoiosText(fontSize: .small, color: .graySix)
                  }
                }
              }
              Spacer()
            }
            
            Spacer()
            
          }
          Spacer()
          
          
          if let created = workout.created_at.intoDate?.timeAgo() {
            VStack(alignment: .leading) {
              Spacer()
              Text(created)
                .DemoiosText(fontSize: .small)
              Text("Created")
                .DemoiosText(fontSize: .xs, color: .graySix)
            }
            .padding(.bottom)
          } // created
        }
        .padding(.horizontal)
        .padding(.top, 4)
      }
    }// main section
    .padding(.horizontal)
  }
} // lead pending view top level

func getInviteStatusIcon(status: String) -> some View {
  
  switch status {
  case "pending":
    return Image(systemName: "circle.fill")
      .foregroundColor(.orange)
    
  case "accepted":
    return Image(systemName: "checkmark")
      .foregroundColor(.green)
    
  case "declined":
    return Image(systemName: "xmark")
      .foregroundColor(.red)
    
  default:
    return Image(systemName: "xmark")
      .foregroundColor(.blue)
    
    
    
  }
  
}




// MARK: - Drag archive
/*
 .offset(CGSize(width: min(dragOffset.width, 30), height: 0))
 .gesture(
 DragGesture()
 .onEnded { _ in
 if dragOffset.width < -30 {
 // if the drag ended while it was left a certain amout,
 // sprint back to the default "open" state.
 withAnimation(.spring()) {
 dragOffset = CGSize(width: -50, height: 0)
 }
 } else {
 withAnimation(.spring()) {
 dragOffset = .zero
 }
 }
 }
 .onChanged {
 switch($0.translation.width, $0.translation.height) {
 case (-100...0, -100...100):
 //                  if $0.translation.width <= 10 {
 //                  print("LEFT -> width offset\($0.translation.width)")
 dragOffset = $0.translation
 //                  }
 //                  print("left")
 case (0...100, -100...100):
 // right swipe, should not trigger if
 //                  print("RIGHT -> width offset \($0.translation.width)")
 if dragOffset.width < 0 {
 dragOffset = $0.translation
 }
 
 default:
 //                  print("not left")
 break
 }
 }
 //              .onEnded { _ in dragOffset = .zero }
 )
 .animation(.spring(), value: dragOffset)
 
 //scratch below
 //      .gesture(DragGesture(minimumDistance: 3.0, coordinateSpace: .local)
 //        .onEnded { value in
 //          DispatchQueue.main.async {
 //            dragOffset = value.translation
 //          }
 //          print(value.translation)
 //          switch(value.translation.width, value.translation.height) {
 //          case (...0, -30...30):  print("left swipe")
 //          case (0..., -30...30):  print("right swipe")
 //          case (-100...100, ...0):  print("up swipe")
 //          case (-100...100, 0...):  print("down swipe")
 //          default:  print("no clue")
 //              }
 //          }
 //      ) // gesture
 */

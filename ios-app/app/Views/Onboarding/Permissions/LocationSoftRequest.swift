//
//  LocationSoftRequest.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/24/23.
//

import SwiftUI

struct LocationSoftRequest: View {
  
  @EnvironmentObject var loc: LocationManager
  
  @State var gotLocPerm: Bool = false
    var body: some View {
      NavigationStack {
        VStack {
          
          Spacer()
          HStack {
            
            Spacer()
            Text("Current loc status \(loc.status?.rawValue ?? 123)")
            Image(systemName: "location.north.circle.fill")
              .resizable()
              .frame(width: 100, height: 100)
            
            Image(systemName: "plus")
              .resizable()
              .frame(width: 50, height: 50)
              .padding(.horizontal)
            
            Logo(size: 125, logoColor: .white)
            
            Spacer()
            
          }
          
          Spacer()
          
          Text("Demoios will use your location to add route information to your outdoor workouts and supply weather information.")
            .DemoiosText(fontSize: .large, weight: .bold)
            .multilineTextAlignment(.center)
          
          Spacer()
          Spacer()
          
          
          VStack(spacing: 0) {
            Button {
              
              loc.askForWhenInUsePermission()
              
              if let st = loc.status?.rawValue {
                if st > 2 {
                  gotLocPerm = true
                }
              }
              
            } label: {
              Text("Enable Location")
              
            }
            .buttonStyle(MainCTAButtonStyle(color: .primarySeven))
            .padding(.horizontal)
//            .simultaneousGesture(TapGesture().onEnded {
//              print("Accepted loc!")
//              //          DispatchQueue.main.async {
//              //              self.workoutController.selectedCompetition = comp_choice
//              //          }
//            })
            
            NavigationLink {
              
              NotificationSoftRequest()
            } label: {
              Text("Not Right Now")
              //              .frame(maxWidth: .infinity)
              
            }
            .buttonStyle(CustomizableResponsiveButton(fontSize: .base, fontWeight: .light, bgColor: .bg, fgColor: .graySeven))
            //              .frame(maxWidth: .infinity)
            .padding(.top)
            .simultaneousGesture(TapGesture().onEnded {
              print("Denied loc")
              //          DispatchQueue.main.async {
              //              self.workoutController.selectedCompetition = comp_choice
              //          }
            })
          }
        }
        .onAppear {
          if let st = loc.status?.rawValue {
            if st > 2 {
              gotLocPerm = true
            }
          }
        }
        .navigationDestination(isPresented: $gotLocPerm) {
          NotificationSoftRequest()
        }
        
        .background(Color.bg.ignoresSafeArea())
      }
    }
}

struct LocationSoftRequest_Previews: PreviewProvider {
    static var previews: some View {
        LocationSoftRequest()
        .setupPreview()
        .environmentObject(LocationManager())
    }
}

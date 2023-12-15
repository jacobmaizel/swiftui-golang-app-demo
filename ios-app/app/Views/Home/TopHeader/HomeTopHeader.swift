//
//  HomeTopHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/14/23.
//

import SwiftUI

struct HomeTopHeader: View {
  
  @EnvironmentObject var data: Datamodel
//  @EnvironmentObject var weather: WeatherManager
  
  
  var body: some View {
    
    GeometryReader { reader in
      HStack {
        // MARK: - Left side Profile icon
        HStack(alignment: .center) {
          Button {
            dispatchAsync {
              data.currentTab = .profile
            }
          } label: {
            CircleImage(url: data.profile.picture, size: 40)
          }
          .padding(.leading)
          
//          Text(weather.temp)
          CurrentWeatherStack(stack: .v)
       
          
          Spacer()
          
          
        }
//        .padding([.leading, .trailing], 10) // Added padding to leading and trailing edges
        .frame(width: reader.size.width / 3)
//        .border(.white)
        
        // MARK: - search view
          
        Text("Demoios")
          .DemoiosText(fontSize: .xxl, color: .grayOne, weight: .bold)
          .padding(.trailing)
//        Text(weather.temp)
        
//        NavigationLink {
//          SearchView()
//        } label: {
//          SectionItem(fixedBy: .both, cornerRadius: .large) {
//            HStack(alignment: .center) {
//              Text("Search")
//              Image(systemName: "magnifyingglass")
//            }
//            .padding()
//            .padding(.horizontal)
////            .padding(.leading)
//          }
//          .opacity(0.5)
//        }
//        .padding(.trailing)
        
//        .foregroundColor(.grayFour)
//        .buttonStyle(ResponsiveButtonBase())
        .frame(width: reader.size.width / 3)
        
        // MARK: - Right side options
        HStack(alignment: .center) {
       
          NavigationLink(value: HomePaths.compCreate){
              Image(systemName: "plus")
                .resizable()
                .foregroundColor(.grayOne)
                .frame(width: 30, height: 30)
            }
          
          
          NavigationLink(value: HomePaths.notificationList) {
            ZStack {
              Image(systemName: "bell.fill")
                .resizable()
                .foregroundColor(.grayOne)
                .frame(width: 30, height: 30)
              
              if !data.notifications.filter({ $0.opened == nil }).isEmpty {
                ZStack {
                  Circle()
                    .fill(Color.primarySeven)
                    .frame(width: 20, height: 20)
                  
                  Text("\(data.notifications.filter({ $0.opened == nil }).count)")
                    .DemoiosText(fontSize: .xs, color: .grayOne)
                }
                .offset(x: 10, y: 10)
              } else {
                ZStack {
                  Circle()
                    .fill(Color.clear)
                    .frame(width: 20, height: 20)
                }
                .offset(x: 10, y: 10)
              }
            }
          }
//          .border(.white)
        .padding([.leading, .trailing], 10) // Added padding to leading and trailing edges
        }
        .frame(width: reader.size.width / 3)
       
//        .border(.red)
      }
//      .task {
//       await weather.getWeather()
//        print(weather.temp)
//      }
    }
    
    .frame(height: 60)
    .DemoiosBackgroundColor()
    //    .padding(.horizontal)
  }
}

struct HomeTopHeader_Previews: PreviewProvider {
  static var previews: some View {
    
    
    HomeTopHeader()
      .environmentObject(Datamodel())
      .environmentObject(WeatherManager())
    
      .setupPreview()
  }
}

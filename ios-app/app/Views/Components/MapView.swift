//
//  MapView.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/23/23.
//

import SwiftUI
import MapKit
import CoreLocation

struct MapView: View {
  
  var route: [CLLocation]
  
  @State var coordinates: [CLLocationCoordinate2D] = []
  @State var poly: MKPolyline?
  
  @State var city: String = ""
  
  var body: some View {
    SectionItem(fixedBy: .v, cornerRadius: .medium) {
      
      VStack(spacing: 0) {
        
        UIKitMapView(route: $poly, coords: $coordinates)
        //        Map(coordinateRegion: $region, interactionModes: [])
          .frame(height: 200)
          .clipShape(RoundedCorner(radius: 8, corners: [.topLeft, .topRight]))
        
        ZStack {
          
          RoundedCorner(radius: 8, corners: [.bottomLeft, .bottomRight])
            .frame(height: 40)
            .foregroundColor(.grayTen)
          
          HStack {
            Text("Route")
            Spacer()
            Image(systemName: "location.fill")
              .foregroundColor(.blue)
            Text(city)
          }
          .DemoiosText(fontSize: .base)
          .padding(.horizontal)
          
        }
      }
      
    }
    //    .onChange(of: route) { newRoute in
    //      if !newRoute.isEmpty {
    //        Task {
    //          self.city = await getCityFromRoutePoint(newRoute[0])
    //        }
    //      }
    //
    //    }
    .task {
      if !route.isEmpty {
        Task {
          self.city = await getCityFromRoutePoint(route[0])
        }
      }
      
    }
//    .onChange(of: route, perform: { newValue in
//      self.coordinates = newValue.map( { $0.coordinate })
//    })
    .onAppear {
      self.coordinates = route.map( { $0.coordinate } )
      self.poly = MKPolyline(coordinates: coordinates, count: coordinates.count)
    }
  }
}

struct MapView_Previews: PreviewProvider {
  static var previews: some View {
    MapView(route: CLLocation.testData)
      .setupPreview()
  }
}



struct UIKitMapView: UIViewRepresentable {
  @Binding var route: MKPolyline?
  //  @Binding var region: MKCoordinateRegion?
  @Binding var coords: [CLLocationCoordinate2D]
  
  
  
  let mapViewDelegate = UIKitMapViewDelegate()
  
  func makeUIView(context: Context) -> MKMapView {
    MKMapView(frame: .zero)
  }
  
  func updateUIView(_ view: MKMapView, context: Context) {
    view.delegate = mapViewDelegate
    view.translatesAutoresizingMaskIntoConstraints = false
    addRoute(to: view)
    setRegion(to: view)
    
  }
}

private extension UIKitMapView {
  func calculateRegionFromRouteCoords(_ coords: [CLLocationCoordinate2D]) -> MKCoordinateRegion {
    var minLatitude: CLLocationDegrees = 90.0
    var maxLatitude: CLLocationDegrees = -90.0
    var minLongitude: CLLocationDegrees = 180.0
    var maxLongitude: CLLocationDegrees = -180.0
    
    for coordinate in coords {
      minLatitude = min(minLatitude, coordinate.latitude)
      maxLatitude = max(maxLatitude, coordinate.latitude)
      minLongitude = min(minLongitude, coordinate.longitude)
      maxLongitude = max(maxLongitude, coordinate.longitude)
    }
    let center = CLLocationCoordinate2D(latitude: (minLatitude + maxLatitude) / 2, longitude: (minLongitude + maxLongitude) / 2)
    let span = MKCoordinateSpan(latitudeDelta: (maxLatitude - minLatitude) * 1.1, longitudeDelta: (maxLongitude - minLongitude) * 1.1)  // Multiplying by 1.1 to add a little padding
    
    return MKCoordinateRegion(center: center, span: span)
  }
  
  func addRoute(to view: MKMapView) {
    if !view.overlays.isEmpty {
      view.removeOverlays(view.overlays)
    }
    
    guard let route = route else {
      
      return
    }
    let mapRect = route.boundingMapRect
    view.setVisibleMapRect(mapRect, edgePadding: UIEdgeInsets(top: 20, left: 20, bottom: 20, right: 20), animated: true)
    view.addOverlay(route)
  }
  
  func setRegion(to view: MKMapView) {
    //    guard let region = region else { return }
    guard !coords.isEmpty else {
      return
    }
    //    view.setRegion(calculateRegionFromRouteCoords(coords), animated: true)
  }
}

class UIKitMapViewDelegate: NSObject, MKMapViewDelegate {
  func mapView(_ mapView: MKMapView, rendererFor overlay: MKOverlay) -> MKOverlayRenderer {
    let renderer = MKPolylineRenderer(overlay: overlay)
    renderer.fillColor = UIColor.red.withAlphaComponent(0.5)
    renderer.strokeColor = UIColor.red.withAlphaComponent(0.8)
    return renderer
  }
}

func getCityFromRoutePoint(_ location: CLLocation) async -> String {
  let cg = CLGeocoder()
  
  do {
    let res = try await cg.reverseGeocodeLocation(location)
    //    print(res.first)
    return res.first?.locality ?? ""
  } catch {
    print(error)
    return ""
  }
  
}

extension CLLocation {
  static let testData: [CLLocation] = [
    CLLocation(latitude: 39.7105738055676, longitude: -73.01698857690056),
    CLLocation(latitude: 40.7105738055676, longitude: -74.01698857690056),
    CLLocation(latitude: 41.7105738055676, longitude: -75.01698857690056)
  ]
}

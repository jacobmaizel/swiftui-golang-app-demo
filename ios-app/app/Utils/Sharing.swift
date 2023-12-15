//
//  Sharing.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/4/23.
//

import Foundation



enum ShareableResource: String, CaseIterable {
  case workouts, competitions, squads, profiles
}
func createShareableLink(type: ShareableResource, id: UUID) -> URL {
  //https://app-site-association.cdn-apple.com/a/v1/6786-2600-4041-7972-c400-88e7-451f-22f2-5814.ngrok-free.app
  //https://app-site-association.cdn-apple.com/a/v1/Demoios.com
  var base = ""
  
  //  let base = "https://Demoios.com/"
  
  if Configuration.environment.absoluteString.lowercased() == "release" {
    base = "Demoios.com"
  } else {
    base = "6a41-2600-4041-7972-c400-acde-afec-f1cb-2223.ngrok-free.app"
  }
  
  let res = "https://" + base + "/app/\(type.rawValue)/"
  
  var resUrl = URL(string: res)
  resUrl?.append(queryItems: [URLQueryItem(name: "id", value: id.uuidString)])
  
  //  print("CREATED SHAARE LINKKKK \(resUrl?.absoluteString)")
  
  return resUrl!
  
  
}

func processOpenedShareableUrl(with url: URL) -> (ShareableResource, UUID)? {
  
  let components: URLComponents? = URLComponents(string: url.absoluteString)
  
  let path = url.pathComponents
  
  guard
    let resource = ShareableResource(rawValue: path.last ?? ""),
    let qp = components?.queryItems?.first(where: { $0.name == "id"} ),
    let resourceId = UUID(uuidString: qp.value ?? "")
  else {
    print("invalid resource type")
    return nil
  }
  
  print("Processed URL:::", resource, qp, resourceId)
  
  return (resource, resourceId)
  
}


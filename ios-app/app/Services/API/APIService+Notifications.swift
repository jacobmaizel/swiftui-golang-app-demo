//
//  APIService+Notifications.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/2/23.
//

import Foundation



extension DemoiosAPIService {
  
  @MainActor
  func openNotification(data: Datamodel = .shared, id: UUID) async {
    
    let endpoint: String = "/notifications/"
    
    let payload: [String: Any] = [
      "id": id.uuidString,
      "opened": Date.now.intoString ?? ""
    ]
    
    do {
      let _: Notification = try await asyncRequest(method: .patch, path: endpoint, payload: payload)
      
    if let idx = data.notifications.firstIndex(where: { $0.id == id }) {
            data.notifications[idx].opened = Date.now.intoString ?? ""
    }
   } catch {
        print(error)
      }
      
    
  }
  
  @MainActor
  func getMyNotifications(data: Datamodel = .shared) async {
    let endpoint: String = "/notifications/"
    
    let res: [Notification]? = try? await asyncRequest(method: .get, path: endpoint)
    
    data.notifications = res ?? []
        
    }
  
}

//
//  NotificationHandler.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/3/23.
//

import Foundation
import SwiftUI


class NotificationHandler: ObservableObject {
  static let shared = NotificationHandler()
  
  @Published var currentNotification: NotificationBannerData?
  
  func new(title: String, body: String, background: Color, icon: NotificationBannerIcons = .info ) {
    self.currentNotification = NotificationBannerData(title: title, body: body, background: background, icon: icon)
  }
}

struct NotificationBannerData: Equatable {
  var title: String
  var body: String
  var background: Color
  var icon: NotificationBannerIcons
}

enum NotificationBannerIcons: String {
  case info = "info.square.fill"
  case warning = "exclamationmark.triangle.fill"
  case ok = "hand.thumbsup.fill"
}

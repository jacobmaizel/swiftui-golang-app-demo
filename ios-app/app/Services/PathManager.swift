//
//  PathManager.swift
//  Demoios
//
//  Created by Jacob Maizel on 8/19/23.
//

import Foundation

import SwiftUI

class PathManager: ObservableObject {
  
  static let shared = PathManager()
  @Published var homePath = NavigationPath()
  @Published var compPath = NavigationPath()
  @Published var workoutPath = NavigationPath()
  @Published var squadPath = NavigationPath()
  @Published var profilePath = NavigationPath()
  
  func clear() {
    homePath.removeLast(PathManager.shared.homePath.count)
    compPath.removeLast(PathManager.shared.compPath.count)
    workoutPath.removeLast(PathManager.shared.workoutPath.count)
    squadPath.removeLast(PathManager.shared.squadPath.count)
    profilePath.removeLast(PathManager.shared.profilePath.count)

  }
  
}

enum WorkoutPaths {
  case `import`
  case create
}

enum SquadPaths {
  case create
  case memberDetail
}

enum ProfilePaths {
  case edit
  case preferences
  case about
}
  enum HomePaths {
//    case complist
    case workoutcreate
    case findsquad
    case viewstats
    case search
    case compCreate
    case notificationList
  }
//
//enum CompTabViewPaths {
//  case detail
//}

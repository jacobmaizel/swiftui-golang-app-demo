//
//  Notification.swift
//  Demoios
//
//  Created by Jacob Maizel on 7/2/23.
//

import Foundation

struct Notification: Identifiable, Decodable, Equatable {
  static func == (lhs: Notification, rhs: Notification) -> Bool {
    if lhs.id == rhs.id {
      return true
    }
    return false
  }
  
  var id = UUID()
  
  var title: String
  var body: String
  
  var sent: String
  var opened: String?
  
  var data: NotificationData
  static let test_comp_unopened = Notification(title: "Comp Unopened", body: "body", sent: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(),  sender_image: Consts.JAKE_M_PROFILE_IMAGE, competition_id: UUID()))
  static let test_squad_unopened = Notification(title: "hi", body: "body", sent: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(),  sender_image: Consts.JAKE_M_PROFILE_IMAGE, squad_id: UUID()))
  static let test_workout_unopened = Notification(title: "Workout UNOpened", body: "Noti Body", sent: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(),  sender_image: Consts.JAKE_M_PROFILE_IMAGE, workout_id: UUID()))
  
  
  static let test_unopened = Notification(title: "hi", body: "body", sent: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(),  sender_image: Consts.JAKE_M_PROFILE_IMAGE, workout_id: UUID()))
  
  static let test_unopened_comp = Notification(title: "hi", body: "body", sent: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(),sender_image: Consts.JAKE_M_PROFILE_IMAGE, competition_id: UUID()))
  
  static let test_opened = Notification(title: "hi opened", body: "body opened", sent: Date.now.adding(days: -1).intoString ?? "", opened: Date.now.intoString ?? "", data: NotificationData(notification_id: UUID(), sender_image: Consts.JAKE_M_PROFILE_IMAGE, workout_id: UUID()))
}


struct NotificationData: Decodable {
  var notification_id: UUID
  
  var sender_id: UUID?
  var sender_image: String = ""
  
  var receiver_id: UUID?
  
  var workout_id: UUID?
  var squad_id: UUID?
  var competition_id: UUID?
  var invite_id: UUID?
}
/*
 type NotificationData struct {
   Notification_id uuid.UUID `json:"notification_id"`
   Sender_id       uuid.UUID `json:"sender_id"`
   Sender_image    string    `json:"sender_image"`
   Receiver_id     uuid.UUID `json:"receiver_id"`

   Workout_id     *uuid.UUID `json:"workout_id"`
   Squad_id       *uuid.UUID `json:"squad_id"`
   Competition_id *uuid.UUID `json:"competition_id"`
   Invite_id      *uuid.UUID `json:"invite_id"`
 }
 */

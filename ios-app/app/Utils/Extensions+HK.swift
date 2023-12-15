//
//  Extensions+HK.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/20/23.
//

import Foundation
import HealthKit



enum LocationChoiceEnum: String, Identifiable, CaseIterable {
  case indoor
  case outdoor
  
  var id: Self {self}
}
extension LocationChoiceEnum {
  
  func getHKSessionLocationType() -> HKWorkoutSessionLocationType {
    switch self {
    case .indoor:
      return .indoor
    case .outdoor:
      return .outdoor
    }
  }
}

extension HKWorkoutSessionLocationType {
  
  func getString() -> String {
    switch self {
    case .unknown:
      return "Unknown"
    case .indoor:
      return "Indoor"
    case .outdoor:
      return "Outdoor"
    @unknown default:
        return "Unknown"
    }
  }
}

extension HKWorkoutActivityType {
    
    /*
     Simple mapping of available workout types to a human readable name.
     */
    var name: String {
        switch self {
          
        case .running:                      return "Running"
        case .walking:                      return "Walking"
        case .cycling:                      return "Cycling"
        case .swimming:                     return "Swimming"


        case .americanFootball:             return "American Football"
        case .archery:                      return "Archery"
        case .australianFootball:           return "Australian Football"
        case .badminton:                    return "Badminton"
        case .baseball:                     return "Baseball"
        case .basketball:                   return "Basketball"
        case .bowling:                      return "Bowling"
        case .boxing:                       return "Boxing"
        case .climbing:                     return "Climbing"
        case .crossTraining:                return "Cross Training"
        case .curling:                      return "Curling"
        case .dance:                        return "Dance"
        case .danceInspiredTraining:        return "Dance Inspired Training"
        case .elliptical:                   return "Elliptical"
        case .equestrianSports:             return "Equestrian Sports"
        case .fencing:                      return "Fencing"
        case .fishing:                      return "Fishing"
        case .functionalStrengthTraining:   return "Functional Strength Training"
        case .golf:                         return "Golf"
        case .gymnastics:                   return "Gymnastics"
        case .handball:                     return "Handball"
        case .hiking:                       return "Hiking"
        case .hockey:                       return "Hockey"
        case .hunting:                      return "Hunting"
        case .lacrosse:                     return "Lacrosse"
        case .martialArts:                  return "Martial Arts"
        case .mindAndBody:                  return "Mind and Body"
        case .mixedMetabolicCardioTraining: return "Mixed Metabolic Cardio Training"
        case .paddleSports:                 return "Paddle Sports"
        case .play:                         return "Play"
        case .preparationAndRecovery:       return "Preparation and Recovery"
        case .racquetball:                  return "Racquetball"
        case .rowing:                       return "Rowing"
        case .rugby:                        return "Rugby"
     
        case .sailing:                      return "Sailing"
        case .skatingSports:                return "Skating Sports"
        case .snowSports:                   return "Snow Sports"
        case .soccer:                       return "Soccer"
        case .softball:                     return "Softball"
        case .squash:                       return "Squash"
        case .stairClimbing:                return "Stair Climbing"
        case .surfingSports:                return "Surfing Sports"
        case .tableTennis:                  return "Table Tennis"
        case .tennis:                       return "Tennis"
        case .trackAndField:                return "Track and Field"
        case .traditionalStrengthTraining:  return "Traditional Strength Training"
        case .volleyball:                   return "Volleyball"
    
        case .waterFitness:                 return "Water Fitness"
        case .waterPolo:                    return "Water Polo"
        case .waterSports:                  return "Water Sports"
        case .wrestling:                    return "Wrestling"
        case .yoga:                         return "Yoga"
            
            // iOS 10
        case .barre:                        return "Barre"
        case .coreTraining:                 return "Core Training"
        case .crossCountrySkiing:           return "Cross Country Skiing"
        case .downhillSkiing:               return "Downhill Skiing"
        case .flexibility:                  return "Flexibility"
        case .highIntensityIntervalTraining:    return "High Intensity Interval Training"
        case .jumpRope:                     return "Jump Rope"
        case .kickboxing:                   return "Kickboxing"
        case .pilates:                      return "Pilates"
        case .snowboarding:                 return "Snowboarding"
        case .stairs:                       return "Stairs"
        case .stepTraining:                 return "Step Training"
        case .wheelchairWalkPace:           return "Wheelchair Walk Pace"
        case .wheelchairRunPace:            return "Wheelchair Run Pace"
            
            // iOS 11
        case .taiChi:                       return "Tai Chi"
        case .mixedCardio:                  return "Mixed Cardio"
        case .handCycling:                  return "Hand Cycling"
            
            // iOS 13
        case .discSports:                   return "Disc Sports"
        case .fitnessGaming:                return "Fitness Gaming"
            
            // Catch-all
        default:                            return "Other"
        }
    }
}

extension String {
    
    /*
     Simple mapping of human readable workout names to their corresponding HKWorkoutActivityType values.
     */
    var workoutActivityType: HKWorkoutActivityType? {
        switch self {
        case "American Football":                          return .americanFootball
        case "Archery":                                    return .archery
        case "Australian Football":                        return .australianFootball
        case "Badminton":                                  return .badminton
        case "Baseball":                                   return .baseball
        case "Basketball":                                 return .basketball
        case "Bowling":                                    return .bowling
        case "Boxing":                                     return .boxing
        case "Climbing":                                   return .climbing
        case "Cross Training":                             return .crossTraining
        case "Curling":                                    return .curling
        case "Cycling":                                    return .cycling
        case "Elliptical":                                 return .elliptical
        case "Equestrian Sports":                          return .equestrianSports
        case "Fencing":                                    return .fencing
        case "Fishing":                                    return .fishing
        case "Functional Strength Training":               return .functionalStrengthTraining
        case "Golf":                                       return .golf
        case "Gymnastics":                                 return .gymnastics
        case "Handball":                                   return .handball
        case "Hiking":                                     return .hiking
        case "Hockey":                                     return .hockey
        case "Hunting":                                    return .hunting
        case "Lacrosse":                                   return .lacrosse
        case "Martial Arts":                               return .martialArts
        case "Mind and Body":                              return .mindAndBody
        case "Paddle Sports":                              return .paddleSports
        case "Play":                                       return .play
        case "Preparation and Recovery":                   return .preparationAndRecovery
        case "Racquetball":                                return .racquetball
        case "Rowing":                                     return .rowing
        case "Rugby":                                      return .rugby
        case "Running":                                    return .running
        case "Sailing":                                    return .sailing
        case "Skating Sports":                             return .skatingSports
        case "Snow Sports":                                return .snowSports
        case "Soccer":                                     return .soccer
        case "Softball":                                   return .softball
        case "Squash":                                     return .squash
        case "Stair Climbing":                             return .stairClimbing
        case "Surfing Sports":                             return .surfingSports
        case "Swimming":                                   return .swimming
        case "Table Tennis":                               return .tableTennis
        case "Tennis":                                     return .tennis
        case "Track and Field":                            return .trackAndField
        case "Traditional Strength Training":              return .traditionalStrengthTraining
        case "Volleyball":                                 return .volleyball
        case "Walking":                                    return .walking
        case "Water Fitness":                              return .waterFitness
        case "Water Polo":  return .waterPolo
        case "Water Sports":                               return .waterSports
        case "Wrestling":                                  return .wrestling
        case "Yoga":                                       return .yoga
            
            // iOS 10
        case "Barre":                                      return .barre
        case "Core Training":                              return .coreTraining
        case "Cross Country Skiing":                        return .crossCountrySkiing
        case "Downhill Skiing":                             return .downhillSkiing
        case "Flexibility":                                return .flexibility
        case "High Intensity Interval Training":            return .highIntensityIntervalTraining
        case "Jump Rope":                                  return .jumpRope
        case "Kickboxing":                                 return .kickboxing
        case "Pilates":                                    return .pilates
        case "Snowboarding":                               return .snowboarding
        case "Stairs":                                     return .stairs
        case "Step Training":                              return .stepTraining
        case "Wheelchair Walk Pace":                        return .wheelchairWalkPace
        case "Wheelchair Run Pace":                         return .wheelchairRunPace
            
            // iOS 11
        case "Tai Chi":                                    return .taiChi
        case "Mixed Cardio":                               return .mixedCardio
        case "Hand Cycling":                               return .handCycling
            
            // iOS 13
        case "Disc Sports":                                return .discSports
        case "Fitness Gaming":                             return .fitnessGaming
            
            // Catch-all
        default:                                            return nil
        }
    }
}

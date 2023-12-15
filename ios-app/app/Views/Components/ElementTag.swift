//
//  ElementTag.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/4/23.
//

import SwiftUI
enum TagSize: CGFloat {
  case small = 0.55
  case medium = 0.70
  case base = 0.9
}

enum StackType {
  case h
  case v
}

enum ElementType: Hashable {
  case privacy
  case participant_types
  case workout_types
  case official
}




// struct WorkoutTypeIconStack: View {
//  var workout_types: [String]
//
//  var body: some View {
//
//    HStack {
//
//    }
//  }
// }

struct ElementTag: View {
  
  var title: String
  var type: ElementType
  var values: [String]
  var scale: TagSize
  var stack: StackType
  var color: Color
  
  //  init(title: String, scale: TagSize, stack: StackType, color: Color, icon: Image) {
  //    self.title = title
  //    self.scale = scale
  //    self.stack = stack
  //    self.color = color
  //    self.icon = icon
  //  }
  
  var baseHeight: CGFloat {
    switch stack {
    case .h:
      return 30
    case .v:
      return 50
    }
  }
  var baseWidth: CGFloat {
    switch stack {
    case .h:
      return 30
    case .v:
      return 60
    }
  }
  
  var height: CGFloat {
    return baseHeight * scale.rawValue
  }
  
  var width: CGFloat {
    return baseWidth * scale.rawValue
  }
  
  var fontSize: DemoiosFontSizes {
    switch scale {
    case .small:
      return .xxs
    case .medium:
      return .small
    case .base:
      return .base
    }
  }
  //  for type: ElementType, with values: [String]
  func calcIconStack() -> some View {
    
    Group {
      if stack == .h {
        HStack {
          ForEach(values, id: \.self) { icon_key in
            Image(systemName: tagInfoLookup[type]?[icon_key] ?? "person.fill")
              .resizable()
              .aspectRatio(contentMode: .fit)
              .frame(width: 20 * scale.rawValue * 1.2, height: 20 * scale.rawValue * 1.2)
        .foregroundColor(color)
          }
        }
      } else {
        
        HStack {
          ForEach(values, id: \.self) { icon_key in
            Image(systemName: tagInfoLookup[type]?[icon_key] ?? "person.fill")
              .resizable()
              .aspectRatio(contentMode: .fit)
              .frame(width: 20 * scale.rawValue * 1.3, height: 20 * scale.rawValue * 1.3)
              .foregroundColor(color)
          }
        }
        
      }
    }
    
    
  }
  
  var body: some View {
    
    SectionItem(fixedBy: .both, cornerRadius: .medium, color: .clear) {
      
      switch stack {
      case .h:
        HStack {
          //          icon
          calcIconStack()
          
          Text(title)
            .DemoiosText(fontSize: fontSize, color: color)
        }
        .padding(EdgeInsets(top: 4, leading: 4, bottom: 4, trailing: 4))
//        .overlay {
//          RoundedRectangle(cornerSize: CGSize(width: 16, height: 16))
//            .stroke(color, lineWidth: 2)
//          //                .frame(minWidth: 75)
//        }
//        .brightness(0.6)
        
      case .v:
        VStack(spacing: 0) {
          calcIconStack()
          
          Text(title)
            .DemoiosText(fontSize: fontSize, color: color)
        }
        .frame(minWidth: width, maxHeight: height)
//        .padding(EdgeInsets(top: 6, leading: 4, bottom: 4, trailing: 4))
        .padding(EdgeInsets(top: 6, leading: 8, bottom: 6, trailing: 8))
//        .overlay {
//          RoundedRectangle(cornerSize: CGSize(width: 16, height: 16))
//            .stroke(color, lineWidth: 2)
//          //                .frame(minWidth: 75)
//        }

//        .brightness(0.6)
      }
    }
    //    .border(.white)
    
    .DemoiosText(fontSize: fontSize)
    
  }
}

struct ElementTag_Previews: PreviewProvider {
  static var previews: some View {
    Group {
      VStack {
        HStack {
          
          ElementTag(title: "Official", type: .official, values: ["official"], scale: .small, stack: .v, color: .primarySeven)
          
          ElementTag(title: "Group", type: .participant_types, values: ["official"], scale: .small, stack: .v, color: .orange)
          
          
          ElementTag(title: "Open", type: .privacy, values: ["official"], scale: .small, stack: .v, color: .blue)
          
          ElementTag(title: "Cardio", type: .workout_types, values: ["official"], scale: .small, stack: .v, color: .purple)
          
        }
        
        
        
        
        HStack {
          ElementTag(title: "Official", type: .workout_types, values: ["official"], scale: .medium, stack: .v, color: .primarySeven)
          
          ElementTag(title: "Group", type: .workout_types, values: ["official"], scale: .medium, stack: .v, color: .orange)
          
          
          ElementTag(title: "Open", type: .workout_types, values: ["official"], scale: .medium, stack: .h, color: .blue)
          
          ElementTag(title: "Cardio", type: .workout_types, values: ["official"], scale: .medium, stack: .h, color: .purple)
        }
        
        HStack {
          ElementTag(title: "Official", type: .workout_types, values: ["official"], scale: .small, stack: .h, color: .primarySeven)
          
          ElementTag(title: "Group", type: .workout_types, values: ["official"], scale: .small, stack: .h, color: .orange)
          
          
          ElementTag(title: "Open", type: .workout_types, values: ["official"], scale: .small, stack: .v, color: .blue)
          
          ElementTag(title: "Cardio", type: .workout_types, values: ["official"], scale: .small, stack: .v, color: .purple)
        }
      }
    }
    .setupPreview()
  }
}

let tagInfoLookup: [ElementType: [String: String]] = [
  .official: [
    "official": "checkmark.seal.fill",
    "community": "person.line.dotted.person.fill"
  ],
  .participant_types: [
    "solo": "person.fill",
    "squad": "person.3.fill"
  ],
  .privacy: [
    "true": "lock.open.fill",
    "false": "lock.fill"
  ],
  .workout_types: [
    "running": "figure.run",
    "cycling": "bicycle",
    "swimming": "figure.pool.swim",
    "walking": "figure.walk"
  ]
]

// let workoutTypeIconMap = [
// ]
//
// let compTypeIconMap = [
//  "official": "checkmark.seal.fill",
//  "community": "person.line.dotted.person.fill"
// ]
//
// let participantTypeIconMap = [
// ]
//
// let privacyIconmap = [
//  "true": "lock.fill",
//  "false": "lock.open.fill"
// ]

/*
 
 "memberType": {
 "squad": {
 "Color": Color(asdf)
 
 }
 }
 
 */
//
// let tagInfolookup: [String: [String: [String: [String]]]] = [
//
//    "memberType": [
//        "Squad": Color(hue: 52, saturation: 32, brightness: 38),
//        "Solo": Color(hue: 208, saturation: 87, brightness: 98),
//        "Any": Color(hue: 36, saturation: 65, brightness: 89)
//    ],
//
//    "creatorType": [
//        "Demoios": .primarySeven,
//        "User": .blue
//    ],
//
//    "joinability": [
//        "InviteOnly": Color(hue: 99, saturation: 35, brightness: 48),
//        "Open": Color(hue: 220, saturation: 34, brightness: 47)
//    ],
//
//    "workoutType": [
//        "AnyCardio": Color(hue: 267, saturation: 53, brightness: 60),
//        "RunWalk": Color(hue: 214, saturation: 98, brightness: 55),
//        "Cycle": Color(hue: 5, saturation: 94, brightness: 86)
//    ]
// ]


// struct IconMap {
//    let workoutTypeIconMap: [String: String]
//    let compTypeIconMap: [String: String]
//    let participantTypeIconMap: [String: String]
//    let privacyIconmap: [Bool: String]
// }
//
// let iconMaps = IconMap(
//    workoutTypeIconMap: [
//        "running": "figure.run",
//        "cycling": "bicycle",
//        "swimming": "figure.pool.swim",
//        "walking": "figure.walk"
//    ],
//    compTypeIconMap: [
//        "official": "checkmark.seal.fill",
//        "community": "person.line.dotted.person.fill"
//    ],
//    participantTypeIconMap: [
//        "solo": "person.fill",
//        "squad": "person.3.fill"
//    ],
//    privacyIconmap: [
//        true: "lock.fill",
//        false: "lock.open.fill"
//    ]
// )
//
//// enum ElementType {
////    case privacy
////    case participant_types
////    case workout_types
////    case official
//// }
//
// let mainDictionary: [ElementType: IconMap] = [
//    .privacy: iconMaps.privacyIconmap,
//    .participant_types: iconMaps.participantTypeIconMap,
//    .workout_types: iconMaps.workoutTypeIconMap,
//    .official: iconMaps.compTypeIconMap
// ]

//
//  Modifiers.swift
//  Demoios
//
//  Created by Jacob Maizel on 2/26/23.
//

import Foundation
import SwiftUI
/*
 
 Approx Built in font sizing
 
 .largeTitle: 34 points
 .title: 28 points
 .headline: 17 points
 .subheadline: 15 points
 .body: 17 points
 .callout: 16 points
 .footnote: 13 points
 .caption: 12 points
 */
#if os(iOS)
struct BannerModifier: ViewModifier {
  @Binding var notiToDisplay: NotificationBannerData?
  
  func body(content: Content) -> some View {
    content.overlay(
      VStack {
        if notiToDisplay != nil {
          let timer = Timer.scheduledTimer(withTimeInterval: 10.0, repeats: false) { _ in
            notiToDisplay = nil
          }
          VStack {
            HStack {
              
              Image(systemName: notiToDisplay?.icon.rawValue ?? "exclamationmark.triangle.fill" )
              
              VStack(alignment: .leading, spacing: 2) {
                
                Text(notiToDisplay?.title ?? "")
                  .DemoiosText(fontSize: .base, weight: .bold)
                
                if let message = notiToDisplay?.body {
                  Text(message)
                    .DemoiosText(fontSize: .xs, color: .grayTwo)
                }
                
              } // vstack
              
              Spacer()
            } // hstack
            .padding(.vertical, 6)
            .padding(.horizontal)
            .frame(minWidth: 0, maxWidth: .infinity)
            //            .foregroundColor(.white)
            .background(notiToDisplay?.background ?? .primarySeven)
            .cornerRadius(DemoiosCornerRadius.large.rawValue)
            .shadow(radius: DemoiosCornerRadius.large.rawValue)
            
            Spacer()
          }
          .padding()
          .transition(AnyTransition.move(edge: .top).combined(with: .opacity))
          .onTapGesture {
            print(notiToDisplay)
            withAnimation {
              notiToDisplay = nil
            }
          }
          .gesture(
            DragGesture()
              .onChanged { _ in
                withAnimation {
                  notiToDisplay = nil
                }
              }
          )
        } // if model != nil
      } // outer vstack
        .animation(.easeInOut(duration: 0.8), value: notiToDisplay)
      
    )
  }
}

#endif

enum FixedBy {
  case both
  case none
  case h
  case v
}

enum DemoiosFontSizes {
  
  /*
   // usage
   Text("Hello, world!")
   .font(.system(size: DemoiosFontSizes.xl.size))
   */
  
  case xxl
  case xl
  case large
  case base
  case small
  case xs
  case xxs
  
  var textStyle: UIFont.TextStyle {
    switch self {
    case .xxl:
      return .title1
    case .xl:
      return .title2
    case .large:
      return .headline
    case .base:
      return .body
    case .small:
      return .footnote
    case .xs:
      return .caption1
    case .xxs:
      return .caption2
    }
  }
  
  var size: CGFloat {
    return UIFont.preferredFont(forTextStyle: textStyle).pointSize
  }
}

enum DemoiosCornerRadius: CGFloat {
  case small = 4
  case medium = 8
  case large = 16
}

struct TextModifier: ViewModifier {
  
  let fontSize: DemoiosFontSizes
  let color: Color
  let weight: Font.Weight
  
  
  func body(content: Content) -> some View {
    content
      .font(.custom("Kanit-Regular", size: fontSize.size))
    //            .font(.system(size: fontSize.size))
      .foregroundColor(color)
      .fontWeight(weight)
  }
}

struct PrimaryButtonModifier: ViewModifier {
  let cornerRadius: DemoiosCornerRadius
  let color: Color
  
  @ViewBuilder func body(content: Content) -> some View {
    content
      .padding()
      .background(color)
      .clipShape(RoundedRectangle(cornerRadius: cornerRadius.rawValue))
      .padding()
  }
}

struct FullScreenBackgroundModifier: ViewModifier {
  func body(content: Content) -> some View {
    content
      .background(Color.bg.ignoresSafeArea())
  }
}

struct TextInputModifier: ViewModifier {
  
  let inputTitle: String
  
  func body(content: Content) -> some View {
    VStack(alignment: .leading, spacing: 0) {
      Text(inputTitle)
        .foregroundColor(Color.grayEight)
        .DemoiosText(fontSize: .xs)
      content
        .padding(.leading, 10)
        .frame(minHeight: 50)
        .foregroundColor(Color.grayOne)
        .overlay {
          RoundedRectangle(cornerRadius: 4)
            .stroke(Color.graySix, lineWidth: 1)
        }
    }
    .padding(10)
  }
}


struct PreviewModifier: ViewModifier {
  
  func body(content: Content) -> some View {
    
    ZStack {
      Rectangle()
        .fill(.black)
      
      content
      //        .environmentObject(Datamodel())
        .environment(\.colorScheme, .dark)
        .previewLayout(.sizeThatFits)
    }
    
  }
  
}

extension View {
  
  func setupPreview() -> some View {
    modifier(PreviewModifier())
  }
  
  func DemoiosBackgroundColor() -> some View {
    modifier(FullScreenBackgroundModifier())
  }
  
  func DemoiosText(fontSize size: DemoiosFontSizes, color: Color = .grayOne, weight: Font.Weight = .regular ) -> some View {
    modifier(TextModifier(fontSize: size, color: color, weight: weight ))
  }
  
  func DemoiosFullWidthButtonLabelStyling(cornerRadius: DemoiosCornerRadius, color: Color) -> some View {
    modifier(PrimaryButtonModifier(cornerRadius: cornerRadius, color: color))
  }
  
  func DemoiosFormInput(title: String) -> some View {
    modifier(TextInputModifier(inputTitle: title))
  }
}

struct RoundedCorner: Shape {
  var radius: CGFloat
  var corners: UIRectCorner
  
  func path(in rect: CGRect) -> Path {
    let path = UIBezierPath(
      roundedRect: rect,
      byRoundingCorners: corners,
      cornerRadii: CGSize(width: radius, height: radius)
    )
    return Path(path.cgPath)
  }
}

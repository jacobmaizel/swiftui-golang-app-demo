//
//  DateHeader.swift
//  Demoios
//
//  Created by Jacob Maizel on 6/14/23.
//

import SwiftUI

struct DateHeader: View {
    
    var from: String
    var to: String
    
    var body: some View {
        HStack {
            Image(systemName: "calendar")
//                .foregroundColor(.gray)
            
            Text("\(from) - \(to)")
        }
    }
}

struct DateHeader_Previews: PreviewProvider {
    static var previews: some View {
        DateHeader(from: "", to: "")
    }
}

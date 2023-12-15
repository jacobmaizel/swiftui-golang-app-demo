//
//  FormMenuBox.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/29/23.
//

import SwiftUI

struct FormMenuBox: View {
    
    var title: String
    var choices: [String]
    var selection: Binding<String>
    
//    var content: () -> Content
    
    var body: some View {
        
        HStack {
            Spacer()
            Picker(title, selection: selection) {
                ForEach(choices, id: \.self) { choice in
                    Text(choice)
                }
            }
            .pickerStyle(.menu)
        }
        .DemoiosFormInput(title: title)
    }
}

struct FormMenuBox_Previews: PreviewProvider {
    static var previews: some View {
        FormMenuBox(title: "test", choices: ["hey"], selection: .constant("hey"))
    }
}

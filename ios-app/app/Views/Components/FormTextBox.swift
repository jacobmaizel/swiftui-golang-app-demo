//
//  FormTextBox.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/17/23.
//

import SwiftUI

struct FormTextBox: View {
    var title: String
    var binding: Binding<String>
    
    var body: some View {
        TextField(title, text: binding, prompt: Text(title))
            .DemoiosFormInput(title: title)
    }
}

struct FormTextBox_Previews: PreviewProvider {
    static var previews: some View {
        FormTextBox(title: "", binding: .constant("hi"))
    }
}

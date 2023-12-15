//
//  FormDateBox.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/17/23.
//

import SwiftUI

struct FormDateBox: View {
    
    var title: String
    var binding: Binding<Date>
    
    var body: some View {
        
        DatePicker("", selection: binding, displayedComponents: .date)
            .padding(.trailing, 10)
            .DemoiosFormInput(title: title)
    }
}

struct FormDateBox_Previews: PreviewProvider {
    static var previews: some View {
        FormDateBox(title: "test", binding: .constant(Date.now))
    }
}

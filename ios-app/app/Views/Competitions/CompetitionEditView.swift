//
//  CompetitionEditView.swift
//  Demoios
//
//  Created by Jacob Maizel on 1/16/23.
//

import SwiftUI

struct CompetitionEditView: View {
    
    enum Level: String, CaseIterable, Identifiable {
        var id: Self {self}
        
        case beginner, intermediate, advanced, pro
    }
    
//    enum Category: String, CaseIterable, Identifiable {
//        var id: Self {self}
//
//        case strength, cardio
//    }
    
    @State private var level: Level = .beginner
//    @State private var category: String = ""
    @State private var tags: String = ""
    @State private var startDate: Date = Date.now
    @State private var endDate: Date = Date.now
    @State private var title: String = ""
    @State private var description: String = ""
    
    var body: some View {
        VStack {
            
            HStack(alignment: .center) {
                Text("Editing " + (title.isEmpty ? "Competition" : title))
                    .font(.title)
                    .padding(.horizontal)
                
                Spacer()
                
                Button("Cancel" ) {
                }
                .padding(.horizontal)
            }
            
            List {
                
                HStack {
                    
            Text("Title")
                
            TextField("", text: $title)
                    .multilineTextAlignment(.trailing)
            }
                
            HStack {
                Text("Description")
                TextField("", text: $description, axis: .vertical )
            }
            
            DatePicker("Start", selection: $startDate, in: ...Date.now)
            DatePicker("End", selection: $startDate, in: ...Date.now)
            
                    Picker("Level", selection: $level) {
                        ForEach(Level.allCases) { level in
                            Text(level.rawValue.capitalized)
                        }
                }
                
//                Picker("Category", selection: $category) {
//                    ForEach(Category.allCases) { category in
//
//                        Text(category.rawValue.capitalized)
//                    }
//                }
            
            HStack {
                Text("Tags")
                TextField("", text: $tags)
            }
        }
            
            Button {
                print("save changes clicked - profile edit")
            } label: {
                Text("Save Changes")
                
                    .bold()
                    .padding()
                    .foregroundColor(.white)
                    .background(.gray)
                    .cornerRadius(10)
            }
        }
    }
}

struct CompetitionEditView_Previews: PreviewProvider {
    static var previews: some View {
        CompetitionEditView()
    }
}

//
//  Protocols.swift
//  Demoios
//
//  Created by Jacob Maizel on 5/7/23.
//

import Foundation

protocol ListableEnum: RawRepresentable, CaseIterable, Identifiable where RawValue == String {}

extension ListableEnum where Self: Hashable {
    var id: Self { self }
    
    var capped: String {
        return self.rawValue.capitalized.replacingOccurrences(of: "_", with: " ")
    }
}

//
//  Configuration.swift
//  Demoios
//
//  Created by Jacob Maizel on 3/14/23.
//
// swiftlint:disable all

import Foundation

enum Configuration {
    
#if os(iOS)
    
    static var apiBaseURL: URL {
        URL(string: string(for: "API_BASE_URL"))!
    }
    
    static var authDomain: URL {
        URL(string: string(for: "AUTH_DOMAIN"))!
    }
    static var authClientId: URL {
        URL(string: string(for: "AUTH_CLIENTID"))!
    }
    static var authAudience: URL {
        URL(string: string(for: "AUTH_AUDIENCE"))!
    }
    
#endif
    
    static var environment: URL {
        URL(string: string(for: "ENVIRONMENT"))!
    }
    
    static private func string(for key: String) -> String {
        Bundle.main.infoDictionary?[key] as! String
    }
}


//swiftlint:enable all

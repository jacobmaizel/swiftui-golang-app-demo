//
//  ConnectivityProvider.swift
//  Demoios
//
//  Created by Jacob Maizel on 4/16/23.
//

import Foundation
import WatchConnectivity

// class ConnectivityProvider: NSObject, WCSessionDelegate {
//    
//    private let session: WCSession
//    
//    init(session: WCSession = .default) {
//        self.session = session
//        super.init()
//        self.session.delegate = self
//        self.connect()
//    }
//    
//// #if os(iOS)
//    func sessionDidBecomeInactive(_ session: WCSession) {
//        print("session did become inactive")
//    }
//    
//    func sessionDidDeactivate(_ session: WCSession) {
//        print("Session did deactivate ")
//    }
//    
//// #endif
//    
//    func connect() {
//        guard WCSession.isSupported() else {
//            print("WCSession is not supported")
//            return
//        }
//        
//        session.activate()
//    }
//    
//    func send(message: [String: Any]) {
//        session.sendMessage(message, replyHandler: nil) { (error) in
//            print(error.localizedDescription)
//        }
//    }
//    
//    func session(_ session: WCSession, activationDidCompleteWith activationState: WCSessionActivationState, error: Error?) {
//        // code
//        
//        print("got activiation state update: \(activationState)")
//    }
//    
//    func session(_ session: WCSession, didReceiveMessage message: [String: Any]) {
//        print("GOT A MESSAGE from WC:::\(message) and platform: \(ProcessInfo.processInfo.operatingSystemVersionString)")
//    }
// }

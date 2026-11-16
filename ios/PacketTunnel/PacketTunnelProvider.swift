import NetworkExtension

/// Network Extension lifecycle boundary.
/// The tunnel intentionally fails closed until a real Xray or direct-TUN
/// PingTunnel adapter is integrated and verified.
final class PacketTunnelProvider: NEPacketTunnelProvider {
    override func startTunnel(options: [String: NSObject]?,
                              completionHandler: @escaping (Error?) -> Void) {
        let error = NSError(
            domain: "com.panha.kxray",
            code: 1001,
            userInfo: [NSLocalizedDescriptionKey:
                "No tunnel engine is integrated in this foundation build."]
        )
        completionHandler(error)
    }

    override func stopTunnel(with reason: NEProviderStopReason,
                             completionHandler: @escaping () -> Void) {
        completionHandler()
    }

    override func handleAppMessage(_ messageData: Data,
                                   completionHandler: ((Data?) -> Void)? = nil) {
        completionHandler?(nil)
    }
}

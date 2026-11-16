# PingTunnel adapter — direct packet mode

The target is full-device VPN over ICMP without a local SOCKS5 intermediary.

This requires more than toggling the upstream client's socks5 option. The mobile side must read IP packets from the OS TUN interface, frame and transport them through ICMP, and the VPS must decode and forward packets (with return traffic routed back through the tunnel).

Required design work:
- Versioned packet framing and flow/session identifiers.
- Strict packet-size and fragmentation policy.
- Authenticated encryption, replay resistance and key rotation.
- Congestion control, retransmission and keepalive.
- Matching server-side adapter and controlled IPv4/IPv6 forwarding tests.

No functional direct-TUN implementation is present yet. Keep this mode unavailable in UI until both client and server paths pass bidirectional traffic and leak tests.

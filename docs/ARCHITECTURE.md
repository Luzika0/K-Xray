# Architecture

## Principles

1. Keep the Android and iOS user interfaces fully native.
2. Share protocol/configuration models and Go engine integration contracts where practical.
3. Let each operating system own its VPN lifecycle and TUN descriptor.
4. Do not route PingTunnel through a local SOCKS5 listener. Its full-device mode must use a packet-oriented adapter.
5. Never report a connection as active until the engine has started and the platform tunnel is established.

## Data path

### Xray

```
Applications
  -> Android VpnService TUN / iOS NEPacketTunnelFlow
  -> platform adapter passes TUN file descriptor to Xray
  -> Xray outbound (VLESS / VMess / Trojan / Shadowsocks, etc.)
  -> remote server
```

Xray documents Android TUN integration through the VpnService file descriptor and iOS integration through the utun descriptor. The native integration must follow the upstream Xray TUN documentation and avoid routing the Xray server connection back into its own TUN.

### PingTunnel (planned direct packet mode)

```
Applications
  -> OS VPN TUN
  -> packet adapter
  -> PingTunnel ICMP carrier
  -> compatible VPS packet adapter
  -> IP forwarding / NAT
  -> Internet
```

The existing PingTunnel SOCKS5 mode is a proxy interface, not a raw-IP TUN interface. Direct mode therefore requires a framed IP-packet transport and a matching server-side adapter. This is a separate engineering milestone; do not simply remove the SOCKS5 setting and assume full VPN works.

## Platform boundaries

- Android: Kotlin owns consent, VpnService lifecycle, foreground notification, routing and descriptor ownership.
- iOS: Swift owns NEVPNManager configuration, Packet Tunnel lifecycle, network settings and descriptor discovery/ownership.
- Go: shared engine orchestration and protocol adapters. Native platform code must pass descriptors using a documented, tested bridge.
- Config: shared JSON schema and import normalization; platform-specific secure storage remains native.

## Engine interface contract

The current Go package intentionally exposes a small not-yet-implemented API. It must fail closed until a real engine adapter is supplied. Never return success for a no-op tunnel.

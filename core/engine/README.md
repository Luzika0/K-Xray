# Shared Go engine

This package is the future gomobile-facing boundary shared by Android and iOS.

Current status: contract only; Start/Stop return ErrNotImplemented. This prevents the native UI from mistaking a no-op for an active VPN.

Planned adapters:
- Xray: pinned Xray-core dependency, built reproducibly for Android/iOS.
- PingTunnel: ICMP carrier plus a packet-oriented TUN adapter. This is not the existing local SOCKS5 mode.

Before binding:
1. Pin upstream versions and review licenses.
2. Verify gomobile compatibility and exported API types.
3. Implement descriptor duplication/ownership carefully per platform.
4. Add unit tests and cross-compile checks.

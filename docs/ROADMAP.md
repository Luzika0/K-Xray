# Implementation roadmap

## Phase 1 — Native foundation (current)

- [x] Repository structure for Android, iOS, Go and shared configuration.
- [x] Native Android Compose application shell with light/dark/system theme.
- [x] Android VpnService declaration and lifecycle placeholder.
- [x] Native SwiftUI app shell and Packet Tunnel extension scaffold.
- [x] Shared Go engine API contract that fails closed.
- [x] Architecture and security documentation.
- [ ] Verify Android build in Android Studio / CI.
- [ ] Generate and open the Xcode project on macOS.

## Phase 2 — Xray integration

- Pin an audited Xray-core release/commit and record its license.
- Build reproducible Android and iOS native artifacts.
- Pass the platform TUN descriptor into Xray using the upstream-supported mechanism.
- Implement routing-loop prevention, DNS, IPv4/IPv6 and MTU handling.
- Add parser/validator tests for supported share links and JSON configs.

## Phase 3 — PingTunnel direct packet transport

- Review and pin upstream PingTunnel source and license.
- Design a versioned packet framing protocol with sequence IDs, fragmentation limits, replay protection and flow control.
- Implement the mobile ICMP carrier adapter and matching Ubuntu server adapter.
- Define packet MTU, retransmission/congestion behavior and keepalive strategy.
- Authenticate and encrypt packets; do not reuse hard-coded example credentials.
- Prove bidirectional IPv4 traffic in a controlled lab before enabling UI connection.

## Phase 4 — Profiles and product features

- Import share links, QR codes, JSON and subscription URLs.
- Securely store secrets in Keychain / Android Keystore.
- Profile editing, grouping, latency checks and diagnostics.
- Connection notifications, reconnect policy and traffic counters.

## Phase 5 — Release readiness

- Android device matrix and iOS device testing.
- Leak tests (DNS, IPv4, IPv6), sleep/wake and network handover tests.
- Battery, memory and throughput profiling.
- Privacy review, license notices, signing and store compliance.

# k-Xray

Native Android and iOS VPN client foundation combining Xray-core and PingTunnel.

> **Status: foundation scaffold.** The native app shells and integration contracts are being established. This repository does not yet contain a working VPN engine, and the Connect action is intentionally not presented as functional until the packet engines are integrated and tested.

## Product goals

- Native Android app: Kotlin, Jetpack Compose and Android `VpnService`.
- Native iOS app: SwiftUI and NetworkExtension `NEPacketTunnelProvider`.
- Xray-core for supported Xray proxy protocols.
- PingTunnel ICMP transport, with a planned direct TUN packet adapter and **no local SOCKS5 intermediary**.
- Shared Go integration layer for engine lifecycle/configuration, with thin native platform adapters.
- Profile import roadmap: share links, QR codes, JSON and subscriptions.
- System/light/dark appearance.

## Repository layout

- `android/` — native Android application.
- `ios/` — native iOS app and Packet Tunnel extension (XcodeGen project).
- `core/` — shared Go engine integration contracts.
- `shared/` — cross-platform configuration schema and documentation.
- `docs/` — architecture, security and implementation roadmap.

## Build prerequisites

### Android (Windows, macOS or Linux)

- Android Studio
- JDK 17 or newer
- Android SDK Platform 36
- Android Gradle Plugin 8.13.2 / Gradle 8.13
- Go toolchain (for later gomobile/native engine integration)

Open `android/` in Android Studio and sync the Gradle project. The initial UI scaffold does not yet start a tunnel.

### iOS (macOS required)

- macOS and Xcode 16 or newer
- XcodeGen
- Apple Developer provisioning with the Network Extension packet-tunnel capability for device deployment

From `ios/` run `xcodegen generate`, open `k-Xray.xcodeproj`, then select a development team and signing settings. The Packet Tunnel provider is a lifecycle scaffold only; it deliberately reports that the engine is not integrated.

## Security notes

- Never commit server passwords, subscription URLs, private keys or REALITY credentials.
- Store secrets using Android Keystore / iOS Keychain when profile persistence is implemented.
- Do not log full configuration payloads or credentials.
- PingTunnel direct-TUN mode requires a compatible server-side packet adapter. The existing SOCKS5 mode is not equivalent to a full-device TUN transport.

## Current implementation status

See [docs/ROADMAP.md](docs/ROADMAP.md) and [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).

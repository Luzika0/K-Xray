# Xray native engine integration

The shared Go package `core/engine` embeds the pinned upstream `github.com/xtls/libxray`
runtime. It exposes gomobile-safe functions:

- `StartXray(tunFD int, configJSON string) error`
- `StopXray() error`
- `XrayVersion() string`
- `IsXrayRunning() bool`

The adapter adds `env.xray.tun.fd` to the supplied Xray JSON and calls the
libXray API v3 `runXray` method. Stop calls `stopXray`. The platform owns the
TUN descriptor; Xray duplicates it internally. Do not close the original
descriptor from the Go adapter.

## Pinned upstream

- libXray commit: `5a0f7d8d1c28b34afb5727b04b26a10f6dbac7e3`
- Go module pseudo-version: `v0.0.0-20261109123211-5a0f7d8d1c28`
- Required Go toolchain: `1.27.3`

libXray itself pins Xray-core. Avoid upgrading the libXray revision without
reviewing its API contract and Xray-core dependency.

## Build Android on Windows

Install Go 1.27.3, Android Studio, Android SDK and Android NDK. Then run from
PowerShell at the repository root:

```powershell
.\scripts\build-xray-android.ps1
```

The output is `android/app/libs/kxray-engine.aar`. Add
`implementation(files("libs/kxray-engine.aar"))` to the Android app module
and use the generated gomobile Java API from the VPN service.

## Build iOS on macOS

Install Xcode (including the iOS Simulator runtime), Go 1.27.3 and command-line
tools, then run:

```bash
./scripts/build-xray-ios.sh
```

The output is `ios/Frameworks/KXrayEngine.xcframework`. Link it to both the
container app and Packet Tunnel extension target.

## Runtime configuration contract

Pass a complete Xray JSON configuration containing valid inbounds/outbounds.
The adapter injects the TUN descriptor into the root `env` object. Android
must call `VpnService.Builder.establish()`; iOS must obtain the utun descriptor
inside `NEPacketTunnelProvider`. The native engine cannot create the OS VPN
permission or configure system routes itself.

For a full-device tunnel, the platform must configure routes/DNS and protect
or bind Xray's own outbound sockets to the physical network to prevent routing
loops. On Android, use `VpnService.protect(socket)` / the appropriate
underlying-network binding. On iOS, configure the Network Extension's
`NEPacketTunnelNetworkSettings` and ensure the provider's transport sockets
use the physical interface.

## Validation status

Go unit tests cover TUN-FD injection and config preservation. The mobile
artifacts still need to be built on the relevant operating systems and tested
on physical devices; this repository change does not claim that step has
already succeeded.

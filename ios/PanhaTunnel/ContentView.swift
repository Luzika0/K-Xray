import SwiftUI

struct ContentView: View {
    @AppStorage("appearance") private var appearance = "system"
    @State private var selectedEngine = "Xray Core"

    var body: some View {
        NavigationStack {
            VStack(spacing: 22) {
                VStack(spacing: 6) {
                    Text("k-Xray")
                        .font(.largeTitle.bold())
                    Text("Native VPN client")
                        .foregroundStyle(.secondary)
                }
                .frame(maxWidth: .infinity)
                .padding(.top, 24)

                VStack(alignment: .leading, spacing: 12) {
                    Text("CONNECTION")
                        .font(.caption.weight(.semibold))
                        .foregroundStyle(.secondary)
                    Text("Not connected")
                        .font(.title2.bold())
                    Text("VPN engines are not integrated in this foundation build.")
                        .font(.subheadline)
                        .foregroundStyle(.secondary)
                    Button("Connect — coming in a later phase") {}
                        .buttonStyle(.borderedProminent)
                        .disabled(true)
                        .frame(maxWidth: .infinity)
                }
                .padding(20)
                .background(.thinMaterial, in: RoundedRectangle(cornerRadius: 24))

                VStack(alignment: .leading, spacing: 12) {
                    Text("SELECT ENGINE")
                        .font(.caption.weight(.semibold))
                        .foregroundStyle(.secondary)
                    engineRow("Xray Core", detail: "VLESS · VMess · Trojan · Shadowsocks")
                    engineRow("PingTunnel", detail: "ICMP transport · direct TUN mode planned")
                }

                Button("Profiles & import (planned)") {}
                    .buttonStyle(.bordered)
                    .disabled(true)
                    .frame(maxWidth: .infinity)

                Spacer()
                Text("Foundation 0.1.0")
                    .font(.caption2)
                    .foregroundStyle(.tertiary)
            }
            .padding(.horizontal, 22)
            .padding(.bottom, 18)
            .navigationBarHidden(true)
        }
        .preferredColorScheme(appearance == "system" ? nil : (appearance == "dark" ? .dark : .light))
    }

    private func engineRow(_ title: String, detail: String) -> some View {
        Button {
            selectedEngine = title
        } label: {
            HStack {
                VStack(alignment: .leading, spacing: 4) {
                    Text(title).font(.headline)
                    Text(detail).font(.caption).foregroundStyle(.secondary)
                }
                Spacer()
                if selectedEngine == title {
                    Image(systemName: "checkmark.circle.fill")
                        .foregroundStyle(.tint)
                }
            }
            .padding(16)
            .background(selectedEngine == title ? Color.accentColor.opacity(0.12) : Color(uiColor: .secondarySystemBackground),
                        in: RoundedRectangle(cornerRadius: 16))
        }
        .buttonStyle(.plain)
    }
}

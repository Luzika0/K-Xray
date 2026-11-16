# Shared configuration

The import layer will normalize these input families into a versioned internal profile model:

- Xray share links (VLESS, VMess, Trojan, Shadowsocks and supported URI variants).
- Xray JSON configuration.
- QR payloads containing supported share links or JSON.
- Subscription URLs (download only after explicit user action; never log URL credentials).

Validation must be strict, size-limited and performed before saving. Secrets must be stored in platform secure storage, not plain preferences or source control.

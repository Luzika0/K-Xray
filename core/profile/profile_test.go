package profile

import "testing"

func TestParseVLESSReality(t *testing.T) {
	p, err := Parse("vless://123e4567-e89b-12d3-a456-426614174000@example.com:443?type=grpc&security=reality&sni=example.org&fp=chrome&flow=xtls-rprx-vision#office")
	if err != nil { t.Fatal(err) }
	if p.Protocol != "vless" || p.Server != "example.com" || p.Port != 443 || p.Network != "grpc" || p.Security != "reality" || p.Name != "office" { t.Fatalf("unexpected profile: %#v", p) }
}

func TestParseTrojan(t *testing.T) {
	p, err := Parse("trojan://secret@example.com:8443?security=tls&sni=cdn.example.org#edge")
	if err != nil { t.Fatal(err) }
	if p.Password != "secret" || p.Port != 8443 || p.SNI != "cdn.example.org" { t.Fatalf("unexpected profile: %#v", p) }
}

func TestRejectInvalidLinks(t *testing.T) {
	for _, raw := range []string{"", "vless://user@host", "unknown://host:443", "trojan://@host:443"} {
		if _, err := Parse(raw); err == nil { t.Errorf("expected error for %q", raw) }
	}
}

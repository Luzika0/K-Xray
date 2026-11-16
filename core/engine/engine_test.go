package engine

import (
	"encoding/json"
	"testing"
)

func TestWithTunFDAddsEnv(t *testing.T) {
	got, err := withTunFD(42, `{"log":{"loglevel":"warning"},"outbounds":[{"protocol":"freedom"}]}`)
	if err != nil { t.Fatal(err) }
	var config map[string]any
	if err := json.Unmarshal([]byte(got), &config); err != nil { t.Fatal(err) }
	env, ok := config["env"].(map[string]any)
	if !ok || env["xray.tun.fd"] != "42" { t.Fatalf("unexpected env: %#v", config["env"]) }
}

func TestWithTunFDPreservesEnv(t *testing.T) {
	got, err := withTunFD(7, `{"env":{"EXISTING":"value"},"outbounds":[]}`)
	if err != nil { t.Fatal(err) }
	var config map[string]any
	_ = json.Unmarshal([]byte(got), &config)
	env := config["env"].(map[string]any)
	if env["EXISTING"] != "value" || env["xray.tun.fd"] != "7" { t.Fatalf("env not preserved: %#v", env) }
}

func TestWithTunFDRejectsInvalidInput(t *testing.T) {
	for _, tc := range []struct{ fd int; config string }{{-1, `{}`}, {1, `not-json`}, {1, `null`}} {
		if _, err := withTunFD(tc.fd, tc.config); err == nil { t.Fatalf("expected error for fd=%d config=%q", tc.fd, tc.config) }
	}
}

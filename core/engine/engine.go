// Package engine is the single Go runtime shared by the Android and iOS apps.
package engine

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"sync"

	libxray "github.com/xtls/libxray"
)

const apiVersion = 3

var (
	mu sync.Mutex
	running bool
	ErrInvalidConfig = errors.New("invalid Xray configuration")
)

// invokeRequest is libXray's stable, versioned mobile entrypoint.
type invokeRequest struct {
	APIVersion int            `json:"apiVersion"`
	Method     string         `json:"method"`
	Payload    map[string]any `json:"payload,omitempty"`
}
type invokeResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Error   string          `json:"error"`
}

// withTunFD injects the OS-owned TUN descriptor into Xray's root env object.
// libXray removed SetTunFd; the supported contract is xray.tun.fd in config env.
func withTunFD(fd int, configJSON string) (string, error) {
	if fd < 0 { return "", errors.New("TUN file descriptor must be non-negative") }
	var config map[string]any
	if err := json.Unmarshal([]byte(configJSON), &config); err != nil {
		return "", fmt.Errorf("%w: %v", ErrInvalidConfig, err)
	}
	if config == nil { return "", ErrInvalidConfig }
	env, ok := config["env"].(map[string]any)
	if !ok {
		env = map[string]any{}
		config["env"] = env
	}
	env["xray.tun.fd"] = strconv.Itoa(fd)
	out, err := json.Marshal(config)
	if err != nil { return "", fmt.Errorf("encode Xray configuration: %w", err) }
	return string(out), nil
}

func invoke(method string, payload map[string]any) error {
	request, err := json.Marshal(invokeRequest{APIVersion: apiVersion, Method: method, Payload: payload})
	if err != nil { return err }
	raw := libxray.Invoke(string(request))
	var response invokeResponse
	if err := json.Unmarshal([]byte(raw), &response); err != nil {
		return fmt.Errorf("decode libXray response: %w", err)
	}
	if !response.Success {
		if response.Error == "" { response.Error = "libXray operation failed" }
		return errors.New(response.Error)
	}
	return nil
}

// StartXray starts the pinned libXray/Xray-core runtime against the platform TUN.
// The caller retains ownership of tunFD; Xray duplicates the descriptor internally.
func StartXray(tunFD int, configJSON string) error {
	mu.Lock()
	defer mu.Unlock()
	if running { return errors.New("Xray is already running") }
	config, err := withTunFD(tunFD, configJSON)
	if err != nil { return err }
	if err := invoke("runXray", map[string]any{"xrayJson": config}); err != nil { return err }
	running = true
	return nil
}

// StopXray stops the managed Xray instance. It is safe to call more than once.
func StopXray() error {
	mu.Lock()
	defer mu.Unlock()
	if err := invoke("stopXray", nil); err != nil { return err }
	running = false
	return nil
}

// XrayVersion reports the embedded upstream core version.
func XrayVersion() string {
	raw, _ := json.Marshal(invokeRequest{APIVersion: apiVersion, Method: "xrayVersion"})
	var response struct {
		Success bool `json:"success"`
		Data struct { Version string `json:"version"` } `json:"data"`
	}
	_ = json.Unmarshal([]byte(libxray.Invoke(string(raw))), &response)
	if !response.Success { return "" }
	return response.Data.Version
}

// IsXrayRunning reports this adapter's managed lifecycle state.
func IsXrayRunning() bool {
	mu.Lock()
	defer mu.Unlock()
	return running
}

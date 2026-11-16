// Package profile parses common Xray share links into a normalized model.
package profile

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type Profile struct {
	Protocol string `json:"protocol"`
	Name string `json:"name"`
	Server string `json:"server"`
	Port uint16 `json:"port"`
	UUID string `json:"uuid,omitempty"`
	Password string `json:"password,omitempty"`
	Method string `json:"method,omitempty"`
	Network string `json:"network,omitempty"`
	Security string `json:"security,omitempty"`
	SNI string `json:"sni,omitempty"`
	ALPN []string `json:"alpn,omitempty"`
	Flow string `json:"flow,omitempty"`
	Path string `json:"path,omitempty"`
	Host string `json:"host,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}

func Parse(raw string) (Profile, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" { return Profile{}, errors.New("empty profile link") }
	u, err := url.Parse(raw)
	if err != nil { return Profile{}, fmt.Errorf("parse profile URI: %w", err) }
	switch strings.ToLower(u.Scheme) {
	case "vless", "trojan": return parseStandard(u)
	case "vmess": return parseVMess(u)
	case "ss", "shadowsocks": return parseShadowsocks(u)
	default: return Profile{}, fmt.Errorf("unsupported profile scheme %q", u.Scheme)
	}
}

func parseStandard(u *url.URL) (Profile, error) {
	host, port, err := hostPort(u); if err != nil { return Profile{}, err }
	p := Profile{Protocol:strings.ToLower(u.Scheme), Server:host, Port:port, Name:fragment(u)}
	q := u.Query()
	p.Network = first(q.Get("type"), "tcp")
	p.Security = q.Get("security")
	p.SNI = first(q.Get("sni"), q.Get("serverName"))
	p.Flow, p.Fingerprint = q.Get("flow"), q.Get("fp")
	p.Path, p.Host = q.Get("path"), first(q.Get("host"), q.Get("authority"))
	p.ALPN = splitCSV(q.Get("alpn"))
	username := ""
	if u.User != nil { username = u.User.Username() }
	if p.Protocol == "vless" {
		p.UUID = username
		if p.UUID == "" { return Profile{}, errors.New("vless link is missing UUID") }
	} else {
		if u.User != nil { p.Password, _ = u.User.Password() }
		if p.Password == "" { p.Password = username }
		if p.Password == "" { return Profile{}, errors.New("trojan link is missing password") }
	}
	return p, nil
}

func parseVMess(u *url.URL) (Profile, error) {
	encoded := strings.TrimPrefix(u.Opaque, "//")
	if encoded == "" { encoded = strings.TrimPrefix(u.Path, "/") }
	if encoded == "" { encoded = u.Host }
	data, err := decodeBase64(encoded)
	if err != nil { return Profile{}, fmt.Errorf("decode vmess payload: %w", err) }
	var v struct {
		Add string `json:"add"`
		Port json.RawMessage `json:"port"`
		ID string `json:"id"`
		PS string `json:"ps"`
		Net string `json:"net"`
		TLS string `json:"tls"`
		SNI string `json:"sni"`
		Host string `json:"host"`
		Path string `json:"path"`
		FP string `json:"fp"`
		ALPN string `json:"alpn"`
	}
	if err := json.Unmarshal(data, &v); err != nil { return Profile{}, fmt.Errorf("decode vmess JSON: %w", err) }
	port, err := strconv.ParseUint(strings.Trim(string(v.Port), `"`), 10, 16)
	if err != nil || port == 0 { return Profile{}, errors.New("vmess link has invalid port") }
	if v.Add == "" || v.ID == "" { return Profile{}, errors.New("vmess link requires add and id") }
	if strings.ContainsAny(v.Add, "/?#@") { return Profile{}, errors.New("vmess link has invalid server host") }
	security := ""; if strings.EqualFold(v.TLS, "tls") { security = "tls" }
	return Profile{Protocol:"vmess", Name:v.PS, Server:v.Add, Port:uint16(port), UUID:v.ID, Network:first(v.Net,"tcp"), Security:security, SNI:v.SNI, Host:v.Host, Path:v.Path, Fingerprint:v.FP, ALPN:splitCSV(v.ALPN)}, nil
}

func parseShadowsocks(u *url.URL) (Profile, error) {
	raw := strings.TrimPrefix(u.Opaque, "//")
	if raw == "" {
		if u.User != nil {
			password, _ := u.User.Password()
			raw = u.User.Username() + ":" + password + "@" + u.Host
		} else { raw = strings.TrimPrefix(u.Host+u.Path, "/") }
	}
	if decoded, err := decodeBase64(raw); err == nil && strings.Contains(string(decoded), "@") { raw = string(decoded) }
	parts := strings.SplitN(raw, "@", 2)
	if len(parts) != 2 { return Profile{}, errors.New("shadowsocks link must contain method:password@host:port") }
	cred := strings.SplitN(parts[0], ":", 2)
	if len(cred) != 2 || cred[0] == "" || cred[1] == "" { return Profile{}, errors.New("shadowsocks link has invalid credentials") }
	host, portText, err := net.SplitHostPort(parts[1])
	if err != nil { return Profile{}, fmt.Errorf("shadowsocks server: %w", err) }
	port, err := strconv.ParseUint(portText,10,16)
	if err != nil || port == 0 { return Profile{}, errors.New("shadowsocks link has invalid port") }
	return Profile{Protocol:"shadowsocks", Name:fragment(u), Server:host, Port:uint16(port), Method:cred[0], Password:cred[1]}, nil
}

func hostPort(u *url.URL) (string, uint16, error) {
	host, portText := u.Hostname(), u.Port()
	if host == "" || portText == "" { return "", 0, errors.New("profile link requires server host and port") }
	port, err := strconv.ParseUint(portText,10,16)
	if err != nil || port == 0 { return "", 0, errors.New("profile link has invalid port") }
	return host, uint16(port), nil
}
func decodeBase64(s string) ([]byte,error) {
	s = strings.TrimSpace(s)
	for _, enc := range []*base64.Encoding{base64.RawStdEncoding,base64.StdEncoding,base64.RawURLEncoding,base64.URLEncoding} {
		if b,err:=enc.DecodeString(s); err==nil { return b,nil }
	}
	return nil, errors.New("invalid base64")
}
func fragment(u *url.URL) string { s,_:=url.PathUnescape(u.Fragment); return s }
func first(a,b string) string { if a!="" { return a }; return b }
func splitCSV(s string) []string { if s=="" { return nil }; out:=strings.Split(s,","); for i:=range out { out[i]=strings.TrimSpace(out[i]) }; return out }

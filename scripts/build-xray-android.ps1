$ErrorActionPreference = "Stop"
$Root = Split-Path -Parent $PSScriptRoot
$Core = Join-Path $Root "core"
$OutDir = Join-Path $Root "android/app/libs"
New-Item -ItemType Directory -Force -Path $OutDir | Out-Null

if (-not (Get-Command go -ErrorAction SilentlyContinue)) { throw "Go 1.27.3+ is required." }
$GoVersion = (go env GOVERSION)
if ($GoVersion -notmatch '^go1\.27\.') { throw "Expected Go 1.27.x, found $GoVersion" }
if (-not (Get-Command gomobile -ErrorAction SilentlyContinue)) { go install golang.org/x/mobile/cmd/gomobile@latest }
gomobile init
Push-Location $Core
try {
    go mod download
    gomobile bind -target=android -androidapi 26 -o (Join-Path $OutDir "kxray-engine.aar") ./engine
} finally { Pop-Location }
Write-Host "Android Xray engine AAR created at android/app/libs/kxray-engine.aar"

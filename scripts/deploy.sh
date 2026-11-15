#!/usr/bin/env bash
# Builds the binary and installs it as the vx-that-for-me systemd service.
#   scripts/deploy.sh
# The first run only writes the config and stops; fill it in and run again to start.
set -euo pipefail
cd "$(dirname "$0")/.."

app=vx-that-for-me
env_file=/etc/$app/$app.env

mise exec -- go test ./...
CGO_ENABLED=0 mise exec -- go build -trimpath -ldflags=-s -o "build/$app" .

sudo install -d -m 755 "/etc/$app"
if ! sudo test -f "$env_file"; then
  sudo install -m 600 "deploy/$app.env.example" "$env_file"
  echo "wrote $env_file; fill it in, then run scripts/deploy.sh again to start the service"
  exit 0
fi

# Same check as main.go. An empty token would crash-loop the service.
required="TELEGRAM_TOKEN"
missing=$(sudo awk -F= -v required="$required" '
  BEGIN { n = split(required, keys, " "); for (i = 1; i <= n; i++) need[keys[i]] = 1 }
  $1 in need && $2 != "" { delete need[$1] }
  END { for (k in need) printf "%s ", k }
' "$env_file")
if [[ -n $missing ]]; then
  sudo systemctl disable --now "$app" 2>/dev/null || true
  echo "$env_file has no value for: $missing" >&2
  echo "fill it in with: sudoedit $env_file" >&2
  exit 1
fi

sudo install -m 755 "build/$app" "/usr/local/bin/$app"
sudo restorecon -F "/usr/local/bin/$app" 2>/dev/null || true
sudo install -m 644 "deploy/$app.service" "/etc/systemd/system/$app.service"
sudo systemctl daemon-reload
sudo systemctl enable "$app"
sudo systemctl restart "$app"
systemctl --no-pager status "$app" | head -5

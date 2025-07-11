#!/bin/bash

set -e

red='\033[0;31m'
green='\033[0;32m'
blue='\033[0;34m'
yellow='\033[0;33m'
plain='\033[0m'

# Check root
[[ $EUID -ne 0 ]] && echo -e "${red}Fatal error: ${plain} Please run this script with root privilege \n " && exit 1

# Check OS
if [[ -f /etc/os-release ]]; then
    source /etc/os-release
    release=$ID
else
    echo "Failed to check the system OS, please contact the author!" >&2
    exit 1
fi

echo "The OS release is: $release"

# Install dependencies
apt-get update && apt-get install -y wget curl tar tzdata git build-essential

# Install Go if not present
if ! command -v go &> /dev/null; then
    echo -e "${yellow}Go not found. Installing Go 1.24...${plain}"
    wget -q https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
    rm go1.24.4.linux-amd64.tar.gz
else
    echo -e "${green}Go is already installed.${plain}"
fi

# Build the project
cd /usr/local/
if [[ -d x-ui ]]; then
    rm -rf x-ui
fi
cp -r $OLDPWD /usr/local/x-ui
cd /usr/local/x-ui

go mod tidy
go build -o /usr/local/x-ui/x-ui main.go
chmod +x /usr/local/x-ui/x-ui

# Create systemd service
cat <<EOF > /etc/systemd/system/x-ui.service
[Unit]
Description=x-ui Service
After=network.target

[Service]
Type=simple
ExecStart=/usr/local/x-ui/x-ui
WorkingDirectory=/usr/local/x-ui
Restart=always
RestartSec=3
User=root

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable x-ui
systemctl restart x-ui

echo -e "${green}x-ui installation from source finished and service started.${plain}"

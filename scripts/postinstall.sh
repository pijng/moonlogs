#!/bin/bash
set -e

SERVICE_NAME="moonlogs"
APP_BINARY="/usr/bin/moonlogs"
SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
INSTALL_USER="${SUDO_USER:-$(whoami)}"
INSTALL_USER_GROUP=$(id -gn "${SUDO_USER:-$(whoami)}")

# Check if systemd is available
if command -v systemctl >/dev/null 2>&1; then
  # Create the service unit file
  cat <<EOF >"${SERVICE_FILE}"
[Unit]
Description=moonlogs
After=network.target

[Service]
ExecStart=${APP_BINARY}
Restart=always
User=${INSTALL_USER}
Group=${INSTALL_USER_GROUP}

[Install]
WantedBy=multi-user.target
EOF

  # Enable and start the service
  systemctl enable "${SERVICE_NAME}.service"
  systemctl start "${SERVICE_NAME}.service"

  echo "Service installed and started successfully."
else
  echo "Systemd not available. Please manually set up your service."
fi

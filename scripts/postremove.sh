#!/bin/sh

SERVICE_FILE="/etc/systemd/system/moonlogs.service"

remove() {
    printf "\033[32m removing moonlogs.service\033[0m\n"

    rm -f "${SERVICE_FILE}"
}

purge() {
    printf "\033[32m purging moonlogs.service\033[0m\n"

    rm -f "${SERVICE_FILE}"
}

upgrade() {
    printf "\033[32m Reloading systemd daemon\033[0m\n"
    systemctl daemon-reload
}

echo "$@"

action="$1"

case "$action" in
  "0" | "remove")
    remove
    ;;
  "1" | "upgrade")
    upgrade
    ;;
  "purge")
    purge
    ;;
  *)
    printf "\033[32m Alpine\033[0m"
    remove
    ;;
esac

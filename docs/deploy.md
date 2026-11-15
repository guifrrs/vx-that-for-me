# Deploy

vx-that-for-me runs on the home server as the `vx-that-for-me` systemd service. It uses long polling, so it only makes outbound requests to Telegram and needs no open port.

## Install or update

```sh
make deploy
```

The script runs the tests, builds a static binary and installs it. It asks for sudo.

| What | Where |
|---|---|
| Binary | `/usr/local/bin/vx-that-for-me` |
| Config | `/etc/vx-that-for-me/vx-that-for-me.env` (mode 600) |
| Unit | `/etc/systemd/system/vx-that-for-me.service` |

The first run only writes the config and stops. Put `TELEGRAM_TOKEN` in it with `sudoedit`, then run `make deploy` again to start the service. While `TELEGRAM_TOKEN` is empty, `make deploy` stops and disables the service and names the empty variable.

Only one copy may poll with a token at a time. A second copy gets `409 Conflict` from Telegram.

## Check

```sh
make status
make logs        # "Received message from @..." for each fixed link
make security    # sandbox exposure score
```

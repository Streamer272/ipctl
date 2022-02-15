# ipctl

Listen to IP change and change your DNS' records dynamically

---

## Features
0. Lightweight - can easily run even on Raspberry Pi!
0. Easy to set up - `ipctl` can set up basic configuration and `systemctl` service itself, no editing needed
0. You don't have to worry about dynamic IP ever again!

## Installation
Check out [Releases](https://github.com/Streamer272/ipctl/releases) for latest versions.

Or, you can download it directly with `Go`
```
go get github.com/Streamer272/ipctl@latest
```
But watch out, you have to add `$HOME/go/bin/` to root path

### Build from source

#### Requirements:

0. Go 17+

Clone the repository with `git`
```
git clone https://github.com/Streamer272/ipctl.git && cd ipctl
```
Install `ipctl` with `install.sh` script
```
bash ./install.sh
```

---

## Quick start

Right after installing, you need to generate config file and `systemctl` service with
```
ipctl config init && ipctl service init
```
Make sure you run this command with `root` privileges

Config file is located in `/etc/ipctl/config`, you can see all current config files in use using
```
ipctl config
```

Service file will be stored in `/lib/systemd/system/ipctl.service`, you can find this path using
```
ipctl service
```

To make `ipctl` start every time you start your computer in background, you need to enable and start it using systemctl
```
systemctl enable --now ipctl
```

### How does it work

Every `interval` (located in config file) milliseconds, a request on `https://api.my-ip.io/ip.json` is made, finding out current IP address. If IP address has changed, `callback_file` will be called with `bash` (be sure to put `#!/usr/bin/bash` on the first line of your callback file). Here, new IP will be provided as environmental variable `IP` (in `python`, you can read this value with `os.getenv("IP")`).
Mind that this service only runs if you are connected to the internet, so you don't have to worry about not having connection in your callback file.


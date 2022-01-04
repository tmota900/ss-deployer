# Super Simple Deployer

## Concept

the vision of this tool is to take the maximum power over github webhooks with the minimum effort, simple.

### How it works

create a deploy.sh file that will be run when an valid request comes from github webhook system.

...and dont forget 😉
```
chmod +x deploy.sh
```

### Install in server (Systemd)

this will help setup a service for the aplication in the background and reload in failure

download to opt dir
```
# cd /opt
```

create file:
```
# vim /etc/systemd/system/ss-deployer.service
```

content:
```
[Unit]
Description=Supper Simple Deployer
ConditionPathExists=/opt
After=network.target
[Service]
Type=simple
User=root
Group=root
WorkingDirectory=/opt
ExecStart=/opt/ss-deployer deployer
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=ss-deployer
[Install]
WantedBy=multi-user.target
```

start service:
```
# systemctl daemon-reload
# service ss-deployer start
```

#### Flags

```
Usage:
   deployer [flags]

Flags:
  -h, --help            help for deployer
  -p, --port string     Target port listner (default "1337")
  -s, --secret string   Configured secret
```
# WEBSITE-HEALTH-CHECKER

A simple cli tool to check health status of any website.

### How to build

``go build .``

This command will create an executable called `website-health-checker`

Run the following command
``$user> ./website-health-checker -h``

This command describes the functionality

```
NAME:
   website-health-checker - Tool to check whether a website is down

USAGE:
   website-health-checker [-d | --domain <domain>] [-p | --port <port>] [-t | --timeout <seconds>]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --domain DOMAIN, -d DOMAIN     pass the DOMAIN of site which needs to be checked
   --port PORT, -p PORT           specifies which PORT to use (default: "80")
   --timeout TIMEOUT, -t TIMEOUT  can be used to specify max timeout in seconds TIMEOUT (default: 5)
   --help, -h                     show help
```

##### Usage example
``website-health-checker -d youtube.com -p 80 -t 10``

- `--domain` or `-d` flag is used to pass the website's domain which needs to be checked
- `--port` or `-p` which is an optional flag can be used to pass any port other `:80` to connect to the domain
- `--timeout` or `t` which is also an optional flag can be used to set timeout for checking the site, default timeout is **5 seconds**

successful result would look like following:
```
[Up] youtube.com is reachable,
From: 192.0.0.1:62220
To: 142.250.183.14:80
```
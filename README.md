# Load Balancer
A simple load balancer / reverse proxy written in GoLang.

Forked from [](https://github.com/d12/Load-Balancer-Golang), modified by tobychui for my own network load-balacning purpose

## Features
- Acts as a reverse proxy, sitting between clients and web servers.
- Distributes load between different web servers based on number of connections currently active
- Falls back on other web-servers if one goes down
- TLS support over proxy and allow non "https" website to be proxy like a "https" one



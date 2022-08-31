---
sidebar_position: 2
---

# DWS DNS Server

DNS Server starts with a node process. By default it listen for UDP/1053, this parameter can be changed in client.toml in config path (`~/.deweb/config/config.toml`).
Parameter: `dns.lport`

DNS server works in coroutine and making requests to main process via GRPC to port TCP/26657 (`rpc.laddr`).
It has 10-seconds cache to prevent constant requests to GRPC server.

To listen on port UDP/53, node process must be executed as a root user. To prevent the whole node running by root user we have a special command which will run only DNS server:

```
sudo dewebd q deweb run-dns-server 53 1.1.1.1:53
```

Currently DNS Server supports only A and MX records.

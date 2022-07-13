Predefined denom `domains` added to genesis.

Set test data for DNS:
```
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
BasicDataWithMX='{"records": [{"type": "A","values": ["192.168.1.10"]},{"type": "MX","values": ["mx.bob.alice.deweb."]}]}'
```

First Bob mint NFT for TLD `deweb`: 
```
./dewebd tx domain register deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Prolongation request send via register command. Allowed only for domain owner. Other user will receive message that 
the domain already registered.

Then create domain `test.deweb`:
```
./dewebd tx domain register test.deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

In Alice will try to register domain in Bob's zone she will receive an error `parent domain check error: domain deweb does not belong to this user`:
```
./dewebd tx domain register alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

But Bob can register domain for Alice (her address `deweb1hl0d8mxd06ph43sfyenl3pu2ecrf7q64rgklz3`):
```
./dewebd tx domain register aliced.deweb --data="$BasicData" --recipient=deweb1hl0d8mxd06ph43sfyenl3pu2ecrf7q64rgklz3 --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then Alice can register domain `www.alice.deweb` because she is the owner of `alice.deweb`:
```
./dewebd tx domain register www.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Alice can register domain for Bob in her zone and then transfer this NFT to Bob, so he will become an owner of `bob.alice.deweb`:
```
./dewebd tx domain register bob.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```
2-step transfer. First domain owner Alice send transaction with expecter domain received and price. Only selected receiver
Bob (his address is `deweb1g5uwu39petj0w32xluz5prqfh78qemf3hjkndu`)can buy the domain. If no receiver determined, anyone can buy this domain.
```
./dewebd tx domain transfer bob.aliced.deweb --recipient=deweb1g5uwu39petj0w32xluz5prqfh78qemf3hjkndu --price=10000 --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
./dewebd tx domain transfer bob.aliced.deweb --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

To cancel created transfer:
```
./dewebd tx domain transfer bob.aliced.deweb --cancel=true --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

After transfer Alice cannot edit Bob's domain, but he can change DNS records: 
```
./dewebd tx domain edit bob.aliced.deweb --data="$BasicDataWithMX" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

## Module parameters 

Parameters to change:

- DomainPrice
- DomainExpiration
- DomainOwnerProlongation
- BlockedTLDs

Submit proposal command:
```
./dewebd tx gov submit-proposal param-change dns_proposal.json --from alice --chain-id deweb-testnet-0 --gas 2000000
```



## DNS Server

DNS Server starts with node process. Listen for UDP/1053, this parameter can be changed in client.toml in config path (~/.deweb/config/config.toml).
Parameter: dns.lport

DNS server works in goroutine and requests main process via GRPC to port TCP/26657 (rpc.laddr).
It has 10-seconds cache to prevent often requests to GRPC server.

To listen on port UDP/53 process must be executed from root user. To prevent all node running from root created command:
```
sudo ./dewebd q deweb run-dns-server 53
```

Now supported only A and MX record types

Predefined denom `domains` added to genesis.

Set test data for DNS:
```
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
BasicDataWithSubPrice='{"records":[{"type": "A","values":["192.168.1.10"]}],"sub_domains_sale": true,"sub_domains_sale_price": 10000}'
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

Bob can edit domain and allow anyone subdomains registration. He will receive payments for these domains:
```
./dewebd tx domain edit deweb --data="BasicDataWithSubPrice" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then alice can register subdomain and pay to BOB:
```
./dewebd tx domain register newalice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

But Bob can register domain for Alice (her address `deweb1x6s67chad4p2rznmclskw7xr3qfppfhjkqs3ee`):
```
./dewebd tx domain register aliced.deweb --data="$BasicData" --recipient=deweb1x6s67chad4p2rznmclskw7xr3qfppfhjkqs3ee --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then Alice can register domain `www.alice.deweb` because she is the owner of `alice.deweb`:
```
./dewebd tx domain register www.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Alice can register domain for Bob in her zone and then transfer this NFT to Bob, so he will become an owner of `bob.alice.deweb`:
```
./dewebd tx domain register bob.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```
2-step transfer. First domain owner Alice send transaction with expecter domain receiver and price. Only selected receiver
Bob (his address is `deweb138sq5w8yxsdzgvs7lse5j3dtr3e2t5z08cw8aj`)can buy the domain. If no receiver determined, anyone can buy this domain.
```
./dewebd tx domain transfer bob.aliced.deweb --recipient=deweb138sq5w8yxsdzgvs7lse5j3dtr3e2t5z08cw8aj --price=10000 --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
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

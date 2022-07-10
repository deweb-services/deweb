Predefined denom `domains` added to genesis.

Set test data for DNS:
```
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
BasicDataWithMX='{"records": [{"type": "A","values": ["192.168.1.10"]},{"type": "MX","values": ["mx.bob.alice.deweb."]}]}'
```

First Bob mint NFT for TLD `deweb`: 
```
./dewebd tx nft register deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Prolongation request send via register command. Allowed only for domain owner. Other user will receive message that 
the domain already registered.

Then create domain `test.deweb`:
```
./dewebd tx nft register test.deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

In Alice will try to register domain in Bob's zone she will receive an error `parent domain check error: domain deweb does not belong to this user`:
```
./dewebd tx nft register alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

But Bob can register domain for Alice (her address `deweb1gq20zy0xryus2zhsy0e8s33k84n0unwgwlpp8d`):
```
./dewebd tx nft register alice.deweb --data="$BasicData" --recipient=deweb1gq20zy0xryus2zhsy0e8s33k84n0unwgwlpp8d --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then Alice can register domain `www.alice.deweb` because she is the owner of `alice.deweb`:
```
./dewebd tx nft register www.alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Alice can register domain for Bob in her zone and then transfer this NFT to Bob, so he will become an owner of `bob.alice.deweb`:
```
./dewebd tx nft register bob.alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```
2-step transfer. First domain owner Bob send transaction with expecter domain received and price. Only selected receiver
can buy the domain. If no receiver determined, anyone can buy this domain.
```
./dewebd tx nft transfer bob.alice.deweb --recipient=deweb179xaeytshtdsjqw527ejrehdnfld3mmtz85cyq --price=10000 --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
./dewebd tx nft transfer bob.alice.deweb --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

To cacnel created transfer:
```
./dewebd tx nft transfer bob.alice.deweb --cancel=true --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

After transfer Alice cannot edit Bob's domain, but he can change DNS records: 
```
./dewebd tx nft edit domains `bob.alice.deweb` --data="$BasicDataWithMX" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
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

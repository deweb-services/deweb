First specify the nft Denom (nft classification) for domains. Only denom `domains` allowed. Here we set that everyone can mint NFT of this classification. And allows edit.
```
./dewebd tx nft issue domains --from alice --name=dns_registry --symbol=DwDNS --mint-restricted=false --update-restricted=false --description="DNS NFT" --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Set test data for DNS:
```
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
BasicDataWithMX='{"records": [{"type": "A","values": ["192.168.1.10"]},{"type": "MX","values": ["mx.bob.alice.deweb."]}]}'
```

First Bob mint NFT for TLD `deweb`.: 
```
./dewebd tx nft mint domains deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then create domain `test.deweb`:
```
./dewebd tx nft mint domains test.deweb --data="$BasicData" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

In Alice will try to register domain in Bob's zone she will receive an error `parent domain check error: domain deweb does not belong to this user`:
```
./dewebd tx nft mint domains alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

But Bob can register domain for Alice (her address `deweb132p7gyc32qzs0wex8wstjgap0x2af9vp5r992n`):
```
./dewebd tx nft mint domains alice.deweb --data="$BasicData" --recipient=deweb132p7gyc32qzs0wex8wstjgap0x2af9vp5r992n --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Then Alice can register domain `www.alice.deweb` because she is the owner of `alice.deweb`:
```
./dewebd tx nft mint domains www.alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

Alice can register domain for Bob in her zone and then transfer this NFT to Bob, so he will become an owner of `bob.alice.deweb`:
```
./dewebd tx nft mint domains bob.alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
./dewebd tx nft transfer deweb1982frnx7muvcf849zlun6wtkr3p7f0fd7eep08 domains bob.alice.deweb --from alice --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
```

After transfer Alice cannot edit Bob's domain, but he can change DNS records: 
```
./dewebd tx nft edit domains bob.alice.deweb --data="$BasicDataWithMX" --from bob --chain-id deweb-testnet-0 --gas 2000000 --output json -b block
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

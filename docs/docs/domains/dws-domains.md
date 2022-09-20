---
sidebar_position: 1
---

# DWS Domains

Decentralized Web Services Blockchain has a DWS Domains module. User can buy domain of any level.

Initial Top-level domain (TLD) price is defined in genesis as 100 DWS and can be changed via the governance proposal.

Each domain ownership is 2 years from buy time. After 2 years the domain owner has the 2 weeks window to extend the ownership by pay the full price (100 DWS).

For example Alice can buy domain `hello/` for 100 DWS. Then Alice can say that anyone can buy a subdomain `*.hello/` for 10 DWS. Any subdomain which Alice will sell will belong to the buyer for 2 years.

For example Bob will buy domain `bob.hello` for 10 DWS. 10 DWS will be transfered to Alice and Bob will get `bob.hello` domain ownership for 2 years. Alice will not have any permission to get this domain back from Bob. After 2 years Bob will have a 2 weeks window to extend the ownership for 10 DWS, which will be transfered to Alice. If Alice changed the price of subdomains to 20 DWS, Bob will still pay 10 DWS per 2 years, as it was the initial price.

## Buy domain

If we want to buy a domain, we can set DNS records:

```
BasicDataEmpty='{}'
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
BasicDataWithSubPrice='{"records":[{"type": "A","values":["192.168.1.10"]}],"sub_domains_sale": true,"sub_domains_sale_price": 10000}'
BasicDataWithMX='{"records": [{"type": "A","values": ["192.168.1.10"]},{"type": "MX","values": ["mx.bob.alice.deweb."]}]}'
```

If we want to start selling the subdomains, we should set `"sub_domains_sale": true` and `"sub_domains_sale_price": 10000`, where `sub_domains_sale_price` is in `udws` denom.

First we can mint a TLD `deweb`:

```
# how to:
# dewebd tx domain register <domainName> --data="$BasicData" --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
dewebd tx domain register deweb --data="$BasicData" --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

## Extend domain ownership (Prolongation)

Prolongation request is sent via the register command. It is allowed only for the domain owner. Other users will receive the message that
the domain is already registered.

Let's create a domain `test.deweb`:

```
dewebd tx domain register test.deweb --data="$BasicData" --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

If Alice will try to register domain in Bob's zone she will receive an error `parent domain check error: domain deweb does not belong to this user`:

```
dewebd tx domain register alice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

## Edit domain settings

Bob can edit domain and allow anyone to buy subdomains of the domain he own. He will receive payments for these domains:

```
BasicDataWithSubPrice='{"records":[{"type": "A","values":["192.168.1.10"]}],"sub_domains_sale": true,"sub_domains_sale_price": 10000}'
dewebd tx domain edit deweb --data="BasicDataWithSubPrice" --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

## Subdomain ownership examples

Then Alice can register subdomain and pay the fee to Bob:

```
BasicData='{"records": [{"type": "A","values": ["192.168.1.10"]}]}'
dewebd tx domain register newalice.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

But Bob also can register subdomain for Alice (her address `deweb1x6s67chad4p2rznmclskw7xr3qfppfhjkqs3ee`):

```
dewebd tx domain register alice.deweb --data="$BasicData" --recipient=deweb1x6s67chad4p2rznmclskw7xr3qfppfhjkqs3ee --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

Then Alice can register domain `www.alice.deweb` because she is the owner of `alice.deweb`:

```
dewebd tx domain register www.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

Alice can register domain for Bob in her zone and then transfer this Domain to Bob, so he will become an owner of `bob.alice.deweb`:

```
dewebd tx domain register bob.aliced.deweb --data="$BasicData" --from alice --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

## 2-step transfer

This is used to sell the domain to exact deweb address for exact price.

Firstly the domain owner Alice sends the transaction with expected domain receiver and price. Only selected receiver Bob (his address is `deweb138sq5w8yxsdzgvs7lse5j3dtr3e2t5z08cw8aj`) can buy the domain. If no receiver determined, anyone can buy this domain.

```
dewebd tx domain transfer bob.aliced.deweb --recipient=deweb138sq5w8yxsdzgvs7lse5j3dtr3e2t5z08cw8aj --price=10000 --from alice --chain-id deweb-2 --gas 2000000 --output json -b block
dewebd tx domain transfer bob.aliced.deweb --from bob --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

To cancel created transfer:

```
dewebd tx domain transfer bob.aliced.deweb --cancel=true --from alice --chain-id deweb-testnet-sirius --gas 2000000 --output json -b block
```

## Module parameters

Parameters that can be changed via Government proposal:

- DomainPrice
- DomainExpiration
- DomainOwnerProlongation
- BlockedTLDs

```json title="dns_proposal.json"
{
  "title": "Increase the lease time for domains",
  "description": "Set lease time to 60 minutes for domains",
  "changes": [
    {
      "subspace": "nft",
      "key": "DomainExpiration",
      "value": "60"
    }
  ],
  "deposit": "10000000udws"
}
```

Submit proposal:

```
dewebd tx gov submit-proposal param-change dns_proposal.json --from alice --chain-id deweb-testnet-sirius --gas 2000000
```

# DWS commands notes

## Initialize chain
```
starport scaffold chain gitlab.com/deweb-services/deweb-chain
```

Create messages
```
starport scaffold message save_wallet private_key chain --module deweb
starport scaffold message delete_wallet address --module deweb

starport scaffold message connect_chain chain address --module deweb
starport scaffold message delete_chain_connect chain address --module deweb
```

Create queries
```
starport scaffold query filter_user_key_records owner address chain deleted limit offset --response uuids --module deweb

starport scaffold query filter_chain_mappings_records owner address chain deleted limit offset --response uuids --module deweb
```

## Storage description

Using storage keys:
- RecordsKey: _records-_ - to store external wallet records. Record ID created from <dws_address>_<external_address> 
- UsersRecords: _users-_ - store mappings between users (DWS accounts) and owned wallets records: list of record keys
- ConnectChainRecords: _connections-_ - to store chain mappings used for bridges. Key: <dws_address>_<chain>_<external_address>
- UserConnectChainRecords: _users_connections-_ - store mapping between chai records and DWS owner. List of chain records keys

Stored wallet record fields:
- string address
- string encrypted_key
- string chain
- bool deleted

Chain mapping record:
- string ext_address
- string chain
- bool deleted

## API
```
/deweb/external_wallets/v1beta1/list
?owner=deweb1z3z7lgfawp22mktkhxav5nr8g35m20ddpc4txj // owner address (DWS address), required
&address=deweb1z3z7lgfawp22mktkhxav5nr8g35m20ddpc4txj // external chai address. Optional field
&chain=cosmos // optional, default: filter results by chain
&deleted=true // optional, default: false
&limit=1000 // optional, default: 10
&offset=1 // optional, default: 1
```

Response:
```
{
    "records": [
        {
            "owner": "",
            "address": "",
            "encrypted_key": "",
            "chain": "",
            "deleted": true
        }
    ]
}
```

## Commands
Create record (transaction):
```
dewebd tx deweb save-wallet [address] [encrypted_key] [chain] [flags]
```
- address - address for which private key to store
- encrypted_key - message with key to store
- chain - from which chain is that address
- flags - cosmos SDK flags

After execution will set owner for this record - who sent transaction.

Mark address removed (transaction):
```
dewebd tx deweb delete-wallet [address] [flags]
```
- address - address for which private key to store
  On execution check that this record owner sent transaction

Filter for records:
```
dewebd query deweb filter-user-wallet-records [owner] [address] [chain] [deleted] [limit] [offset] [flags]
```
Mandatory attribute os only owner. To leave string attribute empty set "" on position
- owner - address of owner for which to list records
- address - get record for selected address
- chain - filter addresses from chain
- deleted - add removed records to list (default: false)
- limit - limit for result
- offset - offset of results list

## Test cases from console

**Initial configuration for Cosmos SDK located in folder: cosmos_sdk_config in current repository**

You must scaffold chain and then replace files in config and data home folder (by default located in ~/.deweb).
And then replace generate files of blockchain with files in this repo.
Or you can try to clone this repo, build binary and then set config hole folder via flag --home

First build binary:
```
starport chain build
```
Initialize blockchain
```
starport chain init
```

# Test commands

## Working with wallets

Created account "alice" with address "deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs"

Created account "bob" with address "deweb1w8s7fujk76eqvkm4uxk28hcd5lv4t9u9h2z865"

Creating test records:
```
~/go/bin/dewebd tx deweb save-wallet cosmosaddress private_cosmos_key cosmos --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-wallet siaaddress private_sia_key sia --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-wallet bobaddress private_bob_key sia --from bob --chain-id deweb
```

Get records by owner address
```
~/go/bin/dewebd q deweb filter-user-wallet-records deweb1w8s7fujk76eqvkm4uxk28hcd5lv4t9u9h2z865
records:
- address: bobaddress
  chain: sia
  deleted: false
  encrypted_key: private_bob_key
  owner: deweb1w8s7fujk76eqvkm4uxk28hcd5lv4t9u9h2z865

~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
- address: siaaddress
  chain: sia
  deleted: false
  encrypted_key: private_sia_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
```

Get records for chain cosmos:
```
~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs "" cosmos
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
```

Remove records:
```
~/go/bin/dewebd tx deweb delete-wallet siaaddress --from alice --chain-id deweb
```

Check that records marked as removed:
```
~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs

~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs "" "" true
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
- address: siaaddress
  chain: sia
  deleted: true
  encrypted_key: private_sia_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
```

Limits and offsets:
```
~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs "" "" true 1
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs

~/go/bin/dewebd q deweb filter-user-wallet-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs "" "" true 1 1
records:
- address: siaaddress
  chain: sia
  deleted: true
  encrypted_key: private_sia_key
  owner: deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs
  ```

## Working with chain mappings

### Create records
```
~/go/bin/dewebd tx deweb connect-chain cosmos cosmosaddress --from alice --chain-id deweb
~/go/bin/dewebd tx deweb connect-chain sia siaaddress --from alice --chain-id deweb
```

### Get stored records

```
~/go/bin/dewebd q deweb filter-chain-mappings-records deweb1ln9z5ymgtehwjvf2y7wmv9ddvag5fjqh8p82rs "" "" false 100 0
```

### Delete mapping record

```
~/go/bin/ dewebd tx deweb delete-chain-connect sia siaaddress --from alice --chain-id deweb
```
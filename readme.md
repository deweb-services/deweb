# deweb
**deweb** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Initialize chain
```
starport scaffold chain gitlab.com/deweb-services/deweb-chain
```

Create messages
```
starport scaffold message save_wallet private_key chain --module deweb
starport scaffold message delete_wallet address --module deweb
```

Create queries
```
starport scaffold query filter_user_key_records owner address chain deleted limit offset --response uuids --module deweb
```

## API
```
/deweb/external_wallets/v1beta1/list
?owner=deweb1z3z7lgfawp22mktkhxav5nr8g35m20ddpc4txj // фильтр по овнеру, мандатори
&address=deweb1z3z7lgfawp22mktkhxav5nr8g35m20ddpc4txj // это фильтр по полю address, может быть адресом любой сети, optional
&chain=cosmos // optional, default: return all chains
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

Test commands

Created account "alice" with address "cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz"

Created account "bob" with address "cosmos1z8yef05j2d3h6qzewdl68t48007xrmurztz7wy"

Creating test records:
```
~/go/bin/dewebd tx deweb save-wallet cosmosaddress private_cosmos_key cosmos --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-wallet siaaddress private_sia_key sia --from alice --chain-id deweb
~/go/bin/dewebd tx deweb save-wallet bobaddress private_bob_key sia --from bob --chain-id deweb
```

Get records by owner address
```
~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1z8yef05j2d3h6qzewdl68t48007xrmurztz7wy
records:
- address: bobaddress
  chain: sia
  deleted: false
  encrypted_key: private_bob_key
  owner: cosmos1z8yef05j2d3h6qzewdl68t48007xrmurztz7wy

~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
- address: siaaddress
  chain: sia
  deleted: false
  encrypted_key: private_sia_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
```

Get records for chain cosmos:
```
~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz "" cosmos
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
```

Remove records:
```
~/go/bin/dewebd tx deweb delete-wallet siaaddress --from alice --chain-id deweb
```

Check that records marked as removed:
```
~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz

~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz "" "" true
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
- address: siaaddress
  chain: sia
  deleted: true
  encrypted_key: private_sia_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
```

Limits and offsets:
```
~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz "" "" true 1
records:
- address: cosmosaddress
  chain: cosmos
  deleted: false
  encrypted_key: private_cosmos_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz

~/go/bin/dewebd q deweb filter-user-wallet-records cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz "" "" true 1 1
records:
- address: siaaddress
  chain: sia
  deleted: true
  encrypted_key: private_sia_key
  owner: cosmos1r5xmpu03ne6pg262qy3ds0y7ujnvkpe9lsh9cz
  ```


## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

### Web Frontend

Starport has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Starport front-end development](https://github.com/tendermint/vue).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/deweb-services/deweb@latest! | sudo bash
```
`deweb-services/deweb` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)

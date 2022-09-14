---
sidebar_position: 1
---

# Validator Setup Guide

## Sirius Testnet

**dewebd** is a blockchain application built using Cosmos SDK v.0.45.5 and Tendermint v.0.34.19.

## Step 0 - Run a fullnode / validator by compiling source code

The updated instructions are always in our GitHub Readme page, click on this **[link](https://github.com/deweb-services/deweb)** to go there.

## Step 1 - Setting up the connection

Instructions for setting up the connection with the DWS TestNet Blockchain.

1. Set the chain-id parameter

```bash
dewebd config chain-id deweb-testnet-3
```

2. **Create a wallet:** You may create a wallet with one or more keys (addresses) using `dewebd`; you can choose a name of your own liking (we strongly advice you use one word)

```bash
    dewebd keys add MyFirstAddress

      name: MyFirstAddress
      type: local
      address: deweb1q6wt62l9r4zef7nj97j5xe7q553j7nsllwrmqe
      pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ArhjQuNzZ+lSpIK9RrXK2da2PKAm7A3zpxTMHQnc/v+J"}'
      mnemonic: ""


      **Important** write this mnemonic phrase in a safe place.
      It is the only way to recover your account if you ever forget your password.

      giant favorite breeze resemble kitten surprise palm way jelly version use lucky pony depart napkin favorite slender normal grace always swarm funny hen cage
```

Your address will look something similar like this: `deweb1q6wt62l9r4zef7nj97j5xe7q553j7nsllwrmqe`

3. **Initialize the folders:** change Moniker by your validator name (use quotes for two or more separated words "Royal Queen Seeds")

```bash
dewebd init Moniker --chain-id deweb-testnet-3
```

This will create a `$HOME/.deweb` folder

4. Download the Genesis `genesis.json` file

```bash
cd $HOME
curl -s https://raw.githubusercontent.com/deweb-services/deweb/main/genesis.json > ~/.deweb/config/genesis.json
```

5. Add to `config.toml` file: server SEEDs:

:::info

It is better to get the seeds from **[NodeJumper](https://nodejumper.io/dws-testnet/sync)** validator, as it has the most recent seeds data: **[link](https://nodejumper.io/dws-testnet/sync)**.

:::

```bash
sed -E -i 's/seeds = \".*\"/seeds = \"4e46b666a70387003b927417ed1ac7f2b7e1bea4@31.44.6.80:26656\"/' $HOME/.deweb/config/config.toml
```

6. You can **set the minimum gas prices** for transactions to be accepted into your node’s mempool. This sets a lower bound on gas prices, preventing spam.

```bash
sed -E -i 's/minimum-gas-prices = \".*\"/minimum-gas-prices = \"0.001udws\"/' $HOME/.deweb/config/app.toml
```

7. Open the P2P port **(26656 by default)**

```bash
sudo ufw allow 26656
```

8. Test the connection **(CTRL + C to stop)**

```bash
dewebd start --log_level info
```

```bash
11:11AM INF committed state app_hash=38E6BC30BBEA37D15576F186CCFA2B3D8E30DFD281AB9F1E4BAACA5DD1E45863 height=23 module=state num_txs=0
11:11AM INF indexed block height=23 module=txindex
11:11AM INF Timed out dur=992.079979 height=24 module=consensus round=0 step=1
11:11AM INF received proposal module=consensus proposal={"Type":32,"block_id":{"hash":"B767CACAFEA3F87ACD0310C97079226BC1CA896BB41C71CD6B8B3B7BFCB9E4C7","parts":{"hash":"3E2426205921EDC5EF324FD3EA97ABB98B3A62AE2881B9CBFBBDF90C923A6315","total":1}},"height":24,"pol_round":-1,"round":0,"signature":"rSY4tqX2fhLVyGzLk/A3OaYt0re3/zka2bhMhYLZjVTn4lANV0yi9TbWHwp43SGiwIUVeBHoHdbFLg8mqVVNCA==","timestamp":"2022-01-12T11:11:54.504978728Z"}
11:11AM INF received complete proposal block hash=B767CACAFEA3F87ACD0310C97079226BC1CA896BB41C71CD6B8B3B7BFCB9E4C7 height=24 module=consensus
```

9. **Service creation.** Ensure that you stopped the previous test with CTRL+C. With all configurations ready, you can start your blockchain node with a single command (`dewebd start`). In this tutorial, however, you will find a simple way to set up `systemd` to run the node daemon with auto-restart.

Setup `dewebd` systemd service (copy and paste all to create the file service):

```bash
    cd $HOME
    echo "[Unit]
    Description=DWS Node
    After=network-online.target
    [Service]
    User=${USER}
    ExecStart=$(which dewebd) start
    Restart=always
    RestartSec=3
    LimitNOFILE=4096
    [Install]
    WantedBy=multi-user.target
    " >dewebd.service
```

Enable and activate the `dewebd` service.

```bash
    sudo mv dewebd.service /lib/systemd/system/
    sudo systemctl enable dewebd.service && sudo systemctl start dewebd.service
```

Check the logs to see if it is working: `sudo journalctl -u dewebd -f`

10. Check the synchronisation: If catching_up = true the node is syncing. Also you can compare your current block with the last synced block of another node, or at our **[Explorer](https://explore.deweb.services/)**:

```bash
curl -s localhost:26657/status  | jq .result.sync_info.catching_up
#true output is syncing - false is synced

curl -s localhost:26657/status | jq .result.sync_info.latest_block_height
#this output is your last block synced

curl -s "https://rpc-deweb.deweb.services/status?"  | jq .result.sync_info.latest_block_height
#this output the public node last block synced
```

## Step 2 - Become a validator

To become a validator you need to perform additional steps. Your node must be fully synced in order to send the TX of validator creation and start to validate the network. You can check if your node has fully synced by comparing your logs and the latest block in the explorer (https://explore.deweb.services/)

1. **You will need coins:** Send coins to your new address, you will need roughly 2 DWS to run the validator (1 DWS for self-delegation and a bit more for transactions).
2. **Send the _Create validator_ TX:**
   When you have your node synced and your wallet funded with coins, send the TX to become validator (change _wallet_name_ and _moniker_):
   :::note
   You can use quotes to include spaces and more than two words `--from "Royal Queen Seeds"`
   :::

```bash
dewebd tx staking create-validator \
    --amount 1000000udws \
    --commission-max-change-rate 0.10 \
    --commission-max-rate 0.2 \
    --commission-rate 0.1 \
    --from WALLET_NAME \
    --min-self-delegation 1 \
    --moniker YOUR_MONIKER \
    --pubkey $(dewebd tendermint show-validator) \
    --chain-id deweb-testnet-3 \
    --gas auto \
    --gas-adjustment 1.5 \
    --gas-prices 0.001udws
```

You can check the list of validators (also in **[Explorer](https://explore.deweb.services/)**):

```bash
dewebd query staking validators --output json| jq
```

3. Another **IMPORTANT** but **optional** action is backup your Validator_priv_key:

```bash
tar -czvf validator_key.tar.gz .deweb/config/*_key.json
gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
rm validator_key.tar.gz
```

This will create a GPG encrypted file with both key files.

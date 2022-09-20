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
dewebd config chain-id deweb-testnet-sirius
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
dewebd init Moniker --chain-id deweb-testnet-sirius
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
sed -E -i 's/seeds = \".*\"/seeds = \"c5b45045b0555c439d94f4d81a5ec4d1a578f98c@dws-testnet.nodejumper.io:27656,a8bc5ae440dd0ee13f091cd1b17d104c1a7b498c@95.217.225.214:29656,072683bf98bc3168e7f46a91b77b17bb38c0ae23@139.59.9.55:14656,a8793bb26c86089febec165300be03f0a8627dc8@77.37.176.99:46656,cfc175e44cf206d8175c91bdbef434c5b59f2461@94.130.50.61:28656,a7201d046651471d0560bce4bee4f80f00f7fa83@144.126.146.3:26656,2c50234b5a740899c18b6d1c3f0be83d2c30a8c0@38.242.216.50:26656,151d897c0236fde8c52b9c61bcb37c02dd37c9a2@65.108.199.222:25656,055cf1973026f65a2cee4c266a2dd2fa31f4fd4f@65.108.209.4:26656,ba7cf8cad9a96c1490c1d39af2af9848046d0992@181.41.142.78:11124,767da520aa74b24904070327761add8512540c87@62.171.132.152:26656,a62e7c05b82903470b60b1cb65723e2b6056f3d2@157.90.169.146:26656,5f2c0bc169ffeba76789f6a63bc3717a8eb4bd49@34.142.210.188:26656,4a07e3296fee5bdc2a45f7c5e6348168a1c9b286@144.91.125.55:26656,79a4dc86281be42768d33f25807ff5ab80c3c916@75.119.134.69:26656,0d25212f510f8868b639861de96ccd31fc1ea4dc@65.21.61.242:26656,d23354d5c60b723f315c28d0dc321aff2e7eedcb@5.161.80.30:14656,674088d85d41b5086029a4136625764071f17db8@5.161.86.210:14656,09a5075aa7a1075c90f72d663312ea49f491ebbf@142.132.196.251:26656,0ec737bc2425e744ad391f43b0ea79b18c7546d6@143.198.177.165:26656,af9f5da99dfbccde21e58473825524ab265dd804@65.108.44.149:46656,f6564402e0abe81813295bca5f0a4f286d15eb0a@65.108.133.73:26656,1cb152a929113c2628786e11a4bffe3c1d653fb7@173.249.50.126:14656,cee858f842bb7c5e13db1986cbd09b10553a2848@178.238.225.58:26656,7de0c3874eb05722c8326d05d8f602d1e5e85bdc@65.108.43.116:56132,5eeebfddb96ce09c09c33c2fcad8b8d19d10b6c9@65.21.199.148:26619,294eb22071d922486c363e0a56e179e507d2be21@138.201.91.105:14656,325b79115ca4ce4ff7cca054061cd347d694dadc@159.69.93.251:26656,ad45a4fe1e9e8a6edd8654924b7b6c1638f0214a@54.39.243.226:26656,a3245ded96e3642ff3f1d80f75f60ba7a58f8877@135.181.249.13:14656,054dfabf7ea154d41b085e20d42c7313f2eb2d48@62.171.148.163:26656,b8a77e817bb619c02dcad1ef11463d5ddc090f3b@185.9.144.138:26656,cb881cdf6f3d9fa02034eeb835c174517d960595@65.108.120.21:26101,ae72548f31f409a92fc00e5b62b513f8261ea7ec@144.91.118.61:26656,429231d7fbc695fe0c1e8de451164fc8194c105c@65.21.132.226:26656,aeceb0aa24c00cae8af8d30e0cc0275d04f5316e@162.55.181.95:14656,052d6215402c793361e52acfc6e257be53b35e21@65.21.143.79:26256,24e1eabae2b50a91b9230c52b89de326a44f0b47@161.97.119.183:26656,fed30dbdcec5ec7b4c4708997069ee6bf5456f60@89.163.255.100:14656,63064d9fe6bdffe6a85154592ec36be48cd63b9e@116.202.236.115:21046,fadb42aa4e0ed2183a8d88488f28e44413492882@141.94.254.145:47656,ee85cb0d941dca5759487f908c7339e2eea568ab@141.94.139.233:27656,aede67ebb4e40cb9564abbff2ff1de08cfbb1d6b@167.86.108.159:26656,56a057a6033664214c43ffadcf3c3ecfdfdf5d2a@209.126.81.240:26656,2c048866863ad3d55b015ba6dff49b63e66a011f@185.163.64.143:26656,64befdb7b718951faf9ce6244a96b791b6913594@95.216.70.104:26666,54ac40f4e4f4cca401c003f4065905ce91f5161a@85.10.198.171:16656,06f84e7041b77b470ce564bc69b521be6d1cca2c@116.203.220.212:26656,0b01867f4d90e24fa06e2292f0f43dd3bb342720@178.250.242.94:14656,f68e3850968edd258aab866d7697dd1f99e6e9fb@75.119.138.95:26656,d4ea6c4a7a4ede65d37b3ef5868b821fcf53732e@167.86.87.75:26656,3c29c80c62b17cfc593bf063fc1273f89c8dcb1a@65.109.5.239:12356,3eb14042b13fe7a3970ab12086117c8695833c84@88.99.95.81:36656,cb832c5b3ed839927cbe720db292101e377f13fe@159.69.149.85:36656,dab1a00b3b680ab5716dee1252ed2db65a311889@176.9.106.43:26656,2f40a727a89deca71e7a7605d69720bf47ff92cc@80.254.8.54:14656,e7a71d24a5f9988c5b0d41301db71a07cc146a43@194.233.90.134:26656,69792df8475149b71de899f88e60e9b45788aa5e@65.21.224.26:26656,f7919e6c903de12544f61e5f4a0ca410fb2a0773@195.3.223.11:26686,31cd90f57408c7f19519ab8a32f43b9543c23f44@161.97.91.132:14656,7a51db33edbf8de8573e7ffc9b27d76448409065@45.151.123.97:14656,bfbcc9759593f83b6c34522569de77d60f026c40@62.171.181.252:14656,0b0b72d370b2bd47bd30ac850ceda1735c850e50@62.171.184.181:14656,1fb96e8d9fd32589a0b10b23dd9fc520151d75a3@62.171.182.95:14656,a9bd4bdf7377a169dbe61943432e334a8f28dfe4@109.238.12.121:14656,20937bcc46b1cb585c27884fc680ffd82a9f0c57@65.108.203.219:22656,53059ddbab05f5ae6cb43bbbd7ee2cd208c4052b@65.109.10.249:26656,b591740ab4bf9a98cdff3073465786981142a997@185.144.99.226:26656,bf4f99c6ad6e8638319a82971286e04e2d2d42a3@75.119.129.39:26656,6e774a859b13269619b4258b02d5f0ed8e384685@185.144.99.225:26656,6888c6103d08344e8abfa0474be48e09120cab02@146.19.24.52:14656,53cebbfbe7d60fb35af1fc80be338ccbbddc9526@116.202.132.219:26699,0752197f9cdecf9df906ddbe88b349c4691fa183@95.179.242.150:26656,43e3232df9f1dd82ce32efc124ba035a2e23ef04@213.202.212.185:14656,fd2fa55269a9c9b1efdea72bb8293ba4203c9962@144.76.224.246:16656,225e49af2b5b245208481c6529a3dbfe5306c3bb@185.244.183.194:26656,7ea36ef113c3907b77b9bf8c73096247277fedba@144.91.102.79:26656,aa3a6d80d59947664dd41595830fa416d98f07a5@87.246.173.248:14656,b4803cf806fae48799efb78b8aee1ed2d0b45f1f@176.9.10.239:26656,12f2341874fba7e887a34113158dbb47374ea57f@65.21.131.215:26646,c1fb7595e3922ab805ce4e9010ca8d94521edf6b@144.91.92.219:26656,4206e15a077492ec2d392d4e9142847409b46285@149.102.143.147:26656,0affb5fdbf35d3bb8ed66af35339e9e845683faa@23.88.7.73:29586,4a7af5a67b065505592837135a8c5a5395dfd7b8@49.12.194.54:26656,a26e7be2b8b871ae5dcf7a1435eef721e95478e2@49.12.224.227:14656,ffdd012058df1104d9358b3016ef325794b3db3a@185.197.194.186:26656,da5a76aa7daa801c52887b318380d34617acbbcd@65.21.134.202:26646,d2cf889b2d4abec0e8ad9b2934fe3d09b5fd7649@62.171.171.152:14656,c16affc35507ce5c504906fed1c478595efb4675@86.48.5.144:26656,7236929b1ec0e3f54ae134560d8b0f08f6ca5024@154.53.61.44:14656,bc166a66999b13a883984f89a90d9d7fbd0fbd93@154.12.232.83:36656,58977d602f28afe21e09aa485b5ec520d178dbc3@66.94.108.7:14656,4eed23291dd15c92087527df00d6110f958a5ecc@161.97.166.146:26656,6e94932e91d09425f5cde9ed594f8f3b184d71b2@88.208.57.200:60856,31c2240813374e7819b6bb64ba24ce9990eb0aa6@194.163.130.165:26656,1d97083fcd4a2be02f18adf425d09f13c00effae@209.126.83.57:14656,52b620acc3bb780e589f090a68d10d563fda2f8a@149.154.65.66:26656,61da70005efc69aee1392d880aa34532c06adfdc@164.215.102.44:26656,5a83f1ac4a561a419d528a6b1af9361878f215b9@65.108.77.106:26869,3d323a9b988541ec369347dd15a1a89ce56ac738@95.217.191.34:14656\"/' $HOME/.deweb/config/config.toml
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
    --chain-id deweb-testnet-sirius \
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



import { IndexedTx, SigningStargateClient, StargateClient, defaultRegistryTypes } from "@cosmjs/stargate"
import { readFile } from "fs/promises"
import { DirectSecp256k1HdWallet, OfflineDirectSigner, EncodeObject, Registry } from "@cosmjs/proto-signing"
import { MsgRegisterDomain } from "./client/src/types/generated/dns_module/tx"
import { Tx } from "cosmjs-types/cosmos/tx/v1beta1/tx"

const rpc = "http://127.0.0.1:26657"

const getAliceSignerFromMnemonic = async (): Promise<OfflineDirectSigner> => {
    return DirectSecp256k1HdWallet.fromMnemonic((await readFile("./testnet.alice.mnemonic.key")).toString(), {
        prefix: "deweb",
    })
}

const runAll = async (): Promise<void> => {
    const client = await StargateClient.connect(rpc)
    console.log("With client, chain id:", await client.getChainId(), ", height:", await client.getHeight())

    const aliceSigner: OfflineDirectSigner = await getAliceSignerFromMnemonic()
    const alice = (await aliceSigner.getAccounts())[0].address
    console.log("Alice's address from signer", alice)

    console.log(
        "Alice balances:",
        await client.getAllBalances(alice)
    )

    // Here we get transaction example
    const dnsTX: IndexedTx = (await client.getTx(
        "791EFD071C7A4F3DEDFC52624ABB46E41633F3E855058EEBB5AD3893E971EC45"
    ))!
    console.log("DNS Tx:", dnsTX)
    const decodedTx: Tx = Tx.decode(dnsTX.tx)
    // console.log("DecodedTx:", decodedTx)
    const txMessage: EncodeObject = decodedTx.body!.messages[0]
    console.log("Decoded messages:", txMessage)
    const sendMessage: MsgRegisterDomain = MsgRegisterDomain.decode(txMessage.value)
    console.log("Sent message:", sendMessage)


    const myRegistry = new Registry(defaultRegistryTypes);
    myRegistry.register("/dewebservices.domain.MsgRegisterDomain", MsgRegisterDomain);
    const signingClient = await SigningStargateClient.connectWithSigner(
        rpc,
        aliceSigner,
        { registry: myRegistry }
    )


    console.log(
        "With signing client, chain id:",
        await signingClient.getChainId(),
        ", height:",
        await signingClient.getHeight()
    )
    // const faucet: string = sendMessage.fromAddress
    // console.log("Faucet balances:", await client.getAllBalances(faucet))
    //
    // // Get the faucet address another way
    // {
    //     const rawLog = JSON.parse(faucetTx.rawLog)
    //     console.log("Raw log:", JSON.stringify(rawLog, null, 4))
    //     const faucet: string = rawLog[0].events
    //         .find((eventEl: any) => eventEl.type === "coin_spent")
    //         .attributes.find((attribute: any) => attribute.key === "spender").value
    //     console.log("Faucet address from raw log:", faucet)
    // }

    const result = await signingClient.signAndBroadcast(
        // the signerAddress
        alice,
        // the message(s)
        [
            {
                typeUrl: "/dewebservices.domain.MsgRegisterDomain",
                value: {
                    id: 'dewebnne',
                    data: '{"records": [{"type": "A","values": ["192.168.1.10"]}]}',
                    sender: alice,
                    recipient: alice
                },
            },
        ],
        // the fee
        {
            amount: [],
            gas: "200000",
        },
    )

}

runAll()
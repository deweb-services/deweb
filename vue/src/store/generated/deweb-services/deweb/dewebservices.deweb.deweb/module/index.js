// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgDeleteKey } from "./types/deweb/tx";
import { MsgSaveUser } from "./types/deweb/tx";
const types = [
    ["/dewebservices.deweb.deweb.MsgDeleteKey", MsgDeleteKey],
    ["/dewebservices.deweb.deweb.MsgSaveUser", MsgSaveUser],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgDeleteKey: (data) => ({ typeUrl: "/dewebservices.deweb.deweb.MsgDeleteKey", value: data }),
        msgSaveUser: (data) => ({ typeUrl: "/dewebservices.deweb.deweb.MsgSaveUser", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };

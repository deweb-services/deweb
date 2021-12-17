import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "dewebservices.deweb.deweb";
export interface MsgSaveUser {
    creator: string;
    message: string;
    chain: string;
}
export interface MsgSaveUserResponse {
}
export interface MsgDeleteKey {
    creator: string;
    uuid: string;
}
export interface MsgDeleteKeyResponse {
}
export declare const MsgSaveUser: {
    encode(message: MsgSaveUser, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveUser;
    fromJSON(object: any): MsgSaveUser;
    toJSON(message: MsgSaveUser): unknown;
    fromPartial(object: DeepPartial<MsgSaveUser>): MsgSaveUser;
};
export declare const MsgSaveUserResponse: {
    encode(_: MsgSaveUserResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgSaveUserResponse;
    fromJSON(_: any): MsgSaveUserResponse;
    toJSON(_: MsgSaveUserResponse): unknown;
    fromPartial(_: DeepPartial<MsgSaveUserResponse>): MsgSaveUserResponse;
};
export declare const MsgDeleteKey: {
    encode(message: MsgDeleteKey, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteKey;
    fromJSON(object: any): MsgDeleteKey;
    toJSON(message: MsgDeleteKey): unknown;
    fromPartial(object: DeepPartial<MsgDeleteKey>): MsgDeleteKey;
};
export declare const MsgDeleteKeyResponse: {
    encode(_: MsgDeleteKeyResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteKeyResponse;
    fromJSON(_: any): MsgDeleteKeyResponse;
    toJSON(_: MsgDeleteKeyResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteKeyResponse>): MsgDeleteKeyResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    SaveUser(request: MsgSaveUser): Promise<MsgSaveUserResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    DeleteKey(request: MsgDeleteKey): Promise<MsgDeleteKeyResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    SaveUser(request: MsgSaveUser): Promise<MsgSaveUserResponse>;
    DeleteKey(request: MsgDeleteKey): Promise<MsgDeleteKeyResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

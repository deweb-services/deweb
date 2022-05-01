import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "dewebservices.deweb.deweb";
export interface UserWalletRec {
    address: string;
    encryptedKey: string;
    chain: string;
    deleted: boolean;
}
export interface RecordsToUser {
    records: string[];
}
export declare const UserWalletRec: {
    encode(message: UserWalletRec, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): UserWalletRec;
    fromJSON(object: any): UserWalletRec;
    toJSON(message: UserWalletRec): unknown;
    fromPartial(object: DeepPartial<UserWalletRec>): UserWalletRec;
};
export declare const RecordsToUser: {
    encode(message: RecordsToUser, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RecordsToUser;
    fromJSON(object: any): RecordsToUser;
    toJSON(message: RecordsToUser): unknown;
    fromPartial(object: DeepPartial<RecordsToUser>): RecordsToUser;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

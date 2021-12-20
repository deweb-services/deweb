import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "dewebservices.deweb.deweb";
export interface UserKeyRec {
    creator: string;
    message: string;
    chain: string;
    deleted: boolean;
}
export interface RecordsToUser {
    records: string[];
}
export declare const UserKeyRec: {
    encode(message: UserKeyRec, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): UserKeyRec;
    fromJSON(object: any): UserKeyRec;
    toJSON(message: UserKeyRec): unknown;
    fromPartial(object: DeepPartial<UserKeyRec>): UserKeyRec;
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

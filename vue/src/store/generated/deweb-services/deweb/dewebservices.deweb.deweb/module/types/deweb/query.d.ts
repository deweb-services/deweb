import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "dewebservices.deweb.deweb";
export interface QueryGetKeyRecordRequest {
    uuid: string;
}
export interface QueryGetKeyRecordResponse {
    uuid: string;
    message: string;
    chain: string;
    deleted: boolean;
}
export interface QueryGetUserKeyRecordsRequest {
    address: string;
}
export interface QueryGetUserKeyRecordsResponse {
    uuids: string[];
}
export interface QueryFilterUserKeyRecordsRequest {
    address: string;
    chain: string;
    deleted: boolean;
}
export interface QueryFilterUserKeyRecordsResponse {
    records: QueryGetKeyRecordResponse[];
}
export declare const QueryGetKeyRecordRequest: {
    encode(message: QueryGetKeyRecordRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetKeyRecordRequest;
    fromJSON(object: any): QueryGetKeyRecordRequest;
    toJSON(message: QueryGetKeyRecordRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetKeyRecordRequest>): QueryGetKeyRecordRequest;
};
export declare const QueryGetKeyRecordResponse: {
    encode(message: QueryGetKeyRecordResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetKeyRecordResponse;
    fromJSON(object: any): QueryGetKeyRecordResponse;
    toJSON(message: QueryGetKeyRecordResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetKeyRecordResponse>): QueryGetKeyRecordResponse;
};
export declare const QueryGetUserKeyRecordsRequest: {
    encode(message: QueryGetUserKeyRecordsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUserKeyRecordsRequest;
    fromJSON(object: any): QueryGetUserKeyRecordsRequest;
    toJSON(message: QueryGetUserKeyRecordsRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetUserKeyRecordsRequest>): QueryGetUserKeyRecordsRequest;
};
export declare const QueryGetUserKeyRecordsResponse: {
    encode(message: QueryGetUserKeyRecordsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetUserKeyRecordsResponse;
    fromJSON(object: any): QueryGetUserKeyRecordsResponse;
    toJSON(message: QueryGetUserKeyRecordsResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetUserKeyRecordsResponse>): QueryGetUserKeyRecordsResponse;
};
export declare const QueryFilterUserKeyRecordsRequest: {
    encode(message: QueryFilterUserKeyRecordsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryFilterUserKeyRecordsRequest;
    fromJSON(object: any): QueryFilterUserKeyRecordsRequest;
    toJSON(message: QueryFilterUserKeyRecordsRequest): unknown;
    fromPartial(object: DeepPartial<QueryFilterUserKeyRecordsRequest>): QueryFilterUserKeyRecordsRequest;
};
export declare const QueryFilterUserKeyRecordsResponse: {
    encode(message: QueryFilterUserKeyRecordsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryFilterUserKeyRecordsResponse;
    fromJSON(object: any): QueryFilterUserKeyRecordsResponse;
    toJSON(message: QueryFilterUserKeyRecordsResponse): unknown;
    fromPartial(object: DeepPartial<QueryFilterUserKeyRecordsResponse>): QueryFilterUserKeyRecordsResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a list of getKeyRecord items. */
    GetKeyRecord(request: QueryGetKeyRecordRequest): Promise<QueryGetKeyRecordResponse>;
    /** Queries a list of getUserKeyRecords items. */
    GetUserKeyRecords(request: QueryGetUserKeyRecordsRequest): Promise<QueryGetUserKeyRecordsResponse>;
    /** Queries a list of filterUserKeyRecords items. */
    FilterUserKeyRecords(request: QueryFilterUserKeyRecordsRequest): Promise<QueryFilterUserKeyRecordsResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    GetKeyRecord(request: QueryGetKeyRecordRequest): Promise<QueryGetKeyRecordResponse>;
    GetUserKeyRecords(request: QueryGetUserKeyRecordsRequest): Promise<QueryGetUserKeyRecordsResponse>;
    FilterUserKeyRecords(request: QueryFilterUserKeyRecordsRequest): Promise<QueryFilterUserKeyRecordsResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

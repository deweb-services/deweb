/* eslint-disable */
import { Params } from "./params";
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.deweb";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: Params;
}

export interface WalletRecordResponse {
  owner: string;
  address: string;
  encryptedKey: string;
  chain: string;
  deleted: boolean;
}

export interface QueryFilterUserWalletRecordsRequest {
  owner: string;
  address: string;
  chain: string;
  deleted: boolean;
  limit: number;
  offset: number;
}

export interface QueryFilterUserWalletRecordsResponse {
  records: WalletRecordResponse[];
}

export interface ChainMappingRecordResponse {
  owner: string;
  extAddress: string;
  chain: string;
  deleted: boolean;
}

export interface QueryFilterChainMappingsRecordsRequest {
  owner: string;
  address: string;
  chain: string;
  deleted: boolean;
  limit: number;
  offset: number;
}

export interface QueryFilterChainMappingsRecordsResponse {
  records: ChainMappingRecordResponse[];
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(
    _: QueryParamsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(
    _: I
  ): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(
    object: I
  ): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params =
      object.params !== undefined && object.params !== null
        ? Params.fromPartial(object.params)
        : undefined;
    return message;
  },
};

function createBaseWalletRecordResponse(): WalletRecordResponse {
  return {
    owner: "",
    address: "",
    encryptedKey: "",
    chain: "",
    deleted: false,
  };
}

export const WalletRecordResponse = {
  encode(
    message: WalletRecordResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.encryptedKey !== "") {
      writer.uint32(26).string(message.encryptedKey);
    }
    if (message.chain !== "") {
      writer.uint32(34).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(40).bool(message.deleted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): WalletRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWalletRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.encryptedKey = reader.string();
          break;
        case 4:
          message.chain = reader.string();
          break;
        case 5:
          message.deleted = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WalletRecordResponse {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      address: isSet(object.address) ? String(object.address) : "",
      encryptedKey: isSet(object.encryptedKey)
        ? String(object.encryptedKey)
        : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
    };
  },

  toJSON(message: WalletRecordResponse): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.address !== undefined && (obj.address = message.address);
    message.encryptedKey !== undefined &&
      (obj.encryptedKey = message.encryptedKey);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WalletRecordResponse>, I>>(
    object: I
  ): WalletRecordResponse {
    const message = createBaseWalletRecordResponse();
    message.owner = object.owner ?? "";
    message.address = object.address ?? "";
    message.encryptedKey = object.encryptedKey ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseQueryFilterUserWalletRecordsRequest(): QueryFilterUserWalletRecordsRequest {
  return {
    owner: "",
    address: "",
    chain: "",
    deleted: false,
    limit: 0,
    offset: 0,
  };
}

export const QueryFilterUserWalletRecordsRequest = {
  encode(
    message: QueryFilterUserWalletRecordsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(32).bool(message.deleted);
    }
    if (message.limit !== 0) {
      writer.uint32(40).int32(message.limit);
    }
    if (message.offset !== 0) {
      writer.uint32(48).int32(message.offset);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryFilterUserWalletRecordsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilterUserWalletRecordsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        case 4:
          message.deleted = reader.bool();
          break;
        case 5:
          message.limit = reader.int32();
          break;
        case 6:
          message.offset = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterUserWalletRecordsRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      address: isSet(object.address) ? String(object.address) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
      limit: isSet(object.limit) ? Number(object.limit) : 0,
      offset: isSet(object.offset) ? Number(object.offset) : 0,
    };
  },

  toJSON(message: QueryFilterUserWalletRecordsRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.address !== undefined && (obj.address = message.address);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    message.limit !== undefined && (obj.limit = Math.round(message.limit));
    message.offset !== undefined && (obj.offset = Math.round(message.offset));
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<QueryFilterUserWalletRecordsRequest>, I>
  >(object: I): QueryFilterUserWalletRecordsRequest {
    const message = createBaseQueryFilterUserWalletRecordsRequest();
    message.owner = object.owner ?? "";
    message.address = object.address ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    message.limit = object.limit ?? 0;
    message.offset = object.offset ?? 0;
    return message;
  },
};

function createBaseQueryFilterUserWalletRecordsResponse(): QueryFilterUserWalletRecordsResponse {
  return { records: [] };
}

export const QueryFilterUserWalletRecordsResponse = {
  encode(
    message: QueryFilterUserWalletRecordsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.records) {
      WalletRecordResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryFilterUserWalletRecordsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilterUserWalletRecordsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.records.push(
            WalletRecordResponse.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterUserWalletRecordsResponse {
    return {
      records: Array.isArray(object?.records)
        ? object.records.map((e: any) => WalletRecordResponse.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryFilterUserWalletRecordsResponse): unknown {
    const obj: any = {};
    if (message.records) {
      obj.records = message.records.map((e) =>
        e ? WalletRecordResponse.toJSON(e) : undefined
      );
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<QueryFilterUserWalletRecordsResponse>, I>
  >(object: I): QueryFilterUserWalletRecordsResponse {
    const message = createBaseQueryFilterUserWalletRecordsResponse();
    message.records =
      object.records?.map((e) => WalletRecordResponse.fromPartial(e)) || [];
    return message;
  },
};

function createBaseChainMappingRecordResponse(): ChainMappingRecordResponse {
  return { owner: "", extAddress: "", chain: "", deleted: false };
}

export const ChainMappingRecordResponse = {
  encode(
    message: ChainMappingRecordResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.extAddress !== "") {
      writer.uint32(18).string(message.extAddress);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(32).bool(message.deleted);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): ChainMappingRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChainMappingRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.extAddress = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        case 4:
          message.deleted = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ChainMappingRecordResponse {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      extAddress: isSet(object.extAddress) ? String(object.extAddress) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
    };
  },

  toJSON(message: ChainMappingRecordResponse): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.extAddress !== undefined && (obj.extAddress = message.extAddress);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChainMappingRecordResponse>, I>>(
    object: I
  ): ChainMappingRecordResponse {
    const message = createBaseChainMappingRecordResponse();
    message.owner = object.owner ?? "";
    message.extAddress = object.extAddress ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseQueryFilterChainMappingsRecordsRequest(): QueryFilterChainMappingsRecordsRequest {
  return {
    owner: "",
    address: "",
    chain: "",
    deleted: false,
    limit: 0,
    offset: 0,
  };
}

export const QueryFilterChainMappingsRecordsRequest = {
  encode(
    message: QueryFilterChainMappingsRecordsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(32).bool(message.deleted);
    }
    if (message.limit !== 0) {
      writer.uint32(40).int32(message.limit);
    }
    if (message.offset !== 0) {
      writer.uint32(48).int32(message.offset);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryFilterChainMappingsRecordsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilterChainMappingsRecordsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        case 4:
          message.deleted = reader.bool();
          break;
        case 5:
          message.limit = reader.int32();
          break;
        case 6:
          message.offset = reader.int32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterChainMappingsRecordsRequest {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      address: isSet(object.address) ? String(object.address) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
      limit: isSet(object.limit) ? Number(object.limit) : 0,
      offset: isSet(object.offset) ? Number(object.offset) : 0,
    };
  },

  toJSON(message: QueryFilterChainMappingsRecordsRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.address !== undefined && (obj.address = message.address);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    message.limit !== undefined && (obj.limit = Math.round(message.limit));
    message.offset !== undefined && (obj.offset = Math.round(message.offset));
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<QueryFilterChainMappingsRecordsRequest>, I>
  >(object: I): QueryFilterChainMappingsRecordsRequest {
    const message = createBaseQueryFilterChainMappingsRecordsRequest();
    message.owner = object.owner ?? "";
    message.address = object.address ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    message.limit = object.limit ?? 0;
    message.offset = object.offset ?? 0;
    return message;
  },
};

function createBaseQueryFilterChainMappingsRecordsResponse(): QueryFilterChainMappingsRecordsResponse {
  return { records: [] };
}

export const QueryFilterChainMappingsRecordsResponse = {
  encode(
    message: QueryFilterChainMappingsRecordsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.records) {
      ChainMappingRecordResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryFilterChainMappingsRecordsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryFilterChainMappingsRecordsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.records.push(
            ChainMappingRecordResponse.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterChainMappingsRecordsResponse {
    return {
      records: Array.isArray(object?.records)
        ? object.records.map((e: any) => ChainMappingRecordResponse.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryFilterChainMappingsRecordsResponse): unknown {
    const obj: any = {};
    if (message.records) {
      obj.records = message.records.map((e) =>
        e ? ChainMappingRecordResponse.toJSON(e) : undefined
      );
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial<
    I extends Exact<DeepPartial<QueryFilterChainMappingsRecordsResponse>, I>
  >(object: I): QueryFilterChainMappingsRecordsResponse {
    const message = createBaseQueryFilterChainMappingsRecordsResponse();
    message.records =
      object.records?.map((e) => ChainMappingRecordResponse.fromPartial(e)) ||
      [];
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  FilterUserWalletRecords(
    request: QueryFilterUserWalletRecordsRequest
  ): Promise<QueryFilterUserWalletRecordsResponse>;
  /** Queries a list of FilterChainMappingsRecords items. */
  FilterChainMappingsRecords(
    request: QueryFilterChainMappingsRecordsRequest
  ): Promise<QueryFilterChainMappingsRecordsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.FilterUserWalletRecords = this.FilterUserWalletRecords.bind(this);
    this.FilterChainMappingsRecords =
      this.FilterChainMappingsRecords.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "Params",
      data
    );
    return promise.then((data) =>
      QueryParamsResponse.decode(new _m0.Reader(data))
    );
  }

  FilterUserWalletRecords(
    request: QueryFilterUserWalletRecordsRequest
  ): Promise<QueryFilterUserWalletRecordsResponse> {
    const data = QueryFilterUserWalletRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "FilterUserWalletRecords",
      data
    );
    return promise.then((data) =>
      QueryFilterUserWalletRecordsResponse.decode(new _m0.Reader(data))
    );
  }

  FilterChainMappingsRecords(
    request: QueryFilterChainMappingsRecordsRequest
  ): Promise<QueryFilterChainMappingsRecordsResponse> {
    const data =
      QueryFilterChainMappingsRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "FilterChainMappingsRecords",
      data
    );
    return promise.then((data) =>
      QueryFilterChainMappingsRecordsResponse.decode(new _m0.Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Long
  ? string | number | Long
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

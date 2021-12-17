/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.deweb";

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

const baseQueryGetKeyRecordRequest: object = { uuid: "" };

export const QueryGetKeyRecordRequest = {
  encode(
    message: QueryGetKeyRecordRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetKeyRecordRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetKeyRecordRequest,
    } as QueryGetKeyRecordRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetKeyRecordRequest {
    const message = {
      ...baseQueryGetKeyRecordRequest,
    } as QueryGetKeyRecordRequest;
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    return message;
  },

  toJSON(message: QueryGetKeyRecordRequest): unknown {
    const obj: any = {};
    message.uuid !== undefined && (obj.uuid = message.uuid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetKeyRecordRequest>
  ): QueryGetKeyRecordRequest {
    const message = {
      ...baseQueryGetKeyRecordRequest,
    } as QueryGetKeyRecordRequest;
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    return message;
  },
};

const baseQueryGetKeyRecordResponse: object = {
  uuid: "",
  message: "",
  chain: "",
  deleted: false,
};

export const QueryGetKeyRecordResponse = {
  encode(
    message: QueryGetKeyRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uuid !== "") {
      writer.uint32(10).string(message.uuid);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
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
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetKeyRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetKeyRecordResponse,
    } as QueryGetKeyRecordResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuid = reader.string();
          break;
        case 2:
          message.message = reader.string();
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

  fromJSON(object: any): QueryGetKeyRecordResponse {
    const message = {
      ...baseQueryGetKeyRecordResponse,
    } as QueryGetKeyRecordResponse;
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = String(object.message);
    } else {
      message.message = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = Boolean(object.deleted);
    } else {
      message.deleted = false;
    }
    return message;
  },

  toJSON(message: QueryGetKeyRecordResponse): unknown {
    const obj: any = {};
    message.uuid !== undefined && (obj.uuid = message.uuid);
    message.message !== undefined && (obj.message = message.message);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetKeyRecordResponse>
  ): QueryGetKeyRecordResponse {
    const message = {
      ...baseQueryGetKeyRecordResponse,
    } as QueryGetKeyRecordResponse;
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    if (object.message !== undefined && object.message !== null) {
      message.message = object.message;
    } else {
      message.message = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = object.deleted;
    } else {
      message.deleted = false;
    }
    return message;
  },
};

const baseQueryGetUserKeyRecordsRequest: object = { address: "" };

export const QueryGetUserKeyRecordsRequest = {
  encode(
    message: QueryGetUserKeyRecordsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserKeyRecordsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserKeyRecordsRequest,
    } as QueryGetUserKeyRecordsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserKeyRecordsRequest {
    const message = {
      ...baseQueryGetUserKeyRecordsRequest,
    } as QueryGetUserKeyRecordsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetUserKeyRecordsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserKeyRecordsRequest>
  ): QueryGetUserKeyRecordsRequest {
    const message = {
      ...baseQueryGetUserKeyRecordsRequest,
    } as QueryGetUserKeyRecordsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetUserKeyRecordsResponse: object = { uuids: "" };

export const QueryGetUserKeyRecordsResponse = {
  encode(
    message: QueryGetUserKeyRecordsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.uuids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetUserKeyRecordsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetUserKeyRecordsResponse,
    } as QueryGetUserKeyRecordsResponse;
    message.uuids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uuids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserKeyRecordsResponse {
    const message = {
      ...baseQueryGetUserKeyRecordsResponse,
    } as QueryGetUserKeyRecordsResponse;
    message.uuids = [];
    if (object.uuids !== undefined && object.uuids !== null) {
      for (const e of object.uuids) {
        message.uuids.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: QueryGetUserKeyRecordsResponse): unknown {
    const obj: any = {};
    if (message.uuids) {
      obj.uuids = message.uuids.map((e) => e);
    } else {
      obj.uuids = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetUserKeyRecordsResponse>
  ): QueryGetUserKeyRecordsResponse {
    const message = {
      ...baseQueryGetUserKeyRecordsResponse,
    } as QueryGetUserKeyRecordsResponse;
    message.uuids = [];
    if (object.uuids !== undefined && object.uuids !== null) {
      for (const e of object.uuids) {
        message.uuids.push(e);
      }
    }
    return message;
  },
};

const baseQueryFilterUserKeyRecordsRequest: object = {
  address: "",
  chain: "",
  deleted: false,
};

export const QueryFilterUserKeyRecordsRequest = {
  encode(
    message: QueryFilterUserKeyRecordsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(24).bool(message.deleted);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryFilterUserKeyRecordsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryFilterUserKeyRecordsRequest,
    } as QueryFilterUserKeyRecordsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.chain = reader.string();
          break;
        case 3:
          message.deleted = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterUserKeyRecordsRequest {
    const message = {
      ...baseQueryFilterUserKeyRecordsRequest,
    } as QueryFilterUserKeyRecordsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = String(object.chain);
    } else {
      message.chain = "";
    }
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = Boolean(object.deleted);
    } else {
      message.deleted = false;
    }
    return message;
  },

  toJSON(message: QueryFilterUserKeyRecordsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFilterUserKeyRecordsRequest>
  ): QueryFilterUserKeyRecordsRequest {
    const message = {
      ...baseQueryFilterUserKeyRecordsRequest,
    } as QueryFilterUserKeyRecordsRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.chain !== undefined && object.chain !== null) {
      message.chain = object.chain;
    } else {
      message.chain = "";
    }
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = object.deleted;
    } else {
      message.deleted = false;
    }
    return message;
  },
};

const baseQueryFilterUserKeyRecordsResponse: object = {};

export const QueryFilterUserKeyRecordsResponse = {
  encode(
    message: QueryFilterUserKeyRecordsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.records) {
      QueryGetKeyRecordResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryFilterUserKeyRecordsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryFilterUserKeyRecordsResponse,
    } as QueryFilterUserKeyRecordsResponse;
    message.records = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.records.push(
            QueryGetKeyRecordResponse.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryFilterUserKeyRecordsResponse {
    const message = {
      ...baseQueryFilterUserKeyRecordsResponse,
    } as QueryFilterUserKeyRecordsResponse;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(QueryGetKeyRecordResponse.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: QueryFilterUserKeyRecordsResponse): unknown {
    const obj: any = {};
    if (message.records) {
      obj.records = message.records.map((e) =>
        e ? QueryGetKeyRecordResponse.toJSON(e) : undefined
      );
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryFilterUserKeyRecordsResponse>
  ): QueryFilterUserKeyRecordsResponse {
    const message = {
      ...baseQueryFilterUserKeyRecordsResponse,
    } as QueryFilterUserKeyRecordsResponse;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(QueryGetKeyRecordResponse.fromPartial(e));
      }
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a list of getKeyRecord items. */
  GetKeyRecord(
    request: QueryGetKeyRecordRequest
  ): Promise<QueryGetKeyRecordResponse>;
  /** Queries a list of getUserKeyRecords items. */
  GetUserKeyRecords(
    request: QueryGetUserKeyRecordsRequest
  ): Promise<QueryGetUserKeyRecordsResponse>;
  /** Queries a list of filterUserKeyRecords items. */
  FilterUserKeyRecords(
    request: QueryFilterUserKeyRecordsRequest
  ): Promise<QueryFilterUserKeyRecordsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  GetKeyRecord(
    request: QueryGetKeyRecordRequest
  ): Promise<QueryGetKeyRecordResponse> {
    const data = QueryGetKeyRecordRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "GetKeyRecord",
      data
    );
    return promise.then((data) =>
      QueryGetKeyRecordResponse.decode(new Reader(data))
    );
  }

  GetUserKeyRecords(
    request: QueryGetUserKeyRecordsRequest
  ): Promise<QueryGetUserKeyRecordsResponse> {
    const data = QueryGetUserKeyRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "GetUserKeyRecords",
      data
    );
    return promise.then((data) =>
      QueryGetUserKeyRecordsResponse.decode(new Reader(data))
    );
  }

  FilterUserKeyRecords(
    request: QueryFilterUserKeyRecordsRequest
  ): Promise<QueryFilterUserKeyRecordsResponse> {
    const data = QueryFilterUserKeyRecordsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Query",
      "FilterUserKeyRecords",
      data
    );
    return promise.then((data) =>
      QueryFilterUserKeyRecordsResponse.decode(new Reader(data))
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

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

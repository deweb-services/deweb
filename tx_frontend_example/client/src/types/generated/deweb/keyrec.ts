/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.v1beta1";

export interface UserWalletRec {
  address: string;
  encryptedKey: string;
  chain: string;
  deleted: boolean;
}

export interface RecordsToUser {
  records: string[];
}

export interface ChainAddressMapping {
  extAddress: string;
  chain: string;
  deleted: boolean;
}

function createBaseUserWalletRec(): UserWalletRec {
  return { address: "", encryptedKey: "", chain: "", deleted: false };
}

export const UserWalletRec = {
  encode(
    message: UserWalletRec,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): UserWalletRec {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserWalletRec();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
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

  fromJSON(object: any): UserWalletRec {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      encryptedKey: isSet(object.encryptedKey)
        ? String(object.encryptedKey)
        : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
    };
  },

  toJSON(message: UserWalletRec): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.encryptedKey !== undefined &&
      (obj.encryptedKey = message.encryptedKey);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserWalletRec>, I>>(
    object: I
  ): UserWalletRec {
    const message = createBaseUserWalletRec();
    message.address = object.address ?? "";
    message.encryptedKey = object.encryptedKey ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    return message;
  },
};

function createBaseRecordsToUser(): RecordsToUser {
  return { records: [] };
}

export const RecordsToUser = {
  encode(
    message: RecordsToUser,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.records) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RecordsToUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRecordsToUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.records.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RecordsToUser {
    return {
      records: Array.isArray(object?.records)
        ? object.records.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: RecordsToUser): unknown {
    const obj: any = {};
    if (message.records) {
      obj.records = message.records.map((e) => e);
    } else {
      obj.records = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RecordsToUser>, I>>(
    object: I
  ): RecordsToUser {
    const message = createBaseRecordsToUser();
    message.records = object.records?.map((e) => e) || [];
    return message;
  },
};

function createBaseChainAddressMapping(): ChainAddressMapping {
  return { extAddress: "", chain: "", deleted: false };
}

export const ChainAddressMapping = {
  encode(
    message: ChainAddressMapping,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.extAddress !== "") {
      writer.uint32(10).string(message.extAddress);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    if (message.deleted === true) {
      writer.uint32(24).bool(message.deleted);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ChainAddressMapping {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseChainAddressMapping();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.extAddress = reader.string();
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

  fromJSON(object: any): ChainAddressMapping {
    return {
      extAddress: isSet(object.extAddress) ? String(object.extAddress) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      deleted: isSet(object.deleted) ? Boolean(object.deleted) : false,
    };
  },

  toJSON(message: ChainAddressMapping): unknown {
    const obj: any = {};
    message.extAddress !== undefined && (obj.extAddress = message.extAddress);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ChainAddressMapping>, I>>(
    object: I
  ): ChainAddressMapping {
    const message = createBaseChainAddressMapping();
    message.extAddress = object.extAddress ?? "";
    message.chain = object.chain ?? "";
    message.deleted = object.deleted ?? false;
    return message;
  },
};

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

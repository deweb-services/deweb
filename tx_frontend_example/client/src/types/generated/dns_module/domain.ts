/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.domain";

/** BaseDomain defines a non-fungible domain token */
export interface BaseDomain {
  id: string;
  data: string;
  owner: string;
}

/** Denom defines a type of NFT */
export interface Denom {
  id: string;
  name: string;
  schema: string;
  creator: string;
  symbol: string;
  mintRestricted: boolean;
  updateRestricted: boolean;
  description: string;
  uri: string;
  uriHash: string;
  data: string;
}

/** IDCollection defines a type of collection with specified ID */
export interface IDCollection {
  denomId: string;
  tokenIds: string[];
}

/** Owner defines a type of owner */
export interface Owner {
  address: string;
  idCollections: IDCollection[];
}

/** Collection defines a type of collection */
export interface Collection {
  denom?: Denom;
  nfts: BaseDomain[];
}

function createBaseBaseDomain(): BaseDomain {
  return { id: "", data: "", owner: "" };
}

export const BaseDomain = {
  encode(
    message: BaseDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.data !== "") {
      writer.uint32(18).string(message.data);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): BaseDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseBaseDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.data = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): BaseDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      data: isSet(object.data) ? String(object.data) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
    };
  },

  toJSON(message: BaseDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.data !== undefined && (obj.data = message.data);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<BaseDomain>, I>>(
    object: I
  ): BaseDomain {
    const message = createBaseBaseDomain();
    message.id = object.id ?? "";
    message.data = object.data ?? "";
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseDenom(): Denom {
  return {
    id: "",
    name: "",
    schema: "",
    creator: "",
    symbol: "",
    mintRestricted: false,
    updateRestricted: false,
    description: "",
    uri: "",
    uriHash: "",
    data: "",
  };
}

export const Denom = {
  encode(message: Denom, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.schema !== "") {
      writer.uint32(26).string(message.schema);
    }
    if (message.creator !== "") {
      writer.uint32(34).string(message.creator);
    }
    if (message.symbol !== "") {
      writer.uint32(42).string(message.symbol);
    }
    if (message.mintRestricted === true) {
      writer.uint32(48).bool(message.mintRestricted);
    }
    if (message.updateRestricted === true) {
      writer.uint32(56).bool(message.updateRestricted);
    }
    if (message.description !== "") {
      writer.uint32(66).string(message.description);
    }
    if (message.uri !== "") {
      writer.uint32(74).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(82).string(message.uriHash);
    }
    if (message.data !== "") {
      writer.uint32(90).string(message.data);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Denom {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDenom();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.schema = reader.string();
          break;
        case 4:
          message.creator = reader.string();
          break;
        case 5:
          message.symbol = reader.string();
          break;
        case 6:
          message.mintRestricted = reader.bool();
          break;
        case 7:
          message.updateRestricted = reader.bool();
          break;
        case 8:
          message.description = reader.string();
          break;
        case 9:
          message.uri = reader.string();
          break;
        case 10:
          message.uriHash = reader.string();
          break;
        case 11:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Denom {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      schema: isSet(object.schema) ? String(object.schema) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      mintRestricted: isSet(object.mintRestricted)
        ? Boolean(object.mintRestricted)
        : false,
      updateRestricted: isSet(object.updateRestricted)
        ? Boolean(object.updateRestricted)
        : false,
      description: isSet(object.description) ? String(object.description) : "",
      uri: isSet(object.uri) ? String(object.uri) : "",
      uriHash: isSet(object.uriHash) ? String(object.uriHash) : "",
      data: isSet(object.data) ? String(object.data) : "",
    };
  },

  toJSON(message: Denom): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.schema !== undefined && (obj.schema = message.schema);
    message.creator !== undefined && (obj.creator = message.creator);
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.mintRestricted !== undefined &&
      (obj.mintRestricted = message.mintRestricted);
    message.updateRestricted !== undefined &&
      (obj.updateRestricted = message.updateRestricted);
    message.description !== undefined &&
      (obj.description = message.description);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Denom>, I>>(object: I): Denom {
    const message = createBaseDenom();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.schema = object.schema ?? "";
    message.creator = object.creator ?? "";
    message.symbol = object.symbol ?? "";
    message.mintRestricted = object.mintRestricted ?? false;
    message.updateRestricted = object.updateRestricted ?? false;
    message.description = object.description ?? "";
    message.uri = object.uri ?? "";
    message.uriHash = object.uriHash ?? "";
    message.data = object.data ?? "";
    return message;
  },
};

function createBaseIDCollection(): IDCollection {
  return { denomId: "", tokenIds: [] };
}

export const IDCollection = {
  encode(
    message: IDCollection,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.denomId !== "") {
      writer.uint32(10).string(message.denomId);
    }
    for (const v of message.tokenIds) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): IDCollection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseIDCollection();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denomId = reader.string();
          break;
        case 2:
          message.tokenIds.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): IDCollection {
    return {
      denomId: isSet(object.denomId) ? String(object.denomId) : "",
      tokenIds: Array.isArray(object?.tokenIds)
        ? object.tokenIds.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: IDCollection): unknown {
    const obj: any = {};
    message.denomId !== undefined && (obj.denomId = message.denomId);
    if (message.tokenIds) {
      obj.tokenIds = message.tokenIds.map((e) => e);
    } else {
      obj.tokenIds = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<IDCollection>, I>>(
    object: I
  ): IDCollection {
    const message = createBaseIDCollection();
    message.denomId = object.denomId ?? "";
    message.tokenIds = object.tokenIds?.map((e) => e) || [];
    return message;
  },
};

function createBaseOwner(): Owner {
  return { address: "", idCollections: [] };
}

export const Owner = {
  encode(message: Owner, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.idCollections) {
      IDCollection.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Owner {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOwner();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.idCollections.push(
            IDCollection.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Owner {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      idCollections: Array.isArray(object?.idCollections)
        ? object.idCollections.map((e: any) => IDCollection.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Owner): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.idCollections) {
      obj.idCollections = message.idCollections.map((e) =>
        e ? IDCollection.toJSON(e) : undefined
      );
    } else {
      obj.idCollections = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Owner>, I>>(object: I): Owner {
    const message = createBaseOwner();
    message.address = object.address ?? "";
    message.idCollections =
      object.idCollections?.map((e) => IDCollection.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCollection(): Collection {
  return { denom: undefined, nfts: [] };
}

export const Collection = {
  encode(
    message: Collection,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.denom !== undefined) {
      Denom.encode(message.denom, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.nfts) {
      BaseDomain.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Collection {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCollection();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = Denom.decode(reader, reader.uint32());
          break;
        case 2:
          message.nfts.push(BaseDomain.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Collection {
    return {
      denom: isSet(object.denom) ? Denom.fromJSON(object.denom) : undefined,
      nfts: Array.isArray(object?.nfts)
        ? object.nfts.map((e: any) => BaseDomain.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Collection): unknown {
    const obj: any = {};
    message.denom !== undefined &&
      (obj.denom = message.denom ? Denom.toJSON(message.denom) : undefined);
    if (message.nfts) {
      obj.nfts = message.nfts.map((e) =>
        e ? BaseDomain.toJSON(e) : undefined
      );
    } else {
      obj.nfts = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Collection>, I>>(
    object: I
  ): Collection {
    const message = createBaseCollection();
    message.denom =
      object.denom !== undefined && object.denom !== null
        ? Denom.fromPartial(object.denom)
        : undefined;
    message.nfts = object.nfts?.map((e) => BaseDomain.fromPartial(e)) || [];
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

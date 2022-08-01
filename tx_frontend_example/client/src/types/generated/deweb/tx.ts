/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.deweb";

export interface MsgSaveWallet {
  creator: string;
  address: string;
  encryptedKey: string;
  chain: string;
}

export interface MsgSaveWalletResponse {}

export interface MsgDeleteWallet {
  creator: string;
  address: string;
}

export interface MsgDeleteWalletResponse {}

export interface MsgConnectChain {
  creator: string;
  chain: string;
  address: string;
}

export interface MsgConnectChainResponse {}

export interface MsgDeleteChainConnect {
  creator: string;
  chain: string;
  address: string;
}

export interface MsgDeleteChainConnectResponse {}

function createBaseMsgSaveWallet(): MsgSaveWallet {
  return { creator: "", address: "", encryptedKey: "", chain: "" };
}

export const MsgSaveWallet = {
  encode(
    message: MsgSaveWallet,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSaveWallet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSaveWallet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSaveWallet {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
      encryptedKey: isSet(object.encryptedKey)
        ? String(object.encryptedKey)
        : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
    };
  },

  toJSON(message: MsgSaveWallet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    message.encryptedKey !== undefined &&
      (obj.encryptedKey = message.encryptedKey);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSaveWallet>, I>>(
    object: I
  ): MsgSaveWallet {
    const message = createBaseMsgSaveWallet();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    message.encryptedKey = object.encryptedKey ?? "";
    message.chain = object.chain ?? "";
    return message;
  },
};

function createBaseMsgSaveWalletResponse(): MsgSaveWalletResponse {
  return {};
}

export const MsgSaveWalletResponse = {
  encode(
    _: MsgSaveWalletResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgSaveWalletResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSaveWalletResponse();
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

  fromJSON(_: any): MsgSaveWalletResponse {
    return {};
  },

  toJSON(_: MsgSaveWalletResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSaveWalletResponse>, I>>(
    _: I
  ): MsgSaveWalletResponse {
    const message = createBaseMsgSaveWalletResponse();
    return message;
  },
};

function createBaseMsgDeleteWallet(): MsgDeleteWallet {
  return { creator: "", address: "" };
}

export const MsgDeleteWallet = {
  encode(
    message: MsgDeleteWallet,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteWallet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteWallet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteWallet {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgDeleteWallet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteWallet>, I>>(
    object: I
  ): MsgDeleteWallet {
    const message = createBaseMsgDeleteWallet();
    message.creator = object.creator ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgDeleteWalletResponse(): MsgDeleteWalletResponse {
  return {};
}

export const MsgDeleteWalletResponse = {
  encode(
    _: MsgDeleteWalletResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgDeleteWalletResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteWalletResponse();
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

  fromJSON(_: any): MsgDeleteWalletResponse {
    return {};
  },

  toJSON(_: MsgDeleteWalletResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteWalletResponse>, I>>(
    _: I
  ): MsgDeleteWalletResponse {
    const message = createBaseMsgDeleteWalletResponse();
    return message;
  },
};

function createBaseMsgConnectChain(): MsgConnectChain {
  return { creator: "", chain: "", address: "" };
}

export const MsgConnectChain = {
  encode(
    message: MsgConnectChain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgConnectChain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgConnectChain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chain = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConnectChain {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgConnectChain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chain !== undefined && (obj.chain = message.chain);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgConnectChain>, I>>(
    object: I
  ): MsgConnectChain {
    const message = createBaseMsgConnectChain();
    message.creator = object.creator ?? "";
    message.chain = object.chain ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgConnectChainResponse(): MsgConnectChainResponse {
  return {};
}

export const MsgConnectChainResponse = {
  encode(
    _: MsgConnectChainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgConnectChainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgConnectChainResponse();
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

  fromJSON(_: any): MsgConnectChainResponse {
    return {};
  },

  toJSON(_: MsgConnectChainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgConnectChainResponse>, I>>(
    _: I
  ): MsgConnectChainResponse {
    const message = createBaseMsgConnectChainResponse();
    return message;
  },
};

function createBaseMsgDeleteChainConnect(): MsgDeleteChainConnect {
  return { creator: "", chain: "", address: "" };
}

export const MsgDeleteChainConnect = {
  encode(
    message: MsgDeleteChainConnect,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chain !== "") {
      writer.uint32(18).string(message.chain);
    }
    if (message.address !== "") {
      writer.uint32(26).string(message.address);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgDeleteChainConnect {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteChainConnect();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chain = reader.string();
          break;
        case 3:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteChainConnect {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      chain: isSet(object.chain) ? String(object.chain) : "",
      address: isSet(object.address) ? String(object.address) : "",
    };
  },

  toJSON(message: MsgDeleteChainConnect): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chain !== undefined && (obj.chain = message.chain);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteChainConnect>, I>>(
    object: I
  ): MsgDeleteChainConnect {
    const message = createBaseMsgDeleteChainConnect();
    message.creator = object.creator ?? "";
    message.chain = object.chain ?? "";
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseMsgDeleteChainConnectResponse(): MsgDeleteChainConnectResponse {
  return {};
}

export const MsgDeleteChainConnectResponse = {
  encode(
    _: MsgDeleteChainConnectResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgDeleteChainConnectResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteChainConnectResponse();
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

  fromJSON(_: any): MsgDeleteChainConnectResponse {
    return {};
  },

  toJSON(_: MsgDeleteChainConnectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteChainConnectResponse>, I>>(
    _: I
  ): MsgDeleteChainConnectResponse {
    const message = createBaseMsgDeleteChainConnectResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SaveWallet(request: MsgSaveWallet): Promise<MsgSaveWalletResponse>;
  DeleteWallet(request: MsgDeleteWallet): Promise<MsgDeleteWalletResponse>;
  ConnectChain(request: MsgConnectChain): Promise<MsgConnectChainResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteChainConnect(
    request: MsgDeleteChainConnect
  ): Promise<MsgDeleteChainConnectResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SaveWallet = this.SaveWallet.bind(this);
    this.DeleteWallet = this.DeleteWallet.bind(this);
    this.ConnectChain = this.ConnectChain.bind(this);
    this.DeleteChainConnect = this.DeleteChainConnect.bind(this);
  }
  SaveWallet(request: MsgSaveWallet): Promise<MsgSaveWalletResponse> {
    const data = MsgSaveWallet.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "SaveWallet",
      data
    );
    return promise.then((data) =>
      MsgSaveWalletResponse.decode(new _m0.Reader(data))
    );
  }

  DeleteWallet(request: MsgDeleteWallet): Promise<MsgDeleteWalletResponse> {
    const data = MsgDeleteWallet.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "DeleteWallet",
      data
    );
    return promise.then((data) =>
      MsgDeleteWalletResponse.decode(new _m0.Reader(data))
    );
  }

  ConnectChain(request: MsgConnectChain): Promise<MsgConnectChainResponse> {
    const data = MsgConnectChain.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "ConnectChain",
      data
    );
    return promise.then((data) =>
      MsgConnectChainResponse.decode(new _m0.Reader(data))
    );
  }

  DeleteChainConnect(
    request: MsgDeleteChainConnect
  ): Promise<MsgDeleteChainConnectResponse> {
    const data = MsgDeleteChainConnect.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "DeleteChainConnect",
      data
    );
    return promise.then((data) =>
      MsgDeleteChainConnectResponse.decode(new _m0.Reader(data))
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

/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.domain.v1beta1";

/** MsgTransferNFT defines an SDK message for transferring an NFT to recipient. */
export interface MsgTransferDomain {
  id: string;
  price: Long;
  cancel: boolean;
  sender: string;
  recipient: string;
}

/** MsgTransferNFTResponse defines the Msg/TransferNFT response type. */
export interface MsgTransferDomainResponse {}

/** MsgEditNFT defines an SDK message for editing a nft. */
export interface MsgEditDomain {
  id: string;
  data: string;
  sender: string;
}

/** MsgEditNFTResponse defines the Msg/EditNFT response type. */
export interface MsgEditdomainResponse {}

/** MsgMintNFT defines an SDK message for creating a new NFT. */
export interface MsgRegisterDomain {
  id: string;
  data: string;
  sender: string;
  recipient: string;
}

/** MsgMintNFTResponse defines the Msg/MintNFT response type. */
export interface MsgRegisterDomainResponse {}

/** MsgBurnNFT defines an SDK message for burning a NFT. */
export interface MsgRemoveDomain {
  id: string;
  sender: string;
}

/** MsgBurnNFTResponse defines the Msg/BurnNFT response type. */
export interface MsgRemoveDomainResponse {}

function createBaseMsgTransferDomain(): MsgTransferDomain {
  return {
    id: "",
    price: Long.UZERO,
    cancel: false,
    sender: "",
    recipient: "",
  };
}

export const MsgTransferDomain = {
  encode(
    message: MsgTransferDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (!message.price.isZero()) {
      writer.uint32(16).uint64(message.price);
    }
    if (message.cancel === true) {
      writer.uint32(24).bool(message.cancel);
    }
    if (message.sender !== "") {
      writer.uint32(34).string(message.sender);
    }
    if (message.recipient !== "") {
      writer.uint32(42).string(message.recipient);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgTransferDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.price = reader.uint64() as Long;
          break;
        case 3:
          message.cancel = reader.bool();
          break;
        case 4:
          message.sender = reader.string();
          break;
        case 5:
          message.recipient = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgTransferDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      price: isSet(object.price) ? Long.fromValue(object.price) : Long.UZERO,
      cancel: isSet(object.cancel) ? Boolean(object.cancel) : false,
      sender: isSet(object.sender) ? String(object.sender) : "",
      recipient: isSet(object.recipient) ? String(object.recipient) : "",
    };
  },

  toJSON(message: MsgTransferDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.price !== undefined &&
      (obj.price = (message.price || Long.UZERO).toString());
    message.cancel !== undefined && (obj.cancel = message.cancel);
    message.sender !== undefined && (obj.sender = message.sender);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferDomain>, I>>(
    object: I
  ): MsgTransferDomain {
    const message = createBaseMsgTransferDomain();
    message.id = object.id ?? "";
    message.price =
      object.price !== undefined && object.price !== null
        ? Long.fromValue(object.price)
        : Long.UZERO;
    message.cancel = object.cancel ?? false;
    message.sender = object.sender ?? "";
    message.recipient = object.recipient ?? "";
    return message;
  },
};

function createBaseMsgTransferDomainResponse(): MsgTransferDomainResponse {
  return {};
}

export const MsgTransferDomainResponse = {
  encode(
    _: MsgTransferDomainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgTransferDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgTransferDomainResponse();
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

  fromJSON(_: any): MsgTransferDomainResponse {
    return {};
  },

  toJSON(_: MsgTransferDomainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgTransferDomainResponse>, I>>(
    _: I
  ): MsgTransferDomainResponse {
    const message = createBaseMsgTransferDomainResponse();
    return message;
  },
};

function createBaseMsgEditDomain(): MsgEditDomain {
  return { id: "", data: "", sender: "" };
}

export const MsgEditDomain = {
  encode(
    message: MsgEditDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.data !== "") {
      writer.uint32(18).string(message.data);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEditDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditDomain();
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
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      data: isSet(object.data) ? String(object.data) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
    };
  },

  toJSON(message: MsgEditDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.data !== undefined && (obj.data = message.data);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEditDomain>, I>>(
    object: I
  ): MsgEditDomain {
    const message = createBaseMsgEditDomain();
    message.id = object.id ?? "";
    message.data = object.data ?? "";
    message.sender = object.sender ?? "";
    return message;
  },
};

function createBaseMsgEditdomainResponse(): MsgEditdomainResponse {
  return {};
}

export const MsgEditdomainResponse = {
  encode(
    _: MsgEditdomainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgEditdomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditdomainResponse();
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

  fromJSON(_: any): MsgEditdomainResponse {
    return {};
  },

  toJSON(_: MsgEditdomainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEditdomainResponse>, I>>(
    _: I
  ): MsgEditdomainResponse {
    const message = createBaseMsgEditdomainResponse();
    return message;
  },
};

function createBaseMsgRegisterDomain(): MsgRegisterDomain {
  return { id: "", data: "", sender: "", recipient: "" };
}

export const MsgRegisterDomain = {
  encode(
    message: MsgRegisterDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.data !== "") {
      writer.uint32(18).string(message.data);
    }
    if (message.sender !== "") {
      writer.uint32(26).string(message.sender);
    }
    if (message.recipient !== "") {
      writer.uint32(34).string(message.recipient);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterDomain();
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
          message.sender = reader.string();
          break;
        case 4:
          message.recipient = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      data: isSet(object.data) ? String(object.data) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
      recipient: isSet(object.recipient) ? String(object.recipient) : "",
    };
  },

  toJSON(message: MsgRegisterDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.data !== undefined && (obj.data = message.data);
    message.sender !== undefined && (obj.sender = message.sender);
    message.recipient !== undefined && (obj.recipient = message.recipient);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterDomain>, I>>(
    object: I
  ): MsgRegisterDomain {
    const message = createBaseMsgRegisterDomain();
    message.id = object.id ?? "";
    message.data = object.data ?? "";
    message.sender = object.sender ?? "";
    message.recipient = object.recipient ?? "";
    return message;
  },
};

function createBaseMsgRegisterDomainResponse(): MsgRegisterDomainResponse {
  return {};
}

export const MsgRegisterDomainResponse = {
  encode(
    _: MsgRegisterDomainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgRegisterDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterDomainResponse();
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

  fromJSON(_: any): MsgRegisterDomainResponse {
    return {};
  },

  toJSON(_: MsgRegisterDomainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterDomainResponse>, I>>(
    _: I
  ): MsgRegisterDomainResponse {
    const message = createBaseMsgRegisterDomainResponse();
    return message;
  },
};

function createBaseMsgRemoveDomain(): MsgRemoveDomain {
  return { id: "", sender: "" };
}

export const MsgRemoveDomain = {
  encode(
    message: MsgRemoveDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.sender !== "") {
      writer.uint32(18).string(message.sender);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      sender: isSet(object.sender) ? String(object.sender) : "",
    };
  },

  toJSON(message: MsgRemoveDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveDomain>, I>>(
    object: I
  ): MsgRemoveDomain {
    const message = createBaseMsgRemoveDomain();
    message.id = object.id ?? "";
    message.sender = object.sender ?? "";
    return message;
  },
};

function createBaseMsgRemoveDomainResponse(): MsgRemoveDomainResponse {
  return {};
}

export const MsgRemoveDomainResponse = {
  encode(
    _: MsgRemoveDomainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): MsgRemoveDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveDomainResponse();
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

  fromJSON(_: any): MsgRemoveDomainResponse {
    return {};
  },

  toJSON(_: MsgRemoveDomainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveDomainResponse>, I>>(
    _: I
  ): MsgRemoveDomainResponse {
    const message = createBaseMsgRemoveDomainResponse();
    return message;
  },
};

/** Msg defines the nft Msg service. */
export interface Msg {
  /** MintNFT defines a method for mint a new nft */
  RegisterDomain(
    request: MsgRegisterDomain
  ): Promise<MsgRegisterDomainResponse>;
  /** RefundHTLC defines a method for editing a nft. */
  EditDomain(request: MsgEditDomain): Promise<MsgEditdomainResponse>;
  /** TransferNFT defines a method for transferring a nft. */
  TransferDomain(
    request: MsgTransferDomain
  ): Promise<MsgTransferDomainResponse>;
  /** BurnNFT defines a method for burning a nft. */
  RemoveDomain(request: MsgRemoveDomain): Promise<MsgRemoveDomainResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RegisterDomain = this.RegisterDomain.bind(this);
    this.EditDomain = this.EditDomain.bind(this);
    this.TransferDomain = this.TransferDomain.bind(this);
    this.RemoveDomain = this.RemoveDomain.bind(this);
  }
  RegisterDomain(
    request: MsgRegisterDomain
  ): Promise<MsgRegisterDomainResponse> {
    const data = MsgRegisterDomain.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Msg",
      "RegisterDomain",
      data
    );
    return promise.then((data) =>
      MsgRegisterDomainResponse.decode(new _m0.Reader(data))
    );
  }

  EditDomain(request: MsgEditDomain): Promise<MsgEditdomainResponse> {
    const data = MsgEditDomain.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Msg",
      "EditDomain",
      data
    );
    return promise.then((data) =>
      MsgEditdomainResponse.decode(new _m0.Reader(data))
    );
  }

  TransferDomain(
    request: MsgTransferDomain
  ): Promise<MsgTransferDomainResponse> {
    const data = MsgTransferDomain.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Msg",
      "TransferDomain",
      data
    );
    return promise.then((data) =>
      MsgTransferDomainResponse.decode(new _m0.Reader(data))
    );
  }

  RemoveDomain(request: MsgRemoveDomain): Promise<MsgRemoveDomainResponse> {
    const data = MsgRemoveDomain.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Msg",
      "RemoveDomain",
      data
    );
    return promise.then((data) =>
      MsgRemoveDomainResponse.decode(new _m0.Reader(data))
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

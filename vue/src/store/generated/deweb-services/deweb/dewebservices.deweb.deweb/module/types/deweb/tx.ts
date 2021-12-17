/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.deweb";

export interface MsgSaveUser {
  creator: string;
  message: string;
  chain: string;
}

export interface MsgSaveUserResponse {}

export interface MsgDeleteKey {
  creator: string;
  uuid: string;
}

export interface MsgDeleteKeyResponse {}

const baseMsgSaveUser: object = { creator: "", message: "", chain: "" };

export const MsgSaveUser = {
  encode(message: MsgSaveUser, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.message !== "") {
      writer.uint32(18).string(message.message);
    }
    if (message.chain !== "") {
      writer.uint32(26).string(message.chain);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSaveUser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSaveUser } as MsgSaveUser;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.message = reader.string();
          break;
        case 3:
          message.chain = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSaveUser {
    const message = { ...baseMsgSaveUser } as MsgSaveUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    return message;
  },

  toJSON(message: MsgSaveUser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.message !== undefined && (obj.message = message.message);
    message.chain !== undefined && (obj.chain = message.chain);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSaveUser>): MsgSaveUser {
    const message = { ...baseMsgSaveUser } as MsgSaveUser;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    return message;
  },
};

const baseMsgSaveUserResponse: object = {};

export const MsgSaveUserResponse = {
  encode(_: MsgSaveUserResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSaveUserResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSaveUserResponse } as MsgSaveUserResponse;
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

  fromJSON(_: any): MsgSaveUserResponse {
    const message = { ...baseMsgSaveUserResponse } as MsgSaveUserResponse;
    return message;
  },

  toJSON(_: MsgSaveUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgSaveUserResponse>): MsgSaveUserResponse {
    const message = { ...baseMsgSaveUserResponse } as MsgSaveUserResponse;
    return message;
  },
};

const baseMsgDeleteKey: object = { creator: "", uuid: "" };

export const MsgDeleteKey = {
  encode(message: MsgDeleteKey, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uuid !== "") {
      writer.uint32(18).string(message.uuid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteKey {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteKey } as MsgDeleteKey;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uuid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteKey {
    const message = { ...baseMsgDeleteKey } as MsgDeleteKey;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = String(object.uuid);
    } else {
      message.uuid = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteKey): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uuid !== undefined && (obj.uuid = message.uuid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteKey>): MsgDeleteKey {
    const message = { ...baseMsgDeleteKey } as MsgDeleteKey;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uuid !== undefined && object.uuid !== null) {
      message.uuid = object.uuid;
    } else {
      message.uuid = "";
    }
    return message;
  },
};

const baseMsgDeleteKeyResponse: object = {};

export const MsgDeleteKeyResponse = {
  encode(_: MsgDeleteKeyResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteKeyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteKeyResponse } as MsgDeleteKeyResponse;
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

  fromJSON(_: any): MsgDeleteKeyResponse {
    const message = { ...baseMsgDeleteKeyResponse } as MsgDeleteKeyResponse;
    return message;
  },

  toJSON(_: MsgDeleteKeyResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDeleteKeyResponse>): MsgDeleteKeyResponse {
    const message = { ...baseMsgDeleteKeyResponse } as MsgDeleteKeyResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  SaveUser(request: MsgSaveUser): Promise<MsgSaveUserResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteKey(request: MsgDeleteKey): Promise<MsgDeleteKeyResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SaveUser(request: MsgSaveUser): Promise<MsgSaveUserResponse> {
    const data = MsgSaveUser.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "SaveUser",
      data
    );
    return promise.then((data) => MsgSaveUserResponse.decode(new Reader(data)));
  }

  DeleteKey(request: MsgDeleteKey): Promise<MsgDeleteKeyResponse> {
    const data = MsgDeleteKey.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.deweb.deweb.Msg",
      "DeleteKey",
      data
    );
    return promise.then((data) =>
      MsgDeleteKeyResponse.decode(new Reader(data))
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

/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "dewebservices.deweb.deweb";

export interface UserKeyRec {
  creator: string;
  message: string;
  chain: string;
  deleted: boolean;
}

export interface RecordsToUser {
  records: string[];
}

const baseUserKeyRec: object = {
  creator: "",
  message: "",
  chain: "",
  deleted: false,
};

export const UserKeyRec = {
  encode(message: UserKeyRec, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
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

  decode(input: Reader | Uint8Array, length?: number): UserKeyRec {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseUserKeyRec } as UserKeyRec;
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

  fromJSON(object: any): UserKeyRec {
    const message = { ...baseUserKeyRec } as UserKeyRec;
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
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = Boolean(object.deleted);
    } else {
      message.deleted = false;
    }
    return message;
  },

  toJSON(message: UserKeyRec): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.message !== undefined && (obj.message = message.message);
    message.chain !== undefined && (obj.chain = message.chain);
    message.deleted !== undefined && (obj.deleted = message.deleted);
    return obj;
  },

  fromPartial(object: DeepPartial<UserKeyRec>): UserKeyRec {
    const message = { ...baseUserKeyRec } as UserKeyRec;
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
    if (object.deleted !== undefined && object.deleted !== null) {
      message.deleted = object.deleted;
    } else {
      message.deleted = false;
    }
    return message;
  },
};

const baseRecordsToUser: object = { records: "" };

export const RecordsToUser = {
  encode(message: RecordsToUser, writer: Writer = Writer.create()): Writer {
    for (const v of message.records) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): RecordsToUser {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseRecordsToUser } as RecordsToUser;
    message.records = [];
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
    const message = { ...baseRecordsToUser } as RecordsToUser;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(String(e));
      }
    }
    return message;
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

  fromPartial(object: DeepPartial<RecordsToUser>): RecordsToUser {
    const message = { ...baseRecordsToUser } as RecordsToUser;
    message.records = [];
    if (object.records !== undefined && object.records !== null) {
      for (const e of object.records) {
        message.records.push(e);
      }
    }
    return message;
  },
};

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

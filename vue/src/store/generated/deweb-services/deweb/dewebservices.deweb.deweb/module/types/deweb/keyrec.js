/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "dewebservices.deweb.deweb";
const baseUserWalletRec = {
    address: "",
    encryptedKey: "",
    chain: "",
    deleted: false,
};
export const UserWalletRec = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseUserWalletRec };
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
    fromJSON(object) {
        const message = { ...baseUserWalletRec };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        if (object.encryptedKey !== undefined && object.encryptedKey !== null) {
            message.encryptedKey = String(object.encryptedKey);
        }
        else {
            message.encryptedKey = "";
        }
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = String(object.chain);
        }
        else {
            message.chain = "";
        }
        if (object.deleted !== undefined && object.deleted !== null) {
            message.deleted = Boolean(object.deleted);
        }
        else {
            message.deleted = false;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        message.encryptedKey !== undefined &&
            (obj.encryptedKey = message.encryptedKey);
        message.chain !== undefined && (obj.chain = message.chain);
        message.deleted !== undefined && (obj.deleted = message.deleted);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseUserWalletRec };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        if (object.encryptedKey !== undefined && object.encryptedKey !== null) {
            message.encryptedKey = object.encryptedKey;
        }
        else {
            message.encryptedKey = "";
        }
        if (object.chain !== undefined && object.chain !== null) {
            message.chain = object.chain;
        }
        else {
            message.chain = "";
        }
        if (object.deleted !== undefined && object.deleted !== null) {
            message.deleted = object.deleted;
        }
        else {
            message.deleted = false;
        }
        return message;
    },
};
const baseRecordsToUser = { records: "" };
export const RecordsToUser = {
    encode(message, writer = Writer.create()) {
        for (const v of message.records) {
            writer.uint32(10).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRecordsToUser };
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
    fromJSON(object) {
        const message = { ...baseRecordsToUser };
        message.records = [];
        if (object.records !== undefined && object.records !== null) {
            for (const e of object.records) {
                message.records.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.records) {
            obj.records = message.records.map((e) => e);
        }
        else {
            obj.records = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRecordsToUser };
        message.records = [];
        if (object.records !== undefined && object.records !== null) {
            for (const e of object.records) {
                message.records.push(e);
            }
        }
        return message;
    },
};

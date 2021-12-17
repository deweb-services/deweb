/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
export const protobufPackage = "dewebservices.deweb.deweb";
const baseQueryGetKeyRecordRequest = { uuid: "" };
export const QueryGetKeyRecordRequest = {
    encode(message, writer = Writer.create()) {
        if (message.uuid !== "") {
            writer.uint32(10).string(message.uuid);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetKeyRecordRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetKeyRecordRequest,
        };
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = String(object.uuid);
        }
        else {
            message.uuid = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.uuid !== undefined && (obj.uuid = message.uuid);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetKeyRecordRequest,
        };
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = object.uuid;
        }
        else {
            message.uuid = "";
        }
        return message;
    },
};
const baseQueryGetKeyRecordResponse = {
    uuid: "",
    message: "",
    chain: "",
    deleted: false,
};
export const QueryGetKeyRecordResponse = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetKeyRecordResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetKeyRecordResponse,
        };
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = String(object.uuid);
        }
        else {
            message.uuid = "";
        }
        if (object.message !== undefined && object.message !== null) {
            message.message = String(object.message);
        }
        else {
            message.message = "";
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
        message.uuid !== undefined && (obj.uuid = message.uuid);
        message.message !== undefined && (obj.message = message.message);
        message.chain !== undefined && (obj.chain = message.chain);
        message.deleted !== undefined && (obj.deleted = message.deleted);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetKeyRecordResponse,
        };
        if (object.uuid !== undefined && object.uuid !== null) {
            message.uuid = object.uuid;
        }
        else {
            message.uuid = "";
        }
        if (object.message !== undefined && object.message !== null) {
            message.message = object.message;
        }
        else {
            message.message = "";
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
const baseQueryGetUserKeyRecordsRequest = { address: "" };
export const QueryGetUserKeyRecordsRequest = {
    encode(message, writer = Writer.create()) {
        if (message.address !== "") {
            writer.uint32(10).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetUserKeyRecordsRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetUserKeyRecordsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetUserKeyRecordsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
        }
        return message;
    },
};
const baseQueryGetUserKeyRecordsResponse = { uuids: "" };
export const QueryGetUserKeyRecordsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.uuids) {
            writer.uint32(10).string(v);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryGetUserKeyRecordsResponse,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryGetUserKeyRecordsResponse,
        };
        message.uuids = [];
        if (object.uuids !== undefined && object.uuids !== null) {
            for (const e of object.uuids) {
                message.uuids.push(String(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.uuids) {
            obj.uuids = message.uuids.map((e) => e);
        }
        else {
            obj.uuids = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryGetUserKeyRecordsResponse,
        };
        message.uuids = [];
        if (object.uuids !== undefined && object.uuids !== null) {
            for (const e of object.uuids) {
                message.uuids.push(e);
            }
        }
        return message;
    },
};
const baseQueryFilterUserKeyRecordsRequest = {
    address: "",
    chain: "",
    deleted: false,
};
export const QueryFilterUserKeyRecordsRequest = {
    encode(message, writer = Writer.create()) {
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
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryFilterUserKeyRecordsRequest,
        };
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
    fromJSON(object) {
        const message = {
            ...baseQueryFilterUserKeyRecordsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = "";
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
        message.chain !== undefined && (obj.chain = message.chain);
        message.deleted !== undefined && (obj.deleted = message.deleted);
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryFilterUserKeyRecordsRequest,
        };
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = "";
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
const baseQueryFilterUserKeyRecordsResponse = {};
export const QueryFilterUserKeyRecordsResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.records) {
            QueryGetKeyRecordResponse.encode(v, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = {
            ...baseQueryFilterUserKeyRecordsResponse,
        };
        message.records = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.records.push(QueryGetKeyRecordResponse.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = {
            ...baseQueryFilterUserKeyRecordsResponse,
        };
        message.records = [];
        if (object.records !== undefined && object.records !== null) {
            for (const e of object.records) {
                message.records.push(QueryGetKeyRecordResponse.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.records) {
            obj.records = message.records.map((e) => e ? QueryGetKeyRecordResponse.toJSON(e) : undefined);
        }
        else {
            obj.records = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = {
            ...baseQueryFilterUserKeyRecordsResponse,
        };
        message.records = [];
        if (object.records !== undefined && object.records !== null) {
            for (const e of object.records) {
                message.records.push(QueryGetKeyRecordResponse.fromPartial(e));
            }
        }
        return message;
    },
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    GetKeyRecord(request) {
        const data = QueryGetKeyRecordRequest.encode(request).finish();
        const promise = this.rpc.request("dewebservices.deweb.deweb.Query", "GetKeyRecord", data);
        return promise.then((data) => QueryGetKeyRecordResponse.decode(new Reader(data)));
    }
    GetUserKeyRecords(request) {
        const data = QueryGetUserKeyRecordsRequest.encode(request).finish();
        const promise = this.rpc.request("dewebservices.deweb.deweb.Query", "GetUserKeyRecords", data);
        return promise.then((data) => QueryGetUserKeyRecordsResponse.decode(new Reader(data)));
    }
    FilterUserKeyRecords(request) {
        const data = QueryFilterUserKeyRecordsRequest.encode(request).finish();
        const promise = this.rpc.request("dewebservices.deweb.deweb.Query", "FilterUserKeyRecords", data);
        return promise.then((data) => QueryFilterUserKeyRecordsResponse.decode(new Reader(data)));
    }
}

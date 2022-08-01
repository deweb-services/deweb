/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.domain";

/** Params defines the parameters for the module. */
export interface Params {
  domainPrice: Long;
  domainExpirationHours: Long;
  domainOwnerProlongationHours: Long;
  subDomainPrice: Long;
  blockTLDs: string[];
}

function createBaseParams(): Params {
  return {
    domainPrice: Long.UZERO,
    domainExpirationHours: Long.ZERO,
    domainOwnerProlongationHours: Long.ZERO,
    subDomainPrice: Long.UZERO,
    blockTLDs: [],
  };
}

export const Params = {
  encode(
    message: Params,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (!message.domainPrice.isZero()) {
      writer.uint32(8).uint64(message.domainPrice);
    }
    if (!message.domainExpirationHours.isZero()) {
      writer.uint32(16).int64(message.domainExpirationHours);
    }
    if (!message.domainOwnerProlongationHours.isZero()) {
      writer.uint32(24).int64(message.domainOwnerProlongationHours);
    }
    if (!message.subDomainPrice.isZero()) {
      writer.uint32(32).uint64(message.subDomainPrice);
    }
    for (const v of message.blockTLDs) {
      writer.uint32(42).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domainPrice = reader.uint64() as Long;
          break;
        case 2:
          message.domainExpirationHours = reader.int64() as Long;
          break;
        case 3:
          message.domainOwnerProlongationHours = reader.int64() as Long;
          break;
        case 4:
          message.subDomainPrice = reader.uint64() as Long;
          break;
        case 5:
          message.blockTLDs.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      domainPrice: isSet(object.domainPrice)
        ? Long.fromValue(object.domainPrice)
        : Long.UZERO,
      domainExpirationHours: isSet(object.domainExpirationHours)
        ? Long.fromValue(object.domainExpirationHours)
        : Long.ZERO,
      domainOwnerProlongationHours: isSet(object.domainOwnerProlongationHours)
        ? Long.fromValue(object.domainOwnerProlongationHours)
        : Long.ZERO,
      subDomainPrice: isSet(object.subDomainPrice)
        ? Long.fromValue(object.subDomainPrice)
        : Long.UZERO,
      blockTLDs: Array.isArray(object?.blockTLDs)
        ? object.blockTLDs.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.domainPrice !== undefined &&
      (obj.domainPrice = (message.domainPrice || Long.UZERO).toString());
    message.domainExpirationHours !== undefined &&
      (obj.domainExpirationHours = (
        message.domainExpirationHours || Long.ZERO
      ).toString());
    message.domainOwnerProlongationHours !== undefined &&
      (obj.domainOwnerProlongationHours = (
        message.domainOwnerProlongationHours || Long.ZERO
      ).toString());
    message.subDomainPrice !== undefined &&
      (obj.subDomainPrice = (message.subDomainPrice || Long.UZERO).toString());
    if (message.blockTLDs) {
      obj.blockTLDs = message.blockTLDs.map((e) => e);
    } else {
      obj.blockTLDs = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.domainPrice =
      object.domainPrice !== undefined && object.domainPrice !== null
        ? Long.fromValue(object.domainPrice)
        : Long.UZERO;
    message.domainExpirationHours =
      object.domainExpirationHours !== undefined &&
      object.domainExpirationHours !== null
        ? Long.fromValue(object.domainExpirationHours)
        : Long.ZERO;
    message.domainOwnerProlongationHours =
      object.domainOwnerProlongationHours !== undefined &&
      object.domainOwnerProlongationHours !== null
        ? Long.fromValue(object.domainOwnerProlongationHours)
        : Long.ZERO;
    message.subDomainPrice =
      object.subDomainPrice !== undefined && object.subDomainPrice !== null
        ? Long.fromValue(object.subDomainPrice)
        : Long.UZERO;
    message.blockTLDs = object.blockTLDs?.map((e) => e) || [];
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

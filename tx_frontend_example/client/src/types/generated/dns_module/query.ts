/* eslint-disable */
import { Params } from "./params";
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dewebservices.domain.v1beta1";

export interface TransferOffer {
  price: Long;
  ExpectedOwnerAddress: string;
}

export interface DNSRecords {
  type: string;
  values: string[];
}

export interface ResponseDomain {
  id: string;
  issued: string;
  validTill: string;
  transferOffer?: TransferOffer;
  records: DNSRecords[];
  subDomainsOnSale: boolean;
  subDomainsSalePrice: Long;
  owner: string;
}

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: Params;
}

/** QueryNFTRequest is the request type for the Query/NFT RPC method */
export interface QueryDomainRequest {
  domainName: string;
}

/** QueryNFTResponse is the response type for the Query/NFT RPC method */
export interface QueryDomainResponse {
  domain?: ResponseDomain;
}

/** QueryNFTRequest is the request type for the Query/NFT RPC method */
export interface QueryOwnedDomainsRequest {
  address: string;
  offset: Long;
  count: Long;
}

/** QueryNFTResponse is the response type for the Query/NFT RPC method */
export interface QueryOwnedDomainsResponse {
  total: Long;
  domains: ResponseDomain[];
}

function createBaseTransferOffer(): TransferOffer {
  return { price: Long.UZERO, ExpectedOwnerAddress: "" };
}

export const TransferOffer = {
  encode(
    message: TransferOffer,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (!message.price.isZero()) {
      writer.uint32(8).uint64(message.price);
    }
    if (message.ExpectedOwnerAddress !== "") {
      writer.uint32(18).string(message.ExpectedOwnerAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransferOffer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransferOffer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.price = reader.uint64() as Long;
          break;
        case 2:
          message.ExpectedOwnerAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TransferOffer {
    return {
      price: isSet(object.price) ? Long.fromValue(object.price) : Long.UZERO,
      ExpectedOwnerAddress: isSet(object.ExpectedOwnerAddress)
        ? String(object.ExpectedOwnerAddress)
        : "",
    };
  },

  toJSON(message: TransferOffer): unknown {
    const obj: any = {};
    message.price !== undefined &&
      (obj.price = (message.price || Long.UZERO).toString());
    message.ExpectedOwnerAddress !== undefined &&
      (obj.ExpectedOwnerAddress = message.ExpectedOwnerAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TransferOffer>, I>>(
    object: I
  ): TransferOffer {
    const message = createBaseTransferOffer();
    message.price =
      object.price !== undefined && object.price !== null
        ? Long.fromValue(object.price)
        : Long.UZERO;
    message.ExpectedOwnerAddress = object.ExpectedOwnerAddress ?? "";
    return message;
  },
};

function createBaseDNSRecords(): DNSRecords {
  return { type: "", values: [] };
}

export const DNSRecords = {
  encode(
    message: DNSRecords,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    for (const v of message.values) {
      writer.uint32(18).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DNSRecords {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDNSRecords();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.values.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DNSRecords {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      values: Array.isArray(object?.values)
        ? object.values.map((e: any) => String(e))
        : [],
    };
  },

  toJSON(message: DNSRecords): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    if (message.values) {
      obj.values = message.values.map((e) => e);
    } else {
      obj.values = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DNSRecords>, I>>(
    object: I
  ): DNSRecords {
    const message = createBaseDNSRecords();
    message.type = object.type ?? "";
    message.values = object.values?.map((e) => e) || [];
    return message;
  },
};

function createBaseResponseDomain(): ResponseDomain {
  return {
    id: "",
    issued: "",
    validTill: "",
    transferOffer: undefined,
    records: [],
    subDomainsOnSale: false,
    subDomainsSalePrice: Long.UZERO,
    owner: "",
  };
}

export const ResponseDomain = {
  encode(
    message: ResponseDomain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.issued !== "") {
      writer.uint32(18).string(message.issued);
    }
    if (message.validTill !== "") {
      writer.uint32(26).string(message.validTill);
    }
    if (message.transferOffer !== undefined) {
      TransferOffer.encode(
        message.transferOffer,
        writer.uint32(34).fork()
      ).ldelim();
    }
    for (const v of message.records) {
      DNSRecords.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    if (message.subDomainsOnSale === true) {
      writer.uint32(48).bool(message.subDomainsOnSale);
    }
    if (!message.subDomainsSalePrice.isZero()) {
      writer.uint32(56).uint64(message.subDomainsSalePrice);
    }
    if (message.owner !== "") {
      writer.uint32(66).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ResponseDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseResponseDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.issued = reader.string();
          break;
        case 3:
          message.validTill = reader.string();
          break;
        case 4:
          message.transferOffer = TransferOffer.decode(reader, reader.uint32());
          break;
        case 5:
          message.records.push(DNSRecords.decode(reader, reader.uint32()));
          break;
        case 6:
          message.subDomainsOnSale = reader.bool();
          break;
        case 7:
          message.subDomainsSalePrice = reader.uint64() as Long;
          break;
        case 8:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ResponseDomain {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      issued: isSet(object.issued) ? String(object.issued) : "",
      validTill: isSet(object.validTill) ? String(object.validTill) : "",
      transferOffer: isSet(object.transferOffer)
        ? TransferOffer.fromJSON(object.transferOffer)
        : undefined,
      records: Array.isArray(object?.records)
        ? object.records.map((e: any) => DNSRecords.fromJSON(e))
        : [],
      subDomainsOnSale: isSet(object.subDomainsOnSale)
        ? Boolean(object.subDomainsOnSale)
        : false,
      subDomainsSalePrice: isSet(object.subDomainsSalePrice)
        ? Long.fromValue(object.subDomainsSalePrice)
        : Long.UZERO,
      owner: isSet(object.owner) ? String(object.owner) : "",
    };
  },

  toJSON(message: ResponseDomain): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.issued !== undefined && (obj.issued = message.issued);
    message.validTill !== undefined && (obj.validTill = message.validTill);
    message.transferOffer !== undefined &&
      (obj.transferOffer = message.transferOffer
        ? TransferOffer.toJSON(message.transferOffer)
        : undefined);
    if (message.records) {
      obj.records = message.records.map((e) =>
        e ? DNSRecords.toJSON(e) : undefined
      );
    } else {
      obj.records = [];
    }
    message.subDomainsOnSale !== undefined &&
      (obj.subDomainsOnSale = message.subDomainsOnSale);
    message.subDomainsSalePrice !== undefined &&
      (obj.subDomainsSalePrice = (
        message.subDomainsSalePrice || Long.UZERO
      ).toString());
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ResponseDomain>, I>>(
    object: I
  ): ResponseDomain {
    const message = createBaseResponseDomain();
    message.id = object.id ?? "";
    message.issued = object.issued ?? "";
    message.validTill = object.validTill ?? "";
    message.transferOffer =
      object.transferOffer !== undefined && object.transferOffer !== null
        ? TransferOffer.fromPartial(object.transferOffer)
        : undefined;
    message.records =
      object.records?.map((e) => DNSRecords.fromPartial(e)) || [];
    message.subDomainsOnSale = object.subDomainsOnSale ?? false;
    message.subDomainsSalePrice =
      object.subDomainsSalePrice !== undefined &&
      object.subDomainsSalePrice !== null
        ? Long.fromValue(object.subDomainsSalePrice)
        : Long.UZERO;
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(
    _: QueryParamsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(
    _: I
  ): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(
    object: I
  ): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params =
      object.params !== undefined && object.params !== null
        ? Params.fromPartial(object.params)
        : undefined;
    return message;
  },
};

function createBaseQueryDomainRequest(): QueryDomainRequest {
  return { domainName: "" };
}

export const QueryDomainRequest = {
  encode(
    message: QueryDomainRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.domainName !== "") {
      writer.uint32(10).string(message.domainName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDomainRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDomainRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domainName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDomainRequest {
    return {
      domainName: isSet(object.domainName) ? String(object.domainName) : "",
    };
  },

  toJSON(message: QueryDomainRequest): unknown {
    const obj: any = {};
    message.domainName !== undefined && (obj.domainName = message.domainName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDomainRequest>, I>>(
    object: I
  ): QueryDomainRequest {
    const message = createBaseQueryDomainRequest();
    message.domainName = object.domainName ?? "";
    return message;
  },
};

function createBaseQueryDomainResponse(): QueryDomainResponse {
  return { domain: undefined };
}

export const QueryDomainResponse = {
  encode(
    message: QueryDomainResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.domain !== undefined) {
      ResponseDomain.encode(message.domain, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDomainResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domain = ResponseDomain.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDomainResponse {
    return {
      domain: isSet(object.domain)
        ? ResponseDomain.fromJSON(object.domain)
        : undefined,
    };
  },

  toJSON(message: QueryDomainResponse): unknown {
    const obj: any = {};
    message.domain !== undefined &&
      (obj.domain = message.domain
        ? ResponseDomain.toJSON(message.domain)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDomainResponse>, I>>(
    object: I
  ): QueryDomainResponse {
    const message = createBaseQueryDomainResponse();
    message.domain =
      object.domain !== undefined && object.domain !== null
        ? ResponseDomain.fromPartial(object.domain)
        : undefined;
    return message;
  },
};

function createBaseQueryOwnedDomainsRequest(): QueryOwnedDomainsRequest {
  return { address: "", offset: Long.ZERO, count: Long.ZERO };
}

export const QueryOwnedDomainsRequest = {
  encode(
    message: QueryOwnedDomainsRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (!message.offset.isZero()) {
      writer.uint32(16).int64(message.offset);
    }
    if (!message.count.isZero()) {
      writer.uint32(24).int64(message.count);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryOwnedDomainsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOwnedDomainsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.offset = reader.int64() as Long;
          break;
        case 3:
          message.count = reader.int64() as Long;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOwnedDomainsRequest {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      offset: isSet(object.offset) ? Long.fromValue(object.offset) : Long.ZERO,
      count: isSet(object.count) ? Long.fromValue(object.count) : Long.ZERO,
    };
  },

  toJSON(message: QueryOwnedDomainsRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.offset !== undefined &&
      (obj.offset = (message.offset || Long.ZERO).toString());
    message.count !== undefined &&
      (obj.count = (message.count || Long.ZERO).toString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOwnedDomainsRequest>, I>>(
    object: I
  ): QueryOwnedDomainsRequest {
    const message = createBaseQueryOwnedDomainsRequest();
    message.address = object.address ?? "";
    message.offset =
      object.offset !== undefined && object.offset !== null
        ? Long.fromValue(object.offset)
        : Long.ZERO;
    message.count =
      object.count !== undefined && object.count !== null
        ? Long.fromValue(object.count)
        : Long.ZERO;
    return message;
  },
};

function createBaseQueryOwnedDomainsResponse(): QueryOwnedDomainsResponse {
  return { total: Long.ZERO, domains: [] };
}

export const QueryOwnedDomainsResponse = {
  encode(
    message: QueryOwnedDomainsResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (!message.total.isZero()) {
      writer.uint32(8).int64(message.total);
    }
    for (const v of message.domains) {
      ResponseDomain.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): QueryOwnedDomainsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryOwnedDomainsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.total = reader.int64() as Long;
          break;
        case 2:
          message.domains.push(ResponseDomain.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOwnedDomainsResponse {
    return {
      total: isSet(object.total) ? Long.fromValue(object.total) : Long.ZERO,
      domains: Array.isArray(object?.domains)
        ? object.domains.map((e: any) => ResponseDomain.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QueryOwnedDomainsResponse): unknown {
    const obj: any = {};
    message.total !== undefined &&
      (obj.total = (message.total || Long.ZERO).toString());
    if (message.domains) {
      obj.domains = message.domains.map((e) =>
        e ? ResponseDomain.toJSON(e) : undefined
      );
    } else {
      obj.domains = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryOwnedDomainsResponse>, I>>(
    object: I
  ): QueryOwnedDomainsResponse {
    const message = createBaseQueryOwnedDomainsResponse();
    message.total =
      object.total !== undefined && object.total !== null
        ? Long.fromValue(object.total)
        : Long.ZERO;
    message.domains =
      object.domains?.map((e) => ResponseDomain.fromPartial(e)) || [];
    return message;
  },
};

/** Query defines the gRPC querier service for NFT module */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Query domain info by domain name */
  Domain(request: QueryDomainRequest): Promise<QueryDomainResponse>;
  /** Query domains owned by user */
  OwnedDomains(
    request: QueryOwnedDomainsRequest
  ): Promise<QueryOwnedDomainsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Domain = this.Domain.bind(this);
    this.OwnedDomains = this.OwnedDomains.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Query",
      "Params",
      data
    );
    return promise.then((data) =>
      QueryParamsResponse.decode(new _m0.Reader(data))
    );
  }

  Domain(request: QueryDomainRequest): Promise<QueryDomainResponse> {
    const data = QueryDomainRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Query",
      "Domain",
      data
    );
    return promise.then((data) =>
      QueryDomainResponse.decode(new _m0.Reader(data))
    );
  }

  OwnedDomains(
    request: QueryOwnedDomainsRequest
  ): Promise<QueryOwnedDomainsResponse> {
    const data = QueryOwnedDomainsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "dewebservices.domain.v1beta1.Query",
      "OwnedDomains",
      data
    );
    return promise.then((data) =>
      QueryOwnedDomainsResponse.decode(new _m0.Reader(data))
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

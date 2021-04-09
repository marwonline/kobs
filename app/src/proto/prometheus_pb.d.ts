// package: plugins.prometheus
// file: prometheus.proto

import * as jspb from "google-protobuf";

export class GetVariablesRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getTimestart(): number;
  setTimestart(value: number): void;

  getTimeend(): number;
  setTimeend(value: number): void;

  getResolution(): string;
  setResolution(value: string): void;

  clearVariablesList(): void;
  getVariablesList(): Array<Variable>;
  setVariablesList(value: Array<Variable>): void;
  addVariables(value?: Variable, index?: number): Variable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVariablesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetVariablesRequest): GetVariablesRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVariablesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVariablesRequest;
  static deserializeBinaryFromReader(message: GetVariablesRequest, reader: jspb.BinaryReader): GetVariablesRequest;
}

export namespace GetVariablesRequest {
  export type AsObject = {
    name: string,
    timestart: number,
    timeend: number,
    resolution: string,
    variablesList: Array<Variable.AsObject>,
  }
}

export class GetVariablesResponse extends jspb.Message {
  clearVariablesList(): void;
  getVariablesList(): Array<Variable>;
  setVariablesList(value: Array<Variable>): void;
  addVariables(value?: Variable, index?: number): Variable;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetVariablesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetVariablesResponse): GetVariablesResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetVariablesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetVariablesResponse;
  static deserializeBinaryFromReader(message: GetVariablesResponse, reader: jspb.BinaryReader): GetVariablesResponse;
}

export namespace GetVariablesResponse {
  export type AsObject = {
    variablesList: Array<Variable.AsObject>,
  }
}

export class GetMetricsRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getTimestart(): number;
  setTimestart(value: number): void;

  getTimeend(): number;
  setTimeend(value: number): void;

  getResolution(): string;
  setResolution(value: string): void;

  clearVariablesList(): void;
  getVariablesList(): Array<Variable>;
  setVariablesList(value: Array<Variable>): void;
  addVariables(value?: Variable, index?: number): Variable;

  clearQueriesList(): void;
  getQueriesList(): Array<Query>;
  setQueriesList(value: Array<Query>): void;
  addQueries(value?: Query, index?: number): Query;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMetricsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetMetricsRequest): GetMetricsRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMetricsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMetricsRequest;
  static deserializeBinaryFromReader(message: GetMetricsRequest, reader: jspb.BinaryReader): GetMetricsRequest;
}

export namespace GetMetricsRequest {
  export type AsObject = {
    name: string,
    timestart: number,
    timeend: number,
    resolution: string,
    variablesList: Array<Variable.AsObject>,
    queriesList: Array<Query.AsObject>,
  }
}

export class GetMetricsResponse extends jspb.Message {
  clearMetricsList(): void;
  getMetricsList(): Array<Metrics>;
  setMetricsList(value: Array<Metrics>): void;
  addMetrics(value?: Metrics, index?: number): Metrics;

  clearInterpolatedqueriesList(): void;
  getInterpolatedqueriesList(): Array<string>;
  setInterpolatedqueriesList(value: Array<string>): void;
  addInterpolatedqueries(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetMetricsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetMetricsResponse): GetMetricsResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetMetricsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetMetricsResponse;
  static deserializeBinaryFromReader(message: GetMetricsResponse, reader: jspb.BinaryReader): GetMetricsResponse;
}

export namespace GetMetricsResponse {
  export type AsObject = {
    metricsList: Array<Metrics.AsObject>,
    interpolatedqueriesList: Array<string>,
  }
}

export class Metrics extends jspb.Message {
  getLabel(): string;
  setLabel(value: string): void;

  getMin(): number;
  setMin(value: number): void;

  getMax(): number;
  setMax(value: number): void;

  clearDataList(): void;
  getDataList(): Array<Data>;
  setDataList(value: Array<Data>): void;
  addData(value?: Data, index?: number): Data;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metrics.AsObject;
  static toObject(includeInstance: boolean, msg: Metrics): Metrics.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Metrics, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metrics;
  static deserializeBinaryFromReader(message: Metrics, reader: jspb.BinaryReader): Metrics;
}

export namespace Metrics {
  export type AsObject = {
    label: string,
    min: number,
    max: number,
    dataList: Array<Data.AsObject>,
  }
}

export class Data extends jspb.Message {
  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Data.AsObject;
  static toObject(includeInstance: boolean, msg: Data): Data.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Data, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Data;
  static deserializeBinaryFromReader(message: Data, reader: jspb.BinaryReader): Data;
}

export namespace Data {
  export type AsObject = {
    x: number,
    y: number,
  }
}

export class Spec extends jspb.Message {
  clearVariablesList(): void;
  getVariablesList(): Array<Variable>;
  setVariablesList(value: Array<Variable>): void;
  addVariables(value?: Variable, index?: number): Variable;

  clearChartsList(): void;
  getChartsList(): Array<Chart>;
  setChartsList(value: Array<Chart>): void;
  addCharts(value?: Chart, index?: number): Chart;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Spec.AsObject;
  static toObject(includeInstance: boolean, msg: Spec): Spec.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Spec, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Spec;
  static deserializeBinaryFromReader(message: Spec, reader: jspb.BinaryReader): Spec;
}

export namespace Spec {
  export type AsObject = {
    variablesList: Array<Variable.AsObject>,
    chartsList: Array<Chart.AsObject>,
  }
}

export class Variable extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getLabel(): string;
  setLabel(value: string): void;

  getQuery(): string;
  setQuery(value: string): void;

  getAllowall(): boolean;
  setAllowall(value: boolean): void;

  clearValuesList(): void;
  getValuesList(): Array<string>;
  setValuesList(value: Array<string>): void;
  addValues(value: string, index?: number): string;

  getValue(): string;
  setValue(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Variable.AsObject;
  static toObject(includeInstance: boolean, msg: Variable): Variable.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Variable, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Variable;
  static deserializeBinaryFromReader(message: Variable, reader: jspb.BinaryReader): Variable;
}

export namespace Variable {
  export type AsObject = {
    name: string,
    label: string,
    query: string,
    allowall: boolean,
    valuesList: Array<string>,
    value: string,
  }
}

export class Chart extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): void;

  getType(): string;
  setType(value: string): void;

  getUnit(): string;
  setUnit(value: string): void;

  getStacked(): boolean;
  setStacked(value: boolean): void;

  getSize(): number;
  setSize(value: number): void;

  clearQueriesList(): void;
  getQueriesList(): Array<Query>;
  setQueriesList(value: Array<Query>): void;
  addQueries(value?: Query, index?: number): Query;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Chart.AsObject;
  static toObject(includeInstance: boolean, msg: Chart): Chart.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Chart, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Chart;
  static deserializeBinaryFromReader(message: Chart, reader: jspb.BinaryReader): Chart;
}

export namespace Chart {
  export type AsObject = {
    title: string,
    type: string,
    unit: string,
    stacked: boolean,
    size: number,
    queriesList: Array<Query.AsObject>,
  }
}

export class Query extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): void;

  getLabel(): string;
  setLabel(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Query.AsObject;
  static toObject(includeInstance: boolean, msg: Query): Query.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Query, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Query;
  static deserializeBinaryFromReader(message: Query, reader: jspb.BinaryReader): Query;
}

export namespace Query {
  export type AsObject = {
    query: string,
    label: string,
  }
}

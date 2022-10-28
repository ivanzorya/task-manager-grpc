/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "main";

export interface Task {
  id: string;
  subject: string;
  done: boolean;
}

export interface TasksRequest {
}

export interface TasksResponse {
  tasks: Task[];
}

export interface CreateTaskRequest {
  task: Task | undefined;
}

export interface UpdateTaskRequest {
  task: Task | undefined;
}

export interface DeleteTaskRequest {
  id: string;
}

export interface SuccessResponse {
  success: boolean;
}

function createBaseTask(): Task {
  return { id: "", subject: "", done: false };
}

export const Task = {
  encode(message: Task, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.subject !== "") {
      writer.uint32(18).string(message.subject);
    }
    if (message.done === true) {
      writer.uint32(24).bool(message.done);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Task {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTask();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.subject = reader.string();
          break;
        case 3:
          message.done = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<Task>): Task {
    const message = createBaseTask();
    message.id = object.id ?? "";
    message.subject = object.subject ?? "";
    message.done = object.done ?? false;
    return message;
  },
};

function createBaseTasksRequest(): TasksRequest {
  return {};
}

export const TasksRequest = {
  encode(_: TasksRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TasksRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTasksRequest();
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

  fromPartial(_: DeepPartial<TasksRequest>): TasksRequest {
    const message = createBaseTasksRequest();
    return message;
  },
};

function createBaseTasksResponse(): TasksResponse {
  return { tasks: [] };
}

export const TasksResponse = {
  encode(message: TasksResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.tasks) {
      Task.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TasksResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTasksResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tasks.push(Task.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<TasksResponse>): TasksResponse {
    const message = createBaseTasksResponse();
    message.tasks = object.tasks?.map((e) => Task.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCreateTaskRequest(): CreateTaskRequest {
  return { task: undefined };
}

export const CreateTaskRequest = {
  encode(message: CreateTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.task !== undefined) {
      Task.encode(message.task, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateTaskRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.task = Task.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<CreateTaskRequest>): CreateTaskRequest {
    const message = createBaseCreateTaskRequest();
    message.task = (object.task !== undefined && object.task !== null) ? Task.fromPartial(object.task) : undefined;
    return message;
  },
};

function createBaseUpdateTaskRequest(): UpdateTaskRequest {
  return { task: undefined };
}

export const UpdateTaskRequest = {
  encode(message: UpdateTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.task !== undefined) {
      Task.encode(message.task, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UpdateTaskRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.task = Task.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<UpdateTaskRequest>): UpdateTaskRequest {
    const message = createBaseUpdateTaskRequest();
    message.task = (object.task !== undefined && object.task !== null) ? Task.fromPartial(object.task) : undefined;
    return message;
  },
};

function createBaseDeleteTaskRequest(): DeleteTaskRequest {
  return { id: "" };
}

export const DeleteTaskRequest = {
  encode(message: DeleteTaskRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteTaskRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteTaskRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<DeleteTaskRequest>): DeleteTaskRequest {
    const message = createBaseDeleteTaskRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseSuccessResponse(): SuccessResponse {
  return { success: false };
}

export const SuccessResponse = {
  encode(message: SuccessResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.success === true) {
      writer.uint32(8).bool(message.success);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SuccessResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSuccessResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.success = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromPartial(object: DeepPartial<SuccessResponse>): SuccessResponse {
    const message = createBaseSuccessResponse();
    message.success = object.success ?? false;
    return message;
  },
};

export type TaskServiceDefinition = typeof TaskServiceDefinition;
export const TaskServiceDefinition = {
  name: "TaskService",
  fullName: "main.TaskService",
  methods: {
    queryGetTasks: {
      name: "QueryGetTasks",
      requestType: TasksRequest,
      requestStream: false,
      responseType: TasksResponse,
      responseStream: false,
      options: {},
    },
    queryCreateTask: {
      name: "QueryCreateTask",
      requestType: CreateTaskRequest,
      requestStream: false,
      responseType: SuccessResponse,
      responseStream: false,
      options: {},
    },
    queryUpdateTask: {
      name: "QueryUpdateTask",
      requestType: UpdateTaskRequest,
      requestStream: false,
      responseType: SuccessResponse,
      responseStream: false,
      options: {},
    },
    queryDeleteTask: {
      name: "QueryDeleteTask",
      requestType: DeleteTaskRequest,
      requestStream: false,
      responseType: SuccessResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface TaskServiceServiceImplementation<CallContextExt = {}> {
  queryGetTasks(request: TasksRequest, context: CallContext & CallContextExt): Promise<DeepPartial<TasksResponse>>;
  queryCreateTask(
    request: CreateTaskRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<SuccessResponse>>;
  queryUpdateTask(
    request: UpdateTaskRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<SuccessResponse>>;
  queryDeleteTask(
    request: DeleteTaskRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<SuccessResponse>>;
}

export interface TaskServiceClient<CallOptionsExt = {}> {
  queryGetTasks(request: DeepPartial<TasksRequest>, options?: CallOptions & CallOptionsExt): Promise<TasksResponse>;
  queryCreateTask(
    request: DeepPartial<CreateTaskRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<SuccessResponse>;
  queryUpdateTask(
    request: DeepPartial<UpdateTaskRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<SuccessResponse>;
  queryDeleteTask(
    request: DeepPartial<DeleteTaskRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<SuccessResponse>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

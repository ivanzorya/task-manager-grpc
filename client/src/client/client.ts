import { createChannel, createClient } from "nice-grpc-web";
import { TaskServiceClient, TaskServiceDefinition } from "../app/proto/tasks";


const channel = createChannel("http://localhost:9090");

export const client: TaskServiceClient = createClient(
    TaskServiceDefinition,
  channel,
);
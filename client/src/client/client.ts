import { createChannel, createClient } from "nice-grpc-web";
import { TaskServiceClient, TaskServiceDefinition } from "../app/proto/tasks";


const channel = createChannel(process.env.REACT_APP_API_HOST!);

export const client: TaskServiceClient = createClient(
    TaskServiceDefinition,
  channel,
);
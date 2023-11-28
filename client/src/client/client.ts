import { createChannel, createClient } from "nice-grpc-web";
import { TaskServiceClient, TaskServiceDefinition } from "../app/proto/tasks";
import { UserServiceClient, UserServiceDefinition } from "../app/proto/users";
import { AuthServiceClient, AuthServiceDefinition } from "../app/proto/auth";


const channel = createChannel(process.env.REACT_APP_API_HOST!);

export const clientTasks: TaskServiceClient = createClient(
    TaskServiceDefinition,
    channel,
);

export const clientUsers: UserServiceClient = createClient(
    UserServiceDefinition,
    channel,
);

export const clientAuth: AuthServiceClient = createClient(
    AuthServiceDefinition,
    channel,
);

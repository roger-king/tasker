interface Task {
    taskId: string;
    entryId: number;
    name: string;
    executor: string;
    schedule: string;
    isRepeatable: boolean;
    enabled: boolean;
    complete: boolean;
    args: Record<string, any>;
    createdAt: Date;
    updatedAt: Date;
    deletedAt: Date;
}

interface NewTaskInput {
    name: string;
    description: string;
    schedule: string;
    args: Record<string, any>;
    executor: string;
}

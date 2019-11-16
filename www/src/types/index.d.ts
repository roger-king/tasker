type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
type Optionalize<T extends K, K> = Omit<T, keyof K>;

interface Task {
    taskId: string;
    entryId: number;
    name: string;
    description: string;
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

interface GithubOAuthLoginResponse {
    access_token: string;
    error: string;
    error_description: string;
    error_uri: string;
    scope: string;
    token_type: string;
}

interface OAuthProviderClientIdResponse {
    client_id: string;
}

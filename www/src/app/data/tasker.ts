export const listTasks = async (): Promise<{ data: [Task] }> => {
    const data = await fetch(`/tasker/tasks`);
    return data.json();
};

export const findTask = async (id: string): Promise<{ data: Task }> => {
    const data = await fetch(`/tasker/tasks/${id}`);
    return data.json();
};

export const createTask = async (input: NewTaskInput): Promise<{ data: Task }> => {
    const data = await fetch(`/tasker/tasks`, { method: 'POST', body: JSON.stringify(input) });
    return data.json();
};

export const disableTask = async (id: string): Promise<{ data: string }> => {
    const data = await fetch(`/tasker/tasks/${id}/disable`, { method: 'PATCH' });
    return data.json();
};

export const deleteTask = async (id: string): Promise<{ data: string }> => {
    const data = await fetch(`/tasker/tasks/${id}`, { method: 'DELETE' });
    return data.json();
};

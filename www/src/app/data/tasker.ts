export const listTasks = async (): Promise<{ data: [Task] }> => {
    const data = await fetch(`/tasker/tasks`);
    return data.json();
};

export const findTask = async (id: string): Promise<{ data: Task }> => {
    const data = await fetch(`/tasker/tasks/${id}`);
    return data.json();
};

export const createTask = (input: NewTaskInput) => {
    return fetch(`/tasker/tasks`, { method: 'POST', body: JSON.stringify(input) });
};

export const disableTask = (id: string) => {
    return fetch(`/tasker/tasks/${id}/disable`, { method: 'PATCH' });
};

export const deleteTask = (id: string) => {
    return fetch(`/tasker/tasks/${id}`, { method: 'DELETE' });
};

export const listSettings = async (): Promise<{ data: [Setting] }> => {
    const data = await fetch(`/tasker/settings/plugin`);
    return data.json();
};

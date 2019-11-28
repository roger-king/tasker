export const listSettings = async (filters: Map<string, string> = new Map()): Promise<{ data: [Setting] }> => {
    let q = `?`;
    if (filters.size > 0) {
        filters.forEach((v: string, k: string) => {
            q += `${k}=${v}&`;
        });
    }

    const data = await fetch(`/tasker/settings/plugin${q}`);
    return data.json();
};

export const createSetting = async (input: Setting): Promise<{ data: Setting }> => {
    const data = await fetch(`/tasker/settings/plugin`, { method: 'POST', body: JSON.stringify(input) });
    return data.json();
};

export const toggleActivePluginSetting = async (input: TogglePluginSetting): Promise<boolean> => {
    const data = await fetch(`/tasker/settings/plugin/toggle`, { method: 'PATCH', body: JSON.stringify(input) });
    const response = await data.json();
    return response.data;
};

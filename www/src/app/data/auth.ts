export const authenticate = async (code: string) => {
    const data = await fetch(`/authenticate/${code}`);
    return data.json();
};

export const range = (start: number, end: number) => {
    return Array(end - start + 1)
        .fill(undefined)
        .map((_: number, idx: number) => start + idx);
};

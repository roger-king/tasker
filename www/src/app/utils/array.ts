export const range = (start: number, end: number) => {
    return Array(end - start + 1)
        .fill(undefined)
        .map((_: number, idx: number) => {
            const num = String(start + idx);
            if (num.length === 1) {
                return `0${num}`;
            }

            return num;
        });
};

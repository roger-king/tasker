// Array builder for numbers - probably should rename this.
export const range = (start: number, end: number): string[] => {
    return Array(end - start + 1)
        .fill(undefined)
        .map((_: number, idx: number): string => {
            const num = String(start + idx);
            if (num.length === 1) {
                return `0${num}`;
            }

            return num;
        });
};

export function setForTime<T>(
    setter: (value: T) => void, valueBefore: T, valueAfter: T, waitTime: number,
    ): void {
        setter(valueBefore);

        setTimeout(() => {
            setter(valueAfter);
        }, waitTime);
}

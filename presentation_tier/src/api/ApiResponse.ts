export interface ApiResponse<T> {
    isError: boolean;
    code: number;
    errors: string[];
    data: T;
}
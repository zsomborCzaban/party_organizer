export interface ApiResponse<T> {
    isError: boolean;
    code: number;
    errors: ApiError[];
    data: T;
}

export interface ApiError {
    field: string;
    value: any;
    err: string;
}
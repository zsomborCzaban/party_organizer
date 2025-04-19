export interface ApiResponse<T> {
  is_error: boolean;
  code: number;
  errors: string | ApiError[];
  data: T;
}

export interface ApiError {
  field: string;
  value?: string;
  err: string;
}

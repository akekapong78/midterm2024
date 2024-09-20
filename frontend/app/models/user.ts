export interface User {
  id: number;
  username: string;
  password: string;
  role: "ADMIN" | "USER";
}

export interface UserRequest {
  username: string;
  password: string;
}

export interface UserResponse {
  id: number;
  username: string;
  role: "ADMIN" | "USER";
}
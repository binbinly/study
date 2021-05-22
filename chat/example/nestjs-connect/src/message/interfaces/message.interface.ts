export interface User {
  id: number;
  name: string;
  avatar: string;
}

export interface SendRequest {
  id: string;
  from: User;
  to_id: number;
  chat_type: number;
  type: number;
  options: string;
  content: string;
  t: number;
}

export interface NotifyRequest {
  type: string;
  to_id: number;
}

export interface MomentRequest {
  type: string;
  to_id: number;
  user_id: number;
  avatar: string;
}

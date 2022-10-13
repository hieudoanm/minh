export type WebhookRequestBody = {
  update_id: number;
  message: Message;
};

export type Message = {
  message_id: number;
  from: User;
  chat: Chat;
  date: number;
  text: string;
};

export type User = {
  id: number;
  is_bot: boolean;
  first_name: string;
  last_name: string;
  username: string;
  language_code: string;
};

export type Chat = {
  id: number;
  first_name: string;
  last_name: string;
  username: string;
  type: string;
};

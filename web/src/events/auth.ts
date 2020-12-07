import { EventEmitter } from 'tools/event';

type AuthEventTypes = {
  LogoutAuthEvent: string; // 登出
  UpdateTokenAuthEvent: string; // 更新 token
};

export const authEventEmitter = new EventEmitter<AuthEventTypes>();

export function emitLogoutAuthEvent() {
  authEventEmitter.emit('LogoutAuthEvent');
}

export function emitUpdateTokenAuthEvent(token: string) {
  authEventEmitter.emit('UpdateTokenAuthEvent', token);
}

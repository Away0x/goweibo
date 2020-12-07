import { EventEmitter } from 'tools/event';

type GlobalEventTypes = {
  ShowGlobalHTTPErrorMessage: string;
};

export const globalEventEmitter = new EventEmitter<GlobalEventTypes>();

// 显示 http error message 事件
export function emitShowGlobalHTTPErrorMessageEvent(text?: React.ReactNode) {
  globalEventEmitter.emit('ShowGlobalHTTPErrorMessage', text);
}

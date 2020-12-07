import mitt, { Emitter, Handler } from 'mitt';

export class EventEmitter<EventTypes extends object> {
  private eventEmitter: Emitter;

  constructor() {
    this.eventEmitter = mitt();
  }

  public on(eventType: keyof EventTypes, handler: Handler) {
    this.eventEmitter.on(String(eventType), handler);
  }

  public off(eventType: keyof EventTypes, handler: Handler) {
    this.eventEmitter.off(String(eventType), handler);
  }

  public onoff(eventType: keyof EventTypes, handler: Handler) {
    this.on(eventType, handler);
    return () => {
      this.off(eventType, handler);
    };
  }

  public emit(eventType: keyof EventTypes, eventArg?: any) {
    this.eventEmitter.emit(String(eventType), eventArg);
  }
}

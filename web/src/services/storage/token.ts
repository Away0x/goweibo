import { TOKEN_KEY } from 'config';
import Cookies from 'tools/cookie';

export default class TokenStorage {
  public static readonly key = TOKEN_KEY;
  private static tokenStore: null | string = null;

  public static set(token: string) {
    Cookies.set(this.key, token);
    this.tokenStore = token;
  }

  public static get(): string {
    if (this.tokenStore) {
      return this.tokenStore;
    }

    const token = Cookies.get(this.key);
    this.tokenStore = token;
    return this.tokenStore || '';
  }

  public static clean() {
    Cookies.remove(this.key);
    this.tokenStore = null;
  }
}

(window as any).__TokenStorage = TokenStorage;

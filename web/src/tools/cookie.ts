class Cookies {
  private static cookies = (window as any).Cookies || null;

  public static set(key: string, val: string, attr?: { [key: string]: any }) {
    this.cookies.set(key, val, attr);
  }

  public static get(key: string, json?: boolean): any {
    return this.cookies.get(key, json);
  }

  public static remove(key: string, attr?: { [key: string]: any }): void {
    this.cookies.remove(key, attr);
  }
}

export default Cookies;

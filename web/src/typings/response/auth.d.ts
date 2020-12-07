declare namespace Response {
  declare namespace Auth {
    /** 登录 */
    export interface Login {
      token: string;
      user: Data.UserData;
    }
  }
}

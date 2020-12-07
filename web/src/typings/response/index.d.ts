declare namespace Response {
  /** 处理后的接口通用响应类型 */
  export interface CommonApiResponse<T = any> {
    status: boolean;
    message: string;
    data: T;
  }
}

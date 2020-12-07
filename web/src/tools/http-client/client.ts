import Axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import qs from 'qs';

type ErrorResp<T = any> = {
  resolved: boolean;
  result: T;
};

export interface RequestConfig extends AxiosRequestConfig {
  json?: boolean;
  formData?: boolean;
  getMockData?: (config: RequestConfig) => any;
}

const JSON_CONTENT_TYPE = { 'Content-Type': 'application/json;charset=UTF-8' };
const FORMDATA_CONTENT_TYPE = { 'Content-Type': 'multipart/form-data;charset=UTF-8' };

export class HTTPClient<Req extends RequestConfig = RequestConfig, Resp = any> {
  private axiosInstance: AxiosInstance;
  private baseURL?: string;

  constructor(
    options: RequestConfig,
    // 处理响应
    private responseResolve: (requestConfig: Req, response: any) => Resp,
    // 错误处理
    private errorResolve: (requestConfig?: Req | null, err?: any) => Resp,
    // 请求参数处理
    private requestResolve?: (requestConfig: Req) => Req,
  ) {
    this.baseURL = options.baseURL;
    this.axiosInstance = Axios.create(options);

    this.axiosInstance.interceptors.request.use(
      (config: RequestConfig) => {
        return config;
      },
      (error: any) => {
        return {
          resolved: true,
          result: this.errorResolve(null, error),
        } as ErrorResp<Resp>;
      },
    );

    this.axiosInstance.interceptors.response.use(
      (response: any) => {
        return response;
      },
      (error: any) => {
        return {
          resolved: true,
          result: this.errorResolve(null, error),
        } as ErrorResp<Resp>;
      },
    );
  }

  private async request(config: Req): Promise<Resp> {
    if (!config.baseURL && this.baseURL) {
      config.baseURL = this.baseURL;
    }

    if (this.requestResolve) {
      config = this.requestResolve(config);
    }

    if (config.getMockData) {
      return config.getMockData(config);
    }

    try {
      if (config.method === 'POST') {
        if (!config.data) config.data = {};

        if (config.formData) {
          const form = new FormData();

          for (const key in config.data) {
            if (config.data.hasOwnProperty(key)) {
              const val = config.data[key];
              form.append(key, val);
            }
          }

          config.data = form;
          config.headers = Object.assign({}, config.headers, FORMDATA_CONTENT_TYPE);
        } else {
          if (!config.json) {
            config.data = qs.stringify(config.data);
          } else {
            config.headers = Object.assign({}, config.headers, JSON_CONTENT_TYPE);
          }
        }
      }

      const result = await this.axiosInstance(config);

      if ((result as any).resolved) {
        return (result as any).result;
      }

      return this.responseResolve(config, result);
    } catch (err) {
      console.warn(`[HTTPClient] `, err);
      return this.errorResolve(config, err);
    }
  }

  public updateBaseURL(baseURL: string) {
    this.baseURL = baseURL;
  }

  public get(config: Req) {
    return this.request(Object.assign({}, config, { method: 'GET' }));
  }

  public post(config: Req) {
    return this.request(Object.assign({}, config, { method: 'POST' }));
  }
}

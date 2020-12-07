import { HTTPClient, RequestConfig } from 'tools/http-client';
import { API_ROOT } from 'config';
import MOCK_DATA from 'services/mock';
import { ApiResponseCode } from 'constants/http';
import { emitLogoutAuthEvent } from 'events/auth';
import { emitShowGlobalHTTPErrorMessageEvent } from 'events/global';
import TokenStorage from 'services/storage/token';

export interface CommonRequestConfig extends RequestConfig {
  hideGlobalErrorToast?: boolean; // 隐藏全局 error toast
  dontAutoBringToken?: boolean; // 不用自动携带 token
  mock?: boolean; // 是否使用 mock 数据 (会调用对应的 getMockData)
}

export const commonHttpClient = new HTTPClient<CommonRequestConfig, Response.CommonApiResponse>(
  {
    baseURL: API_ROOT,
    timeout: 1000 * 25,
  },
  // 处理响应
  (requestConfig, response) => {
    if (response.data.code === ApiResponseCode.TOKEN_ERROR) {
      console.warn('用户未登录', response);
      if (!requestConfig.hideGlobalErrorToast) {
        emitShowGlobalHTTPErrorMessageEvent(response.data.msg || '用户未登录');
      }

      emitLogoutAuthEvent();
      return { status: false, message: '用户未登录', data: null };
    }

    return {
      status: true,
      message: '',
      data: response.data,
    };
  },
  // 错误处理
  (requestConfig, err) => {
    console.error(err);
    const errMsg = '网络连接异常，请稍后重试';

    if (!requestConfig || !requestConfig.hideGlobalErrorToast) {
      emitShowGlobalHTTPErrorMessageEvent(errMsg);
    }

    return {
      status: false,
      message: errMsg,
      data: null,
    };
  },
  // 请求参数处理
  (requestConfig) => {
    if (!requestConfig.dontAutoBringToken) {
      const token = TokenStorage.get();

      requestConfig.headers = requestConfig.headers || {};
      if (!requestConfig.headers['React-Cloud-Music-Token']) {
        requestConfig.headers['React-Cloud-Music-Token'] = token;
      }
    }

    if (requestConfig.mock) {
      requestConfig.getMockData = getMockData;
    }

    return requestConfig;
  },
);

function getMockData(config: CommonRequestConfig): Response.CommonApiResponse {
  const key = config.url || '';
  let mockData = MOCK_DATA[key];

  if (!mockData) {
    const msg = `[getMockData] mock data(${getMockData}) not found`;
    console.error(msg);
    return {
      status: false,
      message: msg,
      data: null,
    };
  }

  if (typeof mockData === 'function') {
    mockData = mockData(config);
  }

  return {
    status: true,
    message: '',
    data: mockData,
  };
}

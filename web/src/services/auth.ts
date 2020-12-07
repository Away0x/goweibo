import { commonHttpClient as client } from './common-http-client';

/** 登录 */
export async function loginService(
  username: string,
  password: string,
): Promise<Response.CommonApiResponse<Response.Auth.Login>> {
  const result = await client.post({
    url: '/login',
    mock: true,
    data: { username, password },
  });

  return result;
}

/** 获取用户信息 */
export async function getUserService(): Promise<Response.CommonApiResponse<Data.Auth.UserData | null>> {
  const result = await client.get({
    url: '/get-user',
    mock: true,
  });

  return result;
}

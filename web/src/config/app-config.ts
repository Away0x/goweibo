const W: any = window;

const APP_CONFIG: { [key: string]: any } = W.APP_CONFIG;
if (!APP_CONFIG) {
  console.error('APP-CONFIG.js 不存在，请检查!');
}

// 接口地址
export const API_ROOT: string = APP_CONFIG.API_ROOT || '';

// ---------------------------- key ----------------------------
export const LOCALSTROAGE_PREFIX: string = APP_CONFIG.LOCALSTROAGE_PREFIX || 'goweibo_app_'; // localstroage prefix
export const TOKEN_KEY: string = APP_CONFIG.TOKEN_KEY || 'goweibo_app_token'; // token key

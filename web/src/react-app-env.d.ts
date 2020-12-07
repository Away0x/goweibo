/// <reference types="react-scripts" />

import 'styled-components';

declare namespace NodeJS {
  interface ProcessEnv {
    readonly NODE_ENV: 'development' | 'production' | 'test';
    readonly PUBLIC_URL: string;
    // project
    readonly version: string;
    readonly branche: string;
    readonly buildtime: string;
    readonly desc: string;
  }
}

declare module 'styled-components' {
  export interface DefaultTheme {
    /** color */
    themeColor: string;
  }
}

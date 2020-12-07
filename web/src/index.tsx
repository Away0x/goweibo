import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';

ReactDOM.render(<App />, document.getElementById('root'));

// 存储项目信息
(window as any).__project__ = (() => {
  return {
    version: process.env.version,
    branche: process.env.branche,
    buildtime: process.env.buildtime,
    desc: process.env.desc,
  };
})();

import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root'),
);

// 存储项目信息
(window as any).__project__ = (() => {
  return {
    version: process.env.version,
    branche: process.env.branche,
    buildtime: process.env.buildtime,
    desc: process.env.desc,
  };
})();

reportWebVitals();

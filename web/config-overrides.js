const { override, addBabelPlugins, addBundleVisualizer, fixBabelImports, addLessLoader } = require('customize-cra');
const { overrideProcessEnv } = require('cra-define-override');
const dayjs = require('dayjs');

const envMap = {
  version: '1.0.0',
  branche: 'master',
  buildtime: dayjs().format('YYYY-MM-DD HH:mm:ss'),
  desc: '',
};

Object.keys(envMap).forEach((k) => (envMap[k] = JSON.stringify(envMap[k])));

module.exports = override(
  // styled-components
  addBabelPlugins(['babel-plugin-styled-components']),
  // antd
  fixBabelImports('import', {
    libraryName: 'antd',
    libraryDirectory: 'es',
    style: true,
  }),
  addLessLoader({
    lessOptions: {
      javascriptEnabled: true,
      modifyVars: { '@primary-color': '#f47983' },
    },
  }),
  // define env
  overrideProcessEnv(envMap),
  // bundle analyze
  addBundleVisualizer(
    {
      analyzerMode: 'static',
      reportFilename: 'report.html',
    },
    true,
  ),
);

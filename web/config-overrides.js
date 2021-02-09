const {
  override,
  addBabelPlugins,
  addBundleVisualizer,
  fixBabelImports,
  addLessLoader,
  addPostcssPlugins,
} = require('customize-cra');
const { overrideProcessEnv } = require('cra-define-override');
const dayjs = require('dayjs');

const envMap = {
  version: '1.0.0',
  branche: 'v2',
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
      modifyVars: { '@primary-color': '#d44439' },
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
  // postcss
  addPostcssPlugins([require('tailwindcss'), require('autoprefixer')]),
);

const mix = require('laravel-mix');

const resourcesDir = 'resources'
const scriptsDist = 'public/scripts'
const stylesDist = 'public/styles'
const scriptsPath = (file, ext = 'ts') => `${resourcesDir}/scripts/${file}.${ext}`
const stylesPath = (file, ext = 'scss') => `${resourcesDir}/styles/${file}.${ext}`

mix
  .disableNotifications()
  .setPublicPath('public/')
  .options({
    postCss: [
      require("tailwindcss"),
    ],
  })

mix.ts(scriptsPath('app'), scriptsDist)

mix
  .sass(stylesPath('app'), stylesDist)


let mix = require('laravel-mix');


mix
  .setPublicPath('public/')
  .ts('resources/js/app.ts', 'public/js')
  .sass('resources/sass/app.scss', 'public/css').version();

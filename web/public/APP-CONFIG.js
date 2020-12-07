(function (name, definition, context) {
  if (typeof module !== 'undefined' && module.exports) module.exports = definition();
  else if (typeof context['define'] === 'function' && (context['define']['amd'] || context['define']['cmd']))
    define(definition);
  else context[name] = definition();
})(
  'APP_CONFIG',
  function () {
    return {
      // ---------------------------- api ----------------------------
      API_ROOT: '/api',

      // ---------------------------- key ----------------------------
      LOCALSTROAGE_PREFIX: 'goweibo_app_', // localstroage prefix
      TOKEN_KEY: 'goweibo_app_token' // token key

      // ---------------------------- dev ----------------------------
    };
  },
  this
);

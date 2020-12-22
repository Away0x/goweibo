<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Welcome {{ APP_NAME }} in {{ APP_RUNMODE }} ({% route 'welcome' %})</title>
</head>
<body>
  <h2>Welcome {{ APP_NAME }}</h2>

  <a href="/apidoc/index.html">API document</a>

  <p>{{ time }}</p>
</body>
</html>

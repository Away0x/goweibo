import React, { lazy } from 'react';
import { Route, Switch } from 'react-router-dom';

import { SpecialRoutePath } from 'constants/router';
import { useFirstLoad } from 'containers/AuthContainer';
import Loading from 'components/Loading';
import Suspense from 'components/Suspense';
import { GuestRoute } from 'components/Route';

import NotFound from 'pages/Errors/NotFound';
const Login = lazy(() => import(/* webpackChunkName: 'login-page' */ 'pages/Login'));
const Home = lazy(() => import(/* webpackChunkName: 'home-page' */ 'pages/Home'));

function RootRoutes() {
  const { ready } = useFirstLoad();

  if (!ready) return <Loading full />;

  return (
    <Switch>
      {/* 首页 */}
      <Route exact path={SpecialRoutePath.Root}>
        <Suspense>
          <Home />
        </Suspense>
      </Route>

      {/* 登录 */}
      <GuestRoute exact path={SpecialRoutePath.Login} homeRoutePath={SpecialRoutePath.Root}>
        <Suspense>
          <Login />
        </Suspense>
      </GuestRoute>

      {/* not found */}
      <Route path={SpecialRoutePath.Any} component={NotFound} />
    </Switch>
  );
}

export default React.memo(RootRoutes);

import React from 'react';
import { Route, Redirect, RouteProps } from 'react-router-dom';

import AuthContainer from 'containers/AuthContainer';

interface GuestRouteProps extends RouteProps {
  homeRoutePath: string; // 首页地址
}

/** 必须未登录才可访问 */
function GuestRoute({ homeRoutePath, children, ...rest }: GuestRouteProps) {
  const { unLogin } = AuthContainer.useContainer();

  return <Route {...rest} render={() => (unLogin ? children : <Redirect to={{ pathname: homeRoutePath }} />)} />;
}

export default React.memo(GuestRoute);

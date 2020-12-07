import React from 'react';
import { Button } from 'antd';
import { useHistory } from 'react-router-dom';

import { SpecialRoutePath } from 'constants/router';
import AuthContainer from 'containers/AuthContainer';

import StyledHome from './style';

function Home() {
  const history = useHistory();

  const { logoutAction, logged, userData } = AuthContainer.useContainer();

  return (
    <StyledHome>
      {logged ? (
        <p>
          {JSON.stringify(userData)}
          <Button
            onClick={() => {
              logoutAction();
              history.push(SpecialRoutePath.Login);
            }}>
            登出
          </Button>
        </p>
      ) : (
        <Button
          onClick={() => {
            history.push(SpecialRoutePath.Login);
          }}>
          登录
        </Button>
      )}
    </StyledHome>
  );
}

export default React.memo(Home);

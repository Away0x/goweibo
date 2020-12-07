import { useMemo } from 'react';
import { HashRouter as Router } from 'react-router-dom';

import { ThemeState, ThemeProvider } from 'containers/ThemeContainer';
import GlobalContainer from 'containers/GlobalContainer';
import AuthContainer, { AuthState } from 'containers/AuthContainer';
import StyledGlobal from 'styles/global';
import { GlobalErrorBoundary } from 'components/ErrorBoundary';
import TokenStorage from 'services/storage/token';
import RootRoutes from 'routes';

interface InitState {
  theme: ThemeState;
  auth: AuthState;
}

function getDefaultInitData(): InitState {
  return {
    theme: {},
    auth: {
      token: TokenStorage.get(),
      userData: null,
    },
  };
}

function App() {
  const initData = useMemo(() => {
    const data = getDefaultInitData();
    console.log('[App#initData]', data);
    return data;
  }, []);

  return (
    <GlobalErrorBoundary>
      <GlobalContainer.Provider>
        <ThemeProvider initialState={initData.theme}>
          <AuthContainer.Provider initialState={initData.auth}>
            {/* 全局样式 */}
            <StyledGlobal />
            {/* 路由 */}
            <Router>
              <RootRoutes />
            </Router>
          </AuthContainer.Provider>
        </ThemeProvider>
      </GlobalContainer.Provider>
    </GlobalErrorBoundary>
  );
}

export default App;

import React, { useCallback } from 'react';
import { createContainer } from 'unstated-next';
import { useImmer } from 'use-immer';
import { DefaultTheme, ThemeProvider as StyledThemeProvider } from 'styled-components';

export type ThemeState = Partial<DefaultTheme>;

export const defaultTheme: DefaultTheme = {
  themeColor: '#d44439',
};

interface ThemeActions {
  changeTheme: (theme: ThemeState) => void;
}

type UseTheme = DefaultTheme & ThemeActions;

function useTheme(initialState?: ThemeState | null): UseTheme {
  const [themeState, updateThemeState] = useImmer<DefaultTheme>({ ...initialState, ...defaultTheme });

  const changeTheme = useCallback(
    (theme: ThemeState) => {
      updateThemeState((state) => {
        for (const key in theme) {
          if (theme.hasOwnProperty(key)) {
            (state as any)[key] = (theme as any)[key];
          }
        }
      });
    },
    [updateThemeState],
  );

  return {
    ...themeState,
    changeTheme,
  };
}

const ThemeContainer = createContainer(useTheme);

interface ThemeProviderProps {
  initialState?: ThemeState | null;
  children: React.ReactNode;
}

function IntlStyledThemeProvider({ children }: ThemeProviderProps) {
  const theme = ThemeContainer.useContainer();

  return <StyledThemeProvider theme={theme}>{children}</StyledThemeProvider>;
}

function ThemeProvider({ initialState, children }: ThemeProviderProps) {
  return (
    <ThemeContainer.Provider initialState={initialState}>
      <IntlStyledThemeProvider>{children}</IntlStyledThemeProvider>
    </ThemeContainer.Provider>
  );
}

// 监听 theme 的变化
function useListenThemeChange() {
  return ThemeContainer.useContainer();
}

export default ThemeContainer;

export { ThemeProvider, useListenThemeChange };

import React, { useCallback, useEffect } from 'react';
import { createContainer } from 'unstated-next';
import { message } from 'antd';

import { globalEventEmitter } from 'events/global';

export interface GlobalState {}

interface GlobalComputedState {}

interface GlobalActions {}

type UseGlobal = GlobalState & GlobalComputedState & GlobalActions;

function useGlobal(): UseGlobal {
  const handleShowHTTPErrorMessage = useCallback((text?: React.ReactNode) => {
    message.error(text);
  }, []);

  useEffect(() => {
    const off = globalEventEmitter.onoff('ShowGlobalHTTPErrorMessage', handleShowHTTPErrorMessage);
    return () => off();
  }, [handleShowHTTPErrorMessage]);

  return {};
}

const GlobalContainer = createContainer(useGlobal);

export default GlobalContainer;

import React from 'react';

import { IS_DEV } from 'config';

import StyledGlobalErrorBoundary, { Inner, Button, Code } from './Global.style';
import { netErrorImg } from './global-icon';

interface GlobalErrorBoundaryState {
  hasError: boolean;
  error: null | any;
  errorInfo: null | any;
}

export default class GlobalErrorBoundary extends React.Component<{}, GlobalErrorBoundaryState> {
  public state: GlobalErrorBoundaryState = {
    hasError: false,
    error: null,
    errorInfo: null,
  };

  public componentDidCatch(error: any, errorInfo: any) {
    this.setState({
      hasError: true,
      error,
      errorInfo,
    });
  }

  public reload = () => {
    window.location.reload();
  };

  private renderErrorStack = () => {
    if (!IS_DEV) return null;
    if (!this.state.errorInfo || !this.state.errorInfo.componentStack) return null;
    return (
      <Code>
        <code>{this.state.errorInfo.componentStack}</code>
      </Code>
    );
  };

  public render() {
    if (!this.state.hasError) {
      return this.props.children;
    }

    return (
      <StyledGlobalErrorBoundary>
        {this.renderErrorStack()}

        <Inner>
          <img src={netErrorImg} alt="网络错误" />
          <h3>网络连接异常</h3>
          <p>请检查网络后点击刷新</p>
          <Button onClick={this.reload}>点击刷新</Button>
        </Inner>
      </StyledGlobalErrorBoundary>
    );
  }
}

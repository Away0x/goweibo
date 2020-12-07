import React from 'react';

import { StyledFullLoadingContainer, StyledLoading } from './style';

interface LoadingProps {
  full?: boolean;
}

function Loading({ full = false }: LoadingProps) {
  if (full) {
    return (
      <StyledFullLoadingContainer>
        <StyledLoading>
          <div></div>
          <div></div>
        </StyledLoading>
      </StyledFullLoadingContainer>
    );
  }

  return (
    <StyledLoading>
      <div></div>
      <div></div>
    </StyledLoading>
  );
}

export default React.memo(Loading);

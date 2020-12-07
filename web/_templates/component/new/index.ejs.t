---
to: src/components/<%= name %>/index.tsx
---

import React from 'react';

import Styled<%= name %> from './style';

interface <%= name %>Props {
  children?: React.ReactNode;
}

function <%= name %>({
  children,
  ...rest
}: <%= name %>Props) {
  return (
    <Styled<%= name %> {...rest}>
      {children}
    </Styled<%= name %>>
  );
}

export default React.memo(<%= name %>);

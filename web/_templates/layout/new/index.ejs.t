---
to: src/layouts/<%= name %>/index.tsx
---

import React from 'react';

import Styled<%= name %> from './style';

interface <%= name %>LayoutProps {
  children?: React.ReactNode;
}

function <%= name %>Layout({
  children
}: <%= name %>LayoutProps) {
  return (
    <Styled<%= name %>>
      {children}
    </Styled<%= name %>>
  );
}

export default React.memo(<%= name %>Layout);

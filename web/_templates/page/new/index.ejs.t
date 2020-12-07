---
to: src/pages/<%= name %>/index.tsx
---

import React from 'react';

import Styled<%= name %> from './style';

function <%= name %>() {
  return (
    <Styled<%= name %>>
      <%= name %> page
    </Styled<%= name %>>
  );
}

export default React.memo(<%= name %>);

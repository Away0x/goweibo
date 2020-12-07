---
to: src/components/<%= name %>/<%= h.changeCase.lcFirst(name) %>.stories.tsx
---

import React from 'react';

import <%= name %> from '.';

export default {
  title: 'Components/<%= name %>',
  component: <%= name %>
};

export const Default = () => {
  return <<%= name %>><%= name %> component</<%= name %>>;
};

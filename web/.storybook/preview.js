import React from 'react';
import { addDecorator, addParameters } from '@storybook/react';

import './preview.css';

import { ThemeProvider } from 'containers/ThemeContainer';
import StyledGlobal from '../src/styles/global';

addDecorator((storyFn) => (
  <ThemeProvider>
    <StyledGlobal />
    {storyFn()}
  </ThemeProvider>
));

addParameters({
  options: {
    showRoots: true,
  },
});

import React, { Suspense as ReactSuspense } from 'react';

import Loading from 'components/Loading';

interface SuspenseProps {
  children?: React.ReactNode;
  component?: React.ReactNode;
}

function Suspense({ children, component }: SuspenseProps) {
  return <ReactSuspense fallback={<Loading full />}>{component || children}</ReactSuspense>;
}

export default React.memo(Suspense);

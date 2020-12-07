---
to: src/containers/<%= name %>Container.ts
---
import { createContainer } from 'unstated-next';
import { useImmer } from 'use-immer';

export interface <%= name %>State {}

interface <%= name %>ComputedState {}

interface <%= name %>Actions {}

type Use<%= name %> = <%= name %>State & <%= name %>ComputedState & <%= name %>Actions;

function use<%= name %>(): Use<%= name %> {
  const [<%= h.changeCase.lcFirst(name) %>State, update<%= name %>State] = useImmer<<%= name %>State>({});

  return { ...<%= h.changeCase.lcFirst(name) %>State };
}

const <%= name %>Container = createContainer(use<%= name %>);

export default <%= name %>Container;

import { AgentPool, ManagedCluster, Resource } from '@azure/arm-containerservice';

export interface IManagedCluster extends Resource {
  location?: string;
  properties?: ManagedCluster;
}

export interface INodePool extends Resource {
  properties?: AgentPool;
}
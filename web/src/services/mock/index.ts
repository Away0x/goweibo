import { authData } from './auth';

interface MockData {
  [k: string]: any;
}

const mockData: MockData = {
  ...authData,
};

export default mockData;

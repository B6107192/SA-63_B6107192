import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import SaveData from './components/SaveData';
 
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/savedata', SaveData);
  },
});
 

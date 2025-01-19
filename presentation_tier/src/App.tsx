import { AppRouter } from '../../frontend/src/AppRouter';
import { NavigationBar } from './components/navigation-bar/NavigationBar';

export const App = () => (
  <>
    <NavigationBar />
    <AppRouter />;
  </>
);

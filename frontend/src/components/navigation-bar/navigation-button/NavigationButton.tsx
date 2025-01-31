import { useNavigate } from 'react-router-dom';
import classes from './NavigationButton.module.scss';

interface Props {
  buttonText: string;
  navigateToLink: string;
  isActive?: boolean;
}

export const NavigationButton = ({ buttonText, navigateToLink, isActive }: Props) => {
  const navigate = useNavigate();

  return (
    <button
      className={`${classes.button} ${isActive ? classes.active : ''}`}
      onClick={() => navigate(navigateToLink)}
    >
      {buttonText}
    </button>
  );
};

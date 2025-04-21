import { useNavigate } from 'react-router-dom';
import classes from './NavigationButtonDarkTheme.module.scss';

interface Props {
    buttonText: string;
    navigateToLink: string;
    isActive?: boolean;
}

export const NavigationButtonDarkTheme = ({ buttonText, navigateToLink, isActive }: Props) => {
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

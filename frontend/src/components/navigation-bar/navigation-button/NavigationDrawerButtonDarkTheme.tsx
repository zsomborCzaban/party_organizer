import classes from './NavigationButtonDarkTheme.module.scss';

interface Props {
    buttonText: string;
    onClick: () => void;
    isActive?: boolean;
}

export const NavigationDrawerButtonDarkTheme = ({ buttonText, onClick, isActive }: Props) => {
    return (
        <button
            className={`${classes.button} ${isActive ? classes.active : ''}`}
            onClick={onClick}
        >
            {buttonText}
        </button>
    );
};

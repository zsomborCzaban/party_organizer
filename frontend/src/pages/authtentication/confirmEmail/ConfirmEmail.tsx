import { useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import classes from './ConfirmEmail.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';

export const ConfirmEmail = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const confirmEmail = () => {
    const hash = searchParams.get('hash');
    const username = searchParams.get('username');

    if (!hash || !username) {
      setError('Invalid or expired confirmation link');
      return;
    }

    setIsLoading(true);
    api.authApi.confirmEmail(username, hash)
        .then(resp => {
          if(resp === 'error'){
            toast.error('Unexpected error')
            setError('Failed to confirm email. Please try again.');
            return
          }

          if (resp.is_error && resp.errors === 'Email already confirmed. try logging in!') {
              toast.success('Email already confirmed')
              navigate('/login')
              return;
          }

          if(resp.is_error && resp.code !== 500){
            setError('Invalid or expired confirmation link');
            return;
          }

          if(resp.is_error && resp.code === 500){
            toast.error('Unexpected error')
            setError('Failed to confirm email. Please try again.');
            return
          }

            toast.success('Email confirmed')
            navigate('/login')
        })
        .catch(() => {
          toast.error('Unexpected error')
          setError('Failed to confirm email. Please try again.');
        })
        .finally(() => {
          setIsLoading(false);
        });
  };

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Confirm Email</h2>
        <>
          <p className={classes.description}>Click the button below to confirm your email address.</p>
          <button
            onClick={confirmEmail}
            className={classes.confirmButton}
            disabled={isLoading}
          >
            {'Confirm My Email'}
          </button>
          {error && (
            <>
              <p className={classes.error}>{error}</p>
              <div className={classes.backToLoginContainer}>
                <a
                  href=""
                  onClick={(e) => {
                    e.preventDefault();
                    navigate('/login');
                  }}
                  className={classes.link}
                >
                  Back to Login
                </a>
              </div>
            </>
          )}
        </>
    </div>
  );
};
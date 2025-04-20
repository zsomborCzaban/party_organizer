import { useState, useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import classes from './ConfirmEmail.module.scss';
import { useApi } from '../../../context/ApiContext';
import { toast } from 'sonner';

export const ConfirmEmail = () => {
  const api = useApi();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const confirmEmail = () => {
    const hash = searchParams.get('hash');
    const username = searchParams.get('username');

    if (!hash || !username) {
      setError('Invalid or expired confirmation link');
      setIsLoading(false);
      return;
    }
    api.authApi.confirmEmail(username, hash)
        .then(resp => {
          if(resp === 'error'){
            toast.error('Unexpected error')
            setError('Failed to confirm email. Please try again.');
            return
          }

          if (resp.is_error && resp.errors === 'Email already confirmed. try logging in!') {
            //todo: emial already confirmed
            return;
          }

          if(resp.is_error && resp.code !== 500){
            console.log('inhere')
            setError('Invalid or expired confirmation link');
            setIsLoading(false);
            return;
          }

          if(resp.is_error && resp.code === 500){
            toast.error('Unexpected error')
            setError('Failed to confirm email. Please try again.');
            return
          }

          //todo: email confrim success


        })
        .catch(() => {
          toast.error('Unexpected error')
          setError('Failed to confirm email. Please try again.');
        })
        .finally(() =>{
          setIsLoading(false);
        })
    
  };


  useEffect(() => {
    confirmEmail();
  }, []);

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Confirm Email</h2>
      {isLoading ? (
        <p className={classes.description}>Confirming your email...</p>
      ) : error ? (
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
      ) : (
        <>
          <p className={classes.success}>Email confirmed successfully!</p>
          <p className={classes.description}>Redirecting to login page...</p>
        </>
      )}
    </div>
  );
};
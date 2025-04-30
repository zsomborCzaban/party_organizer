import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import classes from './ForgotPassword.module.scss'
import { useApi } from '../../../context/ApiContext'
import { toast } from 'sonner'

interface Feedbacks {
  Username?: string
  ButtonError?: string
  IsSuccess?: string

  [key: string]: string | undefined
}

export const ForgotPassword = () => {
  const api = useApi()
  const navigate = useNavigate()
  const [userName, setUsername] = useState('')
  const [feedbacks, setFeedbacks] = useState<Feedbacks>({})
  const [isLoading, setIsLoading] = useState(false)

  const validate = (): boolean => {
    const newFeedbacks: Feedbacks = {}
    let isValid = true

    if (!userName) {
      newFeedbacks.Username = 'Username is required'
      isValid = false
    }

    setFeedbacks(newFeedbacks)
    return isValid
  }

  const isFeedbacksEmpty = () => {
    if(feedbacks.Username) return false
    return true
  }

  const resetPasswordClicked = () => {
    if (!validate()) return

    setIsLoading(true)
    api.authApi.forgotPassword(userName)
        .then(resp => {
          if(resp === 'error'){
            toast.error('Unexpected error')
            return
          }

          if (resp.is_error) {
            if(typeof resp.errors === "string" && resp.errors.includes("user not found")){
              setFeedbacks({Username: "Username doesn't exist"})
              return
            }
            setFeedbacks({ButtonError: "Username doesn't exist"})
            return
          }

          setFeedbacks({IsSuccess: 'true'})
          toast.success('Email sent')
          return

        })
        .catch(() => toast.error('Unexpected error'))
        .finally(() => setIsLoading(false))
  }

  return (
    <div className={classes.container}>
      <h2 className={classes.title}>Forgot Password</h2>
      <p className={classes.description}>Don't worry, you can set a new password in no time.</p>
        <div className={classes.inputGroup}>
          <label htmlFor="email" className={classes.inputLabel}>
            Username
          </label>
          <input
            type="text"
            id="email"
            className={classes.input}
            value={userName}
            onChange={(e) => {
              setUsername(e.target.value)
              setFeedbacks({})
            }}
            placeholder="Enter your username"
          />
          {feedbacks.Username && <p className={classes.error}>{feedbacks.Username}</p>}
        </div>

        <button
          className={classes.resetPasswordButton}
          onClick={resetPasswordClicked}
          disabled={isLoading || !isFeedbacksEmpty()}
        >
          Reset Password
        </button>

        {feedbacks.ButtonError && (
          <p className={classes.error}>{feedbacks.ButtonError}</p>
        )}
        {feedbacks.IsSuccess && (
            <p className={classes.success}>
              Check you emails to proceed further
            </p>
        )}

        <div className={classes.backToLoginContainer}>
          <p>Remember your password?</p>
          <a
              onClick={() => navigate('/login')}
              className={classes.link}
          >
            Back to Login
          </a>
        </div>
    </div>
  )
}
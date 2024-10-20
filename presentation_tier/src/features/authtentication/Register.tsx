import React, { useState, CSSProperties } from 'react';
import {register, RegisterRequestBody} from "./AuthenticationApi";
import {AxiosError} from "axios";
import {ApiError, ApiResponse} from "../../api/ApiResponse";


interface Feedbacks{
    username?: string;
    password?: string;
    email?: string;
    confirmPassword?: string;
    buttonError?: string;
    buttonSuccess?: string;
}

const Register: React.FC = () => {
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [email, setEmail] = useState('')
    const [confirmPassword, setConfirmPassword] = useState<string>('');
    const [feedbacks, setFeedbackds] = useState<Feedbacks>({});

    // Validation logic
    const validate = (): boolean => {
        let valid = true;
        const newFeedbacks: Feedbacks = {};

        if (!username) {
            newFeedbacks.username = 'Username is required.';
            valid = false;
        }

        if (!password) {
            newFeedbacks.password = 'Password is required.';
            valid = false;
        }

        if (!email) {
            newFeedbacks.email = 'Email is required.';
            valid = false;
        }

        if (!confirmPassword) {
            newFeedbacks.confirmPassword = '\'Passwords do not match.';
            valid = false;
        } else if (password !== confirmPassword) {
            newFeedbacks.confirmPassword = 'Passwords do not match.';
            valid = false;
        }

        setFeedbackds(newFeedbacks);
        return valid;
    };

    const handleErrors = (errs: ApiError[])=> {
        let newFeedbacks: Feedbacks = {};
        const keysOfFeedbacks = ["username", "password", "confirmpassword", "email"]

        errs.forEach((err: ApiError) => {
            // @ts-ignore
            const key: keyof Feedbacks = err.field.toLowerCase()
            if(keysOfFeedbacks.includes(key)){
                newFeedbacks[key] = err.err
            } else {
                const unexpectedErr: Feedbacks = {}
                unexpectedErr.buttonError = "unexpected error while registering"
                newFeedbacks = unexpectedErr
                return
            }
        })
        setFeedbackds(newFeedbacks)
    }

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        if (validate()) {
            let registerBody: RegisterRequestBody = {
                username: username,
                email: email,
                password: password,
                confirm_password: confirmPassword,
            }
            register(registerBody)
                .then(() => {
                    const newFeedbacks: Feedbacks = {};
                    newFeedbacks.buttonSuccess = "registered successfully"
                    setFeedbackds(newFeedbacks)
                })
                .catch((err: AxiosError<ApiResponse<void>>) => {
                    if(err.response){
                        let errors = err.response.data.errors
                        handleErrors(errors)
                    } else {
                        const newFeedbacks: Feedbacks = {};
                        newFeedbacks.buttonError = "unexpected error while registering"
                        setFeedbackds(newFeedbacks)
                    }
                })
        }
    };

    // Inline styles
    const styles: { [key: string]: CSSProperties } = {
        container: {
            maxWidth: '400px',
            margin: '0 auto',
            padding: '20px',
            border: '1px solid #ccc',
            borderRadius: '10px',
        },
        formGroup: {
            marginBottom: '15px',
        },
        label: {
            display: 'block',
            marginBottom: '5px',
        },
        input: {
            width: '100%',
            padding: '8px',
            boxSizing: 'border-box',
        },
        error: {
            color: 'red',
            fontSize: '0.875em',
        },
        success: {
            color: 'green',
            fontSize: '0.875em',
        },
        button: {
            width: '100%',
            padding: '10px',
            backgroundColor: '#4CAF50',
            color: 'white',
            border: 'none',
            borderRadius: '5px',
            cursor: 'pointer',
        },
        buttonHover: {
            backgroundColor: '#45a049',
        },
        textCenter: {
            textAlign: 'center',
        },
    };

    return (
        <div style={styles.container}>
            <form onSubmit={handleSubmit}>
                <h2>Register</h2>

                {/* Username Field */}
                <div style={styles.formGroup}>
                    <label htmlFor="username" style={styles.label}>
                        Username
                    </label>
                    <input
                        type="text"
                        id="username"
                        style={styles.input}
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                    {feedbacks.username && <p style={styles.error}>{feedbacks.username}</p>}
                </div>

                {/* Email Field */}
                <div style={styles.formGroup}>
                    <label htmlFor="email" style={styles.label}>
                        Email
                    </label>
                    <input
                        type="text"
                        id="email"
                        style={styles.input}
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    {feedbacks.email && <p style={styles.error}>{feedbacks.email}</p>}
                </div>

                {/* Password Field */}
                <div style={styles.formGroup}>
                    <label htmlFor="password" style={styles.label}>
                        Password
                    </label>
                    <input
                        type="password"
                        id="password"
                        style={styles.input}
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                    {feedbacks.password && <p style={styles.error}>{feedbacks.password}</p>}
                </div>

                {/* Confirm Password Field */}
                <div style={styles.formGroup}>
                    <label htmlFor="confirmPassword" style={styles.label}>
                        Confirm Password
                    </label>
                    <input
                        type="password"
                        id="confirmPassword"
                        style={styles.input}
                        value={confirmPassword}
                        onChange={(e) => setConfirmPassword(e.target.value)}
                    />
                    {feedbacks.confirmPassword && (
                        <p style={styles.error}>{feedbacks.confirmPassword}</p>
                    )}
                </div>

                {/* Submit Button */}
                <button type="submit" style={styles.button}>
                    Send
                </button>

                {feedbacks.buttonError && (<p style={styles.error}>{feedbacks.buttonError}</p>)}

                {feedbacks.buttonSuccess && (<p style={styles.success}>{feedbacks.buttonSuccess}</p>)}


                {/* Sign In Link */}
                <p style={styles.textCenter}>
                    Already have an account? <a href="/login2">Sign In</a>
                </p>
            </form>
        </div>
    );
};

export default Register;

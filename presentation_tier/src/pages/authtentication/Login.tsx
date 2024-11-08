import {useState} from "react";
import {useNavigate} from "react-router-dom";
import {login} from "../../data/apis/AuthenticationApi";
import {Alert, Button, Card, CardBody, Container, Form, FormGroup} from "react-bootstrap";

const Login: React.FC = () => {
    const [username, setUsername] = useState('');
    const [usernameError, setUsernameError] = useState('');
    const [password, setPassword] = useState('');
    const [passwordError, setPasswordError] = useState('');

    const [isLoading, setIsLoading] = useState(false);

    const [apiError, setApiError] = useState('');

    const navigate = useNavigate();

    const validateForm = () => {
        let valid = true;
        if(!username) {
            setUsernameError('Username is required')
            valid = false
        } else {
            setUsernameError('')
        }

        if(!password) {
            setPasswordError('Password is required')
            valid = false
        } else {
            setPasswordError('')
        }

        return valid
    }

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!validateForm()) {
            return;
        }
        try {
            setIsLoading(true);
            await login(username, password);
            navigate('/overview/discover');
        } catch (error) {
            console.log(`api error: ${error}`)
            setApiError(`login failed with error: ${error}`);
        }
        setIsLoading(false);
    };

    return (
        <div className="background-container">
             <Container className="d-flex flex-column justify-content-center align-items-center">
                <h1 id="page-title-header" className="md-5 page-title">
                    Party Organizer
                </h1>
                 <Card className="login-card">
                     <CardBody>
                          <Form>
                              <FormGroup className="mb-3" controlId="forUsername">
                                  <Form.Label>Username</Form.Label>
                                  <Form.Control
                                    type="text"
                                    placeholder="Enter your username"
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                    isInvalid={usernameError.length !== 0}
                                  />
                                  <Form.Control.Feedback
                                    type="invalid"
                                    className={`animate__animated animate__fadeIn ${usernameError.length !== 0 ? '' : 'd-none'}`}>
                                      {usernameError}
                                  </Form.Control.Feedback>
                              </FormGroup>
                              <FormGroup className="mb-3" controlId="forPassword">
                                  <Form.Label>Password</Form.Label>
                                  <Form.Control
                                      type="password"
                                      placeholder="Enter your password"
                                      value={password}
                                      onChange={(e) => setPassword(e.target.value)}
                                      isInvalid={passwordError.length !== 0}
                                  />
                                  <Form.Control.Feedback
                                      type="invalid"
                                      className={`animate__animated animate__fadeIn ${passwordError.length !== 0 ? '' : 'd-none'}`}>
                                      {passwordError}
                                  </Form.Control.Feedback>
                              </FormGroup>
                              {apiError && (
                                  <Alert
                                      variant="dange"
                                      className="mt-3 animate__animated animate__fadeIn"
                                  >
                                      {apiError}
                                  </Alert>
                              )}
                              <div className="d-grid gap-2">
                                  <Button
                                      variant="primary"
                                      type="submit"
                                      className="login-button"
                                      disabled={
                                        isLoading
                                      }
                                      onClick={(e) => handleLogin(e)}
                                  >
                                      {isLoading ? 'Logging in...' : 'Login'}
                                  </Button>
                              </div>
                          </Form>
                     </CardBody>
                 </Card>
             </Container>
        </div>
    );
};

export default Login;
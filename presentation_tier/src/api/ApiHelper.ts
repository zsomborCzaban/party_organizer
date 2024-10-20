export const getBackendUrl = () => {
    return process.env.REACT_APP_BACKEND_URL
        ? process.env.REACT_APP_BACKEND_URL
        : '';
}
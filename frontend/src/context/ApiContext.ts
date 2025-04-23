import { createContext, useContext } from 'react';
import { Api } from '../api/Api';

export const ApiContext = createContext<Api>(new Api());

export const useApi = () => useContext(ApiContext);

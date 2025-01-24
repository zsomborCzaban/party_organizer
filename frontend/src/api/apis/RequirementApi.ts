import {DELETE, get,post} from '../../api/Api';
import {Requirement} from '../types/Requirement';

import {getApiUrl} from '../../api/ApiHelper';

const DRINK_REQUIREMENT_PATH = `${getApiUrl()  }/drinkRequirement`;
const FOOD_REQUIREMENT_PATH = `${getApiUrl()  }/foodRequirement`;


export const getDrinkRequirements = async (partyId: number): Promise<Requirement[]> => new Promise<Requirement[]>((resolve, reject) => {
        get<Requirement[]>(`${DRINK_REQUIREMENT_PATH  }/getByPartyId/${  partyId.toString()}`)
            .then((requirements: Requirement[]) => resolve(requirements))
            .catch((err) => reject(err));
    });

export const getFoodRequirements = async (partyId: number): Promise<Requirement[]> => new Promise<Requirement[]>((resolve, reject) => {
        get<Requirement[]>(`${FOOD_REQUIREMENT_PATH  }/getByPartyId/${  partyId.toString()}`)
            .then((requirements: Requirement[]) => resolve(requirements))
            .catch((err) => reject(err));
    });

export const createDrinkRequirement = async (requestBody: Requirement): Promise<Requirement> => new Promise<Requirement>((resolve, reject) => {
        post<Requirement>(DRINK_REQUIREMENT_PATH, requestBody)
            .then((requirement: Requirement) => resolve(requirement))
            .catch((err) => reject(err));
    });

export const createFoodRequirement = async (requestBody: Requirement): Promise<Requirement> => new Promise<Requirement>((resolve, reject) => {
        post<Requirement>(FOOD_REQUIREMENT_PATH, requestBody)
            .then((requirement: Requirement) => resolve(requirement))
            .catch((err) => reject(err));
    });

export const deleteDrinkRequirement = async (requirementId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        DELETE<void>(`${DRINK_REQUIREMENT_PATH  }/${  requirementId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const deleteFoodRequirement = async (requirementId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        DELETE<void>(`${FOOD_REQUIREMENT_PATH  }/${  requirementId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });
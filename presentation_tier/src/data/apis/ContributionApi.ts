import {Contribution} from '../types/Contribution';
import {DELETE, get, post} from '../../api/Api';
import {getApiUrl} from '../../api/ApiHelper';


const DRINK_CONTRIBUTION_PATH = `${getApiUrl()  }/drinkContribution`;
const FOOD_CONTRIBUTION_PATH = `${getApiUrl()  }/foodContribution`;

export const getDrinkContributions = async (partyId: number): Promise<Contribution[]> => new Promise<Contribution[]>((resolve, reject) => {
        get<Contribution[]>(`${DRINK_CONTRIBUTION_PATH  }/getByParty/${ partyId}`)
            .then((contributions: Contribution[]) => resolve(contributions))
            .catch((err) => reject(err));
    });

export const getFoodContributions = async (partyId: number): Promise<Contribution[]> => new Promise<Contribution[]>((resolve, reject) => {
        get<Contribution[]>(`${FOOD_CONTRIBUTION_PATH  }/getByParty/${  partyId}`)
            .then((contributions: Contribution[]) => resolve(contributions))
            .catch((err) => reject(err));
    });

export const createDrinkContribution = async (requestBody: Contribution): Promise<Contribution> => new Promise<Contribution>((resolve, reject) => {
        post<Contribution>(DRINK_CONTRIBUTION_PATH, requestBody)
            .then((contribution: Contribution) => resolve(contribution))
            .catch((err) => reject(err));
    });

export const createFoodContribution = async (requestBody: Contribution): Promise<Contribution> => new Promise<Contribution>((resolve, reject) => {
        post<Contribution>(FOOD_CONTRIBUTION_PATH, requestBody)
            .then((contribution: Contribution) => resolve(contribution))
            .catch((err) => reject(err));
    });

export const deleteDrinkContribution = async (contributionId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        DELETE<void>(`${DRINK_CONTRIBUTION_PATH  }/${  contributionId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });

export const deleteFoodContribution = async (contributionId: number): Promise<void> => new Promise<void>((resolve, reject) => {
        DELETE<void>(`${FOOD_CONTRIBUTION_PATH  }/${  contributionId.toString()}`)
            .then(() => resolve())
            .catch((err) => reject(err));
    });



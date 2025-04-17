import { DELETE, get, post } from '../Api.ts';
import { getApiUrl } from '../ApiHelper.ts';
import {Contribution, ContributionPopulated} from '../../data/types/Contribution';
import axios, {AxiosInstance, AxiosResponse} from "axios";

const DRINK_CONTRIBUTION_PATH = `${getApiUrl()}/drinkContribution`;
const FOOD_CONTRIBUTION_PATH = `${getApiUrl()}/foodContribution`;


const handleApiResponse = <T>(response: AxiosResponse<T>): T => {
    return response.data;
};

const handleApiError = (error: unknown) => {
    // TODO: handle errors as needed
    if (axios.isAxiosError(error)) {
        console.error(`Axios error: ${error.message}`);
    } else {
        console.error(`Unexpected error: ${error}`);
    }
};

export type ContributionsResponse = {
    data: ContributionPopulated[]
}

export type ContributionResponse = {
    data: ContributionPopulated
}

export class ContributionApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async getDrinkContributionsByParty(partyId: number): Promise<ContributionsResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<ContributionsResponse>(`${DRINK_CONTRIBUTION_PATH}/getByParty/${partyId.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async getFoodContributionsByParty(partyId: number): Promise<ContributionsResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<ContributionsResponse>(`${FOOD_CONTRIBUTION_PATH}/getByParty/${partyId.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }
}

export const getDrinkContributions = async (partyId: number): Promise<Contribution[]> =>
  new Promise<Contribution[]>((resolve, reject) => {
    get<Contribution[]>(`${DRINK_CONTRIBUTION_PATH}/getByParty/${partyId}`)
      .then((contributions: Contribution[]) => resolve(contributions))
      .catch((err) => reject(err));
  });

export const getFoodContributions = async (partyId: number): Promise<Contribution[]> =>
  new Promise<Contribution[]>((resolve, reject) => {
    get<Contribution[]>(`${FOOD_CONTRIBUTION_PATH}/getByParty/${partyId}`)
      .then((contributions: Contribution[]) => resolve(contributions))
      .catch((err) => reject(err));
  });

export const createDrinkContribution = async (requestBody: Contribution): Promise<Contribution> =>
  new Promise<Contribution>((resolve, reject) => {
    post<Contribution>(DRINK_CONTRIBUTION_PATH, requestBody)
      .then((contribution: Contribution) => resolve(contribution))
      .catch((err) => reject(err));
  });

export const createFoodContribution = async (requestBody: Contribution): Promise<Contribution> =>
  new Promise<Contribution>((resolve, reject) => {
    post<Contribution>(FOOD_CONTRIBUTION_PATH, requestBody)
      .then((contribution: Contribution) => resolve(contribution))
      .catch((err) => reject(err));
  });

export const deleteDrinkContribution = async (contributionId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    DELETE<void>(`${DRINK_CONTRIBUTION_PATH}/${contributionId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const deleteFoodContribution = async (contributionId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    DELETE<void>(`${FOOD_CONTRIBUTION_PATH}/${contributionId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

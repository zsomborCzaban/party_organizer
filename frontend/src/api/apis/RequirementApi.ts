import { DELETE, get, post } from '../Api.ts';

import { getApiUrl } from '../ApiHelper.ts';
import {Requirement, RequirementPopulated} from '../../data/types/Requirement';
import axios, {AxiosInstance, AxiosResponse} from "axios";

const DRINK_REQUIREMENT_PATH = `${getApiUrl()}/drinkRequirement`;
const FOOD_REQUIREMENT_PATH = `${getApiUrl()}/foodRequirement`;

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

export type RequirementsResponse = {
    data: RequirementPopulated[]
}

export type RequirementResponse = {
    data: RequirementPopulated
}

export class RequirementApi {
    private axiosInstance: AxiosInstance;

    constructor(axiosInstance: AxiosInstance) {
        this.axiosInstance = axiosInstance;
    }

    async getDrinkRequirementsByPartyId(partyId: number): Promise<RequirementsResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<RequirementsResponse>(`${DRINK_REQUIREMENT_PATH}/getByPartyId/${partyId.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async createDrinkRequirement(requestBody: Requirement): Promise<RequirementResponse | 'error'> {
        try {
            const response = await this.axiosInstance.post<RequirementResponse>(`${DRINK_REQUIREMENT_PATH}`, requestBody)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async deleteDrinkRequirement(requirement: number): Promise<RequirementResponse | 'error'> {
        try {
            const response = await this.axiosInstance.delete<RequirementResponse>(`${DRINK_REQUIREMENT_PATH}/${requirement.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async getFoodRequirementsByPartyId(partyId: number): Promise<RequirementsResponse | 'error'> {
        try {
            const response = await this.axiosInstance.get<RequirementsResponse>(`${FOOD_REQUIREMENT_PATH}/getByPartyId/${partyId.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async createFoodRequirement(requestBody: Requirement): Promise<RequirementResponse | 'error'> {
        try {
            const response = await this.axiosInstance.post<RequirementResponse>(`${FOOD_REQUIREMENT_PATH}`, requestBody)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }

    async deleteFoodRequirement(requirement: number): Promise<RequirementResponse | 'error'> {
        try {
            const response = await this.axiosInstance.delete<RequirementResponse>(`${FOOD_REQUIREMENT_PATH}/${requirement.toString()}`)
            return handleApiResponse(response)
        } catch (error) {
            handleApiError(error)
            return 'error'
        }
    }
}

export const getDrinkRequirements = async (partyId: number): Promise<Requirement[]> =>
  new Promise<Requirement[]>((resolve, reject) => {
    get<Requirement[]>(`${DRINK_REQUIREMENT_PATH}/getByPartyId/${partyId.toString()}`)
      .then((requirements: Requirement[]) => resolve(requirements))
      .catch((err) => reject(err));
  });

export const getFoodRequirements = async (partyId: number): Promise<Requirement[]> =>
  new Promise<Requirement[]>((resolve, reject) => {
    get<Requirement[]>(`${FOOD_REQUIREMENT_PATH}/getByPartyId/${partyId.toString()}`)
      .then((requirements: Requirement[]) => resolve(requirements))
      .catch((err) => reject(err));
  });

export const createDrinkRequirement = async (requestBody: Requirement): Promise<Requirement> =>
  new Promise<Requirement>((resolve, reject) => {
    post<Requirement>(DRINK_REQUIREMENT_PATH, requestBody)
      .then((requirement: Requirement) => resolve(requirement))
      .catch((err) => reject(err));
  });

export const createFoodRequirement = async (requestBody: Requirement): Promise<Requirement> =>
  new Promise<Requirement>((resolve, reject) => {
    post<Requirement>(FOOD_REQUIREMENT_PATH, requestBody)
      .then((requirement: Requirement) => resolve(requirement))
      .catch((err) => reject(err));
  });

export const deleteDrinkRequirement = async (requirementId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    DELETE<void>(`${DRINK_REQUIREMENT_PATH}/${requirementId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const deleteFoodRequirement = async (requirementId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    DELETE<void>(`${FOOD_REQUIREMENT_PATH}/${requirementId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

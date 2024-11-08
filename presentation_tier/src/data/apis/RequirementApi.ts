import {DELETE, get} from "../../api/Api";
import {Requirement} from "../types/Requirement";
import {Contribution} from "../types/Contribution";
import {User} from "../types/User";
import {post} from "../../api/Api";
import {PartyInvite} from "../types/PartyInvite";
import {BACKEND_URL} from "../constants/backend_url";

const DRINK_REQUIREMENT_PATH = BACKEND_URL + "/drinkRequirement"
const FOOD_REQUIREMENT_PATH = BACKEND_URL + "/foodRequirement"


export const getDrinkRequirements = async (partyId: number): Promise<Requirement[]> => {
    return new Promise<Requirement[]>((resolve, reject) => {
        get<Requirement[]>(DRINK_REQUIREMENT_PATH + "/getByPartyId/" + partyId.toString())
            .then((requirements: Requirement[]) => {
                return resolve(requirements);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getFoodRequirements = async (partyId: number): Promise<Requirement[]> => {
    return new Promise<Requirement[]>((resolve, reject) => {
        get<Requirement[]>(FOOD_REQUIREMENT_PATH + "/getByPartyId/" + partyId.toString())
            .then((requirements: Requirement[]) => {
                return resolve(requirements);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const createDrinkRequirement = async (requestBody: Requirement): Promise<Requirement> => {
    return new Promise<Requirement>((resolve, reject) => {
        post<Requirement>(DRINK_REQUIREMENT_PATH, requestBody)
            .then((requirement: Requirement) => {
                return resolve(requirement);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const createFoodRequirement = async (requestBody: Requirement): Promise<Requirement> => {
    return new Promise<Requirement>((resolve, reject) => {
        post<Requirement>(FOOD_REQUIREMENT_PATH, requestBody)
            .then((requirement: Requirement) => {
                return resolve(requirement);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const deleteDrinkRequirement = async (requirementId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        DELETE<void>(DRINK_REQUIREMENT_PATH + "/" + requirementId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
}

export const deleteFoodRequirement = async (requirementId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        DELETE<void>(FOOD_REQUIREMENT_PATH + "/" + requirementId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
}
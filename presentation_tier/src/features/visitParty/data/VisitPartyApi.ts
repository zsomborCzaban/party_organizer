import {DELETE, get} from "../../../api/Api";
import {Requirement} from "./Requirement";
import {Contribution} from "./Contribution";
import {User} from "../../overView/User";
import {post} from "../../../api/Api";

const DRINK_REQUIREMENT_PATH = "http://localhost:8080/api/v0/drinkRequirement/getByPartyId/"
const FOOD_REQUIREMENT_PATH = "http://localhost:8080/api/v0/foodRequirement/getByPartyId/"
const DRINK_CONTRIBUTION_PATH = "http://localhost:8080/api/v0/drinkContribution"
const FOOD_CONTRIBUTION_PATH = "http://localhost:8080/api/v0/foodContribution"
const PARTICIPANTS_PATH = "http://localhost:8080/api/v0/party/getParticipants/"

export const getDrinkRequirements = async (partyId: number): Promise<Requirement[]> => {
    return new Promise<Requirement[]>((resolve, reject) => {
        get<Requirement[]>(DRINK_REQUIREMENT_PATH + partyId.toString())
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
        get<Requirement[]>(FOOD_REQUIREMENT_PATH + partyId.toString())
            .then((requirements: Requirement[]) => {
                return resolve(requirements);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getDrinkContributions = async (partyId: number): Promise<Contribution[]> => {
    return new Promise<Contribution[]>((resolve, reject) => {
        get<Contribution[]>(DRINK_CONTRIBUTION_PATH + '/getByParty/' +partyId)
            .then((contributions: Contribution[]) => {
                return resolve(contributions);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getFoodContributions = async (partyId: number): Promise<Contribution[]> => {
    return new Promise<Contribution[]>((resolve, reject) => {
        get<Contribution[]>(FOOD_CONTRIBUTION_PATH + '/getByParty/' + partyId)
            .then((contributions: Contribution[]) => {
                return resolve(contributions);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const getPartyParticipants = async (partyId: number): Promise<User[]> => {
    return new Promise<User[]>((resolve, reject) => {
        get<User[]>(PARTICIPANTS_PATH + partyId)
            .then((participants: User[]) => {
                return resolve(participants);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const createDrinkContribution = async (requestBody: Contribution): Promise<Contribution> => {
    return new Promise<Contribution>((resolve, reject) => {
        post<Contribution>(DRINK_CONTRIBUTION_PATH, requestBody)
            .then((contribution: Contribution) => {
                return resolve(contribution);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const createFoodContribution = async (requestBody: Contribution): Promise<Contribution> => {
    return new Promise<Contribution>((resolve, reject) => {
        post<Contribution>(FOOD_CONTRIBUTION_PATH, requestBody)
            .then((contribution: Contribution) => {
                return resolve(contribution);
            })
            .catch(err => {
                return reject(err);
            });
    });
};

export const deleteDrinkContribution = async (contributionId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        DELETE<void>(DRINK_CONTRIBUTION_PATH + "/" + contributionId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
}

export const deleteFoodContribution = async (contributionId: number): Promise<void> => {
    return new Promise<void>((resolve, reject) => {
        DELETE<void>(FOOD_CONTRIBUTION_PATH + "/" + contributionId.toString())
            .then(() => {
                return resolve();
            })
            .catch(err => {
                return reject(err);
            });
    });
}
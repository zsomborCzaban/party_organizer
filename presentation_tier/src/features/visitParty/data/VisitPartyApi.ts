import {get} from "../../../api/Api";
import {Requirement} from "./Requirement";
import {Contribution} from "./Contribution";
import {User} from "../../overView/User";

const DRINK_REQUIREMENT_PATH = "http://localhost:8080/api/v0/drinkRequirement/getByPartyId/"
const FOOD_REQUIREMENT_PATH = "http://localhost:8080/api/v0/foodRequirement/getByPartyId/"
const DRINK_CONTRIBUTION_PATH = "http://localhost:8080/api/v0/drinkContribution/getByParty/"
const FOOD_CONTRIBUTION_PATH = "http://localhost:8080/api/v0/foodContribution/getByParty/"
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
        get<Contribution[]>(DRINK_CONTRIBUTION_PATH + partyId)
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
        get<Contribution[]>(FOOD_CONTRIBUTION_PATH + partyId)
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
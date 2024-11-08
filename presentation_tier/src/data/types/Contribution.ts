import {User} from "./User";
import {Requirement} from "./Requirement";

export interface Contribution {
    ID?: number;
    contributor?: User;
    contributor_id?: number,
    requirement?: Requirement;
    requirement_id: number
    quantity: number;
    description?: string;
}
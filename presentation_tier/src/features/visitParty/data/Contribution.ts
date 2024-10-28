import {User} from "../../overView/User";
import {Requirement} from "./Requirement";

export interface Contribution {
    ID?: number;
    contributor?: User;
    requirement?: Requirement;
    requirement_id: number
    quantity: number;
    description?: string;
}
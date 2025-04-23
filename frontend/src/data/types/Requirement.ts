export interface Requirement {
    ID?: number;
    party_id: number
    type: string;
    target_quantity: number;
    quantity_mark: string;
}

export interface RequirementPopulated {
    ID: number;
    party_id: number
    type: string;
    target_quantity: number;
    quantity_mark: string;
}

export const EMPTY_REQUIREMENT_POPULATED: RequirementPopulated = {
    target_quantity: 0,
    ID: 0,
    type: '',
    party_id: 0,
    quantity_mark: '',
}
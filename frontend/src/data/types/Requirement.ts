export interface Requirement {
    ID?: number;
    party_id: number
    type: string;
    target_quantity: number;
    quantity_mark: string;
    description?: string;
}
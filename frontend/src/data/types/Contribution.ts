import {EMPTY_USER, User} from './User';
import {EMPTY_REQUIREMENT_POPULATED, Requirement} from './Requirement';

export interface Contribution {
  ID?: number;
  contributor?: User;
  contributor_id?: number;
  requirement?: Requirement;
  requirement_id: number;
  quantity: number;
  description?: string;
}

export interface ContributionPopulated {
  ID: number;
  contributor: User;
  contributor_id: number;
  requirement: Requirement;
  requirement_id: number;
  quantity: number;
  description?: string;
}

export const EMPTY_CONTRIBUTION_POPULATED: ContributionPopulated = {
  ID: 0,
  contributor: EMPTY_USER,
  contributor_id: 0,
  description: '',
  quantity: 0,
  requirement: EMPTY_REQUIREMENT_POPULATED,
  requirement_id: 0,
}
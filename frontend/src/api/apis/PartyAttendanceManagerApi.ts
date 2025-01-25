import { get } from '../../api/Api';

import { getApiUrl } from '../../api/ApiHelper';
import { Party } from '../../data/types/Party';
import { PartyInvite } from '../../data/types/PartyInvite';

const PARTY_ATTENDANCE_MANAGER_PATH = `${getApiUrl()}/partyAttendanceManager`;

export const getPartyPendingInvites = async (partyId: number): Promise<PartyInvite[]> =>
  new Promise<PartyInvite[]>((resolve, reject) => {
    get<PartyInvite[]>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPartyPendingInvites/${partyId.toString()}`)
      .then((invites: PartyInvite[]) => resolve(invites))
      .catch((err) => reject(err));
  });

export const inviteToParty = async (partyId: number, username: string): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/invite/${partyId.toString()}/${username}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const kickFromParty = async (partyId: number, userId: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/kick/${partyId.toString()}/${userId.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const getPartyInvites = async (): Promise<PartyInvite[]> =>
  new Promise<PartyInvite[]>((resolve, reject) => {
    get<PartyInvite[]>(`${PARTY_ATTENDANCE_MANAGER_PATH}/getPendingInvites`)
      .then((invites: PartyInvite[]) => resolve(invites))
      .catch((err) => reject(err));
  });

export const acceptInvite = async (partyID: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/accept/${partyID.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const declineInvite = async (partyID: number): Promise<void> =>
  new Promise<void>((resolve, reject) => {
    get<void>(`${PARTY_ATTENDANCE_MANAGER_PATH}/decline/${partyID.toString()}`)
      .then(() => resolve())
      .catch((err) => reject(err));
  });

export const joinPublicParty = async (partyId: number): Promise<Party> =>
  new Promise<Party>((resolve, reject) => {
    get<Party>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPublicParty/${partyId.toString()}`)
      .then((party: Party) => resolve(party))
      .catch((err) => reject(err));
  });

export const joinPrivateParty = async (accessCode: string): Promise<Party> =>
  new Promise<Party>((resolve, reject) => {
    get<Party>(`${PARTY_ATTENDANCE_MANAGER_PATH}/joinPrivateParty/${accessCode}`)
      .then((party: Party) => resolve(party))
      .catch((err) => reject(err));
  });

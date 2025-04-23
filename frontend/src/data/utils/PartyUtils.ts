import {NavigateFunction} from "react-router-dom";

export const NavigateToPartyHome = (navigate: NavigateFunction, partyName: string, partyId: number, organizerName: string) => {
    localStorage.setItem('partyName', partyName)
    localStorage.setItem('partyId', partyId.toString())
    localStorage.setItem('partyOrganizerName', organizerName)
    navigate('/partyHome')
}
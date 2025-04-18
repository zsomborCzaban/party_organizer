import {authService} from "../../auth/AuthService.ts";
import {NavigateFunction} from "react-router-dom";
import {handleProfilePictureUpload} from "./imageUtils.ts";
import {ChangeEvent} from "react";
import {Api} from "../../api/Api.ts";
import {toast} from "sonner";

export const handleLogoutUtil = (navigate: NavigateFunction) => {
    authService.userLoggedOut()
    navigate('/')
}

export const handleUploadProfilePictureUtil = (event: ChangeEvent<HTMLInputElement>) => {
    handleProfilePictureUpload(event)
}

export const handleLeavePartyUtils = (api: Api, navigate: NavigateFunction, partyId: number) => {
    api.partyAttendanceApi.leaveParty(partyId)
        .then((resp) => {
            if(resp === 'error'){
                toast.error('error while leaveing party, todo:handle this') //todo: handle error here
                return
            }

            navigate('/')
            toast.success('Party left')
        })
        .catch(() => {
            toast.error('Unexpected error')
        })
}
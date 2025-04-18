import {authService} from "../../auth/AuthService.ts";
import {NavigateFunction} from "react-router-dom";
import {handleProfilePictureUpload} from "./imageUtils.ts";
import {ChangeEvent} from "react";
import {Api} from "../../api/Api.ts";
import {toast} from "sonner";
import {ThunkDispatch, UnknownAction} from "@reduxjs/toolkit";
import {closeDefaultProfileDrawer, closePartyProfileDrawer} from "../../store/slices/profileDrawersSlice.ts";

export const handleLogoutUtil = (navigate: NavigateFunction, dispatch: ThunkDispatch<never, undefined, UnknownAction>, navigateTo: string) => {
    navigate(navigateTo)
    dispatch(closeDefaultProfileDrawer())
    dispatch(closePartyProfileDrawer())
    authService.userLoggedOut()
}

export const handleUploadProfilePictureUtil = (event: ChangeEvent<HTMLInputElement>) => {
    handleProfilePictureUpload(event)
}

export const handleLeavePartyUtils = (api: Api, navigate: NavigateFunction, partyId: number) => {
    api.partyAttendanceApi.leaveParty(partyId)
        .then((resp) => {
            console.log(resp)
            if(!resp.data){
                toast.error(resp.errors) //todo: handle error here
                return
            }

            navigate('/')
            toast.success('Party left')
        })
        .catch(() => {
            toast.error('Unexpected error')
        })
}
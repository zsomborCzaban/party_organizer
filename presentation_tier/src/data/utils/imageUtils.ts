import {ChangeEvent} from "react";
import {uploadPicture} from "../apis/UserApi";
import {authService} from "../../auth/AuthService";
import {getUserProfilePicture} from "../../auth/AuthUserUtil";

export const handleProfilePictureUpload = (event: ChangeEvent<HTMLInputElement>, profilePictureUrlSetter: React.Dispatch<React.SetStateAction<string>>, errorMessageSetter: React.Dispatch<React.SetStateAction<string>>) => {
    const file = event.target.files && event.target.files[0];

    if (!file) return;
    if (!file.type.startsWith("image/")) {
        errorMessageSetter("Upload failed. Please select a valid image file (PNG or JPG).");
        setTimeout(() => {
            errorMessageSetter("")
        }, 4000); // 4000 milliseconds = 4 seconds
        return;
    }

    const formData = new FormData()
    formData.append("image", file)

    uploadPicture(formData)
        .then(resp => {
            authService.userLoggedIn(resp.jwt)
            const newPicture = getUserProfilePicture()
            if(!newPicture) {
                errorMessageSetter("Upload successful but failed to fetch new profile picture. Reload the page to see the new profile picture");
                setTimeout(() => {
                    errorMessageSetter("")
                }, 4000);
                return
            }
            profilePictureUrlSetter(newPicture)
        })
        .catch(err => {
            errorMessageSetter("Something unexpected happened");
            setTimeout(() => {
                errorMessageSetter("")
            }, 4000); // 4000 milliseconds = 4 seconds
        })
}

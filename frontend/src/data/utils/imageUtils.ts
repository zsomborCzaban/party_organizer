import {ChangeEvent} from 'react';
import { uploadPicture } from '../../api/apis/UserApi';
import {toast} from "sonner";

export const handleProfilePictureUpload = (event: ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files && event.target.files[0];

    if (!file) return;
    if (!file.type.startsWith('image/')) {
        toast.error('Upload failed. Please select a valid image file (PNG or JPG).')
    }

    const formData = new FormData();
    formData.append('image', file);

    uploadPicture(formData)
        .then(resp => {
            localStorage.setItem('profile_picture_url', resp.profile_picture_url)
            toast.success('Profile picture uploaded')
        })
        .catch(() => {
            toast.error('Unexpected error')
        });
};

import {ConfigProvider, Modal, theme} from "antd";
import classes from "../deleteContribution/DeleteContributionModal.module.scss";
import {useApi} from "../../../context/ApiContext.ts";
import {toast} from "sonner";
import {useNavigate} from "react-router-dom";

interface DeletePartyModalProps {
    visible: boolean;
    onClose: () => void;
    partyId: number
}


export const DeletePartyModal: React.FC<DeletePartyModalProps> = ({
    visible,
    onClose,
    partyId,
}) => {
    const api = useApi()
    const navigate = useNavigate()

    const handleDelete = (partyIdToDelete: number) => {
            api.partyApi.deleteParty(partyIdToDelete)
                .then(resp => {
                    if(resp.data){
                        if(resp.data === 'delete_success'){
                            toast.success('Party deleted')
                            onClose()
                            navigate('/parties')
                            return;
                        }
                    }

                    toast.error('Unexpected error')
                    return
                })
                .catch(() => toast.error('Unexpected error'))
    }

    return (
    <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
        <Modal
            title='Delete Party'
            open={visible}
            onCancel={onClose}
            footer={null}
            className={classes.modal}
        >
            <div className={classes.modalContent}>
                <label className={classes.label}>Are you sure you want to delete this party? All the requirements and contributions will be lost</label>

                <div className={classes.buttonContainer}>
                    <button
                        onClick={() => handleDelete(partyId)}
                        className={classes.cancelButton}
                    >
                        Delete
                    </button>
                    <button
                        onClick={onClose}
                        className={classes.submitButton}
                    >
                        Cancel
                    </button>
                </div>
            </div>
        </Modal>
    </ConfigProvider>
    );
}
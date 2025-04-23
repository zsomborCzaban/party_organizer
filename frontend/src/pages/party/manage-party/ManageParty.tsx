import { Button, ConfigProvider, Table, theme } from 'antd'
import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import CreateRequirementModal from '../../../components/modal/createRequirement/CreateRequirementModal.tsx'
import DeleteRequirementModal from '../../../components/modal/deleteRequirement/DeleteRequirementModal.tsx'
import {EMPTY_USER, User} from '../../../data/types/User.ts'
import { setForTime } from '../../../data/utils/TimeoutSetterUtils.ts'
import {Requirement, RequirementPopulated} from '../../../data/types/Requirement.ts'
import { invitedTableColumnsLegacy, requirementTableColumnsLegacy, userTableColumnsLegacy } from '../../../data/constants/TableColumns.tsx'
import { inviteToParty } from '../../../api/apis/PartyAttendanceManagerApi.ts'
import {EMPTY_PARTY_POPULATED, PartyPopulated} from "../../../data/types/Party.ts"
import {PartyInvite} from "../../../data/types/PartyInvite.ts"
import {useApi} from "../../../context/ApiContext.ts"
import {toast} from "sonner"
import classes from './ManageParty.module.scss'
import KickUserModal from "../../../components/modal/kickUserModal/KickUserModal.tsx"


const ManageParty = () => {
  const navigate = useNavigate()
  const api = useApi()
  const partyId = Number(localStorage.getItem('partyId') || '-1')

  const [usernameInput, setUsernameInput] = useState('')
  const [inviteFeedbackSuccess, setInviteFeedbackSuccess] = useState('')
  const [inviteFeedbackError, setInviteFeedbackError] = useState('')
  const [requirementModalVisible, setRequirementModalVisible] = useState(false)
  const [requirementModalMode, setRequirementModalMode] = useState('')
  const [kickUserModalVisible, setKickUserModalVisible] = useState(false)
  const [userToKick, setUserToKick] = useState<User>(EMPTY_USER)
  const [deleteModalVisible, setDeleteModalVisible] = useState(false)
  const [deleteModalMode, setDeleteModalMode] = useState('')
  const [requirementToDelete, setRequirementToDelete] = useState(-1)

  const [party, setParty] = useState<PartyPopulated>(EMPTY_PARTY_POPULATED)
  const [drinkReqs, setDrinkReqs] = useState<RequirementPopulated[]>([])
  const [foodReqs, setFoodReqs] = useState<RequirementPopulated[]>([])
  const [participants, setParticipants] = useState<User[]>([])
  const [pendingInvites, setPendingInvites] = useState<PartyInvite[]>([])
  const [reloadDrinkReqs, setReloadDrinkReqs] = useState(0)
  const [reloadFoodReqs, setReloadFoodReqs] = useState(0)
  const [reloadParticipants, setReloadParticipants] = useState(0)
  const [reloadPendingInvites, setReloadPendingInvites] = useState(0)

  useEffect(() => {
    api.partyApi.getParty(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party')
            return
          }
          if(result === 'private party'){
            toast.error('Navigation error')
            navigate('/partyHome')
            return
          }
          setParty(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.partyApi, navigate, partyId])

  useEffect(() => {
    api.requirementApi.getDrinkRequirementsByPartyId(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load drink requirements')
            return
          }
          setDrinkReqs(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.requirementApi, partyId, reloadDrinkReqs])

  useEffect(() => {
    api.requirementApi.getFoodRequirementsByPartyId(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load food requirements')
            return
          }
          setFoodReqs(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.requirementApi, partyId, reloadFoodReqs])

  useEffect(() => {
    api.partyApi.getPartyParticipants(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party participants')
            return
          }
          setParticipants(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
    
  }, [api.partyApi, partyId, reloadParticipants])

  useEffect(() => {
    api.partyAttendanceApi.getPartyPendingInvites(partyId)
        .then(result => {
          if(result === 'error'){
            toast.error('Unable to load party invites')
            return
          }
          setPendingInvites(result.data)
        })
        .catch(() => {
          toast.error('Unexpected error')
        })
  }, [api.partyAttendanceApi, partyId, reloadPendingInvites])


  const handleInviteToParty = (username: string) => {
    inviteToParty(party.ID, username)
      .then(() => {
        setReloadPendingInvites((reloadPendingInvites+1)%2)
        setForTime<string>(setInviteFeedbackSuccess, 'Invite sent!', '', 3000)
      })
      .catch((err) => {
        console.log(err)
        if(err && err.response && err.response.data && err.response.data.errors){
          const errorMsg: string = err.response.data.errors
          if(errorMsg.includes('user not found')) {
            setForTime<string>(setInviteFeedbackError, "User doesn't exist!", '', 3000)
            return
          }
        }
        setForTime<string>(setInviteFeedbackError, 'Something went wrong!', '', 3000)
        toast.error('Unexpected error')
      })
  }

  const handleAddRequirement = (mode: string) => {
    if (mode === 'drink') setRequirementModalMode('drink')
    if (mode === 'food') setRequirementModalMode('food')
    setRequirementModalVisible(true)
  }

  const handleDeleteRequirement = (requirement: Requirement, mode: string) => {
    if (mode === 'drink') {
      setDeleteModalMode('drink')
      setRequirementToDelete(requirement.ID || -1)
      setDeleteModalVisible(true)
    }

    if (mode === 'food') {
      setDeleteModalMode('food')
      setRequirementToDelete(requirement.ID || -1)
      setDeleteModalVisible(true)
    }
  }

  const handleKickParticipant = (kickedUser: User) => {
    setUserToKick(kickedUser)
    setKickUserModalVisible(true)
  }

  const participantColumns = [
    ...userTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: User) => (
        <Button
          className={classes.errorButton}
          onClick={() => handleKickParticipant(record)}
        >
          Kick
        </Button>
      ),
    },
  ]

  const drinkRequirementColumns = [
    ...requirementTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: Requirement) => (
        <Button
          className={classes.errorButton}
          onClick={() => handleDeleteRequirement(record, 'drink')}
        >
          Delete
        </Button>
      ),
    },
  ]

  const foodRequirementColumns = [
    ...requirementTableColumnsLegacy,
    {
      title: '',
      key: 'action 1',
      render: (record: Requirement) => (
        <Button
          className={classes.errorButton}
          onClick={() => handleDeleteRequirement(record, 'food')}
        >
          Delete
        </Button>
      ),
    },
  ]

  const renderReqs = (requirements: Requirement[], mode: string) => {
    if (!requirements || requirements.length === 0) {
      return <div>There&#39s no drink requirements yet!</div>
    }
    return (
      <Table
        dataSource={requirements.map((req) => ({ ...req, key: req.ID }))}
        columns={mode === 'drink' ? drinkRequirementColumns : foodRequirementColumns}
        pagination={false}
        scroll={{ y: 200 }}
      />
    )
  }

  const renderParticipants = () => {
    if (!participants || participants.length === 0) {
      return <div>There&#39s no drink requirements yet!</div>
    }
    return (
      <Table
        dataSource={participants.map((person) => ({ ...person, key: person.ID }))}
        columns={participantColumns}
        pagination={false}
        scroll={{ y: 200 }}
      />
    )
  }

  const renderPendingInvites = () => {
    if (!pendingInvites || pendingInvites.length === 0) {
      return <div>There&#39s no pending invites at the moment</div>
    }
    return (
      <Table
        dataSource={pendingInvites.map((invite) => ({ ...invite, key: invite.ID }))}
        columns={invitedTableColumnsLegacy}
        pagination={false}
        scroll={{ y: 200 }}
      />
    )
  }

  return (
    <div className={classes.background}>
      <div className={classes.outerContainer}>
        <ConfigProvider theme={{ algorithm: theme.darkAlgorithm }}>
          <CreateRequirementModal
            visible={requirementModalVisible}
            onClose={() => {
              setRequirementModalVisible(false)
              setReloadDrinkReqs((reloadDrinkReqs+1)%2) //todo: find a better solution instead of refreshing on every modal close
              setReloadFoodReqs((reloadFoodReqs+1)%2)
            }}
            mode={requirementModalMode}
          />
          <DeleteRequirementModal
            visible={deleteModalVisible}
            onClose={() => {
              setDeleteModalVisible(false)
              setReloadDrinkReqs((reloadDrinkReqs+1)%2) //todo: find a better solution instead of refreshing on every modal close
              setReloadFoodReqs((reloadFoodReqs+1)%2)
            }}
            mode={deleteModalMode}
            requirementId={requirementToDelete}
          />
          <KickUserModal
              visible={kickUserModalVisible}
              onClose={() => {
                setKickUserModalVisible(false)
                setReloadParticipants((reloadDrinkReqs+1)%2) //todo: find a better solution instead of refreshing on every modal close
              }}
              user={userToKick}
              partyId={partyId}
          />
          <div className={classes.container}>
            <h2>Invite</h2>
            <div className={classes.inputContainer}>
              <input
                type='text'
                id='username'
                value={usernameInput}
                placeholder='Enter username'
                onChange={(e) => setUsernameInput(e.target.value)}
                className={classes.input}
              />
              <Button
                type='primary'
                className={classes.button}
                onClick={() => handleInviteToParty(usernameInput)}
              >
                Invite
              </Button>
              {inviteFeedbackSuccess && <p className={classes.success}>{inviteFeedbackSuccess}</p>}
              {inviteFeedbackError && <p className={classes.error}>{inviteFeedbackError}</p>}
            </div>

            <h2>Drink Requirements</h2>
            <div className={classes.requirementContainer}>
              <Button
                type='primary'
                className={classes.button}
                onClick={() => handleAddRequirement('drink')}
              >
                Add
              </Button>
              <div className={classes.requirementTable}>
                {!drinkReqs && <div>Loading...</div>}
                {drinkReqs && renderReqs(drinkReqs, 'drink')}
              </div>
            </div>

            <h2>Food Requirements</h2>
            <div className={classes.requirementContainer}>
              <Button
                type='primary'
                className={classes.button}
                onClick={() => handleAddRequirement('food')}
              >
                Add
              </Button>
              <div className={classes.requirementTable}>
                {!foodReqs && <div>Loading...</div>}
                {foodReqs && renderReqs(foodReqs, 'food')}
              </div>
            </div>

            <h2>Participants</h2>
            <div className={classes.requirementContainer}>
              <div className={classes.requirementTable}>
                {!participants && <div>Loading...</div>}
                {participants && renderParticipants()}
              </div>
            </div>

            <h2>Pending Invites</h2>
            <div className={classes.requirementContainer}>
              <div className={classes.requirementTable}>
                {!pendingInvites && <div>Loading...</div>}
                {pendingInvites && renderPendingInvites()}
              </div>
            </div>
          </div>
        </ConfigProvider>
      </div>
    </div>
  )
}

export default ManageParty

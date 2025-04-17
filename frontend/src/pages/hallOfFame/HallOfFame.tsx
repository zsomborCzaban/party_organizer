import {useApi} from "../../context/ApiContext.ts";
import {useNavigate} from "react-router-dom";
import {getUserName} from "../../auth/AuthUserUtil.ts";
import {useEffect, useState} from "react";
import {ContributionPopulated} from "../../data/types/Contribution.ts";
import {RequirementPopulated} from "../../data/types/Requirement.ts";
import {User} from "../../data/types/User.ts";
import {toast} from "sonner";
import { DownOutlined, DeleteOutlined } from '@ant-design/icons';
import classes from './HallOfFame.module.scss';

export const HallOfFame = () => {
    const api = useApi()
    const navigate = useNavigate()
    const userName = getUserName() || ''
    const organizerName = localStorage.getItem('partyOrganizerName') || ''
    const partyId = Number(localStorage.getItem('partyId') || '-1')

    const [participants, setParticipants] = useState<User[]>([])
    const [drinkContributions, setDrinkContributions] = useState<ContributionPopulated[]>([])
    const [foodContributions, setFoodContributions] = useState<ContributionPopulated[]>([])
    const [drinkRequirements, setDrinkRequirements] = useState<RequirementPopulated[]>([])
    const [foodRequirements, setFoodRequirements] = useState<RequirementPopulated[]>([])
    const [expandedUsers, setExpandedUsers] = useState<Set<number>>(new Set())

    useEffect(() => {
        if(userName === '' || organizerName === '' || partyId === -1){
            toast.error('Navigation error')
            navigate('/partyHome')
        }
    }, [navigate, organizerName, partyId, userName]);

    useEffect(() => {
        api.partyApi.getPartyParticipants(partyId)
            .then(response => {
                if(response === 'error'){
                    toast.error('Unable to load participants')
                    return
                }
                setParticipants(response.data)
            })
            .catch(() => {
                toast.error('Unexpected error')
            })
    }, [api.partyApi, partyId]);

    useEffect(() => {
        api.requirementApi.getDrinkRequirementsByPartyId(partyId)
            .then(response => {
                if(response === 'error'){
                    toast.error('Unable to load drink requirements')
                    return
                }
                setDrinkRequirements(response.data)
            })
            .catch(() => {
                toast.error('Unexpected error')
            })
    }, [api.requirementApi, partyId]);

    useEffect(() => {
        api.requirementApi.getFoodRequirementsByPartyId(partyId)
            .then(response => {
                if(response === 'error'){
                    toast.error('Unable to load food requirements')
                    return
                }
                setFoodRequirements(response.data)
            })
            .catch(() => {
                toast.error('Unexpected error')
            })
    }, [api.requirementApi, partyId]);

    useEffect(() => {
        api.contributionApi.getDrinkContributionsByParty(partyId)
            .then(response => {
                if(response === 'error'){
                    toast.error('Unable to load drink contributions')
                    return
                }
                setDrinkContributions(response.data)
            })
            .catch(() => {
                toast.error('Unexpected error')
            })
    }, [api.contributionApi, partyId]);

    useEffect(() => {
        api.contributionApi.getFoodContributionsByParty(partyId)
            .then(response => {
                if(response === 'error'){
                    toast.error('Unable to load food contributions')
                    return
                }
                setFoodContributions(response.data)
            })
            .catch(() => {
                toast.error('Unexpected error')
            })
    }, [api.contributionApi, partyId]);

    const toggleUser = (userId: number) => {
        setExpandedUsers(prev => {
            const newSet = new Set(prev)
            if (newSet.has(userId)) {
                newSet.delete(userId)
            } else {
                newSet.add(userId)
            }
            return newSet
        })
    }

    const getContributionsForUser = (userId: number) => {
        const allContributions = [...drinkContributions, ...foodContributions]
        return allContributions.filter(contribution => contribution.contributor.ID === userId)
    }

    const getTotalContributionsForUser = (userId: number) => {
        return getContributionsForUser(userId).length
    }

    const getRequirementById = (requirementId: number) => {
        const allRequirements = [...drinkRequirements, ...foodRequirements]
        return allRequirements.find(req => req.ID === requirementId)
    }

    const renderUser = (user: User) => {
        const isExpanded = expandedUsers.has(user.ID)
        const userContributions = getContributionsForUser(user.ID)
        const totalContributions = getTotalContributionsForUser(user.ID)

        return (
            <div key={user.ID} className={classes.userCard}>
                <div 
                    className={classes.userHeader}
                    onClick={() => toggleUser(user.ID)}
                >
                    <div className={classes.userInfo}>
                        <img 
                            src={user.profile_picture_url}
                            alt={user.username}
                            className={classes.profilePicture}
                        />
                        <div className={classes.userDetails}>
                            <div className={classes.userName}>{user.username}</div>
                            <div className={classes.contributionCount}>
                                Total Contributions: {totalContributions}
                            </div>
                        </div>
                    </div>
                    <DownOutlined 
                        className={`${classes.expandIcon} ${isExpanded ? classes.expanded : ''}`}
                    />
                </div>
                {isExpanded && (
                    <div className={classes.contributionsList}>
                        {userContributions.map(contribution => {
                            const requirement = getRequirementById(contribution.requirement_id)
                            return (
                                <div key={contribution.ID} className={classes.contribution}>
                                    <div className={classes.contributionDetails}>
                                        <div className={classes.requirementName}>
                                            {requirement?.type}
                                        </div>
                                    </div>

                                    <div className={classes.contributionQuantity}>
                                        {contribution.quantity} {requirement?.quantity_mark}
                                    </div>
                                    {contribution.description && (
                                        <div className={classes.contributionDescription}>
                                            {contribution.description}
                                        </div>
                                    )}
                                    {(contribution.contributor.username === userName || userName === organizerName) && (
                                        <button
                                            className={classes.deleteButton}
                                            onClick={(e) => {
                                                e.stopPropagation();
                                                // Delete functionality will be implemented later
                                            }}
                                        >
                                            <DeleteOutlined/>
                                        </button>
                                    )}
                                </div>
                            )
                        })}
                    </div>
                )}
            </div>
        )
    }

    return (
        <div className={classes.outerContainer}>
            <div className={classes.pageContainer}>
                <h1 className={classes.title}>Hall of Fame</h1>
                <p className={classes.description}>
                    See who contributed the most to the party! Click on a participant to view their contributions.
                </p>

                <div className={classes.usersList}>
                    {[...participants]
                        .sort((a, b) => {
                            const aContributions = getTotalContributionsForUser(a.ID);
                            const bContributions = getTotalContributionsForUser(b.ID);
                            return bContributions - aContributions; // Sort in descending order
                        })
                        .map(user => renderUser(user))}
                </div>
            </div>
        </div>
    )
}
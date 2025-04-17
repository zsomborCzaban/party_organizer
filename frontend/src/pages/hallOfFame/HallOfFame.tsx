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

interface ContributionWithRequirement {
    contribution: ContributionPopulated;
    requirementType: string;
    requirementMark: string;
}

interface UserContributionData {
    user: User;
    totalContributions: number;
    contributions: ContributionWithRequirement[];
}

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

    const [userContributionData, setUserContributionData] = useState<UserContributionData[]>([])

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

    useEffect(() => {
        const allRequirements = [...drinkRequirements, ...foodRequirements];
        const allContributions = [...drinkContributions, ...foodContributions];

        const requirementsMap = new Map(
            allRequirements.map(req => [req.ID, { type: req.type, mark: req.quantity_mark }])
        );

        const userData = participants.map(user => {
            const userContributions = allContributions
                .filter(c => c.contributor.ID === user.ID)
                .map(contribution => {
                    const requirement = requirementsMap.get(contribution.requirement_id);
                    return {
                        contribution,
                        requirementType: requirement?.type || '',
                        requirementMark: requirement?.mark || ''
                    };
                });

            return {
                user,
                totalContributions: userContributions.length,
                contributions: userContributions
            };
        });

        setUserContributionData(userData);
    }, [participants, drinkContributions, foodContributions, drinkRequirements, foodRequirements]);

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

    const renderUser = (userData: UserContributionData) => {
        const isExpanded = expandedUsers.has(userData.user.ID)

        return (
            <div key={userData.user.ID} className={classes.userCard}>
                <div 
                    className={classes.userHeader}
                    onClick={() => toggleUser(userData.user.ID)}
                >
                    <div className={classes.userInfo}>
                        <img 
                            src={userData.user.profile_picture_url}
                            alt={userData.user.username}
                            className={classes.profilePicture}
                        />
                        <div className={classes.userDetails}>
                            <div className={classes.userName}>{userData.user.username}</div>
                            <div className={classes.contributionCount}>
                                Total Contributions: {userData.totalContributions}
                            </div>
                        </div>
                    </div>
                    <DownOutlined 
                        className={`${classes.expandIcon} ${isExpanded ? classes.expanded : ''}`}
                    />
                </div>
                {isExpanded && (
                    <div className={classes.contributionsList}>
                        {userData.contributions.map(({contribution, requirementType, requirementMark}) => (
                            <div key={contribution.ID} className={classes.contribution}>
                                <div className={classes.contributionDetails}>
                                    <div className={classes.requirementName}>
                                        {requirementType}
                                    </div>
                                    <div className={classes.contributionQuantity}>
                                        {contribution.quantity} {requirementMark}
                                    </div>
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
                        ))}
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
                    Give the top contributor some praise when the party starts :) . He/She deserves it!
                </p>

                <div className={classes.usersList}>
                    {userContributionData
                        .sort((a, b) => b.totalContributions - a.totalContributions)
                        .map(userData => renderUser(userData))}
                </div>
            </div>
        </div>
    )
}
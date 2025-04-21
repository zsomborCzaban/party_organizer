import {useApi} from "../../context/ApiContext.ts";
import {useNavigate} from "react-router-dom";
import {getUserName} from "../../auth/AuthUserUtil.ts";
import {useEffect, useState} from "react";
import {ContributionPopulated, EMPTY_CONTRIBUTION_POPULATED} from "../../data/types/Contribution.ts";
import {EMPTY_REQUIREMENT_POPULATED, RequirementPopulated} from "../../data/types/Requirement.ts";
import {User} from "../../data/types/User.ts";
import {toast} from "sonner";
import { DownOutlined, DeleteOutlined } from '@ant-design/icons';
import classes from './HallOfFame.module.scss';
import DeleteContributeModal from '../../components/modal/deleteContribution/DeleteContributionModal.tsx';

interface ContributionWithRequirement {
    contribution: ContributionPopulated;
    requirement: RequirementPopulated
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
    const [refreshDrinkContributions, setRefreshDrinkContributions] = useState(0)
    const [refreshFoodContributions, setRefreshFoodContributions] = useState(0)

    const [isDeleteModalVisible, setIsDeleteModalVisible] = useState(false);
    const [deleteModalMode, setDeleteModalMode] = useState<'drink' | 'food'>('drink');
    const [selectedContribution, setSelectedContribution] = useState<ContributionPopulated>(EMPTY_CONTRIBUTION_POPULATED);
    const [selectedRequirement, setSelectedRequirement] = useState<RequirementPopulated>(EMPTY_REQUIREMENT_POPULATED);

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
    }, [api.contributionApi, partyId, refreshDrinkContributions]);

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
    }, [api.contributionApi, partyId, refreshFoodContributions]);

    useEffect(() => {
        const drinkRequirementsMap = new Map(
            drinkRequirements.map(req => [req.ID,  req ])
        );

        const foodRequirementsMap = new Map(
            foodRequirements.map(req => [req.ID,  req ])
        );

        //could be improved
        const userData = participants.map(user => {
            const userDrinkContributions = drinkContributions
                .filter(c => c.contributor.ID === user.ID)
                .map(contribution => {
                    const requirement = drinkRequirementsMap.get(contribution.requirement_id) || EMPTY_REQUIREMENT_POPULATED;
                    return {
                        contribution,
                        requirement,
                    };
                });

            const userFoodContributions = foodContributions
                .filter(c => c.contributor.ID === user.ID)
                .map(contribution => {
                    const requirement = foodRequirementsMap.get(contribution.requirement_id) || EMPTY_REQUIREMENT_POPULATED;
                    return {
                        contribution,
                        requirement,
                    };
                });

            return {
                user,
                totalContributions: userDrinkContributions.length + userFoodContributions.length,
                contributions: [...userDrinkContributions, ...userFoodContributions]
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

    const handleDeleteClick = (contribution: ContributionPopulated, requirement: RequirementPopulated, mode: 'drink' | 'food') => {
        setSelectedContribution(contribution);
        setSelectedRequirement(requirement);
        setDeleteModalMode(mode);
        setIsDeleteModalVisible(true);
    };

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
                        {userData.contributions.map(({contribution, requirement}) => {
                            return (
                                <div key={contribution.ID} className={classes.contribution}>
                                    <div className={classes.contributionDetails}>
                                        {contribution.description && (
                                            <div className={classes.contributionDetails}>
                                                <div className={classes.requirementName}>
                                                    {requirement.type}
                                                </div>
                                            <div className={classes.contributionInfo}>
                                                <div className={classes.contributionQuantity}>
                                                    {contribution.quantity} {requirement.quantity_mark}
                                                </div>
                                                <div className={classes.contributionDescription}>
                                                    {contribution.description}
                                                </div>
                                            </div>
                                            </div>
                                        )}
                                        {!contribution.description && (
                                            <div className={classes.contributionDetails}>
                                                <div className={classes.requirementName}>
                                                    {requirement.type}, {contribution.quantity} {requirement.quantity_mark}
                                                </div>
                                            </div>
                                        )}
                                    </div>
                                    {(contribution.contributor.username === userName || userName === organizerName) && (
                                        <button
                                            className={classes.deleteButton}
                                            onClick={(e) => {
                                                e.stopPropagation();
                                                if (requirement) {
                                                    handleDeleteClick(
                                                        contribution,
                                                        requirement,
                                                        drinkRequirements.some(req => req.ID === requirement.ID && req.type === requirement.type && req.target_quantity === requirement.target_quantity && req.quantity_mark === requirement.quantity_mark) ? 'drink' : 'food'
                                                    );
                                                }
                                            }}
                                        >
                                            <DeleteOutlined/>
                                        </button>
                                    )}
                                </div>
                            );
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
                    Give the top contributor some praise when the party starts :) . He/She deserves it!
                </p>

                <div className={classes.usersList}>
                    {userContributionData
                        .sort((a, b) => b.totalContributions - a.totalContributions)
                        .map(userData => renderUser(userData))}
                </div>
            </div>

            <DeleteContributeModal
                visible={isDeleteModalVisible}
                onClose={() => setIsDeleteModalVisible(false)}
                mode={deleteModalMode}
                contributionId={selectedContribution.ID}
                contribution={selectedContribution}
                requirement={selectedRequirement}
                onFoodSuccess={() => setRefreshFoodContributions(prevState => (prevState+1)%2)}
                onDrinkSuccess={() => setRefreshDrinkContributions(prevState => (prevState+1)%2)}
            />
        </div>
    )
}
import {useApi} from "../../context/ApiContext.ts";
import {useEffect, useState} from "react";
import {ContributionPopulated, EMPTY_CONTRIBUTION_POPULATED} from "../../data/types/Contribution.ts";
import {EMPTY_REQUIREMENT_POPULATED, RequirementPopulated} from "../../data/types/Requirement.ts";
import {getUserName} from "../../auth/AuthUserUtil.ts";
import {useNavigate} from "react-router-dom";
import {toast} from "sonner";
import { DownOutlined, DeleteOutlined } from '@ant-design/icons';
import classes from './Contributions.module.scss';
import { ContributeModal } from '../../components/modal/createContribution/ContributeModal.tsx';
import DeleteContributeModal from '../../components/modal/deleteContribution/DeleteContributionModal.tsx';

export const Contributions = () => {
    const api = useApi()
    const navigate = useNavigate()
    const userName = getUserName() || ''
    const organizerName = localStorage.getItem('partyOrganizerName') || ''
    const partyId = Number(localStorage.getItem('partyId') || '-1')

    const [drinkContributions, setDrinkContributions] = useState<ContributionPopulated[]>([])
    const [foodContributions, setFoodContributions] = useState<ContributionPopulated[]>([])
    const [drinkRequirements, setDrinkRequirements] = useState<RequirementPopulated[]>([])
    const [foodRequirements, setFoodRequirements] = useState<RequirementPopulated[]>([])
    const [refreshDrinkContributions, setRefreshDrinkContributions] = useState(0)
    const [refreshFoodContributions, setRefreshFoodContributions] = useState(0)
    const [expandedDrinkRequirements, setExpandedDrinkRequirements] = useState<Set<number>>(new Set())
    const [expandedFoodRequirements, setExpandedFoodRequirements] = useState<Set<number>>(new Set())

    const [isModalVisible, setIsModalVisible] = useState(false);
    const [modalMode, setModalMode] = useState<'drink' | 'food'>('drink');
    
    const [isDeleteModalVisible, setIsDeleteModalVisible] = useState(false);
    const [deleteModalMode, setDeleteModalMode] = useState<'drink' | 'food'>('drink');
    const [selectedContribution, setSelectedContribution] = useState<ContributionPopulated>(EMPTY_CONTRIBUTION_POPULATED);
    const [selectedRequirement, setSelectedRequirement] = useState<RequirementPopulated>(EMPTY_REQUIREMENT_POPULATED);

    useEffect(() => {
        if(userName === '' || organizerName === '' || partyId === -1){
            toast.error('Navigation error')
            navigate('/partyHome')
        }
    }, [navigate, organizerName, partyId, userName]);

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

    const toggleRequirement = (requirementId: number, isDrink: boolean) => {
        if (isDrink) {
            setExpandedDrinkRequirements(prev => {
                const newSet = new Set(prev)
                if (newSet.has(requirementId)) {
                    newSet.delete(requirementId)
                } else {
                    newSet.add(requirementId)
                }
                return newSet
            })
        } else {
            setExpandedFoodRequirements(prev => {
                const newSet = new Set(prev)
                if (newSet.has(requirementId)) {
                    newSet.delete(requirementId)
                } else {
                    newSet.add(requirementId)
                }
                return newSet
            })
        }
    }

    const getContributionsForRequirement = (requirementId: number, contributions: ContributionPopulated[]) => {
        return contributions.filter(contribution => contribution.requirement_id === requirementId)
    }

    const getTotalContributedQuantity = (requirementId: number, contributions: ContributionPopulated[]) => {
        return contributions
            .filter(contribution => contribution.requirement_id === requirementId)
            .reduce((sum, contribution) => sum + contribution.quantity, 0)
    }

    const handleDeleteClick = (contribution: ContributionPopulated, requirement: RequirementPopulated, mode: 'drink' | 'food') => {
        setSelectedContribution(contribution);
        setSelectedRequirement(requirement);
        setDeleteModalMode(mode);
        setIsDeleteModalVisible(true);
    };

    const handleContributeClick = (mode: 'drink' | 'food') => {
        setModalMode(mode);
        setIsModalVisible(true);
    };

    const getModalOptions = (mode: 'drink' | 'food') => {
        const requirements = mode === 'drink' ? drinkRequirements : foodRequirements;
        return requirements.map(req => ({
            value: req.ID,
            label: `${req.type} (${req.target_quantity} ${req.quantity_mark})`
        }));
    };

    const renderRequirement = (requirement: RequirementPopulated, contributions: ContributionPopulated[], isDrink: boolean) => {
        const isExpanded = isDrink 
            ? expandedDrinkRequirements.has(requirement.ID)
            : expandedFoodRequirements.has(requirement.ID)
        const requirementContributions = getContributionsForRequirement(requirement.ID, contributions)
        const totalContributed = getTotalContributedQuantity(requirement.ID, contributions)

        return (
            <div key={requirement.ID} className={classes.requirement}>
                <div 
                    className={classes.requirementHeader}
                    onClick={() => toggleRequirement(requirement.ID, isDrink)}
                >
                    <div className={classes.requirementInfo}>
                        <div className={classes.requirementName}>{requirement.type}</div>
                        <div className={classes.requirementQuantity}>
                            {totalContributed} / {requirement.target_quantity} {requirement.quantity_mark}
                        </div>
                    </div>
                    <DownOutlined 
                        className={`${classes.expandIcon} ${isExpanded ? classes.expanded : ''}`}
                    />
                </div>
                {isExpanded && (
                    <div className={classes.requirementContent}>
                        {requirementContributions.map(contribution => (
                            <div key={contribution.ID} className={classes.contribution}>
                                <div className={classes.contributorInfo}>
                                    <img 
                                        src={contribution.contributor.profile_picture_url}
                                        alt={contribution.contributor.username}
                                        className={classes.profilePicture}
                                    />
                                    <span className={classes.contributorName}>
                                        {contribution.contributor.username}
                                    </span>
                                </div>
                                <div className={classes.contributionDetails}>
                                    <div className={classes.contributionQuantity}>
                                        {contribution.quantity} {requirement.quantity_mark}
                                    </div>
                                    {contribution.description && (
                                        <div className={classes.contributionDescription}>
                                            {contribution.description}
                                        </div>
                                    )}
                                </div>
                                {(contribution.contributor.username === userName || userName === organizerName) && (
                                    <button 
                                        className={classes.deleteButton}
                                        onClick={(e) => {
                                            e.stopPropagation();
                                            handleDeleteClick(contribution, requirement, isDrink ? 'drink' : 'food');
                                        }}
                                    >
                                        <DeleteOutlined />
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
                <h1 className={classes.title}>Contributions</h1>

                <div className={classes.contributionsSection}>
                    <div className={classes.section}>
                        <h2 className={classes.sectionTitle}>Drink Contributions</h2>
                        <div className={classes.buttonContainer}>
                            <button 
                                className={classes.addButton}
                                onClick={() => handleContributeClick('drink')}
                            >
                                Contribute
                            </button>
                        </div>
                        {drinkRequirements.map(requirement =>
                            renderRequirement(requirement, drinkContributions, true)
                        )}
                    </div>

                    <div className={classes.section}>
                        <h2 className={classes.sectionTitle}>Food Contributions</h2>
                        <div className={classes.buttonContainer}>
                            <button 
                                className={classes.addButton}
                                onClick={() => handleContributeClick('food')}
                            >
                                Contribute
                            </button>
                        </div>
                        {foodRequirements.map(requirement =>
                            renderRequirement(requirement, foodContributions, false)
                        )}
                    </div>
                </div>
            </div>

            <ContributeModal
                visible={isModalVisible}
                onClose={() => setIsModalVisible(false)}
                options={getModalOptions(modalMode)}
                mode={modalMode}
                onFoodSuccess={() => setRefreshFoodContributions(prevState => (prevState+1)%2)}
                onDrinkSuccess={() => setRefreshDrinkContributions(prevState => (prevState+1)%2)}
            />

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
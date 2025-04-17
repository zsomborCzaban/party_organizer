import {useApi} from "../../context/ApiContext.ts";
import {useEffect, useState} from "react";
import {ContributionPopulated} from "../../data/types/Contribution.ts";
import {RequirementPopulated} from "../../data/types/Requirement.ts";
import {getUserName} from "../../auth/AuthUserUtil.ts";
import {useNavigate} from "react-router-dom";
import {toast} from "sonner";
import { DownOutlined } from '@ant-design/icons';
import classes from './Contributions.module.scss';

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
    const [expandedRequirements, setExpandedRequirements] = useState<Set<number>>(new Set())

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

    const toggleRequirement = (requirementId: number) => {
        setExpandedRequirements(prev => {
            const newSet = new Set(prev)
            if (newSet.has(requirementId)) {
                newSet.delete(requirementId)
            } else {
                newSet.add(requirementId)
            }
            return newSet
        })
    }

    const getContributionsForRequirement = (requirementId: number, contributions: ContributionPopulated[]) => {
        return contributions.filter(contribution => contribution.requirement_id === requirementId)
    }

    const getTotalContributedQuantity = (requirementId: number, contributions: ContributionPopulated[]) => {
        return contributions
            .filter(contribution => contribution.requirement_id === requirementId)
            .reduce((sum, contribution) => sum + contribution.quantity, 0)
    }

    const renderRequirement = (requirement: RequirementPopulated, contributions: ContributionPopulated[]) => {
        const isExpanded = expandedRequirements.has(requirement.id)
        const requirementContributions = getContributionsForRequirement(requirement.id, contributions)
        const totalContributed = getTotalContributedQuantity(requirement.id, contributions)

        return (
            <div key={requirement.id} className={classes.requirement}>
                <div 
                    className={classes.requirementHeader}
                    onClick={() => toggleRequirement(requirement.id)}
                >
                    <div className={classes.requirementInfo}>
                        <div className={classes.requirementName}>{requirement.name}</div>
                        <div className={classes.requirementQuantity}>
                            {totalContributed} / {requirement.quantity} {requirement.quantity_mark}
                        </div>
                    </div>
                    <DownOutlined 
                        className={`${classes.expandIcon} ${isExpanded ? classes.expanded : ''}`}
                    />
                </div>
                {isExpanded && (
                    <div className={classes.requirementContent}>
                        {requirementContributions.map(contribution => (
                            <div key={contribution.id} className={classes.contribution}>
                                <div className={classes.contributorInfo}>
                                    <img 
                                        src={contribution.contributor.profile_picture} 
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
                            </div>
                        ))}
                    </div>
                )}
            </div>
        )
    }

    return (
        <div className={classes.pageContainer}>
            <h1 className={classes.title}>Manage Contributions</h1>
            <p className={classes.description}>
                View and manage all contributions for the party. You can see what has been contributed and what is still needed.
            </p>

            <div className={classes.contributionsSection}>
                <div className={classes.section}>
                    <h2 className={classes.sectionTitle}>Drink Contributions</h2>
                    <button className={classes.addButton}>Add Drink Contribution</button>
                    {drinkRequirements.map(requirement => 
                        renderRequirement(requirement, drinkContributions)
                    )}
                </div>

                <div className={classes.section}>
                    <h2 className={classes.sectionTitle}>Food Contributions</h2>
                    <button className={classes.addButton}>Add Food Contribution</button>
                    {foodRequirements.map(requirement => 
                        renderRequirement(requirement, foodContributions)
                    )}
                </div>
            </div>
        </div>
    )
}
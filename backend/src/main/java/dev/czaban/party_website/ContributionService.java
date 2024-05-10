package dev.czaban.party_website;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class ContributionService {

    @Autowired //auto init
    private ContributionRepository contributionRepository;

    public List<Contribution> allContribution(){
        return contributionRepository.findAll();
    }

    public List<Contribution> allContributionWithType(String type){
        return contributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList());
    }

    public void createContribution(String type, double quantity, String contributorName, String description){ //should be bool
        Contribution contribution = new Contribution(type, quantity, contributorName, description);
        contributionRepository.insert(contribution);
    }
}

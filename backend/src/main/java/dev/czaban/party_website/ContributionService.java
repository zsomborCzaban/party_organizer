package dev.czaban.party_website;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Service
public class ContributionService {

    @Autowired //auto init
    private ContributionRepository contributionRepository;

    Logger logger = LoggerFactory.getLogger(ContributionService.class);

    private final Set<String> types = new HashSet<>(){{ //todo: get from db
        add("beer");
        add("wine");
        add("spirit");
    }};;

    public List<Contribution> allContribution(){
        return contributionRepository.findAll();
    }

    public List<Contribution> allContributionWithType(String type){
        return contributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList());
    }

    public Boolean createContribution(Contribution contribution){ //should be bool
        logger.info("inserting to DB");
        try{
            contributionRepository.insert(contribution);
            logger.info("insertion to DB Successful");
            return true;
        } catch (Exception e){
            logger.error("Error while inserting to DB: {}", e.getMessage());
            return false;
        }


    }

    public boolean isValidType(String type){
        return types.contains(type);
    }
}

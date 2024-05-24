package dev.czaban.party_website.services;

import dev.czaban.party_website.models.DrinkContribution;
import dev.czaban.party_website.repositories.ContributionRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Service
public class DrinkContributionService {

    private final ContributionRepository contributionRepository;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);

    public DrinkContributionService(ContributionRepository contributionRepository) {
        this.contributionRepository = contributionRepository;
    }

    private final Set<String> types = new HashSet<>(){{ //todo: get from db
        add("beer");
        add("wine");
        add("spirit");
    }};

    public List<DrinkContribution> allContribution(){
        return contributionRepository.findAll();
    }

    public List<DrinkContribution> allContributionWithType(String type){
        logger.info("return of allcontributions with type: {}", contributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList()));
        return contributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList());
    }

    public Boolean createContribution(DrinkContribution drinkContribution){ //should be bool
        logger.info("inserting to DB");
        try{
            contributionRepository.insert(drinkContribution);
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

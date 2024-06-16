package dev.czaban.party_website.services;

import com.fasterxml.jackson.databind.node.ObjectNode;
import dev.czaban.party_website.models.drink.DrinkContribution;
import dev.czaban.party_website.models.drink.DrinkTypes;
import dev.czaban.party_website.repositories.DrinkContributionRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Service
public class DrinkContributionService {

    private final DrinkContributionRepository drinkContributionRepository;
    private final DrinkTypes drinkTypes;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);

    public DrinkContributionService(DrinkContributionRepository drinkContributionRepository, DrinkTypes drinkTypes) {
        this.drinkContributionRepository = drinkContributionRepository;
        this.drinkTypes = drinkTypes;
    }

    private final Set<String> types = new HashSet<>(){{ //todo: get from db
        add("beer");
        add("wine");
        add("spirit");
    }};

    public List<DrinkContribution> allContribution(){
        return drinkContributionRepository.findAll();
    }

    public List<DrinkContribution> allContributionWithType(String type){
        logger.info("return of allcontributions with type: {}", drinkContributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList()));
        System.out.println(drinkTypes.getDrinkType("beer"));
        return drinkContributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList());
    }

    public Boolean createContribution(ObjectNode json){ //should be bool
        DrinkContribution dc = new DrinkContribution(drinkTypes.getDrinkType(json.get("type").asText()), json.get("quantity").asDouble(), json.get("description").asText(), json.get("contributorName").asText());
        logger.info("inserting to DB");
        try{
            drinkContributionRepository.save(dc);
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

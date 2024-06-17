package dev.czaban.party_website.services.drink;

import com.fasterxml.jackson.databind.node.ObjectNode;
import dev.czaban.party_website.models.drink.DrinkContribution;
import dev.czaban.party_website.repositories.DrinkContributionRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class DrinkContributionService {

    private final DrinkContributionRepository drinkContributionRepository;
    private final DrinkTypeService drinkTypeService;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);

    public DrinkContributionService(DrinkContributionRepository drinkContributionRepository, DrinkTypeService drinkTypeService) {
        this.drinkContributionRepository = drinkContributionRepository;
        this.drinkTypeService = drinkTypeService;
    }

    public List<DrinkContribution> allContribution(){
        return drinkContributionRepository.findAll();
    }

    public List<DrinkContribution> allContributionWithType(String type){
        logger.info("return of allcontributions with type: {}", drinkContributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList()));
        System.out.println(drinkTypeService.getDrinkType("beer"));
        return drinkContributionRepository.findAll().stream().filter(c -> c.getType().equals(type)).collect(Collectors.toList());
    }

    public Boolean createContribution(ObjectNode json){ //should be bool
        DrinkContribution dc = new DrinkContribution(drinkTypeService.getDrinkType(json.get("type").asText()), json.get("quantity").asDouble(), json.get("description").asText(), json.get("contributorName").asText());
        logger.info("inserting to DB");
        try{
            drinkContributionRepository.insert(dc);
            logger.info("insertion to DB Successful");
            return true;
        } catch (Exception e){
            logger.error("Error while inserting to DB: {}", e.getMessage());
            return false;
        }


    }
}

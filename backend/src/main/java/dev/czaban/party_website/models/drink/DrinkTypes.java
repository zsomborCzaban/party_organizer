package dev.czaban.party_website.models.drink;

import dev.czaban.party_website.repositories.DrinkTypeRepository;
import dev.czaban.party_website.services.DrinkContributionService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.HashSet;
import java.util.List;

@Component   //by default components and beans are singleton, so no need to implement it :(
public class DrinkTypes {

    //private static DrinkTypes instance = new DrinkTypes();
    private final DrinkTypeRepository drinkTypeRepository;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);
    private final HashMap<String, DrinkType> drinkTypes = new HashMap<>();

    private DrinkTypes(DrinkTypeRepository drinkTypeRepository){
        this.drinkTypeRepository = drinkTypeRepository;
        drinkTypeRepository.findAll().forEach(c -> drinkTypes.put(c.getDrinkType(), c));
    }

    public boolean isValidDrinkType(String drinkType){
        return drinkTypes.containsKey(drinkType);
    }

    public DrinkType getDrinkType(String drinkType){
        return drinkTypes.get(drinkType);
    }

    public List<DrinkType> getAllDrinkType(){
        return drinkTypeRepository.findAll();
    }
}

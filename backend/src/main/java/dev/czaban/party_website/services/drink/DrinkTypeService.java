package dev.czaban.party_website.services.drink;


import dev.czaban.party_website.models.drink.DrinkType;
import dev.czaban.party_website.repositories.DrinkTypeRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.HashMap;

@Service
public class DrinkTypeService {

    //private static DrinkTypes instance = new DrinkTypes();
    private final DrinkTypeRepository drinkTypeRepository;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);
    private final HashMap<String, DrinkType> drinkTypes = new HashMap<>();

    private DrinkTypeService(DrinkTypeRepository drinkTypeRepository){
        this.drinkTypeRepository = drinkTypeRepository;
        drinkTypeRepository.findAll().forEach(c -> drinkTypes.put(c.getDrinkType(), c));
    }

    public boolean isValidDrinkType(String drinkType){
        return drinkTypes.containsKey(drinkType);
    }

    public DrinkType getDrinkType(String drinkType){
        return drinkTypes.get(drinkType);
    }

    public Collection<DrinkType> getAllDrinkType(){
        return drinkTypes.values();
    }

}

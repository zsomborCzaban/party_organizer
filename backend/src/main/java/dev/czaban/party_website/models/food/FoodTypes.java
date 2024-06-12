package dev.czaban.party_website.models.food;

import dev.czaban.party_website.models.drink.DrinkType;
import dev.czaban.party_website.repositories.DrinkTypeRepository;
import dev.czaban.party_website.repositories.FoodTypeRepository;
import dev.czaban.party_website.services.DrinkContributionService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.HashMap;
import java.util.HashSet;
import java.util.stream.Collectors;

public class FoodTypes {
    private final FoodTypeRepository foodTypeRepository;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);
    private final HashMap<String, FoodType> foodTypes = new HashMap<>();

    private FoodTypes(FoodTypeRepository foodTypeRepository){
        this.foodTypeRepository = foodTypeRepository;
        foodTypeRepository.findAll().forEach(c ->  foodTypes.put(c.getFoodType(), c));
    }

    public boolean isValidDrinkType(String foodType){
        return foodTypes.containsKey(foodType);
    }

    public FoodType getFoodType(String foodType){
        return foodTypes.get(foodType);
    }
}

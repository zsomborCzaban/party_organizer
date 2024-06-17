package dev.czaban.party_website.services.food;

import dev.czaban.party_website.models.food.FoodType;
import dev.czaban.party_website.repositories.FoodTypeRepository;
import dev.czaban.party_website.services.drink.DrinkContributionService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.Collection;
import java.util.HashMap;

@Service
public class FoodTypeService {
    private final FoodTypeRepository foodTypeRepository;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionService.class);
    private final HashMap<String, FoodType> foodTypes = new HashMap<>();

    private FoodTypeService(FoodTypeRepository foodTypeRepository){
        this.foodTypeRepository = foodTypeRepository;
        foodTypeRepository.findAll().forEach(c ->  foodTypes.put(c.getFoodType(), c));
    }

    public boolean isValidFoodType(String foodType){
        return foodTypes.containsKey(foodType);
    }

    public FoodType getFoodType(String foodType){
        return foodTypes.get(foodType);
    }

    public Collection<FoodType> getAllDrinkType(){
        return foodTypes.values();
    }
}

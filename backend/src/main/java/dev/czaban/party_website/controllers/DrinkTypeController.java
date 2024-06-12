package dev.czaban.party_website.controllers;

import dev.czaban.party_website.models.drink.DrinkContribution;
import dev.czaban.party_website.models.drink.DrinkType;
import dev.czaban.party_website.models.drink.DrinkTypes;
import dev.czaban.party_website.services.DrinkContributionService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/api/drink_type")
public class DrinkTypeController {

    private final DrinkTypes drinkTypes;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionController.class);

    public DrinkTypeController(DrinkTypes drinkTypes) {
        this.drinkTypes = drinkTypes;
    }

    @GetMapping("/")
    public ResponseEntity<List<DrinkType>> AllDrinkType(){
        return  new ResponseEntity<>(drinkTypes.getAllDrinkType(), HttpStatus.OK);
    }

}

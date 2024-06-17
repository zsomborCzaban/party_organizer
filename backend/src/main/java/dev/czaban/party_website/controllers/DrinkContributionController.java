package dev.czaban.party_website.controllers;

import com.fasterxml.jackson.databind.node.ObjectNode;
import dev.czaban.party_website.services.drink.DrinkContributionService;
import dev.czaban.party_website.models.drink.DrinkContribution;
import dev.czaban.party_website.services.drink.DrinkTypeService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/contributions")
public class DrinkContributionController {

    private final DrinkContributionService drinkContributionService;
    private final DrinkTypeService drinkTypeService;
    private final Logger logger = LoggerFactory.getLogger(DrinkContributionController.class);

    public DrinkContributionController(DrinkContributionService drinkContributionService, DrinkTypeService drinkTypeService) {
        this.drinkContributionService = drinkContributionService;
        this.drinkTypeService = drinkTypeService;
    }

    @GetMapping("/{type}")
    public ResponseEntity<List<DrinkContribution>> getContributionsByType(@PathVariable String type){
        if(!drinkTypeService.isValidDrinkType(type)){
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST); //todo: return message
        }
        return  new ResponseEntity<>(drinkContributionService.allContributionWithType(type), HttpStatus.OK);
    }

    @CrossOrigin //todo: allow only for frontend
    @GetMapping()
    public ResponseEntity<List<DrinkContribution>> getContributions(){
        System.out.println("requested");
        System.out.println(drinkContributionService.allContribution());
        return  new ResponseEntity<>(drinkContributionService.allContribution(), HttpStatus.OK);
    }

    //todo: make post with data validation and with auth
    @CrossOrigin
    @PostMapping(path ="/contribution", consumes = "application/json" /*produces = MediaType.APPLICATION_JSON_VALUE --- responseEntity will use string as raw value even if this is specified*/)    //solution: create wrapper class for the string :/
    public ResponseEntity<String> createContribution(@RequestBody ObjectNode json){    //todo: send all errors in the same response. (solution?: Maybe make a costum bean validator)

//        if(!drinkContributionService.isValidType(drinkContribution.getType())){
//            return new ResponseEntity<>("{type: 'incorrect type, choose from the available options'}", HttpStatus.BAD_REQUEST);
//        }
//
        //return drinkContributionService.createContribution(json);
        if(drinkContributionService.createContribution(json)){
            return new ResponseEntity<>("Contribution created", HttpStatus.OK);
        }
        return new ResponseEntity<>("DB error", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public Map<String, String> handleValidationExceptions(MethodArgumentNotValidException ex){
        Map<String, String> errors = new HashMap<>();

        ex.getBindingResult().getAllErrors().forEach(error -> {
            String fieldName = ((FieldError) error).getField();
            String errorMessage = error.getDefaultMessage();
            errors.put(fieldName, errorMessage);
        });

        return errors;
    }
}

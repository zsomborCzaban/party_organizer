package dev.czaban.party_website.controllers;

import dev.czaban.party_website.services.ContributionService;
import dev.czaban.party_website.models.Contribution;
import jakarta.validation.Valid;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
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
public class ContributionController {

    @Autowired
    private ContributionService contributionService;

    Logger logger = LoggerFactory.getLogger(ContributionController.class);

    @GetMapping("/beers")
    public ResponseEntity<List<Contribution>> beers(){
        return  new ResponseEntity<>(contributionService.allContributionWithType("beer"), HttpStatus.OK);
    }

    @CrossOrigin //todo: allow only for frontend
    @GetMapping("/")
    public ResponseEntity<List<Contribution>> contributions(){
        System.out.println("requested");
        System.out.println(contributionService.allContribution());
        return  new ResponseEntity<>(contributionService.allContribution(), HttpStatus.OK);
    }

    //todo: make post with data validation and with auth
    @CrossOrigin
    @PostMapping(path ="/contribution", consumes = "application/json" /*produces = MediaType.APPLICATION_JSON_VALUE --- responseEntity will use string as raw value even if this is specified*/)    //solution: create wrapper class for the string :/
    public ResponseEntity<String> createContribution(@Valid @RequestBody Contribution contribution){    //todo: send all errors in the same response. (solution?: Maybe make a costum bean validator)
        System.out.println(contribution);

        if(!contributionService.isValidType(contribution.getType())){
            return new ResponseEntity<>("{type: 'incorrect type, choose from the available options'}", HttpStatus.BAD_REQUEST);
        }
        if(contributionService.createContribution(contribution)){
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
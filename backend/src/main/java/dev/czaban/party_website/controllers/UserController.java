package dev.czaban.party_website.controllers;

import dev.czaban.party_website.models.DrinkContribution;
import dev.czaban.party_website.models.User;
import dev.czaban.party_website.services.UserService;
import jakarta.validation.Valid;
import org.bson.types.ObjectId;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpStatusCode;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;

//this class it to create/delete users through the api
@RestController
@RequestMapping("/api/admin")
public class UserController {

    private static final Logger logger = LoggerFactory.getLogger(UserController.class);
    private final UserService userService;

    public UserController(UserService userService) {
        this.userService = userService;
    }

    @GetMapping("/users")
    public ResponseEntity<String> testUser(){

            return new ResponseEntity<>("test message", HttpStatus.OK);
    }

    @PostMapping("/users")
    public ResponseEntity<String> createUser(@RequestBody User user){ //todo: validate maybe
        logger.info("user creation started");

        if(userService.createUser(user)){
            return new ResponseEntity<>("user created", HttpStatus.OK);
        }
        return new ResponseEntity<>("DB error", HttpStatus.INTERNAL_SERVER_ERROR);
    }

    @DeleteMapping("/users/{username}")
    public ResponseEntity<String> deleteUser(@PathVariable String username){
        userService.deleteUser(username);
        logger.info("User deleted with name: {}", username);
        return new ResponseEntity<>("deletion endpoint called", HttpStatus.OK);
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

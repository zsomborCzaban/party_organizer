package dev.czaban.party_website.controllers;

import dev.czaban.party_website.services.TokenService;
import org.apache.el.parser.Token;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AuthController {

    private static final Logger logger = LoggerFactory.getLogger(AuthController.class);

    private final TokenService tokenService;

    public AuthController(TokenService tokenService){
        this.tokenService = tokenService;
    }

    @PostMapping("/token")
    public String token(Authentication authentication){
        logger.debug("Token requested for user: {}", authentication.getName());
        String token = tokenService.generateToken(authentication);
        logger.debug("Token granted: {}", token);
        return token;
    }
}

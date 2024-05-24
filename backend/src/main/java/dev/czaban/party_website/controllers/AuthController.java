package dev.czaban.party_website.controllers;

import dev.czaban.party_website.models.LoginRequest;
import dev.czaban.party_website.services.TokenService;
import org.apache.el.parser.Token;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AuthController {

    private static final Logger logger = LoggerFactory.getLogger(AuthController.class);
    private final TokenService tokenService;
    private final AuthenticationManager authManager;

    public AuthController(TokenService tokenService, AuthenticationManager authManager){
        this.tokenService = tokenService;
        this.authManager = authManager;
    }

    @PostMapping("/token")
    public String token(@RequestBody LoginRequest userLogin){
        Authentication authentication = authManager.authenticate(new UsernamePasswordAuthenticationToken(userLogin.username(), userLogin.password()));
        logger.debug("Token requested for user: {}", authentication.getName());
        String token = tokenService.generateToken(authentication);
        logger.debug("Token granted: {}", token);
        return token;
    }
}

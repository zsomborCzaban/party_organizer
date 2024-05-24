package dev.czaban.party_website.services;


import dev.czaban.party_website.models.User;
import dev.czaban.party_website.repositories.UserRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import java.util.Optional;

//do not put in prod
@Service
public class UserService {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;
    private final Logger logger = LoggerFactory.getLogger(UserService.class);

    public UserService(UserRepository userRepository, PasswordEncoder passwordEncoder) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
    }

//    public boolean createUser(String username, String password, String contributorName, String roles){
//        User user = new User(username, passwordEncoder.encode(password), contributorName, roles);
//        logger.info("creating user: {} in DB", user.getUserName());
//        try{
//            userRepository.insert(user);
//            logger.info("creating user: {} in DB successful", user.getUserName());
//            return true;
//        } catch (Exception e){
//            logger.error("Error while creating user: {} in DB {}", user.getUserName(), e.getMessage());
//            return false;
//        }
//    }

    public boolean createUser(User user){
        user.setPassword(passwordEncoder.encode(user.getPassword()));
        logger.info("creating user: {} in DB", user.getUsername());
        try{
            userRepository.insert(user);
            logger.info("creating user: {} in DB successful", user.getUsername());
            return true;
        } catch (Exception e){
            logger.error("Error while creating user: {} in DB {}", user.getUsername(), e.getMessage());
            return false;
        }
    }

    public void deleteUser(String username){
        Optional<User> result = userRepository.findByUsername(username);
        result.ifPresent(user -> userRepository.deleteById(user.getId()));

    }
}

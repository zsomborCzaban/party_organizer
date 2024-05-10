package dev.czaban.party_website;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/contributions")
public class ContributionController {

    @Autowired
    private ContributionService contributionService;

    @GetMapping("/beers")
    public ResponseEntity<List<Contribution>> beers(){
        return  new ResponseEntity<>(contributionService.allContributionWithType("beer"), HttpStatus.OK);
    }

    @CrossOrigin //todo: allow only for frontend
    @GetMapping()
    public ResponseEntity<List<Contribution>> contributions(){
        System.out.println("requested");
        System.out.println(contributionService.allContribution());
        return  new ResponseEntity<>(contributionService.allContribution(), HttpStatus.OK);
    }

    //todo: make post with data validation and with auth
    @PostMapping()
    public ResponseEntity<String> createContribution(@RequestBody Map<String, String> payload){
        return new ResponseEntity<>("hehe,  forbided ^^", HttpStatus.FORBIDDEN);
    }
}

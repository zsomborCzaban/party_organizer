package dev.czaban.party_website.controllers;

import dev.czaban.party_website.configs.SecurityConfig;
import dev.czaban.party_website.repositories.UserRepository;
import dev.czaban.party_website.services.DrinkContributionService;
import dev.czaban.party_website.services.JpaUserDetailsService;
import dev.czaban.party_website.services.TokenService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.boot.test.mock.mockito.MockBean;
import org.springframework.context.annotation.Import;
import org.springframework.http.MediaType;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest({DrinkContributionController.class, AuthController.class})
@Import({SecurityConfig.class, TokenService.class, JpaUserDetailsService.class})
class DrinkContributionControllerTest {

    @Autowired
    private MockMvc mvc;

    @MockBean
    DrinkContributionService drinkContributionService;
    @MockBean
    UserRepository userRepository;

    private final String CONTRIBUTIONS_END_POINT_PATH = "/api/contributions";
    private final String TOKEN_END_POINT_PATH = "/token";

    @Test
    void contributionsUnauthenticated() throws Exception{
        this.mvc.perform(get(CONTRIBUTIONS_END_POINT_PATH))
                .andExpect(status().isUnauthorized());
    }

    @Test
    void contributionsAuthenticated() throws Exception{
        MvcResult result = this.mvc.perform(post(TOKEN_END_POINT_PATH)
                .contentType(MediaType.APPLICATION_JSON)
                .content("{ \"username\": \"zsombor\", \"password\": \"strong_password_haha\" }"))
                .andExpect(status().isOk())
                .andReturn();

        String token = result.getResponse().getContentAsString();

        this.mvc.perform(get(CONTRIBUTIONS_END_POINT_PATH)
                .header("Authorization", "Bearer " + token))
                .andExpect(status().isOk());
    }
}
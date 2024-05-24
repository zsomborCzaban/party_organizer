package dev.czaban.party_website.models;


import jakarta.persistence.Entity;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "drink_types")
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class DrinkType {

    private int id;
    private String drinkType;

}

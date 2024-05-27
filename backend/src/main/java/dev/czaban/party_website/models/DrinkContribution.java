package dev.czaban.party_website.models;

import jakarta.validation.constraints.*;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "drink_contributions")
@AllArgsConstructor
@NoArgsConstructor
public class DrinkContribution {

    @Id
    private ObjectId id;

    @NotEmpty(message = "type is mandatory")
    private String type; //maybe enum
    public String getType() {
        return type;
    }

    @NotNull(message = "quantity is mandatory")
    @Digits(integer = 20, fraction = 20, message="quantity must be a number with less than 20 digits on both sides of the decimal point") //make a better error message lol
    private double quantity;

    @NotEmpty(message = "contributor_name is mandatory")
    @Size(min = 1, max = 100, message="contributor_name must be between 1 and 100 characters")
    private String contributor_name;

    @Size(max = 300, message="Description cannot be longer than 300 characters")
    private String description;

    public DrinkContribution(String type, double quantity, String description, String contributor_name) {
        this.type = type;
        this.quantity = quantity;
        this.description = description;
        this.contributor_name = contributor_name;
    }

    public void setType(String type) {
        this.type = type;
    }

    public double getQuantity() {
        return quantity;
    }

    public void setQuantity(double quantity) {
        this.quantity = quantity;
    }

    public String getContributor_name() {
        return contributor_name;
    }

    public void setContributor_name(String contributor_name) {
        this.contributor_name = contributor_name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public ObjectId getId() {
        return id;
    }

    public void setId(ObjectId id) {
        this.id = id;
    }
}

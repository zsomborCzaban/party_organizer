package dev.czaban.party_website.models.drink;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.validation.constraints.*;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@Entity
@AllArgsConstructor
@NoArgsConstructor
public class DrinkContribution {

    @Id
    @GeneratedValue
    private Long id;

    @NotNull(message = "type is mandatory")
    private DrinkType type; //maybe enum

    @NotNull(message = "quantity is mandatory")
    @Digits(integer = 20, fraction = 20, message="quantity must be a number with less than 20 digits on both sides of the decimal point") //make a better error message lol
    private double quantity;

    @NotEmpty(message = "contributorName is mandatory")
    @Size(min = 1, max = 100, message="contributorName must be between 1 and 100 characters")
    @Column(name = "contributor_name")
    private String contributorName;

    @Size(max = 300, message="Description cannot be longer than 300 characters")
    private String description;

    public DrinkContribution(DrinkType type, double quantity, String description, String contributorName) {
        this.type = type;
        this.quantity = quantity;
        this.description = description;
        this.contributorName = contributorName;
    }

    public DrinkType getType() {
        return type;
    }

    public void setType(DrinkType type) {
        this.type = type;


    }

    public double getQuantity() {
        return quantity;
    }

    public void setQuantity(double quantity) {
        this.quantity = quantity;
    }

    public String getContributor_name() {
        return contributorName;
    }

    public void setContributor_name(String contributorName) {
        this.contributorName = contributorName;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }
}

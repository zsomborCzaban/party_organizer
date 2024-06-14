package dev.czaban.party_website.models.food;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@Entity
@AllArgsConstructor
@NoArgsConstructor
public class FoodContribution {

    @Id
    @GeneratedValue
    private Long id;

    private String type;

    @Column(name = "contributor_name")
    private String contributorName;

    private double quantity;

    private String quantityMark;

    private String description;

    private boolean isMainDish;

    public FoodContribution(String type, String contributorName, double quantity, String quantityMark, String description, boolean isMainDish) {
        this.type = type;
        this.contributorName = contributorName;
        this.quantity = quantity;
        this.quantityMark = quantityMark;
        this.description = description;
        this.isMainDish = isMainDish;
    }

    public Long getId() {
        return id;
    }

    public void setIf(Long id) {
        this.id = id;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getContributor_name() {
        return contributorName;
    }

    public void setContributor_name(String contributorName) {
        this.contributorName = contributorName;
    }

    public double getQuantity() {
        return quantity;
    }

    public void setQuantity(double quantity) {
        this.quantity = quantity;
    }

    public String getQuantityMark() {
        return quantityMark;
    }

    public void setQuantityMark(String quantityMark) {
        this.quantityMark = quantityMark;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public boolean isMainDish() {
        return isMainDish;
    }

    public void setMainDish(boolean mainDish) {
        isMainDish = mainDish;
    }
}

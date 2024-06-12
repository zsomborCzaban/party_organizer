package dev.czaban.party_website.models.food;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "food_contributions")
@AllArgsConstructor
@NoArgsConstructor
public class FoodContribution {

    @Id
    private ObjectId objectId;

    private String type;

    private String contributor_name;

    private double quantity;

    private String quantityMark;

    private String description;

    private boolean isMainDish;

    public FoodContribution(String type, String contributor_name, double quantity, String quantityMark, String description, boolean isMainDish) {
        this.type = type;
        this.contributor_name = contributor_name;
        this.quantity = quantity;
        this.quantityMark = quantityMark;
        this.description = description;
        this.isMainDish = isMainDish;
    }

    public ObjectId getObjectId() {
        return objectId;
    }

    public void setObjectId(ObjectId objectId) {
        this.objectId = objectId;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getContributor_name() {
        return contributor_name;
    }

    public void setContributor_name(String contributor_name) {
        this.contributor_name = contributor_name;
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

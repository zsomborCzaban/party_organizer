package dev.czaban.party_website.models.drink;


import jakarta.persistence.Column;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.Field;

@Document(collection = "drink_types")
@AllArgsConstructor
@NoArgsConstructor
public class DrinkType {

    @Id
    private ObjectId id;
    private String type;
    private String name;
    @Field("quantity_mark")
    private String quantityMark;

    public DrinkType(String type, String name, String quantityMark) {
        this.type = type;
        this.name = name;
        this.quantityMark = quantityMark;
    }

    public ObjectId getId() {
        return id;
    }

    public void setId(ObjectId id) {
        this.id = id;
    }

    public String getDrinkType() {
        return type;
    }

    public void setDrinkType(String type) {
        this.type = type;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getQuantityMark() {
        return quantityMark;
    }

    public void setQuantityMark(String quantityMark) {
        this.quantityMark = quantityMark;
    }
}

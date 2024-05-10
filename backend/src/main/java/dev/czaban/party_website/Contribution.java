package dev.czaban.party_website;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "contributions")
@Data
@AllArgsConstructor
@NoArgsConstructor
public class Contribution {

    @Id
    private ObjectId id;

    private String type; //maybe enum
    public String getType() {
        return type;
    }

    private double quantity;

    private String contributor_name;

    private String description;

    public Contribution(String type, double quantity, String description, String contributor_name) {
        this.type = type;
        this.quantity = quantity;
        this.description = description;
        this.contributor_name = contributor_name;
    }

}

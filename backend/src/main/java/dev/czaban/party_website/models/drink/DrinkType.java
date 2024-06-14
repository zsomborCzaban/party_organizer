package dev.czaban.party_website.models.drink;


import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@Entity
@AllArgsConstructor
@NoArgsConstructor
public class DrinkType {

    @Id
    @GeneratedValue
    private Long id;
    private String type;
    private String name;
    @Column(name = "quantity_mark")
    private String quantityMark;

    public DrinkType(String type, String name, String quantityMark) {
        this.type = type;
        this.name = name;
        this.quantityMark = quantityMark;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
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

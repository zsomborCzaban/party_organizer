package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.food.FoodType;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodTypeRepository  extends CrudRepository<FoodType, Long> {
}

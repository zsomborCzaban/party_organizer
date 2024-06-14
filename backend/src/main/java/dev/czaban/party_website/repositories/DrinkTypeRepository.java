package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.drink.DrinkType;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DrinkTypeRepository  extends CrudRepository<DrinkType, Long> {
}

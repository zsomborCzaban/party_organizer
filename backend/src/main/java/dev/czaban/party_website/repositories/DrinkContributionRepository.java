package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.drink.DrinkContribution;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DrinkContributionRepository extends CrudRepository<DrinkContribution, Long> {
}

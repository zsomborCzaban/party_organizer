package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.drink.DrinkContribution;
import dev.czaban.party_website.models.drink.DrinkType;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface DrinkContributionRepository extends CrudRepository<DrinkContribution, Long> {
    @Override
    List<DrinkContribution> findAll(); //could also use a jparepository
}

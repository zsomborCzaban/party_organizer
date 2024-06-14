package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.food.FoodContribution;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodContributionRepository extends CrudRepository<FoodContribution, Long> {
}

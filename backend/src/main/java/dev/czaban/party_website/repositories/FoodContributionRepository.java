package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.food.FoodContribution;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface FoodContributionRepository extends MongoRepository<FoodContribution, ObjectId> {
}

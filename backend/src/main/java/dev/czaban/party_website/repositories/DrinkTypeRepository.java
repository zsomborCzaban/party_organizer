package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.drink.DrinkType;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface DrinkTypeRepository  extends MongoRepository<DrinkType, ObjectId> {
}

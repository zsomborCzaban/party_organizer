package dev.czaban.party_website;

import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface ContributionRepository extends MongoRepository<Contribution, ObjectId> {
}

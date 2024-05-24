package dev.czaban.party_website.repositories;

import dev.czaban.party_website.models.User;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface UserRepository extends MongoRepository<User, ObjectId> {
    Optional<User> findByUsername(String username); //this is a query method. derrived by the name (select * from users  where username = param)
}

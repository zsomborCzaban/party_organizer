package dev.czaban.party_website.models;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;
import org.bson.types.ObjectId;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.annotation.Id;

@Document(collection = "users")
@AllArgsConstructor
@NoArgsConstructor
public class User {

    @Id // @generatedValue is not available for mongoDB
    private ObjectId id;
    private String username;
    private String password;
    private String contributorName;
    private String roles;

    public User(String username, String password, String contributorName, String roles) {
        this.username = username;
        this.password = password;
        this.contributorName = contributorName;
        this.roles = roles;
    }

    public ObjectId getId() {
        return id;
    }

    public void setId(ObjectId id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getContributorName() {
        return contributorName;
    }

    public void setContributorName(String contributorName) {
        this.contributorName = contributorName;
    }

    public String getRoles() {
        return roles;
    }

    public void setRoles(String roles) {
        this.roles = roles;
    }

    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", username='" + username + '\'' +
                ", password='" + password + '\'' +
                ", contributorName='" + contributorName + '\'' +
                ", roles='" + roles + '\'' +
                '}';
    }
}

package dev.czaban.party_website.models;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;


@Entity
@AllArgsConstructor
@NoArgsConstructor
public class User {

    @Id // @generatedValue is not available for mongoDB
    @GeneratedValue
    private Long id;
    private String username;
    private String password;
    @Column(name = "contributor_name")
    private String contributorName;
    private String roles;

    public User(String username, String password, String contributorName, String roles) {
        this.username = username;
        this.password = password;
        this.contributorName = contributorName;
        this.roles = roles;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
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

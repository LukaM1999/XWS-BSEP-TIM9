package com.agent.agent.repository;

import com.agent.agent.model.RegisteredUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

public interface RegisteredUserRepository extends JpaRepository<RegisteredUser, String> {

    RegisteredUser findByUsername(String username);

    @Query("UPDATE RegisteredUser u SET u.firstName = ?1, u.lastName = ?2, u.email = ?3, u.phone = ?4, u.address = ?5, u.city = ?6, u.country = ?7, u.dislinktUsername = ?8, u.dislinktToken = ?9 WHERE u.username = ?10")
    RegisteredUser updateUser(String firstName, String lastName, String email, String phoneNumber, String address,
                              String city, String country, String dislinktUsername, String dislinktToken, String username);
}

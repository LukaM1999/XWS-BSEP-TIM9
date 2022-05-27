package com.agent.agent.repository;

import com.agent.agent.model.RegisteredUser;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

public interface RegisteredUserRepository extends JpaRepository<RegisteredUser, String> {

    RegisteredUser findByUsername(String username);

//    RegisteredUser updateUser(String firstName, String lastName, String email, String phoneNumber, String address,
//                              String city, String country, String dislinktToken);
    //Query to update firstName, lastName, email, phoneNumber, address, city, country, dislinktToken for a user with username
    @Query("UPDATE RegisteredUser u SET u.firstName = ?1, u.lastName = ?2, u.email = ?3, u.phone = ?4, u.address = ?5, u.city = ?6, u.country = ?7, u.dislinktToken = ?8 WHERE u.username = ?9")
    RegisteredUser updateUser(String firstName, String lastName, String email, String phoneNumber, String address,
                              String city, String country, String dislinktToken, String username);
}

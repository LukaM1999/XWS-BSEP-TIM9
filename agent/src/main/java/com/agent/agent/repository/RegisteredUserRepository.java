package com.agent.agent.repository;

import com.agent.agent.model.RegisteredUser;
import org.springframework.data.jpa.repository.JpaRepository;

public interface RegisteredUserRepository extends JpaRepository<RegisteredUser, String> {

    RegisteredUser findByUsername(String username);
}

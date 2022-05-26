package com.agent.agent.service;

import com.agent.agent.model.RegisteredUser;
import org.springframework.security.core.userdetails.UserDetailsService;

public interface RegisteredUserService extends UserDetailsService {

    RegisteredUser registerUser(String username, String password);
}

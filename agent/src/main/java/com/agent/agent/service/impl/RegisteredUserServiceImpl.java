package com.agent.agent.service.impl;

import com.agent.agent.model.RegisteredUser;
import com.agent.agent.repository.RegisteredUserRepository;
import com.agent.agent.service.RegisteredUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
public class RegisteredUserServiceImpl implements RegisteredUserService {

    @Autowired
    private RegisteredUserRepository userRepository;

    @Override
    public RegisteredUser registerUser(String username, String password) {
        RegisteredUser user = new RegisteredUser(username, password);
        return userRepository.save(user);
    }

    @Override
    public RegisteredUser getUserByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public RegisteredUser updateUser(RegisteredUser user) {
        return userRepository.save(user);
    }

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        return userRepository.findByUsername(username);
    }
}

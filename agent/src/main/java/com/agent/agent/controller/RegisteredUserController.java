package com.agent.agent.controller;

import com.agent.agent.dto.ProfileDTO;
import com.agent.agent.model.RegisteredUser;
import com.agent.agent.service.RegisteredUserService;
import com.agent.agent.util.ProfileDTOMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/user")
public class RegisteredUserController {

    @Autowired
    private RegisteredUserService registeredUserService;

    @Autowired
    private ProfileDTOMapper profileDTOMapper;

    @PatchMapping("")
    @PreAuthorize("hasAnyAuthority('USER', 'COMPANY_OWNER')")
    public ResponseEntity<RegisteredUser> updateUser(@RequestBody ProfileDTO profileDTO) {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        RegisteredUser loggedIn = (RegisteredUser) auth.getPrincipal();
        RegisteredUser user = registeredUserService.getUserByUsername(profileDTO.getUsername());
        if(user == null || !user.getUsername().equals(loggedIn.getUsername())){
            return new ResponseEntity<>(HttpStatus.FORBIDDEN);
        }
        user = profileDTOMapper.updateWithNullAsNoChange(profileDTO, user);
        return ResponseEntity.ok(registeredUserService.updateUser(user));
    }
}

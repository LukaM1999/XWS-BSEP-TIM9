package com.agent.agent.controller;

import com.agent.agent.dto.JwtDTO;
import com.agent.agent.dto.LoginDTO;
import com.agent.agent.dto.RegistrationDTO;
import com.agent.agent.model.RegisteredUser;
import com.agent.agent.service.RegisteredUserService;
import com.agent.agent.util.TokenUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.*;

//Kontroler zaduzen za autentifikaciju korisnika
@RestController
@RequestMapping(value = "/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthController {

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private RegisteredUserService userService;

    // Prvi endpoint koji pogadja korisnik kada se loguje.
    // Tada zna samo svoje korisnicko ime i lozinku i to prosledjuje na backend.
    @PostMapping("/login")
    public ResponseEntity<JwtDTO> createAuthenticationToken(
            @RequestBody LoginDTO loginDto) {

        // Ukoliko kredencijali nisu ispravni, logovanje nece biti uspesno, desice se
        // AuthenticationException

        Authentication authentication =
                authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(loginDto.getUsername(), loginDto.getPassword()));

        // Ukoliko je autentifikacija uspesna, ubaci korisnika u trenutni security
        // kontekst
        SecurityContextHolder.getContext().setAuthentication(authentication);

        // Kreiraj token za tog korisnika
        RegisteredUser user = (RegisteredUser) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUsername());
        int expiresIn = tokenUtils.getExpiredIn();

        //user.setPassword("");
        // Vrati token kao odgovor na uspesnu autentifikaciju
        return ResponseEntity.ok(new JwtDTO(user, jwt, expiresIn));
    }

    // Endpoint za registraciju novog korisnika
    @PostMapping("/signup")
    public ResponseEntity<RegisteredUser> addUser(@RequestBody RegistrationDTO registrationDTO) {
        UserDetails existUser = this.userService.loadUserByUsername(registrationDTO.getUsername());
        if (existUser != null) {
            throw new NullPointerException("Username already exists: " + registrationDTO.getUsername());
        }
        RegisteredUser user = this.userService.registerUser(registrationDTO.getUsername(), registrationDTO.getPassword());
        return new ResponseEntity<>(user, HttpStatus.CREATED);
    }

//    @GetMapping("/verify")
//    public ResponseEntity<Object> addUser(@RequestParam String token) throws IllegalAccessException, URISyntaxException {
//        Customer customer = customerService.findByToken(token);
//        if (customer == null) throw new NullPointerException("Username with this token doesn't exist!");
//        if (customer.isEnabled()) throw new IllegalAccessException("Account is already verified!");
//        customerService.verifyCustomer(customer.getUsername());
//        HttpHeaders httpHeaders = new HttpHeaders();
//        httpHeaders.setLocation(new URI("http://localhost:7000/login"));
//        return new ResponseEntity<>(httpHeaders, HttpStatus.SEE_OTHER);
//    }
//
//    @PostMapping("/confirmPassword")
//    public boolean confirmPassword(@RequestBody LoginDTO loginDto) {
//        return userService.isPasswordValid(loginDto);
//    }
}

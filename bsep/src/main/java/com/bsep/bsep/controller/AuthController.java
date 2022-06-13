package com.bsep.bsep.controller;

import com.bsep.bsep.data.Account;
import com.bsep.bsep.dto.JwtDTO;
import com.bsep.bsep.dto.LoginDTO;
import com.bsep.bsep.service.AccountService;
import com.bsep.bsep.util.TokenUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

//Kontroler zaduzen za autentifikaciju korisnika
@RestController
@RequestMapping(value = "/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthController {

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private AccountService userService;

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
        Account user = (Account) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUsername());
        int expiresIn = tokenUtils.getExpiredIn();

        //user.setPassword("");
        // Vrati token kao odgovor na uspesnu autentifikaciju
        return ResponseEntity.ok(new JwtDTO(user, jwt, expiresIn));
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

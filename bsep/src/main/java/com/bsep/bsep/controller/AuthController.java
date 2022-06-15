package com.bsep.bsep.controller;

import com.bsep.bsep.data.Account;
import com.bsep.bsep.dto.JwtDTO;
import com.bsep.bsep.dto.LoginDTO;
import com.bsep.bsep.service.AccountService;
import com.bsep.bsep.util.TokenUtils;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.message.StringMapMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
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

    private final Logger logger = LogManager.getLogger("XML_ROLLING_FILE_APPENDER");

    // Prvi endpoint koji pogadja korisnik kada se loguje.
    // Tada zna samo svoje korisnicko ime i lozinku i to prosledjuje na backend.
    @PostMapping("/login")
    public ResponseEntity<JwtDTO> createAuthenticationToken(
            @RequestBody LoginDTO loginDto) {

        // Ukoliko kredencijali nisu ispravni, logovanje nece biti uspesno, desice se
        // AuthenticationException

        StringMapMessage mapMessage = new StringMapMessage();
        mapMessage.put("msg", "Login attempt");

        logger.info(mapMessage);

        Authentication authentication;
        try {
            authentication = authenticationManager.authenticate(new UsernamePasswordAuthenticationToken(loginDto.getUsername(), loginDto.getPassword()));
        } catch (AuthenticationException e) {
            mapMessage.put("msg", "Authentication failed: " + e.getMessage());
            logger.error(mapMessage);
            throw new RuntimeException(e);
        }

        // Ukoliko je autentifikacija uspesna, ubaci korisnika u trenutni security
        // kontekst
        SecurityContextHolder.getContext().setAuthentication(authentication);

        // Kreiraj token za tog korisnika
        Account user = (Account) authentication.getPrincipal();
        String jwt = tokenUtils.generateToken(user.getUsername());
        int expiresIn = tokenUtils.getExpiredIn();

        //user.setPassword("");
        // Vrati token kao odgovor na uspesnu autentifikaciju
        mapMessage.put("msg", "Login successful");
        mapMessage.put("username", loginDto.getUsername());
        logger.info(mapMessage);

        return ResponseEntity.ok(new JwtDTO(user, jwt, expiresIn));
    }
}

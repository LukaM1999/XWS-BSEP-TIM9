package com.bsep.bsep.security;

import com.bsep.bsep.service.AccountService;
import com.bsep.bsep.util.TokenUtils;
import io.jsonwebtoken.ExpiredJwtException;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import org.apache.logging.log4j.message.StringMapMessage;
import org.apache.logging.log4j.util.StringMap;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.web.filter.OncePerRequestFilter;

import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

// Filter koji ce presretati SVAKI zahtev klijenta ka serveru
// (sem nad putanjama navedenim u WebSecurityConfig.configure(WebSecurity web))
// Filter proverava da li JWT token postoji u Authorization header-u u zahtevu koji stize od klijenta
// Ukoliko token postoji, proverava se da li je validan. Ukoliko je sve u redu, postavlja se autentifikacija
// u SecurityContext holder kako bi podaci o korisniku bili dostupni u ostalim delovima aplikacije gde su neophodni
public class TokenAuthFilter extends OncePerRequestFilter {

    private TokenUtils tokenUtils;

    private UserDetailsService userDetailsService;

    private final Logger logger = LogManager.getLogger("XML_ROLLING_FILE_APPENDER");

    public TokenAuthFilter(TokenUtils tokenHelper, AccountService userDetailsService) {
        this.tokenUtils = tokenHelper;
        this.userDetailsService = userDetailsService;
    }

    @Override
    public void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain)
            throws IOException, ServletException {

        StringMapMessage mapMessage = new StringMapMessage();
        mapMessage.put("msg", "Authenticating request");
        mapMessage.put("method", request.getMethod());
        mapMessage.put("url", request.getRequestURL().toString());
        String origin = request.getHeader("Origin");
        if (origin != null) {
            mapMessage.put("origin", origin);
        }
        String userAgent = request.getHeader("User-Agent");
        if (userAgent != null) {
            mapMessage.put("userAgent", userAgent);
        }

        logger.info(mapMessage);

        String username;

        // 1. Preuzimanje JWT tokena iz zahteva
        String authToken = tokenUtils.getToken(request);

        if(authToken == null) {
            mapMessage.put("msg", "No token found");
            logger.warn(mapMessage);
            chain.doFilter(request, response);
            return;
        }

        try {

            // 2. Citanje korisnickog imena iz tokena
            username = tokenUtils.getUsernameFromToken(authToken);

            if(username == null) {
                mapMessage.put("msg", "No username found");
                logger.warn(mapMessage);
                chain.doFilter(request, response);
                return;
            }

            // 3. Preuzimanje korisnika na osnovu username-a
            UserDetails userDetails = userDetailsService.loadUserByUsername(username);

            if(userDetails == null) {
                mapMessage.put("username", username);
                mapMessage.put("msg", "No user found");
                logger.warn(mapMessage);
                chain.doFilter(request, response);
                return;
            }

            // 4. Provera da li je prosledjeni token validan
            if (tokenUtils.validateToken(authToken, userDetails)) {

                // 5. Kreiraj autentifikaciju
                TokenBasedAuth authentication = new TokenBasedAuth(userDetails);
                authentication.setToken(authToken);
                SecurityContextHolder.getContext().setAuthentication(authentication);

            }
            else {
                mapMessage.put("msg", "Invalid token");
                logger.warn(mapMessage);
            }

        } catch (Exception e) {
            mapMessage.put("msg", e.getMessage());
            logger.error(mapMessage);
        }

        // prosledi request dalje u sledeci filter
        chain.doFilter(request, response);
    }

}

package com.agent.agent.dto;

import com.agent.agent.model.RegisteredUser;
import lombok.Getter;
import lombok.Setter;

// DTO koji enkapsulira generisani JWT i njegovo trajanje koji se vracaju klijentu
public class JwtDTO {

    @Getter
    @Setter
    private RegisteredUser user;
    @Getter
    @Setter
    private String accessToken;
    @Getter
    @Setter
    private Long expiresIn;

    public JwtDTO() {
        this.user = null;
        this.accessToken = null;
        this.expiresIn = null;
    }

    public JwtDTO(RegisteredUser user, String accessToken, long expiresIn) {
        this.user = user;
        this.accessToken = accessToken;
        this.expiresIn = expiresIn;
    }

}

package com.bsep.bsep.dto;

import com.bsep.bsep.data.Account;
import lombok.Getter;
import lombok.Setter;

// DTO koji enkapsulira generisani JWT i njegovo trajanje koji se vracaju klijentu
public class JwtDTO {

    @Getter
    @Setter
    private Account user;
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

    public JwtDTO(Account user, String accessToken, long expiresIn) {
        this.user = user;
        this.accessToken = accessToken;
        this.expiresIn = expiresIn;
    }

}

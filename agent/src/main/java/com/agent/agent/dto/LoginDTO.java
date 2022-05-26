package com.agent.agent.dto;

import lombok.Getter;
import lombok.Setter;

public class LoginDTO {
    @Getter
    @Setter
    private String username;
    @Getter
    @Setter
    private String password;

    public LoginDTO() {
        super();
    }

    public LoginDTO(String username, String password) {
        this.setUsername(username);
        this.setPassword(password);
    }
}

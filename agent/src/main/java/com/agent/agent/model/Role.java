package com.agent.agent.model;

import lombok.Getter;
import lombok.Setter;
import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;
import java.io.Serializable;
import java.util.List;

// POJO koji implementira Spring Security GrantedAuthority kojim se mogu definisati role u aplikaciji

@Entity
@Table(name = "role")
public class Role implements GrantedAuthority, Serializable {

    private static final long serialVersionUID = 1L;

    @Id
    @Column(name="role_name")
    private String roleName;

    //@JsonIgnore
    @Override
    public String getAuthority() {
        return roleName;
    }

    public void setName(String name) {
        this.roleName = name;
    }

    public Role(String name) {
        this.roleName = name;
    }

    public Role(){}
}


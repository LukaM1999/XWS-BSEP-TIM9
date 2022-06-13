package com.bsep.bsep.data;

import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table
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

    public Role(String roleName) {
    	this.roleName = roleName;
    }

    public static long getSerialVersionUID() {
        return serialVersionUID;
    }

    public String getRoleName() {
        return roleName;
    }

    public void setRoleName(String roleName) {
        this.roleName = roleName;
    }

    public Role() {
    }
}

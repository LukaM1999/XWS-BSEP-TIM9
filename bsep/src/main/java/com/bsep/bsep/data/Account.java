package com.bsep.bsep.data;

import com.fasterxml.jackson.annotation.JsonIdentityInfo;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.ObjectIdGenerators;
import lombok.Getter;
import lombok.Setter;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import javax.persistence.*;
import java.sql.Timestamp;
import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

@Entity
@Table(name = "account")
public class Account implements UserDetails {

    private static final long serialVersionUID = 1L;

    @Id
    @Column
    private String username;

    @Column
    @JsonProperty(access = JsonProperty.Access.WRITE_ONLY)
    private String password;

    @ManyToOne(cascade = CascadeType.ALL)
    @JoinColumn(name="roleName")
    private Role role;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return new ArrayList<Role>(List.of(this.role));
    }

    @Column(name = "last_password_reset_date")
    private Timestamp lastPasswordResetDate;

    @Column(name = "enabled")
    private boolean enabled;

    @Override
    public String getPassword() {
        return password;
    }

    @Override
    public String getUsername() {
        return username;
    }

    @Override
    @JsonIgnore
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    @JsonIgnore
    public boolean isAccountNonLocked() {
        return true;
    }

    @Override
    @JsonIgnore
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return enabled;
    }

    public Account(){
        super();
    }

    public Account(String username, String password, String role){
        this.username = username;
        this.password = password;
        this.role = new Role(role);
    }

    public static long getSerialVersionUID() {
        return serialVersionUID;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public Role getRole() {
        return role;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    public Timestamp getLastPasswordResetDate() {
        return lastPasswordResetDate;
    }

    public void setLastPasswordResetDate(Timestamp lastPasswordResetDate) {
        this.lastPasswordResetDate = lastPasswordResetDate;
    }
}

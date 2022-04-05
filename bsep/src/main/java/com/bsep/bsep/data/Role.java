package com.bsep.bsep.data;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;
import java.io.Serializable;


@Entity
@Table
public class Role implements GrantedAuthority, Serializable {

    private static final long serialVersionUID = 1L;

    @Column(name="roleName")
    private String roleName;

    @Id
    @Column(name="id")
    @SequenceGenerator(name = "role_id_gen", sequenceName = "role_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "role_id_gen")
    Long id;

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

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Role() {
    }


}

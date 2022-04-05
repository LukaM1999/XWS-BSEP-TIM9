package com.bsep.bsep.data;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
@Entity
public class UserCertificate {

    @Id
    @Column
    private String certificateSerialNumber;
    @Column
    private String email;
    @Column
    private boolean revoked;
}

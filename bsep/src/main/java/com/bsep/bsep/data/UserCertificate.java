package com.bsep.bsep.data;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
@Entity
public class UserCertificate {

    @Id
    @Column
    @SequenceGenerator(name = "crt_id_gen", sequenceName = "crt_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "crt_id_gen")
    private Long certificateSerialNumber;
    @Column
    private String username;
    @Column
    private boolean revoked;
}

package com.bsep.bsep.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.bouncycastle.asn1.x500.X500Name;

import java.security.PrivateKey;
import java.security.PublicKey;
import java.util.Date;

@NoArgsConstructor
@Getter
@Setter
public class CertificateDTO {

    private String commonNameSubject;
    private String nameSubject;
    private String surnameSubject;
    private String emailSubject;
    private String countrySubject;
    private Date startDate;
    private Date endDate;

    private String commonNameIssuer;
    private String nameIssuer;
    private String surnameIssuer;
    private String emailIssuer;
    private String countryIssuer;
    private String serialNumberIssuer;

    private String authority;
}

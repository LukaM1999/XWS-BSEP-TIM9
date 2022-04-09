package com.bsep.bsep.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.bouncycastle.asn1.x500.X500Name;

import java.security.PrivateKey;
import java.security.PublicKey;
import java.util.Date;
import java.util.List;

@NoArgsConstructor
@Getter
@Setter
public class CertificateDTO {

    private String commonNameSubject;
    private String nameSubject;
    private String surnameSubject;
    private String usernameSubject;
    private String countrySubject;
    private String serialNumberSubject;
    private Date startDate;
    private Date endDate;
    private String authoritySubject;

    private String commonNameIssuer;
    private String nameIssuer;
    private String surnameIssuer;
    private String usernameIssuer;
    private String countryIssuer;
    private String serialNumberIssuer;
    private String authorityIssuer;

    private List<Integer> keyUsages;

    public CertificateDTO(String authorityIssuer, String authoritySubject, List<Integer> keyUsages, String serialNumberIssuer){
     this.authorityIssuer = authorityIssuer;
     this.authoritySubject = authoritySubject;
     this.keyUsages = keyUsages;
     this.serialNumberIssuer = serialNumberIssuer;
    }
}

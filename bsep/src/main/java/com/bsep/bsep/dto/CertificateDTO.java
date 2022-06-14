package com.bsep.bsep.dto;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.bouncycastle.asn1.x500.X500Name;

import javax.validation.constraints.NotNull;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.util.Date;
import java.util.List;

@NoArgsConstructor
@Getter
@Setter
public class CertificateDTO {

    @NotNull(message = "CommonName is mandatory")
    private String commonNameSubject;
    @NotNull(message = "Name is mandatory")
    private String nameSubject;
    @NotNull(message = "Name is mandatory")
    private String surnameSubject;
    @NotNull(message = "Name is mandatory")
    private String usernameSubject;
    @NotNull(message = "Name is mandatory")
    private String countrySubject;
    private String serialNumberSubject;
    @JsonFormat(pattern="yyyy-MM-dd")
    private Date startDate;
    @JsonFormat(pattern="yyyy-MM-dd")
    private Date endDate;
    @NotNull(message = "Authority is mandatory")
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

    public CertificateDTO(String authorityIssuer, String authoritySubject, List<Integer> keyUsages, String serialNumberIssuer, String commonNameSubject){
        this.authorityIssuer = authorityIssuer;
        this.authoritySubject = authoritySubject;
        this.keyUsages = keyUsages;
        this.serialNumberIssuer = serialNumberIssuer;
        this.commonNameSubject = commonNameSubject;
    }
}

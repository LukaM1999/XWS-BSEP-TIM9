package com.bsep.bsep.service.impl;

import com.bsep.bsep.data.UserCertificate;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.keystores.KeyStoreReader;
import com.bsep.bsep.repository.UserCertificateRepository;
import com.bsep.bsep.service.UserCertificateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.couchbase.CouchbaseProperties;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Service;

import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.util.ArrayList;
import java.util.List;

@Service
public class UserCertificateServiceImpl implements UserCertificateService {

    @Autowired
    private Environment env;
    @Autowired
    private UserCertificateRepository userCertificateRepository;
    @Autowired
    private CertificateService certificateService;
    private final KeyStoreReader keyStoreReader = new KeyStoreReader();

    public List<CertificateDTO> getUserCertificates(String username) throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        List<X509Certificate> certificates = new ArrayList<>();
        List<UserCertificate> userCertificates = userCertificateRepository.findByUsername(username);
        for(UserCertificate uc: userCertificates){
            X509Certificate crt;
            if(uc.isRevoked()) continue;
            crt = (X509Certificate) keyStoreReader.readCertificate(env.getProperty("keystore.path") + "root.jks", "12345", uc.getCertificateSerialNumber().toString());
            if(crt == null)
                crt = (X509Certificate) keyStoreReader.readCertificate(env.getProperty("keystore.path") + "ca.jks", "12345", uc.getCertificateSerialNumber().toString());
            if(crt == null)
                crt = (X509Certificate) keyStoreReader.readCertificate(env.getProperty("keystore.path") + "endEntity.jks", "12345", uc.getCertificateSerialNumber().toString());
            if(crt != null)
                certificates.add(crt);
        }
        if(certificates.size() == 0) return null;
        List<CertificateDTO> certificateDTOs = certificateService.certificateToDTO(certificates);
        CertificateDTO certificateDTO = new CertificateDTO();
        certificateDTO.setSerialNumberSubject(certificateDTOs.get(0).getSerialNumberSubject());
        certificateDTO.setAuthoritySubject("ca");
        List<CertificateDTO> issuedCertificates = certificateService.getIssuedCertificates(certificateDTO);
        if(issuedCertificates.size() == 0) {
            certificateDTO.setAuthoritySubject("root");
            issuedCertificates = certificateService.getIssuedCertificates(certificateDTO);
        }
        if(issuedCertificates.size() == 0) return certificateDTOs;
        certificateDTOs.addAll(issuedCertificates);
        return certificateDTOs;
    }

}

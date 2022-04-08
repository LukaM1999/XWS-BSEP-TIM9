package com.bsep.bsep.controller;

import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.service.impl.CertificateService;
import org.bouncycastle.crypto.tls.CertificateType;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.util.List;

@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private CertificateService certificateService;

    @PostMapping("/createCertificate")
    public void createCertificate(@RequestBody CertificateDTO certificateDTO){
        certificateService.createCertificate(certificateDTO);
    }

    @GetMapping("/getEndCertificates")
    public List<CertificateDTO> getAllEndUserCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.certificateToDTO(certificateService.getAllEndUserCertificates());
    }

    @GetMapping("/getRootCertificates")
    public List<CertificateDTO> getAllRootCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.certificateToDTO(certificateService.getAllRootCertificates());
    }

    @GetMapping("/getCACertificates")
    public List<CertificateDTO> getAllCaCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.certificateToDTO(certificateService.getAllCACertificates());
    }

    @PostMapping("/getCertificateChain")
    public List<CertificateDTO> getCertificateChain(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, NoSuchAlgorithmException, ParseException,
            InvalidKeyException, NoSuchProviderException {
        return certificateService.getCertificateChain(certificateDTO);
    }

    @PostMapping("/revokeCertificate")
    public boolean revokeCertificate(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, NoSuchAlgorithmException, ParseException,
            InvalidKeyException, NoSuchProviderException {
        return certificateService.revokeCertificate(certificateDTO);
    }

}

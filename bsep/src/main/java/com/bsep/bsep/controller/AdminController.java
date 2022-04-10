package com.bsep.bsep.controller;

import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.service.impl.CertificateService;
import org.bouncycastle.crypto.tls.CertificateType;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.InputStreamResource;
import org.springframework.core.io.Resource;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.security.InvalidKeyException;
import java.security.KeyStoreException;
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
        return certificateService.certificateToDTO(certificateService.getAllActiveEndUserCertificates());
    }

    @GetMapping("/getRootCertificates")
    public List<CertificateDTO> getAllRootCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.certificateToDTO(certificateService.getAllActiveRootCertificates());
    }

    @GetMapping("/getCACertificates")
    public List<CertificateDTO> getAllCaCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.certificateToDTO(certificateService.getAllActiveCACertificates());
    }

    @PostMapping("/getCertificateChain")
    public List<CertificateDTO> getCertificateChain(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, NoSuchAlgorithmException, ParseException,
            InvalidKeyException, NoSuchProviderException {
        return certificateService.getCertificateChain(certificateDTO);
    }

    @PostMapping("/revokeCertificate")
    public boolean revokeCertificate(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.revokeCertificate(certificateDTO);
    }

    @PostMapping("/getIssuedCertificates")
    public List<CertificateDTO> getIssuedCertificates(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.getIssuedCertificates(certificateDTO);
    }

    @PostMapping("/downloadCertificate")
    public ResponseEntity<Resource> downloadCertificate(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException, KeyStoreException, IOException {
        certificateService.extractCertificate(certificateDTO);
        File file = new File(certificateDTO.getSerialNumberSubject() + ".crt");
        InputStreamResource resource = new InputStreamResource(new FileInputStream(file));

        return ResponseEntity.ok()
                .contentLength(file.length())
                .contentType(MediaType.APPLICATION_OCTET_STREAM)
                .body(resource);
    }

}

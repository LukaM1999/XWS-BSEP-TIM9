package com.bsep.bsep.controller;

import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.service.impl.CertificateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.InputStreamResource;
import org.springframework.core.io.Resource;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private CertificateService certificateService;

    @PostMapping("/createCertificate")
    @PreAuthorize("!hasAuthority('endEntity')")
    public CertificateDTO createCertificate(@Valid @RequestBody CertificateDTO certificateDTO) {
        X509Certificate created = certificateService.createCertificate(certificateDTO);
        if(created == null) return null;
        return certificateService.certificateToDTO(new ArrayList<>(Collections.singletonList(created))).get(0);
    }

    @GetMapping("/getAllCertificates")
    public List<CertificateDTO> getAllCertificates() {
        return certificateService.getAllCertificates();
    }

    @GetMapping("/getEndCertificates")
    @PreAuthorize("hasAuthority('endEntity')")
    public List<CertificateDTO> getAllEndUserCertificates() {
        return certificateService.certificateToDTO(certificateService.getAllActiveEndUserCertificates());
    }

    @GetMapping("/getRootCertificates")
    @PreAuthorize("hasAuthority('admin')")
    public List<CertificateDTO> getAllRootCertificates() {
        return certificateService.certificateToDTO(certificateService.getAllActiveRootCertificates());
    }

    @GetMapping("/getCACertificates")
    @PreAuthorize("hasAuthority('ca')")
    public List<CertificateDTO> getAllCaCertificates() {
        return certificateService.certificateToDTO(certificateService.getAllActiveCACertificates());
    }

    @PostMapping("/getCertificateChain")
    public List<CertificateDTO> getCertificateChain(@RequestBody CertificateDTO certificateDTO)
            throws CertificateException, NoSuchAlgorithmException,
            InvalidKeyException, NoSuchProviderException {
        return certificateService.getCertificateChain(certificateDTO);
    }

    @PostMapping("/revokeCertificate")
    @PreAuthorize("hasAuthority('admin')")
    public boolean revokeCertificate(@RequestBody CertificateDTO certificateDTO) {
        return certificateService.revokeCertificate(certificateDTO);
    }

    @PostMapping("/getIssuedCertificates")
    public List<CertificateDTO> getIssuedCertificates(@RequestBody CertificateDTO certificateDTO) {
        return certificateService.getIssuedCertificates(certificateDTO);
    }

    @PostMapping("/downloadCertificate")
    public ResponseEntity<Resource> downloadCertificate(@RequestBody CertificateDTO certificateDTO) throws FileNotFoundException {
        certificateService.extractCertificate(certificateDTO);
        File file = new File(certificateDTO.getSerialNumberSubject() + ".crt");
        InputStreamResource resource = new InputStreamResource(new FileInputStream(file));
        file.deleteOnExit();

        return ResponseEntity.ok()
                .contentLength(file.length())
                .contentType(MediaType.APPLICATION_OCTET_STREAM)
                .body(resource);
    }

}

package com.bsep.bsep.controller;

import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.service.impl.CertificateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.security.cert.X509Certificate;

@RestController
@RequestMapping("/admin")
public class AdminController {

    @Autowired
    private CertificateService certificateService;

    @PostMapping("/createCertificate")
    public void createCertificate(@RequestBody CertificateDTO certificateDTO){
        certificateService.createCertificate(certificateDTO);
    }

}

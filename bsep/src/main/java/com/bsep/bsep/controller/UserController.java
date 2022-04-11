package com.bsep.bsep.controller;

import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.service.UserCertificateService;
import com.bsep.bsep.service.impl.CertificateService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.websocket.server.PathParam;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.text.ParseException;
import java.util.List;

@RestController
@RequestMapping("/user")
public class UserController {

    @Autowired
    private UserCertificateService certificateService;

    @GetMapping("/login/{username}")
    public List<CertificateDTO> login(@PathVariable String username) throws CertificateException, ParseException,
            NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return certificateService.getUserCertificates(username);
    }
}

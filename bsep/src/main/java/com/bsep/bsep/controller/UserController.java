package com.bsep.bsep.controller;

import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.dto.PasswordDTO;
import com.bsep.bsep.service.UserCertificateService;
import com.bsep.bsep.service.impl.CertificateService;
import com.enzoic.client.Enzoic;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.websocket.server.PathParam;
import java.io.IOException;
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

    @PostMapping("/isPasswordCompromised")
    public boolean isPasswordCompromised(@RequestBody PasswordDTO passwordDTO) throws IOException {
        Enzoic enzoic = new Enzoic("a1e69f4971a943ea832249204668ad36", "+7!QCvGc69$Sz#@n2w@egxZufxadszs?");
        return enzoic.CheckPassword(passwordDTO.getPassword());
    }
}

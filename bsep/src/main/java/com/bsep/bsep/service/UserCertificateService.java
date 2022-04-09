package com.bsep.bsep.service;

import com.bsep.bsep.dto.CertificateDTO;

import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.security.cert.CertificateException;
import java.text.ParseException;
import java.util.List;

public interface UserCertificateService {

    List<CertificateDTO> getUserCertificates(String username) throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException;
}

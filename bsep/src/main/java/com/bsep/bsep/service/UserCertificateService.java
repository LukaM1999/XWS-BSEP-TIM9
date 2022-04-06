package com.bsep.bsep.service;

import java.security.cert.X509Certificate;

public interface UserCertificateService {

    boolean isValid(String email, String certificateSerialNumber);
    //X509Certificate createCertificate(CertificateDTO certificateDTO);
}

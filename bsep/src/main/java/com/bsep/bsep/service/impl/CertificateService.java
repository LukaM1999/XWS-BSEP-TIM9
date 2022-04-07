package com.bsep.bsep.service.impl;

import com.bsep.bsep.certificates.CertificateGenerator;
import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.data.SubjectData;
import com.bsep.bsep.data.UserCertificate;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.keystores.KeyStoreWriter;
import com.bsep.bsep.repository.UserCertificateRepository;
import com.bsep.bsep.util.CertificateChainGenerator;
import org.bouncycastle.asn1.ASN1String;
import org.bouncycastle.asn1.x500.RDN;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.PropertySource;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Service;

import javax.security.auth.Subject;
import java.security.KeyPair;
import java.security.KeyStore;
import java.security.PrivateKey;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

@PropertySource("classpath:application.properties")
@Service
public class CertificateService {
    @Autowired
    private Environment env;

    final private CertificateChainGenerator certificateChainGenerator = new CertificateChainGenerator();


    @Autowired
    private UserCertificateRepository userCertificateRepository;

    public X509Certificate createCertificate(CertificateDTO certificateDTO) {
        CertificateGenerator certificateGenerator = new CertificateGenerator();
        KeyPair keyPairIssuer = new CertificateChainGenerator().generateKeyPair();
        KeyStoreWriter keystore = new KeyStoreWriter();

        char[] password = "12345".toCharArray();

        UserCertificate userCertificate = userCertificateRepository.save(new UserCertificate(null, certificateDTO.getEmailSubject(), false));

        SubjectData subjectData = generateSubjectData(certificateDTO, userCertificate.getCertificateSerialNumber().toString());
        IssuerData issuerData = generateIssuerData(certificateDTO, keyPairIssuer.getPrivate());

        X509Certificate x509Certificate = certificateGenerator.generateCertificate(subjectData, issuerData);

        keystore.loadKeyStore( env.getProperty("keystore.path") + certificateDTO.getAuthority() + ".jks", password);
        keystore.write(userCertificate.getCertificateSerialNumber().toString(), keyPairIssuer.getPrivate(), password, x509Certificate);
        keystore.saveKeyStore(env.getProperty("keystore.path") + certificateDTO.getAuthority() + ".jks", password);

        return x509Certificate;
    }

    private SubjectData generateSubjectData(CertificateDTO certificateDTO, String serialNumber) {
        KeyPair keyPairSubject = new CertificateChainGenerator().generateKeyPair();
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, certificateDTO.getCommonNameSubject());
        x500NameBuilder.addRDN(BCStyle.NAME, certificateDTO.getNameSubject());
        x500NameBuilder.addRDN(BCStyle.SURNAME, certificateDTO.getSurnameSubject());
        x500NameBuilder.addRDN(BCStyle.EmailAddress, certificateDTO.getEmailSubject());
        x500NameBuilder.addRDN(BCStyle.C, certificateDTO.getCountrySubject());
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, serialNumber);

        return new SubjectData(keyPairSubject.getPublic(), x500NameBuilder.build(), serialNumber, certificateDTO.getStartDate(), certificateDTO.getEndDate());

    }

    private IssuerData generateIssuerData(CertificateDTO certificateDTO, PrivateKey privateKey) {
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, certificateDTO.getCommonNameIssuer());
        x500NameBuilder.addRDN(BCStyle.NAME, certificateDTO.getNameIssuer());
        x500NameBuilder.addRDN(BCStyle.SURNAME, certificateDTO.getSurnameIssuer());
        x500NameBuilder.addRDN(BCStyle.EmailAddress, certificateDTO.getEmailIssuer());
        x500NameBuilder.addRDN(BCStyle.C, certificateDTO.getCountryIssuer());
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, certificateDTO.getSerialNumberIssuer());

        return new IssuerData(privateKey, x500NameBuilder.build());

    }
}

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
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.asn1.x500.style.IETFUtils;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.PropertySource;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Service;

import javax.security.auth.Subject;
import java.io.*;
import java.security.*;
import java.security.cert.*;
import java.security.cert.Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.Enumeration;
import java.util.List;

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

    public List<X509Certificate> getAllEndUserCertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./src/main/resources/keystores/endEntity.jks", "12345");
        for (X509Certificate certificate : certificates) {
            //if (_ocspListService.checkCertificateValidity(certificate)) {
            retList.add(certificate);
            //}
        }

        return retList;
    }

    public List<X509Certificate> getAllRootCertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./src/main/resources/keystores/root.jks", "12345");
        for (X509Certificate certificate : certificates) {
            //if (_ocspListService.checkCertificateValidity(certificate)) {
            retList.add(certificate);
            //}
        }

        return retList;
    }

    public List<X509Certificate> getAllCACertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./src/main/resources/keystores/ca.jks", "12345");
        for (X509Certificate certificate : certificates) {
            //if (_ocspListService.checkCertificateValidity(certificate)) {
            retList.add(certificate);
            //}
        }

        return retList;
    }

    private List<X509Certificate> readAllCertificate(String keyStoreFile, String keyStorePass) {
        List<String> aliases = readAliases(keyStoreFile, keyStorePass);
        List<X509Certificate> certificates = new ArrayList<>();
        for (String a : aliases){
            certificates.add(readCertificate(keyStoreFile, keyStorePass,a));
        }
        return certificates;
    }

    private List<String> readAliases(String keyStoreFile, String keyStorePass) {
        List<String> temp = new ArrayList();
        try{
            KeyStore ks;
            ks = KeyStore.getInstance("JKS", "SUN");
            BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
            ks.load(in, keyStorePass.toCharArray());
            //Enumeration interface generates a series of elements
            Enumeration<String> keys = ks.aliases();
            while(keys.hasMoreElements()){
                String key = keys.nextElement();
                temp.add(key);
            }
        } catch (KeyStoreException | NoSuchProviderException | NoSuchAlgorithmException | CertificateException | IOException e) {
            e.printStackTrace();
        }
        return temp;
    }

    private X509Certificate readCertificate(String keyStoreFile, String keyStorePass, String alias) {
        try {
            KeyStore ks = KeyStore.getInstance("JKS", "SUN");
            BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
            ks.load(in, keyStorePass.toCharArray());

            if(ks.isKeyEntry(alias)) {
                Certificate cert = ks.getCertificate(alias);
                CertificateFactory certFactory = CertificateFactory.getInstance("X.509");
                InputStream inp = new ByteArrayInputStream(cert.getEncoded());
                return (X509Certificate)certFactory.generateCertificate(inp);
            }
        } catch (KeyStoreException | NoSuchProviderException | NoSuchAlgorithmException | CertificateException | IOException e) {
            e.printStackTrace();
        }
        return null;
    }

    public List<CertificateDTO> certificateToDTO(List<X509Certificate> certificateList, String authority) throws CertificateEncodingException, ParseException {
        List<CertificateDTO> dto = new ArrayList<>();

        for(X509Certificate certificate : certificateList){
            CertificateDTO certDto = new CertificateDTO();
            JcaX509CertificateHolder certHolder = new JcaX509CertificateHolder((X509Certificate) certificate);
            X500Name subject = certHolder.getSubject();
            X500Name issuer = certHolder.getIssuer();
            String temp;
            RDN cn;
            if(subject.getRDNs(BCStyle.CN).length > 0) {
                cn = subject.getRDNs(BCStyle.CN)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setCommonNameSubject(temp);
            }
            if(subject.getRDNs(BCStyle.NAME).length > 0) {
                cn = subject.getRDNs(BCStyle.NAME)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setNameSubject(temp);
            }
            if(subject.getRDNs(BCStyle.SURNAME).length > 0) {
                cn = subject.getRDNs(BCStyle.SURNAME)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setSurnameSubject(temp);
            }
            if(subject.getRDNs(BCStyle.EmailAddress).length > 0) {
                cn = subject.getRDNs(BCStyle.EmailAddress)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setEmailSubject(temp);
            }
            if(subject.getRDNs(BCStyle.C).length > 0) {
                cn = subject.getRDNs(BCStyle.C)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setCountrySubject(temp);
            }
            if(subject.getRDNs(BCStyle.SERIALNUMBER).length > 0) {
                cn = subject.getRDNs(BCStyle.SERIALNUMBER)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setSerialNumberSubject(temp);
            }

            //--------------------------------------------------------------------------

            if(issuer.getRDNs(BCStyle.CN).length > 0) {
                cn = issuer.getRDNs(BCStyle.CN)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setCommonNameIssuer(temp);
            }
            if(issuer.getRDNs(BCStyle.NAME).length > 0) {
                cn = issuer.getRDNs(BCStyle.NAME)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setNameIssuer(temp);
            }
            if(issuer.getRDNs(BCStyle.SURNAME).length > 0) {
                cn = issuer.getRDNs(BCStyle.SURNAME)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setSurnameIssuer(temp);
            }
            if(issuer.getRDNs(BCStyle.EmailAddress).length > 0) {
                cn = issuer.getRDNs(BCStyle.EmailAddress)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setEmailIssuer(temp);
            }
            if(issuer.getRDNs(BCStyle.C).length > 0) {
                cn = issuer.getRDNs(BCStyle.C)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setCountryIssuer(temp);
            }
            if(issuer.getRDNs(BCStyle.SERIALNUMBER).length > 0) {
                cn = issuer.getRDNs(BCStyle.SERIALNUMBER)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setSerialNumberIssuer(temp);
            }
            certDto.setStartDate(certificate.getNotBefore());
            certDto.setEndDate(certificate.getNotAfter());
            certDto.setAuthority(authority);
            dto.add(certDto);
        }

        return dto;
    }
}

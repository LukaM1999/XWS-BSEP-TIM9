package com.bsep.bsep.service.impl;

import com.bsep.bsep.certificates.CertificateGenerator;
import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.data.SubjectData;
import com.bsep.bsep.keystores.KeyStoreWriter;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.springframework.stereotype.Service;

import java.security.*;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class CertificateService {

    public CertificateService() {
        Security.addProvider(new BouncyCastleProvider());
    }

    public void generate() {
        SubjectData subjectData = generateSubjectDataRoot();

        KeyPair keyPairIssuer = generateKeyPair();
        IssuerData issuerData = generateIssuerDataRoot(keyPairIssuer.getPrivate());

        CertificateGenerator certificateGenerator = new CertificateGenerator();
        X509Certificate x509Certificate = certificateGenerator.generateCertificate(subjectData, issuerData);

        KeyStoreWriter rootKs = new KeyStoreWriter();
        KeyStoreWriter caKs = new KeyStoreWriter();
        KeyStoreWriter endKs = new KeyStoreWriter();

        char[] password = new char[5];
        password[0] = '1';
        password[1] = '2';
        password[2] = '3';
        password[3] = '4';
        password[4] = '5';

        rootKs.loadKeyStore(null, password);
        caKs.loadKeyStore(null, password);
        endKs.loadKeyStore(null, password);

        rootKs.write("root", keyPairIssuer.getPrivate(), password, x509Certificate);
        rootKs.saveKeyStore("./src/main/resources/keystores/root.jks", password);

        SubjectData subjectData1 = generateSubjectData();
        CertificateGenerator certificateGenerator1 = new CertificateGenerator();
        X509Certificate x509Certificate2 = certificateGenerator1.generateCertificate(subjectData1, issuerData);

        caKs.write("ca", keyPairIssuer.getPrivate(), password, x509Certificate2);
        caKs.saveKeyStore("./src/main/resources/keystores/ca.jks", password);

        SubjectData subjectData2 = generateSubjectDataEndEntity();
        CertificateGenerator certificateGenerator2 = new CertificateGenerator();
        X509Certificate x509Certificate3 = certificateGenerator2.generateCertificate(subjectData2, issuerData);

        endKs.write("end-entity", keyPairIssuer.getPrivate(), password, x509Certificate3);
        endKs.saveKeyStore("./src/main/resources/keystores/endEntity.jks", password);
    }

    public SubjectData generateSubjectData() {
        try {
            KeyPair keyPairSubject = generateKeyPair();

            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("06-04-2022");
            Date endDate = simpleDateFormat.parse("06-04-2042");
            String serialNumber = "2";

            X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
            x500NameBuilder.addRDN(BCStyle.CN, "Luka Miletic");
            x500NameBuilder.addRDN(BCStyle.SURNAME, "Luka");
            x500NameBuilder.addRDN(BCStyle.GIVENNAME, "Miletic");
            x500NameBuilder.addRDN(BCStyle.E, "luka.miletic@gmail.com");
            x500NameBuilder.addRDN(BCStyle.C, "CA");
            x500NameBuilder.addRDN(BCStyle.UID, "0002");

            return new SubjectData(keyPairSubject.getPublic(), x500NameBuilder.build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }
    public SubjectData generateSubjectDataEndEntity() {
        try {
            KeyPair keyPairSubject = generateKeyPair();

            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("06-04-2022");
            Date endDate = simpleDateFormat.parse("06-04-2024");
            String serialNumber = "3";

            X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
            x500NameBuilder.addRDN(BCStyle.CN, "Nemanja Radojcic");
            x500NameBuilder.addRDN(BCStyle.SURNAME, "Nemanja");
            x500NameBuilder.addRDN(BCStyle.GIVENNAME, "Radojcic");
            x500NameBuilder.addRDN(BCStyle.EmailAddress, "nemanja.radojcic@gmail.com");
            x500NameBuilder.addRDN(BCStyle.C, "RS");
            x500NameBuilder.addRDN(BCStyle.UID, "0003");

            return new SubjectData(keyPairSubject.getPublic(), x500NameBuilder.build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }

    public IssuerData generateIssuerDataRoot(PrivateKey privateKey) {
        return new IssuerData(privateKey, getX500Name().build());
    }

    private X500NameBuilder getX500Name() {
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, "Mihajlo Kisic");
        x500NameBuilder.addRDN(BCStyle.SURNAME, "Mihajlo");
        x500NameBuilder.addRDN(BCStyle.GIVENNAME, "Kisic");
        x500NameBuilder.addRDN(BCStyle.EmailAddress, "mihajlo.kisic@gmail.com");
        x500NameBuilder.addRDN(BCStyle.C, "RS");
        x500NameBuilder.addRDN(BCStyle.UID, "0001");
        return x500NameBuilder;
    }

    public SubjectData generateSubjectDataRoot() {
        try {
            KeyPair keyPairSubject = generateKeyPair();

            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("19-05-2020");
            Date endDate = simpleDateFormat.parse("19-06-2020");
            String serialNumber = "1";

            return new SubjectData(keyPairSubject.getPublic(), getX500Name().build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }

    public KeyPair generateKeyPair() { // Bilo je private ja promenio u public
        try {
            KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
            SecureRandom random = SecureRandom.getInstance("SHA1PRNG", "SUN");
            keyPairGenerator.initialize(2048, random);
            return keyPairGenerator.generateKeyPair();
        } catch (NoSuchAlgorithmException | NoSuchProviderException e) {
            e.printStackTrace();
        }
        return null;
    }

    public static void main(String[] args) {
        CertificateService certificateService = new CertificateService();
        certificateService.generate();
    }
}

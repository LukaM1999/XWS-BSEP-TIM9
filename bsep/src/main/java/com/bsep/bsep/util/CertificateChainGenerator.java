package com.bsep.bsep.util;

import com.bsep.bsep.certificates.CertificateGenerator;
import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.data.SubjectData;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.keystores.KeyStoreReader;
import com.bsep.bsep.keystores.KeyStoreWriter;
import org.bouncycastle.asn1.x500.X500NameBuilder;
import org.bouncycastle.asn1.x500.style.BCStyle;
import org.bouncycastle.jce.provider.BouncyCastleProvider;

import java.security.*;
import java.security.cert.X509Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Date;
import java.util.List;

public class CertificateChainGenerator {

    public CertificateChainGenerator() {
        Security.addProvider(new BouncyCastleProvider());
    }

    public void generate() {
        KeyStoreWriter rootKs = new KeyStoreWriter();
        KeyStoreWriter caKs = new KeyStoreWriter();
        KeyStoreWriter endKs = new KeyStoreWriter();
        KeyStoreWriter keys = new KeyStoreWriter();


        char[] password = "12345".toCharArray();
        rootKs.loadKeyStore(null, password);
        caKs.loadKeyStore(null, password);
        endKs.loadKeyStore(null, password);
        keys.loadKeyStore(null, password);


        KeyPair keyPairRoot = generateKeyPair();
        KeyPair keyPairCA = generateKeyPair();
        KeyPair keyPairEE = generateKeyPair();

        SubjectData subjectDataRoot = generateSubjectDataRoot(keyPairRoot.getPublic());
        IssuerData issuerDataRoot = generateIssuerDataRoot(keyPairRoot.getPrivate());

        CertificateGenerator certificateGenerator = new CertificateGenerator();
        CertificateDTO dtoRoot = new CertificateDTO("root","root", new ArrayList<>(Arrays.asList(1, 1, 1, 1, 1, 1, 1)), "1");
        X509Certificate x509Certificate = certificateGenerator.generateCertificate(subjectDataRoot, issuerDataRoot, dtoRoot);

        rootKs.write("1", keyPairRoot.getPrivate(), password, x509Certificate);
        rootKs.saveKeyStore("./src/main/resources/keystores/root.jks", password);

        SubjectData subjectDataCA = generateSubjectDataCA(keyPairCA.getPublic());
        IssuerData issuerDataCA = generateIssuerDataRoot(keyPairRoot.getPrivate());
        CertificateGenerator certificateGenerator1 = new CertificateGenerator();
        CertificateDTO dtoCA = new CertificateDTO("root", "ca", new ArrayList<>(Arrays.asList(1, 1, 1)), "1");
        X509Certificate x509Certificate2 = certificateGenerator1.generateCertificate(subjectDataCA, issuerDataCA, dtoCA);

        caKs.write("2", keyPairRoot.getPrivate(), password, x509Certificate2);
        caKs.saveKeyStore("./src/main/resources/keystores/ca.jks", password);

        IssuerData issuerDataCANew = generateIssuerDataRoot(keyPairCA.getPrivate());
        CertificateGenerator certificateGeneratorCA = new CertificateGenerator();
        dtoCA = new CertificateDTO("ca", "ca", new ArrayList<>(Arrays.asList(1, 1, 1)), "2");
        X509Certificate x509CertificateCA = certificateGeneratorCA.generateCertificate(subjectDataCA, issuerDataCANew, dtoCA);

        // Sacuvaj CA privatni kljuc da bi mogao da se dobavi pri potpisivanju novog sertifikata
        keys.write("2", keyPairCA.getPrivate(), password, x509CertificateCA);
        keys.saveKeyStore("./src/main/resources/keystores/keys.jks", password);


        SubjectData subjectDataEE = generateSubjectDataEndEntity(keyPairEE.getPublic());
        IssuerData issuerDataEE = generateIssuerDataCA(keyPairCA.getPrivate());
        CertificateGenerator certificateGenerator2 = new CertificateGenerator();
        CertificateDTO dtoEE = new CertificateDTO("ca", "endEntity", new ArrayList<>(), "2");
        X509Certificate x509Certificate3 = certificateGenerator2.generateCertificate(subjectDataEE, issuerDataEE, dtoEE);

        endKs.write("3", keyPairCA.getPrivate(), password, x509Certificate3);
        endKs.saveKeyStore("./src/main/resources/keystores/endEntity.jks", password);
    }

    public SubjectData generateSubjectDataCA(PublicKey publicKey) {
        try {
            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("06-04-2022");
            Date endDate = simpleDateFormat.parse("06-04-2042");
            String serialNumber = "2";

            X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
            x500NameBuilder.addRDN(BCStyle.CN, "Luka Miletic");
            x500NameBuilder.addRDN(BCStyle.NAME, "Luka");
            x500NameBuilder.addRDN(BCStyle.SURNAME, "Miletic");
            x500NameBuilder.addRDN(BCStyle.UID, "luka.miletic@gmail.com");
            x500NameBuilder.addRDN(BCStyle.C, "RS");
            x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, "2");

            return new SubjectData(publicKey, x500NameBuilder.build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }
    public SubjectData generateSubjectDataEndEntity(PublicKey publicKey) {
        try {

            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("07-04-2022");
            Date endDate = simpleDateFormat.parse("07-04-2024");
            String serialNumber = "3";

            X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
            x500NameBuilder.addRDN(BCStyle.CN, "Nemanja Radojcic");
            x500NameBuilder.addRDN(BCStyle.NAME, "Nemanja");
            x500NameBuilder.addRDN(BCStyle.SURNAME, "Radojcic");
            x500NameBuilder.addRDN(BCStyle.UID, "nemanja.radojcic@gmail.com");
            x500NameBuilder.addRDN(BCStyle.C, "RS");
            x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, "3");

            return new SubjectData(publicKey, x500NameBuilder.build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }

    public IssuerData generateIssuerDataRoot(PrivateKey privateKey) {
        return new IssuerData(privateKey, getX500NameRoot().build());
    }

    private X500NameBuilder getX500NameRoot() {
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, "Mihajlo Kisic");
        x500NameBuilder.addRDN(BCStyle.NAME, "Mihajlo");
        x500NameBuilder.addRDN(BCStyle.SURNAME, "Kisic");
        x500NameBuilder.addRDN(BCStyle.UID, "mihajlo.kisic@gmail.com");
        x500NameBuilder.addRDN(BCStyle.C, "RS");
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, "1");
        return x500NameBuilder;
    }

    public IssuerData generateIssuerDataCA(PrivateKey privateKey) {
        return new IssuerData(privateKey, getX500NameCA().build());
    }

    private X500NameBuilder getX500NameCA() {
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, "Luka Miletic");
        x500NameBuilder.addRDN(BCStyle.NAME, "Luka");
        x500NameBuilder.addRDN(BCStyle.SURNAME, "Miletic");
        x500NameBuilder.addRDN(BCStyle.UID, "luka.miletic@gmail.com");
        x500NameBuilder.addRDN(BCStyle.C, "RS");
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, "2");
        return x500NameBuilder;
    }

    public SubjectData generateSubjectDataRoot(PublicKey publicKey) {
        try {
            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("dd-MM-yyyy");
            Date startDate = simpleDateFormat.parse("07-04-2022");
            Date endDate = simpleDateFormat.parse("07-04-2042");
            String serialNumber = "1";

            return new SubjectData(publicKey, getX500NameRoot().build(), serialNumber, startDate, endDate);
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return null;
    }

    public KeyPair generateKeyPair() {
        try {
            KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
            keyPairGenerator.initialize(2048, new SecureRandom());
            return keyPairGenerator.generateKeyPair();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }
        return null;
    }

    public static void main(String[] args) {
        CertificateChainGenerator certificateService = new CertificateChainGenerator();
        certificateService.generate();
    }
}

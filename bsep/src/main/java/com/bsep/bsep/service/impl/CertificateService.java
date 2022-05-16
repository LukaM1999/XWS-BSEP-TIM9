package com.bsep.bsep.service.impl;

import com.bsep.bsep.certificates.CertificateGenerator;
import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.data.SubjectData;
import com.bsep.bsep.data.UserCertificate;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.keystores.KeyStoreReader;
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
import org.springframework.core.io.InputStreamResource;
import org.springframework.stereotype.Service;

import javax.security.auth.Subject;
import java.io.*;
import java.nio.charset.StandardCharsets;
import java.security.*;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.security.cert.Certificate;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.time.LocalDateTime;
import java.util.*;

@PropertySource("classpath:application.properties")
@Service
public class CertificateService {
    @Autowired
    private Environment env;

    @Autowired
    private UserCertificateRepository userCertificateRepository;

    public X509Certificate createCertificate(CertificateDTO certificateDTO) {
        if(certificateDTO.getSerialNumberIssuer() != null && !isNewCertificateDTODateValid(certificateDTO)) return null;
        CertificateGenerator certificateGenerator = new CertificateGenerator();
        KeyPair keyPair = new CertificateChainGenerator().generateKeyPair();
        KeyStoreWriter keystore = new KeyStoreWriter();

        char[] password = "12345".toCharArray();

        UserCertificate userCertificate = userCertificateRepository.save(new UserCertificate(null, certificateDTO.getUsernameSubject(), false));

        IssuerData issuerData;
        SubjectData subjectData = generateSubjectData(certificateDTO, userCertificate.getCertificateSerialNumber().toString(), keyPair.getPublic());
        if(certificateDTO.getAuthoritySubject().equals("root")) {
            certificateDTO.setSerialNumberIssuer(userCertificate.getCertificateSerialNumber().toString());
            issuerData = generateIssuerData(certificateDTO, keyPair.getPrivate());
        }
        else
            issuerData = new KeyStoreReader().readIssuerFromStore(env.getProperty("keystore.path") + "keys.jks", certificateDTO.getSerialNumberIssuer(), password, password);

        if(issuerData == null) {
            issuerData = new KeyStoreReader().readIssuerFromStore(env.getProperty("keystore.path") + "root.jks", certificateDTO.getSerialNumberIssuer(), password, password);
        }

        //Cuva privatni kljuc od subjecta ako je CA
        if(certificateDTO.getAuthoritySubject().equals("ca")){
            KeyStoreWriter privateKeys = new KeyStoreWriter();
            IssuerData newIssuerData = generateIssuerData(certificateDTO, keyPair.getPrivate());
            privateKeys.loadKeyStore(env.getProperty("keystore.path") + "keys.jks", password);
            CertificateGenerator certificateGeneratorCA = new CertificateGenerator();
            CertificateDTO dtoCA = new CertificateDTO("ca", "ca", certificateDTO.getKeyUsages(), certificateDTO.getSerialNumberIssuer());
            X509Certificate x509CertificateCA = certificateGeneratorCA.generateCertificate(subjectData, newIssuerData, dtoCA);
            privateKeys.write(userCertificate.getCertificateSerialNumber().toString(), newIssuerData.getPrivateKey(), password, x509CertificateCA);
            privateKeys.saveKeyStore(env.getProperty("keystore.path") + "keys.jks", password);
        }

        X509Certificate x509Certificate = certificateGenerator.generateCertificate(subjectData, issuerData, certificateDTO);
        keystore.loadKeyStore( env.getProperty("keystore.path") + certificateDTO.getAuthoritySubject() + ".jks", password);
        keystore.write(userCertificate.getCertificateSerialNumber().toString(), issuerData.getPrivateKey(), password, x509Certificate);
        keystore.saveKeyStore(env.getProperty("keystore.path") + certificateDTO.getAuthoritySubject() + ".jks", password);


        return x509Certificate;
    }

    public List<CertificateDTO> getAllCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        List<X509Certificate> certificates = getAllActiveRootCertificates();
        certificates.addAll(getAllActiveCACertificates());
        certificates.addAll(getAllActiveEndUserCertificates());
        return certificateToDTO(certificates);
    }

    private SubjectData generateSubjectData(CertificateDTO certificateDTO, String serialNumber, PublicKey publicKey) {

        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, certificateDTO.getCommonNameSubject());
        x500NameBuilder.addRDN(BCStyle.NAME, certificateDTO.getNameSubject());
        x500NameBuilder.addRDN(BCStyle.SURNAME, certificateDTO.getSurnameSubject());
        x500NameBuilder.addRDN(BCStyle.UID, certificateDTO.getUsernameSubject());
        x500NameBuilder.addRDN(BCStyle.C, certificateDTO.getCountrySubject());
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, serialNumber);

        return new SubjectData(publicKey, x500NameBuilder.build(), serialNumber, certificateDTO.getStartDate(), certificateDTO.getEndDate());

    }

    private IssuerData generateIssuerData(CertificateDTO certificateDTO, PrivateKey privateKey) {
        X500NameBuilder x500NameBuilder = new X500NameBuilder(BCStyle.INSTANCE);
        x500NameBuilder.addRDN(BCStyle.CN, certificateDTO.getCommonNameIssuer());
        x500NameBuilder.addRDN(BCStyle.NAME, certificateDTO.getNameIssuer());
        x500NameBuilder.addRDN(BCStyle.SURNAME, certificateDTO.getSurnameIssuer());
        x500NameBuilder.addRDN(BCStyle.UID, certificateDTO.getUsernameIssuer());
        x500NameBuilder.addRDN(BCStyle.C, certificateDTO.getCountryIssuer());
        x500NameBuilder.addRDN(BCStyle.SERIALNUMBER, certificateDTO.getSerialNumberIssuer());

        return new IssuerData(privateKey, x500NameBuilder.build());

    }

    public List<X509Certificate> getAllEndUserCertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./keystores/endEntity.jks", "12345");
        for (X509Certificate certificate : certificates) {
            retList.add(certificate);
        }
        return retList;
    }

    public List<X509Certificate> getAllRootCertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./keystores/root.jks", "12345");
        for (X509Certificate certificate : certificates) {
            retList.add(certificate);
        }
        return retList;
    }

    public List<X509Certificate> getAllCACertificates() {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> certificates = readAllCertificate("./keystores/ca.jks", "12345");
        for (X509Certificate certificate : certificates) {
            retList.add(certificate);
        }
        return retList;
    }

    public List<X509Certificate> getAllActiveCACertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return getNotRevokedCertificates(readAllCertificate("./keystores/ca.jks", "12345"));
    }

    public List<X509Certificate> getAllActiveEndUserCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return getNotRevokedCertificates(readAllCertificate("./keystores/endEntity.jks", "12345"));
    }

    public List<X509Certificate> getAllActiveRootCertificates() throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        return getNotRevokedCertificates(readAllCertificate("./keystores/root.jks", "12345"));
    }

    private List<X509Certificate> getNotRevokedCertificates(List<X509Certificate> certificates) throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        List<X509Certificate> retList = new ArrayList<>();
        List<X509Certificate> toRevoke = new ArrayList<>();

        for (X509Certificate certificate : certificates) {
            if(!userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(certificate))).isRevoked()) {
                if (!isCertificateDateValid(certificate)) {
                    toRevoke.add(certificate);
                }
            }
        }
        for(CertificateDTO cert : certificateToDTO(toRevoke)){
            revokeCertificate(cert);
        }
        for (X509Certificate certificate : certificates) {
            if(!userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(certificate))).isRevoked()) {
                if (isCertificateDateValid(certificate)) {
                    retList.add(certificate);
                }
            }
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

    public List<CertificateDTO> certificateToDTO(List<X509Certificate> certificateList) throws CertificateException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        List<CertificateDTO> dto = new ArrayList<>();

        for(X509Certificate certificate : certificateList){
            CertificateDTO certDto = new CertificateDTO();
            JcaX509CertificateHolder certHolder = new JcaX509CertificateHolder(certificate);
            X500Name subject = certHolder.getSubject();
            X500Name issuer = certHolder.getIssuer();
            String authority = "ca";
            if(isSelfSigned(certificate)) authority = "root";
            else if(certificate.getBasicConstraints() == -1) authority = "endEntity";
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
            if(subject.getRDNs(BCStyle.UID).length > 0) {
                cn = subject.getRDNs(BCStyle.UID)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setUsernameSubject(temp);
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
            if(issuer.getRDNs(BCStyle.UID).length > 0) {
                cn = issuer.getRDNs(BCStyle.UID)[0];
                temp = IETFUtils.valueToString(cn.getFirst().getValue());
                certDto.setUsernameIssuer(temp);
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
            certDto.setAuthoritySubject(authority);
            List<Integer> keyUsages = new ArrayList<>();
            if(certificate.getKeyUsage() != null){
                for(int i = 8; i >= 0; i-- ){
                    if(i == 8 && certificate.getKeyUsage()[i]) {
                        keyUsages.add(32768);
                        continue;
                    }
                    if(certificate.getKeyUsage()[i])
                        keyUsages.add((int)Math.pow(2, 7 - i));
                }
            }
            certDto.setKeyUsages(keyUsages);
            dto.add(certDto);
        }

        return dto;
    }

    public List<CertificateDTO> getCertificateChain(
            CertificateDTO chainStart
    ) throws NoSuchAlgorithmException, InvalidKeyException,
            NoSuchProviderException, CertificateException, ParseException {

        X509Certificate startingPoint = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + chainStart.getAuthoritySubject() + ".jks", "12345", chainStart.getSerialNumberSubject());
        List<X509Certificate> certificates = new ArrayList<>(getAllActiveRootCertificates());
        certificates.addAll(getAllActiveCACertificates());
        certificates.addAll(getAllActiveEndUserCertificates());
        LinkedList path = new LinkedList();
        path.add(startingPoint);
        boolean nodeAdded = true;
        // Keep looping until an iteration happens where we don't add any nodes
        // to our path.
        while (nodeAdded) {
            // We'll start out by assuming nothing gets added.  If something
            // gets added, then nodeAdded will be changed to "true".
            nodeAdded = false;
            X509Certificate top = (X509Certificate) path.getLast();
            if (isSelfSigned(top)) {
                // We're self-signed, so we're done!
                break;
            }

            // Not self-signed.  Let's see if we're signed by anyone in the
            // collection.
            Iterator it = certificates.iterator();
            while (it.hasNext()) {
                X509Certificate x509 = (X509Certificate) it.next();
                if (verifySignatures(top, x509.getPublicKey())) {
                    // We're signed by this guy!  Add him to the chain we're
                    // building up.
                    path.add(x509);
                    nodeAdded = true;
                    it.remove(); // Not interested in this guy anymore!
                    break;
                }
                // Not signed by this guy, let's try the next guy.
            }
        }
        X509Certificate[] results = new X509Certificate[path.size()];

        path.toArray(results);
        return certificateToDTO(List.of(results));
    }

    public static boolean isSelfSigned(X509Certificate cert)
            throws CertificateException, InvalidKeyException,
            NoSuchAlgorithmException, NoSuchProviderException {

        return verifySignatures(cert, cert.getPublicKey());
    }

    private static boolean verifySignatures(X509Certificate cert, PublicKey key)
            throws CertificateException, InvalidKeyException,
            NoSuchAlgorithmException, NoSuchProviderException {

        String sigAlg = cert.getSigAlgName();
        String keyAlg = key.getAlgorithm();
        sigAlg = sigAlg != null ? sigAlg.trim().toUpperCase() : "";
        keyAlg = keyAlg != null ? keyAlg.trim().toUpperCase() : "";
        if (keyAlg.length() >= 2 && sigAlg.endsWith(keyAlg)) {
            try {
                cert.verify(key);
                return true;
            } catch (SignatureException se) {
                return false;
            }
        } else {
            return false;
        }
    }

    public boolean revokeCertificate(CertificateDTO chainStart) throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        if(userCertificateRepository.findBySerialNum(Long.parseLong(chainStart.getSerialNumberSubject())).isRevoked()) return false;

        X509Certificate startingPoint = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + chainStart.getAuthoritySubject() + ".jks", "12345", chainStart.getSerialNumberSubject());
        List<X509Certificate> certificates = new ArrayList<>(getAllCACertificates());
        certificates.addAll(getAllEndUserCertificates());

        LinkedList toRevoke = new LinkedList();
        toRevoke.add(startingPoint);

        List<X509Certificate> issuers = new ArrayList<>();

        //one kojima sam izdao
        for (X509Certificate cert: certificates) {
            if(getSubjectSerialNum(startingPoint).equals(getIssuerSerialNum(cert))){
                if(userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(cert))).isRevoked()) continue;
                issuers.add(cert);
                toRevoke.add(cert);
            }
        }
        //oni kojima je izdato od strane mojih izdatih
        for (int i = 0; i < issuers.size(); i++) {
            for (X509Certificate cert: certificates) {
                if(getSubjectSerialNum(issuers.get(i)).equals(getIssuerSerialNum(cert))){
                    if(userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(cert))).isRevoked()) continue;
                    issuers.add(cert);
                    toRevoke.add(cert);
                }
            }
        }

        X509Certificate[] results = new X509Certificate[toRevoke.size()];
        toRevoke.toArray(results);
        RevokeCertificateOCSP(results);
        return true;
    }

    private String getIssuerSerialNum(X509Certificate cert) throws CertificateEncodingException {
        JcaX509CertificateHolder holder = new JcaX509CertificateHolder(cert);
        RDN cn = holder.getIssuer().getRDNs(BCStyle.SERIALNUMBER)[0];
        return IETFUtils.valueToString(cn.getFirst().getValue());

    }

    private String getSubjectSerialNum(X509Certificate cert) throws CertificateEncodingException {
        JcaX509CertificateHolder holder = new JcaX509CertificateHolder(cert);
        RDN cn = holder.getSubject().getRDNs(BCStyle.SERIALNUMBER)[0];
        return IETFUtils.valueToString(cn.getFirst().getValue());

    }

    private void RevokeCertificateOCSP(X509Certificate[] results) throws CertificateEncodingException {
        for (X509Certificate cert: results) {
            UserCertificate certificate = userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(cert)));
            if(certificate != null) {
                certificate.setRevoked(true);
                userCertificateRepository.save(certificate);
            }
        }
    }

    public  List<CertificateDTO> getIssuedCertificates(CertificateDTO chainStart) throws CertificateException, ParseException, NoSuchAlgorithmException, InvalidKeyException, NoSuchProviderException {
        if(userCertificateRepository.findBySerialNum(Long.parseLong(chainStart.getSerialNumberSubject())).isRevoked()) return null;

        X509Certificate startingPoint = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + chainStart.getAuthoritySubject() + ".jks", "12345", chainStart.getSerialNumberSubject());
        if(startingPoint == null) return new ArrayList<>();
        List<X509Certificate> certificates = new ArrayList<>(getAllActiveCACertificates());
        certificates.addAll(getAllActiveEndUserCertificates());

        LinkedList issued = new LinkedList();

        //one kojima sam izdao
        for (X509Certificate cert: certificates) {
            if(getSubjectSerialNum(startingPoint).equals(getIssuerSerialNum(cert))){
                if(!userCertificateRepository.findBySerialNum(Long.parseLong(getSubjectSerialNum(cert))).isRevoked())
                    issued.add(cert);
            }
        }
        X509Certificate[] results = new X509Certificate[issued.size()];
        issued.toArray(results);
        return certificateToDTO(List.of(results));
    }

    private boolean isCertificateDateValid(X509Certificate certificate) throws CertificateEncodingException {
        X509Certificate issuer = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + "ca.jks", "12345", getIssuerSerialNum(certificate));

        boolean isRoot = false;
        if(issuer == null) {
            issuer = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + "root.jks", "12345", getIssuerSerialNum(certificate));
            isRoot = true;
        }
        Date now = new Date();
        if(certificate.getNotAfter().before(now)){
            return false;
        }

        if(isRoot) return true;

        return issuer.getNotBefore().before(certificate.getNotBefore())
                && issuer.getNotAfter().after(certificate.getNotAfter());
    }

    private boolean isNewCertificateDTODateValid(CertificateDTO dto) {
        X509Certificate issuer = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + "ca.jks", "12345", dto.getSerialNumberIssuer());

        if(issuer == null) {
            issuer = (X509Certificate) new KeyStoreReader().readCertificate(env.getProperty("keystore.path") + "root.jks", "12345", dto.getSerialNumberIssuer());
        }
        Date now = new Date();
        //if(dto.getStartDate().before(now) || dto.getEndDate().before(now)) return false;
        if(dto.getStartDate().after(dto.getEndDate())) return false;

        return issuer.getNotBefore().before(dto.getStartDate())
                && issuer.getNotAfter().after(dto.getEndDate());
    }

    public boolean extractCertificate(CertificateDTO certificateDto) throws CertificateException, IOException {
        String authority = "";
        if(certificateDto.getAuthoritySubject().equals("root"))
            authority = certificateDto.getAuthoritySubject();
        if(certificateDto.getAuthoritySubject().equals("ca"))
            authority = certificateDto.getAuthoritySubject();
        if(certificateDto.getAuthoritySubject().equals("endEntity"))
            authority = certificateDto.getAuthoritySubject();
        X509Certificate certificate = readCertificate(env.getProperty("keystore.path") + authority + ".jks", "12345", certificateDto.getSerialNumberSubject());
        FileOutputStream os = new FileOutputStream(certificateDto.getSerialNumberSubject() + ".crt");
        os.write("-----BEGIN CERTIFICATE-----\n".getBytes(StandardCharsets.US_ASCII));
        os.write(Base64.getEncoder().encode(certificate.getEncoded()));
        os.write("\n-----END CERTIFICATE-----\n".getBytes(StandardCharsets.US_ASCII));
        os.close();
        if(!certificateDto.getAuthoritySubject().equals("ca"))
            return true;
        PrivateKey key = new KeyStoreReader().readPrivateKey(env.getProperty("keystore.path") + "keys.jks", "12345", certificateDto.getSerialNumberSubject(), "12345");
        os = new FileOutputStream(certificateDto.getSerialNumberSubject() + "-key" + ".pem");
        os.write("-----BEGIN PRIVATE KEY-----\n".getBytes(StandardCharsets.US_ASCII));
        os.write(Base64.getEncoder().encode(key.getEncoded()));
        os.write("\n-----END PRIVATE KEY-----\n".getBytes(StandardCharsets.US_ASCII));
        os.close();
        return true;
    }
}

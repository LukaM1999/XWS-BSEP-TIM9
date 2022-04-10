package com.bsep.bsep.certificates;

import com.bsep.bsep.data.IssuerData;
import com.bsep.bsep.data.SubjectData;
import com.bsep.bsep.dto.CertificateDTO;
import com.bsep.bsep.keystores.KeyStoreReader;
import org.bouncycastle.asn1.x509.*;
import org.bouncycastle.cert.CertIOException;
import org.bouncycastle.cert.X509CertificateHolder;
import org.bouncycastle.cert.X509v3CertificateBuilder;
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter;
import org.bouncycastle.cert.jcajce.JcaX509ExtensionUtils;
import org.bouncycastle.cert.jcajce.JcaX509v3CertificateBuilder;
import org.bouncycastle.operator.ContentSigner;
import org.bouncycastle.operator.OperatorCreationException;
import org.bouncycastle.operator.jcajce.JcaContentSignerBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.stereotype.Component;

import java.math.BigInteger;
import java.security.NoSuchAlgorithmException;
import java.security.cert.CertificateEncodingException;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;

@Component
public class CertificateGenerator {

	public CertificateGenerator() {}
	
	public X509Certificate generateCertificate(SubjectData subjectData, IssuerData issuerData, CertificateDTO certificateDTO) {
		try {
			//Posto klasa za generisanje sertifiakta ne moze da primi direktno privatni kljuc pravi se builder za objekat
			//Ovaj objekat sadrzi privatni kljuc izdavaoca sertifikata i koristiti se za potpisivanje sertifikata
			//Parametar koji se prosledjuje je algoritam koji se koristi za potpisivanje sertifiakta
			JcaContentSignerBuilder builder = new JcaContentSignerBuilder("SHA256WithRSAEncryption");
			//Takodje se navodi koji provider se koristi, u ovom slucaju Bouncy Castle
			builder = builder.setProvider("BC");

			//Formira se objekat koji ce sadrzati privatni kljuc i koji ce se koristiti za potpisivanje sertifikata
			ContentSigner contentSigner = builder.build(issuerData.getPrivateKey());

			//Postavljaju se podaci za generisanje sertifiakta
			X509v3CertificateBuilder certGen = new JcaX509v3CertificateBuilder(issuerData.getX500name(),
					new BigInteger(subjectData.getSerialNumber()),
					subjectData.getStartDate(),
					subjectData.getEndDate(),
					subjectData.getX500name(),
					subjectData.getPublicKey());

			boolean isCA = !certificateDTO.getAuthoritySubject().equals("endEntity");
			certGen.addExtension(Extension.basicConstraints, true, new BasicConstraints(isCA));

			int allKeyUsages = 0;
			for (int i = 0; i < certificateDTO.getKeyUsages().size(); i++) {
				allKeyUsages = allKeyUsages + certificateDTO.getKeyUsages().get(i);
			}
			certGen.addExtension(Extension.keyUsage, true, new KeyUsage(allKeyUsages));

			JcaX509ExtensionUtils utils = new JcaX509ExtensionUtils();

			SubjectKeyIdentifier ski = utils.createSubjectKeyIdentifier(subjectData.getPublicKey());
			certGen.addExtension(Extension.subjectKeyIdentifier, false, ski);

			if (!certificateDTO.getAuthoritySubject().equals("root")) {
				java.security.cert.Certificate certIssuer = new KeyStoreReader().readCertificate("./keystores/root.jks", "12345", certificateDTO.getSerialNumberIssuer());
				if(certIssuer == null) certIssuer = new KeyStoreReader().readCertificate("./keystores/ca.jks", "12345", certificateDTO.getSerialNumberIssuer());
				AuthorityKeyIdentifier authorityKey = utils.createAuthorityKeyIdentifier(certIssuer.getPublicKey());
				certGen.addExtension(Extension.authorityKeyIdentifier, false, authorityKey);
			}
			//Generise se sertifikat
			X509CertificateHolder certHolder = certGen.build(contentSigner);

			//Builder generise sertifikat kao objekat klase X509CertificateHolder
			//Nakon toga je potrebno certHolder konvertovati u sertifikat, za sta se koristi certConverter
			JcaX509CertificateConverter certConverter = new JcaX509CertificateConverter();
			certConverter = certConverter.setProvider("BC");

			//Konvertuje objekat u sertifikat
			return certConverter.getCertificate(certHolder);
		} catch (IllegalArgumentException | IllegalStateException | OperatorCreationException | CertificateException | CertIOException | NoSuchAlgorithmException e) {
			e.printStackTrace();
		}
		return null;
	}
}

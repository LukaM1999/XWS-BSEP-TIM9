package com.bsep.bsep.keystores;

import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.lang.reflect.Array;
import java.security.*;
import java.security.cert.*;
import java.security.cert.Certificate;
import java.util.Arrays;
import java.util.List;

public class KeyStoreWriter {
	//KeyStore je Java klasa za citanje specijalizovanih datoteka koje se koriste za cuvanje kljuceva
	//Tri tipa entiteta koji se obicno nalaze u ovakvim datotekama su:
	// - Sertifikati koji ukljucuju javni kljuc
	// - Privatni kljucevi
	// - Tajni kljucevi, koji se koriste u simetricnima siframa
	private KeyStore keyStore;
	
	public KeyStoreWriter() {
		try {
			keyStore = KeyStore.getInstance("JKS", "SUN");
		} catch (KeyStoreException | NoSuchProviderException e) {
			e.printStackTrace();
		}
	}

	//Check if certificate chain is valid using bouncy castle OCSP
	public boolean isCertificateChainValid(Certificate[] chain) {
		try {
			CertificateFactory certificateFactory = CertificateFactory.getInstance("X.509", "BC");
			CertPath certPath = certificateFactory.generateCertPath(Arrays.asList(chain));
			CertPathValidator validator = CertPathValidator.getInstance("PKIX", "BC");
			PKIXParameters params = new PKIXParameters(keyStore);
			params.setRevocationEnabled(false);
			validator.validate(certPath, params);
			return true;
		} catch (CertPathValidatorException | CertificateException | NoSuchProviderException | KeyStoreException | InvalidAlgorithmParameterException | NoSuchAlgorithmException e) {
			e.printStackTrace();
			return false;
		}
	}


	
	public void loadKeyStore(String fileName, char[] password) {
		try {
			if(fileName != null) {
				keyStore.load(new FileInputStream(fileName), password);
			} else {
				//Ako je cilj kreirati novi KeyStore poziva se i dalje load, pri cemu je prvi parametar null
				keyStore.load(null, password);
			}
		} catch (NoSuchAlgorithmException | CertificateException | IOException e) {
			e.printStackTrace();
		}
	}
	
	public void saveKeyStore(String fileName, char[] password) {
		try {
			keyStore.store(new FileOutputStream(fileName), password);
		} catch (KeyStoreException | NoSuchAlgorithmException | CertificateException | IOException e) {
			e.printStackTrace();
		}
	}
	
	public void write(String alias, PrivateKey privateKey, char[] password, Certificate certificate) {
		try {
			keyStore.setKeyEntry(alias, privateKey, password, new Certificate[] {certificate});
		} catch (KeyStoreException e) {
			e.printStackTrace();
		}
	}
}

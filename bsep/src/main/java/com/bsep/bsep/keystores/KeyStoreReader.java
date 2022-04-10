package com.bsep.bsep.keystores;

import com.bsep.bsep.data.IssuerData;
import org.bouncycastle.asn1.x500.X500Name;
import org.bouncycastle.cert.jcajce.JcaX509CertificateHolder;

import java.io.BufferedInputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.security.*;
import java.security.cert.Certificate;
import java.security.cert.CertificateException;
import java.security.cert.X509Certificate;
import java.util.Collection;
import java.util.Iterator;
import java.util.LinkedList;

public class KeyStoreReader {
	//KeyStore je Java klasa za citanje specijalizovanih datoteka koje se koriste za cuvanje kljuceva
	//Tri tipa entiteta koji se obicno nalaze u ovakvim datotekama su:
	// - Sertifikati koji ukljucuju javni kljuc
	// - Privatni kljucevi
	// - Tajni kljucevi, koji se koriste u simetricnima siframa
	private KeyStore keyStore;
	
	public KeyStoreReader() {
		try {
			keyStore = KeyStore.getInstance("JKS", "SUN");
		} catch (KeyStoreException | NoSuchProviderException e) {
			e.printStackTrace();
		}
	}
	/**
	 * Zadatak ove funkcije jeste da ucita podatke o izdavaocu i odgovarajuci privatni kljuc.
	 * Ovi podaci se mogu iskoristiti da se novi sertifikati izdaju.
	 * 
	 * @param keyStoreFile - datoteka odakle se citaju podaci
	 * @param alias - alias putem kog se identifikuje sertifikat izdavaoca
	 * @param password - lozinka koja je neophodna da se otvori key store
	 * @param keyPass - lozinka koja je neophodna da se izvuce privatni kljuc
	 * @return - podatke o izdavaocu i odgovarajuci privatni kljuc
	 */
	public IssuerData readIssuerFromStore(String keyStoreFile, String alias, char[] password, char[] keyPass) {
		try {
			//Datoteka se ucitava
			BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
			keyStore.load(in, password);
			//Iscitava se sertifikat koji ima dati alias
			Certificate cert = keyStore.getCertificate(alias);
			if(cert == null) return null;
			//Iscitava se privatni kljuc vezan za javni kljuc koji se nalazi na sertifikatu sa datim aliasom
			PrivateKey privKey = (PrivateKey) keyStore.getKey(alias, keyPass);

			X500Name issuerName = new JcaX509CertificateHolder((X509Certificate) cert).getSubject();
			return new IssuerData(privKey, issuerName);
		} catch (KeyStoreException e) {
			e.printStackTrace();
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (CertificateException e) {
			e.printStackTrace();
		} catch (UnrecoverableKeyException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		}
		return null;
	}
	
	/**
	 * Ucitava sertifikat is KS fajla
	 */
    public Certificate readCertificate(String keyStoreFile, String keyStorePass, String alias) {
		try {
			//kreiramo instancu KeyStore
			KeyStore ks = KeyStore.getInstance("JKS", "SUN");
			//ucitavamo podatke

			BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
			ks.load(in, keyStorePass.toCharArray());
			if(ks.isKeyEntry(alias)) {
				return ks.getCertificate(alias);
			}
		} catch (KeyStoreException | NoSuchProviderException | NoSuchAlgorithmException | CertificateException | IOException e) {
			e.printStackTrace();
		}
		return null;
	}

	public static X509Certificate[] buildPath(
			X509Certificate startingPoint, Collection certificates
	) throws NoSuchAlgorithmException, InvalidKeyException,
			NoSuchProviderException, CertificateException {

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
				if (verify(top, x509.getPublicKey())) {
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
		return results;
	}

	public static boolean isSelfSigned(X509Certificate cert)
			throws CertificateException, InvalidKeyException,
			NoSuchAlgorithmException, NoSuchProviderException {

		return verify(cert, cert.getPublicKey());
	}

	public static boolean verify(X509Certificate cert, PublicKey key)
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

	public Certificate[] readCertificateChain(String keyStoreFile, String keyStorePass, String alias) {
		try {
			//kreiramo instancu KeyStore
			KeyStore ks = KeyStore.getInstance("JKS", "SUN");
			//ucitavamo podatke
			BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
			ks.load(in, keyStorePass.toCharArray());
			if(ks.isKeyEntry(alias)) {
				return ks.getCertificateChain(alias);
			}
		} catch (KeyStoreException | NoSuchProviderException | NoSuchAlgorithmException | CertificateException | IOException e) {
			e.printStackTrace();
		}
		return null;
	}
	
	/**
	 * Ucitava privatni kljuc is KS fajla
	 */
	public PrivateKey readPrivateKey(String keyStoreFile, String keyStorePass, String alias, String pass) {
		try {
			//kreiramo instancu KeyStore
			KeyStore ks = KeyStore.getInstance("JKS", "SUN");
			//ucitavamo podatke
			BufferedInputStream in = new BufferedInputStream(new FileInputStream(keyStoreFile));
			ks.load(in, keyStorePass.toCharArray());
			
			if(ks.isKeyEntry(alias)) {
				PrivateKey pk = (PrivateKey) ks.getKey(alias, pass.toCharArray());
				return pk;
			}
		} catch (KeyStoreException e) {
			e.printStackTrace();
		} catch (NoSuchProviderException e) {
			e.printStackTrace();
		} catch (FileNotFoundException e) {
			e.printStackTrace();
		} catch (NoSuchAlgorithmException e) {
			e.printStackTrace();
		} catch (CertificateException e) {
			e.printStackTrace();
		} catch (IOException e) {
			e.printStackTrace();
		} catch (UnrecoverableKeyException e) {
			e.printStackTrace();
		}
		return null;
	}
}

package com.bsep.bsep;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication(scanBasePackages = {"com.bsep.bsep"})
public class BsepApplication {

	public static void main(String[] args) {
		SpringApplication.run(BsepApplication.class, args);
	}

}

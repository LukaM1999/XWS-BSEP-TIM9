package com.agent.agent.controller;

import com.agent.agent.model.Company;
import com.agent.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/admin")
@PreAuthorize("hasAuthority('ADMIN')")
public class AdminController {

    @Autowired
    private CompanyService companyService;

    @GetMapping("/company")
    public ResponseEntity<List<Company>> getAllUnapproved() {
        List<Company> companies = companyService.getAllUnapproved();
        if(companies.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(companies, HttpStatus.OK);
    }

    @PatchMapping("/company/{companyName}")
    public ResponseEntity<Company> approveCompany(@PathVariable String companyName) {
        if(!companyService.approveCompany(companyName))
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        return new ResponseEntity<>(HttpStatus.OK);
    }

    @DeleteMapping("/company/{companyName}")
    public ResponseEntity<Company> declineCompany(@PathVariable String companyName) {
        if(companyService.declineCompany(companyName) < 1)
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        return new ResponseEntity<>(HttpStatus.OK);
    }
}

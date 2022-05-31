package com.agent.agent.service;

import com.agent.agent.model.Company;
import com.agent.agent.model.RegisteredUser;

import java.util.List;

public interface CompanyService {


    boolean approveCompany(String companyName);
    Company createCompany(Company company, RegisteredUser owner);

    List<Company> getAllUnapproved();

    List<Company> getAllApproved();

    List<Company> searchCompanies(String name);

    Company getCompanyByName(String name);

    Company updateCompany(Company company);
}

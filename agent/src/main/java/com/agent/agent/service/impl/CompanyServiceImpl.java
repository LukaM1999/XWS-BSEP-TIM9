package com.agent.agent.service.impl;

import com.agent.agent.model.Company;
import com.agent.agent.model.RegisteredUser;
import com.agent.agent.model.Role;
import com.agent.agent.repository.CompanyRepository;
import com.agent.agent.repository.RegisteredUserRepository;
import com.agent.agent.repository.RoleRepository;
import com.agent.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CompanyServiceImpl implements CompanyService {

    @Autowired
    private CompanyRepository companyRepository;

    @Autowired
    private RegisteredUserRepository registeredUserRepository;

    @Autowired
    private RoleRepository roleRepository;

    @Override
    public boolean approveCompany(String companyName) {
        Company company = companyRepository.findByName(companyName);
        if(company == null) return false;
        RegisteredUser owner = registeredUserRepository.findByUsername(company.getOwnerUsername());
        if(owner == null) return false;
        owner.setRole(roleRepository.findByRoleName("COMPANY_OWNER"));
        registeredUserRepository.save(owner);
        company.setApproved(true);
        companyRepository.save(company);
        return true;
    }

    @Override
    public Company createCompany(Company company, RegisteredUser owner) {
        if(!owner.getUsername().equals(company.getOwnerUsername()))
            return null;
        company.setCompanyOwner(owner);
        return companyRepository.save(company);
    }

    @Override
    public List<Company> getAllUnapproved() {
        return companyRepository.findAllUnapproved();
    }

    @Override
    public List<Company> getAllApproved() {
        return companyRepository.findAllApproved();
    }

    @Override
    public List<Company> searchCompanies(String name) {
        return companyRepository.searchCompanies(name);
    }

    @Override
    public Company getCompanyByName(String name) {
        return companyRepository.findByNameAndApproved(name);
    }

    @Override
    public Company updateCompany(Company company) {
        return companyRepository.save(company);
    }

    @Override
    public void updateCompanyRating(double rating, String companyName) {
        companyRepository.updateCompanyRating(rating, companyName);
    }


}

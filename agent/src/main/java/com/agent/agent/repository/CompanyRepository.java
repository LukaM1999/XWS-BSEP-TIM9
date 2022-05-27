package com.agent.agent.repository;

import com.agent.agent.model.Company;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface CompanyRepository extends JpaRepository<Company, Long> {

    @Query("SELECT c FROM Company c WHERE c.name = ?1 and c.isApproved = false")
    Company findByName(String companyName);

    @Query("SELECT c FROM Company c WHERE c.isApproved = false")
    List<Company> findAllUnapproved();


    @Query("SELECT c FROM Company c WHERE c.name LIKE %?1% and c.isApproved = true")
    List<Company> searchCompanies(String name);

    @Query("SELECT c FROM Company c WHERE c.isApproved = true and c.name = ?1")
    Company findByNameAndApproved(String name);
}

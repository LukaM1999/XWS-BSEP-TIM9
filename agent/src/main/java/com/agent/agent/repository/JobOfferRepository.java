package com.agent.agent.repository;

import com.agent.agent.model.JobOffer;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface JobOfferRepository extends JpaRepository<JobOffer, Long> {

    List<JobOffer> findAllByCompanyName(String companyName);

    JobOffer findByCompanyName(String companyName);

}

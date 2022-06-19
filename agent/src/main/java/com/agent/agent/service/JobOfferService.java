package com.agent.agent.service;

import com.agent.agent.model.JobOffer;

import java.util.List;

public interface JobOfferService {
    JobOffer addJobOffer(JobOffer jobOffer);
    List<JobOffer> getCompanyJobOffers(String companyName);
    void promoteJobOffer(Long jobOfferId);
}

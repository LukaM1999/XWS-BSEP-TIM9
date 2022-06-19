package com.agent.agent.service.impl;

import com.agent.agent.model.JobOffer;
import com.agent.agent.repository.JobOfferRepository;
import com.agent.agent.service.JobOfferService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.List;

@Service
public class JobOfferServiceImpl implements JobOfferService {

    @Autowired
    private JobOfferRepository jobOfferRepository;

    @Override
    public JobOffer addJobOffer(JobOffer jobOffer) {
        jobOffer.setCreatedAt(LocalDateTime.now().atZone(ZoneId.of("UTC")).toLocalDateTime());
        return jobOfferRepository.save(jobOffer);
    }

    @Override
    public List<JobOffer> getCompanyJobOffers(String companyName) {
        return jobOfferRepository.findAllByCompanyName(companyName);
    }

    @Override
    public void promoteJobOffer(Long jobOfferId) {
        JobOffer jobOffer = jobOfferRepository.findById(jobOfferId).get();
        jobOffer.setPromoted(true);
        jobOfferRepository.save(jobOffer);
    }
}
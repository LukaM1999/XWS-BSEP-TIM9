package com.agent.agent.service.impl;

import com.agent.agent.model.Interview;
import com.agent.agent.repository.InterviewRepository;
import com.agent.agent.service.InterviewService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;

@Service
public class InterviewServiceImpl implements InterviewService {

    @Autowired
    private InterviewRepository interviewRepository;

    @Override
    public Interview addInterview(Interview interview) {
        interview.setDateCreated(LocalDate.now());
        return interviewRepository.save(interview);
    }

    @Override
    public List<Interview> getCompanyInterviews(String companyName) {
        return interviewRepository.findAllByCompanyName(companyName);
    }

    @Override
    public Double getCompanyInterviewAverageRating(String companyName) {
        return interviewRepository.findAverageRatingByCompany(companyName);
    }
}

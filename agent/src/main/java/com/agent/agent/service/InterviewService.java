package com.agent.agent.service;

import com.agent.agent.model.Interview;

import java.util.List;

public interface InterviewService {
    Interview addInterview(Interview interview);
    List<Interview> getCompanyInterviews(String companyName);
    Double getCompanyInterviewAverageRating(String companyName);
}

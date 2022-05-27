package com.agent.agent.repository;

import com.agent.agent.model.Interview;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface InterviewRepository extends JpaRepository<Interview, Long> {

    List<Interview> findAllByCompanyName(String companyName);

    @Query("SELECT AVG(i.rating) FROM Interview i WHERE i.companyName = ?1")
    Double findAverageRatingByCompany(String companyName);
}

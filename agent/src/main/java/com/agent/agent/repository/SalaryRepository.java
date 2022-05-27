package com.agent.agent.repository;

import com.agent.agent.model.Salary;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.List;

public interface SalaryRepository extends JpaRepository<Salary, Long> {

    List<Salary> findAllByCompanyName(String companyName);

    @Query("SELECT DISTINCT s.position FROM Salary s WHERE s.companyName = ?1")
    List<String> findAllPositions(String companyName);


    @Query("SELECT AVG(s.monthlyNetSalary) FROM Salary s WHERE s.companyName = ?1 and s.position = ?2")
    Double findAverageSalaryByCompanyAndPosition(String companyName, String position);
}

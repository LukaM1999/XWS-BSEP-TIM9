package com.agent.agent.service;

import com.agent.agent.model.Salary;

import java.util.List;
import java.util.Map;

public interface SalaryService {
    Salary addSalary(Salary salary);
    List<Salary> getCompanySalaries(String companyName);
    Map<String, Double> getCompanySalaryAveragePerPosition(String companyName);
}

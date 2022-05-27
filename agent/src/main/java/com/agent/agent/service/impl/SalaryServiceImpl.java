package com.agent.agent.service.impl;

import com.agent.agent.model.Salary;
import com.agent.agent.repository.SalaryRepository;
import com.agent.agent.service.SalaryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class SalaryServiceImpl implements SalaryService {

    @Autowired
    private SalaryRepository salaryRepository;

    @Override
    public List<Salary> getCompanySalaries(String companyName) {
        return salaryRepository.findAllByCompanyName(companyName);
    }

    @Override
    public Map<String, Double> getCompanySalaryAveragePerPosition(String companyName) {
        List<String> positions = salaryRepository.findAllPositions(companyName);
        Map<String, Double> averagePerPosition = new HashMap<>(positions.size());
        for(String position : positions) {
            averagePerPosition.put(position, salaryRepository.findAverageSalaryByCompanyAndPosition(companyName, position));
        }
        return averagePerPosition;
    }

    @Override
    public Salary addSalary(Salary salary) {
        return salaryRepository.save(salary);
    }
}

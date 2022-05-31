package com.agent.agent.controller;

import com.agent.agent.dto.CompanyDTO;
import com.agent.agent.model.*;
import com.agent.agent.service.*;
import com.agent.agent.util.CompanyDTOMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = "/company", produces = MediaType.APPLICATION_JSON_VALUE)
public class CompanyController {

    @Autowired
    private CompanyService companyService;

    @Autowired
    private CommentService commentService;

    @Autowired
    private SalaryService salaryService;

    @Autowired
    private InterviewService interviewService;

    @Autowired
    private RegisteredUserService registeredUserService;

    @Autowired
    private CompanyDTOMapper companyDTOMapper;

    @PostMapping("")
    @PreAuthorize("hasAuthority('USER')")
    public ResponseEntity<Company> createCompany(@RequestBody Company company) {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        Company createdCompany = companyService.createCompany(company, (RegisteredUser) auth.getPrincipal());
        if(createdCompany == null)
            return new ResponseEntity<>(HttpStatus.FORBIDDEN);
        return new ResponseEntity<>(createdCompany, HttpStatus.CREATED);
    }

    @GetMapping("/search")
    public ResponseEntity<List<Company>> searchCompanies(@RequestParam(value = "name", required = false) String name) {
        List<Company> companies = companyService.searchCompanies(name);
        if(companies.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(companies, HttpStatus.OK);
    }

    @GetMapping("")
    public ResponseEntity<List<Company>> getAllApproved() {
        List<Company> companies = companyService.getAllApproved();
        if(companies.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(companies, HttpStatus.OK);
    }

    @GetMapping("/{companyName}")
    public ResponseEntity<Company> getCompany(@PathVariable String companyName) {
        Company company = companyService.getCompanyByName(companyName);
        if(company == null)
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(company, HttpStatus.OK);
    }

    @PostMapping("/comment")
    @PreAuthorize("hasAuthority('USER')")
    public ResponseEntity<Comment> addComment(@RequestBody Comment comment) {
        if(commentService.addComment(comment) == null)
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        return new ResponseEntity<>(comment, HttpStatus.OK);
    }

    @GetMapping("/{companyName}/comment")
    public ResponseEntity<List<Comment>> getCompanyComments(@PathVariable String companyName) {
        List<Comment> comments = commentService.getCompanyComments(companyName);
        if(comments.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(comments, HttpStatus.OK);
    }

    @PostMapping("/salary")
    @PreAuthorize("hasAuthority('USER')")
    public ResponseEntity<Salary> addSalary(@RequestBody Salary salary) {
        if(salaryService.addSalary(salary) == null)
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        return new ResponseEntity<>(salary, HttpStatus.OK);
    }

    @GetMapping("/{companyName}/salary")
    public ResponseEntity<List<Salary>> getCompanySalaries(@PathVariable String companyName) {
        List<Salary> salaries = salaryService.getCompanySalaries(companyName);
        if(salaries.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(salaries, HttpStatus.OK);
    }

    @GetMapping("/{companyName}/salary/average")
    public ResponseEntity<Map<String, Double>> getCompanySalaryAveragePerPosition(@PathVariable String companyName) {
        Map<String, Double> averages = salaryService.getCompanySalaryAveragePerPosition(companyName);
        if(averages.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(averages, HttpStatus.OK);
    }

    @PostMapping("/interview")
    @PreAuthorize("hasAuthority('USER')")
    public ResponseEntity<Interview> addInterview(@RequestBody Interview interview) {
        if(interviewService.addInterview(interview) == null)
            return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
        return new ResponseEntity<>(interview, HttpStatus.OK);
    }

    @GetMapping("/{companyName}/interview")
    public ResponseEntity<List<Interview>> getCompanyInterviews(@PathVariable String companyName) {
        List<Interview> interviews = interviewService.getCompanyInterviews(companyName);
        if(interviews.isEmpty())
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(interviews, HttpStatus.OK);
    }

    @GetMapping("/{companyName}/interview/average")
    public ResponseEntity<Double> getCompanyInterviewAverageRating(@PathVariable String companyName) {
        Double average = interviewService.getCompanyInterviewAverageRating(companyName);
        if(average == null)
            return new ResponseEntity<>(HttpStatus.NO_CONTENT);
        return new ResponseEntity<>(average, HttpStatus.OK);
    }

    @PatchMapping("/{companyName}")
    @PreAuthorize("hasAuthority('COMPANY_OWNER')")
    public ResponseEntity<Company> updateCompany(@PathVariable String companyName, @RequestBody CompanyDTO companyDTO) {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        if(!((RegisteredUser)auth.getPrincipal()).getUsername().equals(companyDTO.getOwnerUsername()))
            return new ResponseEntity<>(HttpStatus.FORBIDDEN);
        Company company = companyService.getCompanyByName(companyName);
        if(company == null)
            return new ResponseEntity<>(HttpStatus.NOT_FOUND);
        Company updatedCompany = companyDTOMapper.updateWithNullAsNoChange(companyDTO, company);
        return ResponseEntity.ok(companyService.updateCompany(updatedCompany));
    }
}

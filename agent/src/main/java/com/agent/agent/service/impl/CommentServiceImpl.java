package com.agent.agent.service.impl;

import com.agent.agent.model.Comment;
import com.agent.agent.repository.CommentRepository;
import com.agent.agent.service.CommentService;
import com.agent.agent.service.CompanyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.List;

@Service
public class CommentServiceImpl implements CommentService {

    @Autowired
    private CommentRepository commentRepository;

    @Autowired
    private CompanyService companyService;

    @Override
    public Comment addComment(Comment comment) {
        comment.setDateCreated(LocalDate.now());
        companyService.updateCompanyRating(comment.getRating(), comment.getCompanyName());
        return commentRepository.save(comment);
    }

    @Override
    public List<Comment> getCompanyComments(String companyName) {
        return commentRepository.findAllByCompanyName(companyName);
    }
}
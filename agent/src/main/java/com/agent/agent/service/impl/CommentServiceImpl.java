package com.agent.agent.service.impl;

import com.agent.agent.model.Comment;
import com.agent.agent.repository.CommentRepository;
import com.agent.agent.service.CommentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CommentServiceImpl implements CommentService {

    @Autowired
    private CommentRepository commentRepository;

    @Override
    public Comment addComment(Comment comment) {
        return commentRepository.save(comment);
    }

    @Override
    public List<Comment> getCompanyComments(String companyName) {
        return commentRepository.findAllByCompanyName(companyName);
    }
}
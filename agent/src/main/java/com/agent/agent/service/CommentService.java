package com.agent.agent.service;

import com.agent.agent.model.Comment;

import java.util.List;

public interface CommentService {

    Comment addComment(Comment comment);

    List<Comment> getCompanyComments(String companyName);
}

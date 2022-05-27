package com.agent.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table
public class Comment {
    @Id
    @SequenceGenerator(name = "comment_id_gen", sequenceName = "comment_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "comment_id_gen")
    @Getter
    private Long id;

    @Column
    @Getter
    @Setter
    private String companyName;

    @Column
    @Getter
    @Setter
    private String position;

    @Column
    @Getter
    @Setter
    private String seniority;

    @Column
    @Getter
    @Setter
    private String engagement;

    @Column
    @Getter
    @Setter
    private boolean currentlyEmployed;

    @Column
    @Getter
    @Setter
    private String subject;

    @Column
    @Getter
    @Setter
    private String content;

    @Column
    @Getter
    @Setter
    private double rating;

    //Constructors
    public Comment() {
    }

    public Comment(String companyName, String position, String seniority, String engagement, boolean currentlyEmployed,
                   String subject, String content, double rating) {
        this.companyName = companyName;
        this.position = position;
        this.seniority = seniority;
        this.engagement = engagement;
        this.currentlyEmployed = currentlyEmployed;
        this.subject = subject;
        this.content = content;
        this.rating = rating;
    }
}

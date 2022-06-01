package com.agent.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDate;

@Entity
@Table
public class Interview {
    @Id
    @SequenceGenerator(name = "interview_id_gen", sequenceName = "interview_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "interview_id_gen")
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
    private String year;

    @Column
    @Getter
    @Setter
    private String subject;

    @Column
    @Getter
    @Setter
    private String hr;

    @Column
    @Getter
    @Setter
    private String technical;

    @Column
    @Getter
    @Setter
    private int duration;

    @Column
    @Getter
    @Setter
    private int difficulty;

    @Column
    @Getter
    @Setter
    private double rating;

    @Column
    @Getter
    @Setter
    private LocalDate dateCreated;

    @Column(columnDefinition = "boolean default false")
    @Getter
    @Setter
    private boolean acceptedOffer;

    public Interview() {
    }

    public Interview(String companyName, String position, String year, String subject, String hr, String technical,
                     int duration, int difficulty, double rating) {
        this.companyName = companyName;
        this.position = position;
        this.year = year;
        this.subject = subject;
        this.hr = hr;
        this.technical = technical;
        this.duration = duration;
        this.difficulty = difficulty;
        this.rating = rating;
    }
}

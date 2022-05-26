package com.agent.agent.model;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table
public class Salary {
    @Id
    @SequenceGenerator(name = "salary_id_gen", sequenceName = "salary_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "salary_id_gen")
    @Getter
    private Long id;

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
    private double monthlyNetSalary;

    public Salary() {
    }

    public Salary(String position, String seniority, String engagement, boolean currentlyEmployed, double monthlyNetSalary) {
        this.position = position;
        this.seniority = seniority;
        this.engagement = engagement;
        this.currentlyEmployed = currentlyEmployed;
        this.monthlyNetSalary = monthlyNetSalary;
    }
}

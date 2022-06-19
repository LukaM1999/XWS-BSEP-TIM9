package com.agent.agent.model;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.springframework.format.annotation.DateTimeFormat;

import javax.persistence.*;
import java.time.LocalDate;
import java.time.LocalDateTime;

@NoArgsConstructor
@AllArgsConstructor
@Entity
public class JobOffer {

    @Id
    @SequenceGenerator(name = "job_id_gen", sequenceName = "job_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "job_id_gen")
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
    private String description;

    @Column
    @Getter
    @Setter
    private String criteria;

    @Column
    @Getter
    @Setter
    @JsonFormat(shape = JsonFormat.Shape.STRING)
    private LocalDateTime createdAt;

    @Column
    @Getter
    @Setter
    private boolean isPromoted;

}

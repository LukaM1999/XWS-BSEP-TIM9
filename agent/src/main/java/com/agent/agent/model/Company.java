package com.agent.agent.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;

@Entity
@Table
public class Company {

    @Id
    @Column
    @SequenceGenerator(name = "company_id_gen", sequenceName = "company_id_seq", initialValue = 1, allocationSize = 1)
    @GeneratedValue(strategy = GenerationType.SEQUENCE, generator = "company_id_gen")
    @Getter
    private Long id;

    @JoinColumn(name = "owner_username", insertable = false, updatable = false)
    @ManyToOne(targetEntity = RegisteredUser.class, fetch = FetchType.EAGER)
    @Getter
    @Setter
    @JsonIgnore
    private RegisteredUser companyOwner;

    @Column(name = "owner_username")
    @Getter
    @Setter
    private String ownerUsername;
    @Column(unique = true)
    @Getter
    @Setter
    private String name;
    @Column
    @Getter
    @Setter
    private String country;
    @Column
    @Getter
    @Setter
    private String city;
    @Column
    @Getter
    @Setter
    private String address;
    @Column
    @Getter
    @Setter
    private String phone;
    @Column
    @Getter
    @Setter
    private String email;
    @Column
    @Getter
    @Setter
    private String website;
    @Column(columnDefinition = "VARCHAR(1000) default ''")
    @Getter
    @Setter
    private String description;
    @Column
    @Getter
    @Setter
    private String yearEstablished;
    @Column
    @Getter
    @Setter
    private String size;
    @Column
    @Getter
    @Setter
    private String industry;
    @Column
    @Getter
    @Setter
    private boolean isApproved;
    @Column(columnDefinition = "Decimal(3,2) default '0.00'")
    @Getter
    @Setter
    private double rating;
    @Column(columnDefinition = "INTEGER default '0'")
    @Getter
    @Setter
    private int ratingCount;

    public Company() {
        super();
    }

    public Company(String ownerUsername, String name, String country, String city, String address, String phone, String email, String website, String description, String yearEstablished, String size, String industry) {
        this.ownerUsername = ownerUsername;
        this.name = name;
        this.country = country;
        this.city = city;
        this.address = address;
        this.phone = phone;
        this.email = email;
        this.website = website;
        this.description = description;
        this.yearEstablished = yearEstablished;
        this.size = size;
        this.industry = industry;
    }
}

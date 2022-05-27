package com.agent.agent.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;


@NoArgsConstructor
@AllArgsConstructor
public class CompanyDTO {

    @Getter
    @Setter
    private String ownerUsername;
    @Getter
    @Setter
    private String name;
    @Getter
    @Setter
    private String address;
    @Getter
    @Setter
    private String city;
    @Getter
    @Setter
    private String country;
    @Getter
    @Setter
    private String phone;
    @Getter
    @Setter
    private String email;
    @Getter
    @Setter
    private String website;
    @Getter
    @Setter
    private String description;
    @Getter
    @Setter
    private String yearEstablished;
    @Getter
    @Setter
    private String size;
    @Getter
    @Setter
    private String industry;
}

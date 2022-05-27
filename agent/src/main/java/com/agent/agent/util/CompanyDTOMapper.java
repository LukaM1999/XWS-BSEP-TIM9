package com.agent.agent.util;

import com.agent.agent.dto.CompanyDTO;
import com.agent.agent.model.Company;
import org.mapstruct.BeanMapping;
import org.mapstruct.Mapper;
import org.mapstruct.MappingTarget;
import org.mapstruct.NullValuePropertyMappingStrategy;


@Mapper(componentModel = "spring")

public interface CompanyDTOMapper {

    /**
     * Null values in the fields of the DTO will not be set as null in the target. They will be ignored instead.
     *
     * @return The target Company object
     */
    @BeanMapping(nullValuePropertyMappingStrategy = NullValuePropertyMappingStrategy.IGNORE)
    Company updateWithNullAsNoChange(CompanyDTO companyDTO, @MappingTarget Company company);

}

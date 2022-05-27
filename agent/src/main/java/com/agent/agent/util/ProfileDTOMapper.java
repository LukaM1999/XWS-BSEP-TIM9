package com.agent.agent.util;

import com.agent.agent.dto.ProfileDTO;
import com.agent.agent.model.RegisteredUser;
import com.agent.agent.service.RegisteredUserService;
import org.mapstruct.BeanMapping;
import org.mapstruct.Mapper;
import org.mapstruct.MappingTarget;
import org.mapstruct.NullValuePropertyMappingStrategy;
import org.springframework.context.annotation.Bean;


@Mapper(componentModel = "spring")

public interface ProfileDTOMapper {

    /**
     * Null values in the fields of the DTO will not be set as null in the target. They will be ignored instead.
     *
     * @return The target RegisteredUser object
     */
    @BeanMapping(nullValuePropertyMappingStrategy = NullValuePropertyMappingStrategy.IGNORE)
    RegisteredUser updateWithNullAsNoChange(ProfileDTO profileDTO, @MappingTarget RegisteredUser user);

}

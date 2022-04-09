package com.bsep.bsep.repository;

import com.bsep.bsep.data.UserCertificate;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import java.util.List;

public interface UserCertificateRepository extends JpaRepository<UserCertificate, Long> {

    @Query("select c from UserCertificate c where " +
            "c.certificateSerialNumber = :serial ")
    UserCertificate findBySerialNum(@Param("serial") Long serial);

    List<UserCertificate> findByUsername(String username);
}

package com.bsep.bsep.repository;

import com.bsep.bsep.data.Account;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.security.core.userdetails.UserDetails;

public interface AccountRepository extends JpaRepository<Account, Long> {
    UserDetails findByUsername(String username);
}

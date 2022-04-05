package com.bsep.bsep.repository;

import com.bsep.bsep.data.Account;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AccountRepository extends JpaRepository<Account, Long> {
}

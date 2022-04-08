insert into role(id, role_name) values (nextval('role_id_seq'), 'ADMIN');
insert into account(email, password, role_id) values ('admin@gmail.com', 'admin123', 1);
insert into user_certificate(certificate_serial_number, email, revoked) values (nextval('crt_id_seq'), 'mihajlo.kisic@gmail.com', false);
insert into user_certificate(certificate_serial_number, email, revoked) values (nextval('crt_id_seq'), 'luka.miletic@gmail.com', false);
insert into user_certificate(certificate_serial_number, email, revoked) values (nextval('crt_id_seq'), 'nemanja.radojcic@gmail.com', false);
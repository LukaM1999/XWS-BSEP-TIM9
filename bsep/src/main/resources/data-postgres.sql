insert into role(role_name) values ('admin');
insert into role(role_name) values ('ca');
insert into role(role_name) values ('endEntity');
                                                                                        --jabuka123
insert into account(username, password, role_name, enabled) values ('admin@gmail.com', '$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i', 'admin', true);
insert into account(username, password, role_name, enabled) values ('luka.miletic@gmail.com', '$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i', 'ca', true);
insert into account(username, password, role_name, enabled) values ('nemanja.radojcic@gmail.com', '$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i', 'endEntity', true);
insert into user_certificate(certificate_serial_number, username, revoked) values (nextval('crt_id_seq'), 'mihajlo.kisic@gmail.com', false);
insert into user_certificate(certificate_serial_number, username, revoked) values (nextval('crt_id_seq'), 'luka.miletic@gmail.com', false);
insert into user_certificate(certificate_serial_number, username, revoked) values (nextval('crt_id_seq'), 'nemanja.radojcic@gmail.com', false);
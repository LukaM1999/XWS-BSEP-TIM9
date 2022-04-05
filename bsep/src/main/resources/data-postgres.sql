insert into role(id, role_name) values (nextval('role_id_seq'), 'ADMIN');
insert into account(email, password, role_id) values ('admin@gmail.com', 'admin123', 1);
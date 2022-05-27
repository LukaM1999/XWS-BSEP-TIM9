insert into role(role_name) values ('ADMIN');
insert into role(role_name) values ('USER');
insert into role(role_name) values ('COMPANY_OWNER');
insert into registered_user(username, password, first_name, last_name, email, address, city, country, phone, role_name,
                            enabled) values ('Admin', '$2a$10$UVn74F/yEiUzKWBSGVyzHe2UfpVJ95zY50Q8bz1RFyrAYVfwFAj4i',
                                             'Admin', 'Admincic', 'admin@gmail.com', 'Strazilovska 27', 'Novi Sad',
                                             'Serbia', '066432231', 'ADMIN', true);
insert into registered_user(username, password, role_name, enabled)
values ('imbiamba', '$2a$10$46vcjpM2KOvc76hjcNb9NOgXsNKpWXR1b.tEXetZcWV0l4FQc8i5.', 'COMPANY_OWNER', true);

insert into company(id, owner_username, name, address, city, country, phone, email, website, description, year_established,
                    size, industry, is_approved) values (1, 'imbiamba', 'Code Mime Inc.', 'Bulevar oslobodjenja 18','Novi Sad',
                                            'Serbia', '066433235', 'codemime@gmail.com', 'www.codemime.com', 'Code Mime is a
                                            software development company that specializes in developing web and mobile
                                            applications.', '2000', '100-120', 'IT', true);
insert into salary(id, company_name, position, engagement, currently_employed, monthly_net_salary)
values (1, 'Code Mime Inc.', 'Software Engineer (Junior)', 'Full-time', true, 900);

insert into salary(id, company_name, position, engagement, currently_employed, monthly_net_salary)
values (2, 'Code Mime Inc.', 'Software Engineer (Junior)', 'Full-time', true, 800);

insert into salary(id, company_name, position, engagement, currently_employed, monthly_net_salary)
values (3, 'Code Mime Inc.', 'Software Developer (Junior)', 'Full-time', true, 700);

insert into salary(id, company_name, position, engagement, currently_employed, monthly_net_salary)
values (4, 'Code Mime Inc.', 'Software Developer (Junior)', 'Full-time', true, 650);





create  table Users (
	id serial primary key,
	name varchar(30) not null, 
	email varchar(100) not null,
	passwd varchar(30) not null
);

create table Orders(
	id serial primary key,
	user_id int,
	order_date date default null,
	total_amount float not null,
	FOREIGN KEY (user_id) REFERENCES Users (id) ON DELETE CASCADE
);


create table Products(
	id serial primary key,
	name varchar(100) not null,
	price float not null
);


create table OrderProducts(
orderid INT references Orders ON UPDATE CASCADE,
productid INT references Products ON UPDATE CASCADE,
PRIMARY KEY (orderid, productid)
);

CREATE INDEX idx_order_product
ON OrderProducts(orderid, productid);	

create index idx_name_users
on Users(name);

create index idx_name_producs
on Producs(name);



insert into Users (name, email, passwd)
     values ('alex', 'alex@mail.ru', '11111'),
            ('alex1', 'alex1@mail.ru', '11111'),
            ('alex2', 'alex2@mail.ru', '11111');

update Users set name='Pup' where name='alex1';
update Users set passwd='2222' where name='alex2';

select * from users;

delete from users where id=3;
delete from users where name=alex2;

INSERT into Products (name, price) 
        values ('cpu', 6000.59),
                ('memory', 2000.0),
                ('video_card', 20000);


select * from Products;

update Products set name='Riva64' WHERE name = 'video_card';
delete from Products where name='memory';



insert into Orders (user_id, order_date, total_amount)
        values (3, '2024-05-18', 20000),
                (1, "2024-06-17", 5000);

delete from Orders where id=1;
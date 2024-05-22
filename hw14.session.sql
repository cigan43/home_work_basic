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
        values (2, '2024-05-18', 20000)
                (1, "2024-06-17", 5000);



create table blogs(
                      bid char(20) primary key ,
                      title varchar(50) not null ,
                      author char(20) not null ,
                      ts timestamp,
                      content text
);

create table author(
                       aid char(20),
                       name varchar(10),
                       gender bool,
                       age int8,
                       contact char(20),
                       account varchar(20),
                       primary key (aid,account)
);


create table contact(
                        aid char(20),
                        tel char(20),
                        email varchar(50),
                        primary key (aid,email)
);


create table comment(
                        bid char(20),
                        aid char(20),
                        ts timestamp,
                        msg text,
                        primary key (bid,aid)
);


create table massage(
                        sendId char(20),
                        recvId char(20),
                        msg text,
                        ts timestamp,
                        primary key (sendId,recvId)
);

create table account(
                        account varchar(40) primary key ,
                        password varchar(50) not null,
                        banned bool
);

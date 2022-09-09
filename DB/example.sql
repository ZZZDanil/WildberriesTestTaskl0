drop table if exists order_table_cache;

create table order_table_cache(
    id serial primary key,
    data jsonb);
	
insert into order_table_cache(data) values ('{"test1": "a", "order_uid": "1"}');
select data->'order_uid' as id, data from order_table_cache;

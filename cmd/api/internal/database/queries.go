package database

var DBqueries = []string{
	"CREATE TABLE IF NOT EXISTS users (id serial4 NOT NULL, username text NOT NULL, email text NOT NULL, password  text NOT NULL,CONSTRAINT users_pkey PRIMARY KEY (id));",
	// "CREATE TABLE stock (id serial4 NOT NULL, stock_name varchar(6) NULL, actual_price numeric NULL, CONSTRAINT stock_pkey PRIMARY KEY (id));",
	// "CREATE TABLE wallet (user_id int4 NULL, stock_id int4 NULL, quantity int4 NULL, medium_price numeric NULL);",
	// "ALTER TABLE wallet ADD CONSTRAINT wallet_stock_id_fkey FOREIGN KEY (stock_id) REFERENCES public.stock(id);",
	// "ALTER TABLE wallet ADD CONSTRAINT wallet_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);",
}

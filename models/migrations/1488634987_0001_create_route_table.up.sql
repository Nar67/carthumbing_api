CREATE TABLE route (
		id serial PRIMARY KEY,
		user_uuid char(30) NOT NULL,
		origin point NOT NULL,
		destination point NOT NULL,
		selected_date date NOT NULL DEFAULT current_date,
		type char(30) NOT NULL
);
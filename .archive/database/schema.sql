CREATE TABLE users (
	id INTEGER PRIMARY KEY,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	role_id INTEGER NOT NULL,
	FOREIGN KEY(role_id) REFERENCES roles(id));

CREATE TABLE roles (
	id INTEGER PRIMARY KEY,
	role TEXT NOT NULL);

CREATE TABLE routes (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	status_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY(status_id) REFERENCES statuses(id),
	FOREIGN KEY(user_id) REFERENCES users(id));

CREATE TABLE statuses (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL);

CREATE TABLE points_of_interest (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	longitude DECIMAL(9,6),
	latitude DECIMAL(9,6));

CREATE TABLE types (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	icon TEXT);	

CREATE TABLE entries (
	id INTEGER PRIMARY KEY,
	content TEXT NOT NULL,
	hyperlink TEXT NOT NULL,
	type_id	INTEGER NOT NULL,
	FOREIGN KEY(type_id) REFERENCES types(id));

CREATE TABLE categories (
	id INTEGER PRIMARY KEY,
	category TEXT NOT NULL,
	icon TEXT NOT NULL,
	weight INTEGER NOT NULL,
	type_id INTEGER,
	FOREIGN KEY(type_id) REFERENCES types(id));

CREATE TABLE routes_points_of_interest (
	route_id INTEGER NOT NULL,
	point_of_interest_id INTEGER NOT NULL,
	FOREIGN KEY(route_id) REFERENCES routes(id),
	FOREIGN KEY(point_of_interest_id) REFERENCES points_of_interest(id));

CREATE TABLE entries_routes (
	entry_id INTEGER NOT NULL,
	route_id INTEGER NOT NULL,
	FOREIGN KEY(entry_id) REFERENCES entries(id),
	FOREIGN KEY(route_id) REFERENCES routes(id));

CREATE TABLE entries_points_of_interest (
	entry_id INTEGER NOT NULL,
	point_of_interest_id INTEGER NOT NULL,
	FOREIGN KEY(entry_id) REFERENCES entries(id),
	FOREIGN KEY(point_of_interest_id) REFERENCES points_of_interest(id));

CREATE TABLE categories_routes (
	category_id INTEGER NOT NULL,
	route_id INTEGER NOT NULL,
	FOREIGN KEY(category_id) REFERENCES categories(id),
	FOREIGN KEY(route_id) REFERENCES routes(id));

CREATE TABLE categories_points_of_interest (
	category_id INTEGER NOT NULL,
	point_of_interest_id INTEGER NOT NULL,
	FOREIGN KEY(category_id) REFERENCES categories(id),
	FOREIGN KEY(point_of_interest_id) REFERENCES points_of_interest(id));

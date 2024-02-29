CREATE TABLE authors (
	id int NOT NULL AUTO_INCREMENT,
	first_name varchar NOT NULL,
	last_name varchar,
	date_of_birth date
);


CREATE TABLE users (
	id int NOT NULL AUTO_INCREMENT,
	first_name varchar,
	email varchar NOT NULL,
	password_hash varchar NOT NULL
);


CREATE TABLE books (
	id int NOT NULL AUTO_INCREMENT,
	name varchar NOT NULL,
	author_id int NOT NULL,
	published_date date NOT NULL,
	num_pages int,
	cover_url varchar,

	PRIMARY KEY (book_id),
	FOREIGN KEY (author_id) REFERENCES authors(id)
);


CREATE TABLE read_list_items (
	id int NOT NULL AUTO_INCREMENT,
	book_id int NOT NULL,
	user_id int NOT NULL,

	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE reviews (
	id int NOT NULL AUTO_INCREMENT,

	description varchar NOT NULL,
	rating int NOT NULL,

	book_id int NOT NULL,
	user_id int NOT NULL,


	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE review_comments (
	id int NOT NULL AUTO_INCREMENT,

	comment varchar NOT NULL,

	user_id int NOT NULL,
	review_id int NOT NULL,

	FOREIGN KEY (user_id) REFERENCES users(id)
	FOREIGN KEY (review_id) REFERENCES reviews(id)
	
);

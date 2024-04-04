CREATE TABLE authors (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name varchar(255) NOT NULL,
	last_name varchar(255),
	date_of_birth date
);


CREATE TABLE users (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name varchar(255),
	email varchar(255) NOT NULL UNIQUE,
	password_hash varchar(255) NOT NULL
);

CREATE TABLE user_relationships (
	first_user_id int NOT NULL,
	second_user_id int NOT NULL,
	type ENUM('Blocked', 'Followed'),

	PRIMARY KEY (
		first_user_id,
		second_user_id
	),

	FOREIGN KEY (first_user_id) REFERENCES users(id),
	FOREIGN KEY (second_user_id) REFERENCES users(id)
);


CREATE TABLE books (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	name varchar(255) NOT NULL,
	author_id int NOT NULL,
	published_date date NOT NULL,
	num_pages int,
	cover_url varchar(255),

	FOREIGN KEY (author_id) REFERENCES authors(id)
);


CREATE TABLE read_list_items (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	book_id int NOT NULL,
	user_id int NOT NULL,

	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE reviews (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

	description varchar(255) NOT NULL,
	rating int NOT NULL,

	book_id int NOT NULL,
	user_id int NOT NULL,


	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE review_comments (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

	comment varchar(255) NOT NULL,

	user_id int NOT NULL,
	review_id int NOT NULL,

	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (review_id) REFERENCES reviews(id)
	
);

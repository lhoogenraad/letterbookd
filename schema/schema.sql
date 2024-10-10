CREATE TABLE authors (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name varchar(255) NOT NULL,
	last_name varchar(255),
	date_of_birth date
);


CREATE TABLE users (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	first_name varchar(255),
	last_name varchaer(255),
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
	synopsis text DEFAULT 'No synopsis available.',

	FOREIGN KEY (author_id) REFERENCES authors(id)
);


CREATE TABLE read_list_items (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
	book_id int NOT NULL,
	user_id int NOT NULL,

	status ENUM('Unread', 'Read'),

	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE read_list_items ADD UNIQUE unique_user_book(user_id, book_id);

CREATE TABLE reviews (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

	description text NOT NULL,
	rating int NOT NULL,

	book_id int NOT NULL,
	user_id int NOT NULL,

	FOREIGN KEY (book_id) REFERENCES books(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE reviews ADD UNIQUE user_book_review_unique (user_id, book_id);

CREATE TABLE review_comments (
	id int NOT NULL AUTO_INCREMENT PRIMARY KEY,

	comment text NOT NULL,

	user_id int NOT NULL,
	review_id int NOT NULL,

	archived boolean DEFAULT false,
	edited boolean DEFAULT false,

	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (review_id) REFERENCES reviews(id)
);



-- Insert default test data
INSERT INTO authors
(first_name, last_name, date_of_birth)
VALUES
('Julius', 'Caeser', '0001-01-01'),
('George', 'Martin', '1948-09-20'),
('Jane', 'Austin', '1775-12-16');

INSERT INTO books
(name, author_id, published_date, num_pages)
VALUES
('Commentarii de Bello Gallico', 1, '0001-01-01', 730),
('A Song of Ice and Fire', 2, '2009-12-01', 642),
('Pride and Prejudice', 3, '1813-01-28', 518);

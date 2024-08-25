
-- Create the Authors table
CREATE TABLE Authors (
    author_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Create the Categories table
CREATE TABLE Categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);


-- Create the Users table
CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(15),
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- Create the Books table with Tags as an array
CREATE TABLE Books (
    book_id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    author_id INT,
    category_id INT,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    published_date DATE,
    available_copies INT DEFAULT 0,
    tags TEXT[] DEFAULT '{}', -- Store tags as an array of text
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    image_url TEXT, -- Stores the image URL
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (author_id) REFERENCES Authors(author_id),
    FOREIGN KEY (category_id) REFERENCES Categories(category_id)
);


-- Create the Leases table
CREATE TABLE Leases (
    lease_id SERIAL PRIMARY KEY,
    user_id INT,
    book_id INT,
    lease_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP,
    is_returned BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (book_id) REFERENCES Books(book_id)
);


-- Optional: Trigger function to update UpdatedAt column
-- CREATE OR REPLACE FUNCTION update_timestamp()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     NEW.UpdatedAt = CURRENT_TIMESTAMP;
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- Optional: Create triggers for Users and Books tables
-- CREATE TRIGGER update_users_timestamp
-- BEFORE UPDATE ON Users
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- CREATE TRIGGER update_books_timestamp
-- BEFORE UPDATE ON Books
-- FOR EACH ROW
-- EXECUTE FUNCTION update_timestamp();

-- Sample Inserts into the Authors table (optional)
INSERT INTO Authors (Name) VALUES
('Frank Herbert'),
('Isaac Asimov');

-- Sample Inserts into the Categories table (optional)
INSERT INTO Categories (Name) VALUES
('Science Fiction'),
('Fantasy');

-- Sample Inserts into the Books table with Tags
INSERT INTO Books (Title, AuthorID, CategoryID, ISBN, PublishedDate, AvailableCopies, Tags) VALUES
('Dune', 1, 1, '9780441013593', '1965-08-01', 3, ARRAY['Science Fiction', 'Classic']),
('Foundation', 2, 1, '9780553293357', '1951-06-01', 5, ARRAY['Science Fiction', 'Adventure']);

-- Sample Inserts into the Users table
INSERT INTO Users (FirstName, LastName, Email, PhoneNumber) VALUES
('John', 'Doe', 'john.doe@example.com', '123-456-7890'),
('Jane', 'Smith', 'jane.smith@example.com', '987-654-3210');

-- Sample Inserts into the Leases table
INSERT INTO Leases (UserID, BookID, LeaseDate) VALUES
(1, 1, CURRENT_TIMESTAMP),
(2, 2, CURRENT_TIMESTAMP);

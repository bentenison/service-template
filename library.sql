-- Create the Authors table (optional)
CREATE TABLE Authors (
    AuthorID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL
);

-- Create the Categories table (optional)
CREATE TABLE Categories (
    CategoryID SERIAL PRIMARY KEY,
    Name VARCHAR(50) NOT NULL
);

-- Create the Users table
CREATE TABLE Users (
    UserID SERIAL PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    PhoneNumber VARCHAR(15),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the Books table with Tags as an array
CREATE TABLE Books (
    BookID SERIAL PRIMARY KEY,
    Title VARCHAR(100) NOT NULL,
    AuthorID INT,
    CategoryID INT,
    ISBN VARCHAR(20) UNIQUE NOT NULL,
    PublishedDate DATE,
    AvailableCopies INT DEFAULT 0,
    Tags TEXT[] DEFAULT '{}', -- Store tags as an array of text
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (AuthorID) REFERENCES Authors(AuthorID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- Create the Leases table
CREATE TABLE Leases (
    LeaseID SERIAL PRIMARY KEY,
    UserID INT,
    BookID INT,
    LeaseDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ReturnDate TIMESTAMP,
    IsReturned BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (BookID) REFERENCES Books(BookID)
);

-- Optional: Trigger function to update UpdatedAt column
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.UpdatedAt = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Optional: Create triggers for Users and Books tables
CREATE TRIGGER update_users_timestamp
BEFORE UPDATE ON Users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_books_timestamp
BEFORE UPDATE ON Books
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

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

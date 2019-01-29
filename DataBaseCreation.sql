show databases;
CREATE DATABASE IF NOT EXISTS mercadolibre;
use mercadolibre;

CREATE TABLE IF NOT EXISTS Classifications (
    Class_ID INT AUTO_INCREMENT,
    Class_Description VARCHAR(255) NOT NULL,
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    PRIMARY KEY (Class_ID)
)  ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS Owners (
    Owner_ID VARCHAR(255) NOT NULL,
    Owner_Name VARCHAR(255),
	Owner_Email VARCHAR(255),
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    PRIMARY KEY (Owner_ID)
)  ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS DBase (
    Base_ID INT AUTO_INCREMENT,
    Base_Name VARCHAR(255) NOT NULL,
    Base_ClassID INT NOT NULL,
    Base_OwnerID VARCHAR(255) NOT NULL,
    Base_Time DATETIME,
	FOREIGN KEY fk_class(Base_ClassID)
	REFERENCES Classifications(Class_ID),
	FOREIGN KEY fk_owner(Base_OwnerID)
	REFERENCES Owners(Owner_ID),
    status TINYINT NOT NULL,
    priority TINYINT NOT NULL,
    description TEXT,
    PRIMARY KEY (Base_ID)
)  ENGINE=INNODB;
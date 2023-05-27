CREATE TABLE `users` (
                         `id` INT primary key AUTO_INCREMENT,
                         `name` VARCHAR(191) NOT NULL,
                         `phone_number` VARCHAR(191) NOT NULL UNIQUE,
                         `email` VARCHAR(191) NOT NULL UNIQUE,
                         `password` varchar(191),
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
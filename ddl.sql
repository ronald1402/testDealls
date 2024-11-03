CREATE TABLE `users` (
     `id` int NOT NULL AUTO_INCREMENT,
     `name` varchar(255) NOT NULL,
     `email` varchar(255) NOT NULL,
     `username` varchar(255) NOT NULL,
     `hashed_password` varchar(255) NOT NULL,
     PRIMARY KEY (`id`),
     UNIQUE KEY `email` (`email`),
     UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `profile_likes` (
     `id` int NOT NULL AUTO_INCREMENT,
     `user_id` int NOT NULL,
     `liked_user_id` int NOT NULL,
     `status` enum('like','pass') NOT NULL,
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     KEY `user_id` (`user_id`),
     KEY `liked_user_id` (`liked_user_id`),
     KEY `idx_created_at_desc` (`created_at` DESC),
     CONSTRAINT `profile_likes_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
     CONSTRAINT `profile_likes_ibfk_2` FOREIGN KEY (`liked_user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `user_profiles` (
     `profile_id` int NOT NULL AUTO_INCREMENT,
     `user_id` int NOT NULL,
     `profile_picture` varchar(255) DEFAULT NULL,
     `bio` text,
     `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`profile_id`),
     KEY `user_id` (`user_id`),
     CONSTRAINT `user_profiles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
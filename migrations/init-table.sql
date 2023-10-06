CREATE TABLE `vote_options` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) UNIQUE NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `user_question_votes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `question_id` int NOT NULL,
  `vote_option_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `questions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `body` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `question_statistics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `question_id` int UNIQUE NOT NULL,
  `thumbs_up` int NOT NULL,
  `thumbs_down` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `user_answer_votes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `answer_id` int NOT NULL,
  `vote_option_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `answers` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `questions_id` int NOT NULL,
  `user_id` int NOT NULL,
  `body` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `answer_statistics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `answer_id` int UNIQUE NOT NULL,
  `thumbs_up` int NOT NULL,
  `thumbs_down` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `user_comment_votes` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `comment_id` int NOT NULL,
  `vote_option_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `comments` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `answer_id` int NOT NULL,
  `user_id` int NOT NULL,
  `body` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `comment_statistics` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `comment_id` int UNIQUE NOT NULL,
  `thumbs_up` int NOT NULL,
  `thumbs_down` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `parent_child_comments` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `parent_comment_id` int NOT NULL,
  `child_comment_id` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE UNIQUE INDEX `user_question_votes_index_0` ON `user_question_votes` (`user_id`, `question_id`);

CREATE UNIQUE INDEX `user_answer_votes_index_1` ON `user_answer_votes` (`user_id`, `answer_id`);

CREATE UNIQUE INDEX `user_comment_votes_index_2` ON `user_comment_votes` (`user_id`, `comment_id`);

CREATE UNIQUE INDEX `parent_child_comments_index_3` ON `parent_child_comments` (`parent_comment_id`, `child_comment_id`);

ALTER TABLE `user_question_votes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_question_votes` ADD FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_question_votes` ADD FOREIGN KEY (`vote_option_id`) REFERENCES `vote_options` (`id`) ON DELETE CASCADE;

ALTER TABLE `questions` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `question_statistics` ADD FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_answer_votes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_answer_votes` ADD FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_answer_votes` ADD FOREIGN KEY (`vote_option_id`) REFERENCES `vote_options` (`id`) ON DELETE CASCADE;

ALTER TABLE `answers` ADD FOREIGN KEY (`questions_id`) REFERENCES `questions` (`id`) ON DELETE CASCADE;

ALTER TABLE `answers` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `answer_statistics` ADD FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_comment_votes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_comment_votes` ADD FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE;

ALTER TABLE `user_comment_votes` ADD FOREIGN KEY (`vote_option_id`) REFERENCES `vote_options` (`id`) ON DELETE CASCADE;

ALTER TABLE `comments` ADD FOREIGN KEY (`answer_id`) REFERENCES `answers` (`id`) ON DELETE CASCADE;

ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

ALTER TABLE `comment_statistics` ADD FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE;

ALTER TABLE `parent_child_comments` ADD FOREIGN KEY (`parent_comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE;

ALTER TABLE `parent_child_comments` ADD FOREIGN KEY (`child_comment_id`) REFERENCES `comments` (`id`) ON DELETE CASCADE;
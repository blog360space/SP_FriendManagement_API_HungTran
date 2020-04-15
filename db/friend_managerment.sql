CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL,
  `text` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` time DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1;

CREATE TABLE `relationships` (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user1_id` int(10) unsigned NOT NULL DEFAULT '0',
  `user2_id` int(10) unsigned NOT NULL DEFAULT '0',
  `subscribe` smallint(5) unsigned NOT NULL DEFAULT '0',
  `friend_status` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '0: friend request, 1: friend: 2 block',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

INSERT INTO `users` VALUES
   (1,'andy@example.com',NULL,NULL,NULL),
   (2,'john@example.com',NULL,NULL,NULL),
   (3,'hung.tran@example.com',NULL,NULL,NULL),
   (4,'lisa@example.com',NULL,NULL,NULL),
   (5,'common@example.com',NULL,NULL,NULL),
   (6,'kate@example.com',NULL,NULL,NULL);

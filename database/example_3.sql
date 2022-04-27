-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.11-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping structure for table golang.colors
CREATE TABLE IF NOT EXISTS `colors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) NOT NULL,
  `color` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.colors: ~4 rows (approximately)
DELETE FROM `colors`;
INSERT INTO `colors` (`id`, `product_id`, `color`) VALUES
	(11, 1, 'Blue'),
	(12, 1, 'Red'),
	(13, 2, 'Green'),
	(15, 2, 'Black');

-- Dumping structure for table golang.messages
CREATE TABLE IF NOT EXISTS `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `email` varchar(150) NOT NULL,
  `message` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.messages: ~3 rows (approximately)
DELETE FROM `messages`;
INSERT INTO `messages` (`id`, `name`, `email`, `message`) VALUES
	(1, 'ww', 'ww@ww.com', 'wwwwwww'),
	(2, 'ee', 'ee@ee.com', 'eeee'),
	(3, 'qq', 'qq@qq.com', 'qq');

-- Dumping structure for table golang.pages
CREATE TABLE IF NOT EXISTS `pages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.pages: ~1 rows (approximately)
DELETE FROM `pages`;
INSERT INTO `pages` (`id`, `name`, `content`) VALUES
	(1, 'home', '<p>\r\n<strong>page content from database table [pages]</strong></br> \r\npage content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] - page content from database table [pages] -  \r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p>\r\n<p>\r\nLorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur justo dolor, sagittis vitae elit vel, vehicula dignissim purus. Etiam ut erat interdum arcu ultrices egestas eu fringilla sem. Donec pulvinar fringilla quam, vitae tristique ipsum venenatis eu. Curabitur egestas porttitor dolor, quis suscipit massa euismod eget. Nullam velit ligula, rutrum at pellentesque sit amet, maximus quis leo. Ut laoreet ex sem, vel posuere nisl viverra et. Nullam suscipit volutpat metus. Sed mattis aliquet commodo. Pellentesque cursus elit eget est blandit, ac posuere risus posuere. Nunc lobortis tortor ante, eu sagittis sapien molestie non. Pellentesque ultricies orci at enim pulvinar, et rutrum turpis egestas.\r\n</p> ');

-- Dumping structure for table golang.products
CREATE TABLE IF NOT EXISTS `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(50) NOT NULL,
  `type` varchar(50) NOT NULL,
  `name` varchar(250) NOT NULL,
  `description` text NOT NULL,
  `price` float DEFAULT NULL,
  `images` text DEFAULT NULL,
  `status` text DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.products: ~2 rows (approximately)
DELETE FROM `products`;
INSERT INTO `products` (`id`, `code`, `type`, `name`, `description`, `price`, `images`, `status`) VALUES
	(1, 'F-150', 'Truck', 'Ford F-1500', 'The redesigned 2021 Ford F-150 claims one of the top spots in our full-size pickup truck rankings because of its tremendous capability and spacious, comfortable cabin', 29990, NULL, 'avail'),
	(2, 'F-250', 'Truck', 'Ford F-2500', 'The all new 2022 Ford Super Duty F-250 Crew Cab comes fully ready with interior comforts, advanced technology, and exterior conveniences to make every trip you decide to go on in this powerful 4-door Pickup so enjoyable! With an available 7.3L gas V8 engine, the Super Duty F-250 Crew Cab can easily cut through any terrain and get even the toughest of jobs done! Discover how much its massive available engine can do for you. ', 35990, NULL, NULL);

-- Dumping structure for table golang.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstname` varchar(50) NOT NULL,
  `lastname` varchar(50) NOT NULL,
  `username` text NOT NULL,
  `password` text NOT NULL,
  `auth_hash` text NOT NULL,
  `expiration` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table golang.users: ~1 rows (approximately)
DELETE FROM `users`;
INSERT INTO `users` (`id`, `firstname`, `lastname`, `username`, `password`, `auth_hash`, `expiration`) VALUES
	(8, 'k', 'd', 'admin@example.com', '$2a$08$9L4boKq4wEE8hC1Kx5bxmu79nmg/oJkIFlHWIVv8GUlliUHztoBr6', '8b218dfa5c4bce705410f7b0911e66fcd9f7965d080118fa49b1542e7dcdea407f8a9b350bd5dd5e4e3032df490173213689088e4dfe9b34c1f199cf340ab5f0', '2022-04-27 01:58:54');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;

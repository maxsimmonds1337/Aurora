-- MySQL dump 10.13  Distrib 8.0.32, for Linux (aarch64)
--
-- Host: localhost    Database: aurora
-- ------------------------------------------------------
-- Server version	8.0.32

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `baby_logs`
--

DROP TABLE IF EXISTS `baby_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `baby_logs` (
  `log_id` int NOT NULL AUTO_INCREMENT,
  `baby_id` int NOT NULL,
  `time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `activities` set('peed','pooped','vomited','had_blood') DEFAULT NULL,
  `color` varchar(7) DEFAULT NULL,
  `breast_milk_time` decimal(10,2) DEFAULT NULL,
  `breast_milk_mls` decimal(10,2) DEFAULT NULL,
  `formula_milk_mls` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `baby_logs`
--

LOCK TABLES `baby_logs` WRITE;
/*!40000 ALTER TABLE `baby_logs` DISABLE KEYS */;
INSERT INTO `baby_logs` VALUES (20,1,'2023-03-28 15:10:01','','',20.00,0.00,0.00),(21,1,'2023-03-28 15:20:41','peed','',0.00,0.00,0.00),(22,1,'2023-03-28 15:39:59','','',0.00,33.00,0.00),(23,1,'2023-03-28 19:15:44','peed','',0.00,0.00,0.00),(24,1,'2023-03-28 20:16:27','','',40.00,0.00,0.00),(25,1,'2023-03-28 20:34:54','peed','',0.00,0.00,0.00),(26,1,'2023-03-28 21:16:26','','',0.00,0.00,70.00),(27,1,'2023-03-28 21:26:51','','',0.00,0.00,30.00),(28,1,'2023-03-28 22:39:23','peed','',0.00,0.00,0.00),(29,1,'2023-03-28 22:49:01','','',5.00,0.00,0.00),(31,1,'2023-03-28 17:50:00','','',0.00,0.00,60.00),(32,1,'2023-03-28 20:10:00','peed','',0.00,0.00,0.00),(33,1,'2023-03-28 22:59:38','','',5.00,0.00,0.00),(34,1,'2023-03-28 23:15:43','peed','',0.00,0.00,0.00),(35,1,'2023-03-28 23:27:38','','',8.00,0.00,0.00),(36,1,'2023-03-28 23:46:14','','',0.00,0.00,60.00),(38,1,'2023-03-29 00:13:57','peed','',0.00,0.00,0.00),(39,1,'2023-03-29 01:56:45','peed','',0.00,0.00,0.00),(41,1,'2023-03-29 07:41:35','peed','',0.00,60.00,10.00),(42,1,'2023-03-29 10:13:54','','',10.00,0.00,0.00),(43,1,'2023-03-29 10:49:20','peed','',0.00,0.00,0.00),(44,1,'2023-03-29 11:26:00','','',34.00,0.00,0.00),(45,1,'2023-03-29 12:03:58','peed','',0.00,0.00,30.00),(46,1,'2023-03-29 12:17:19','','',0.00,0.00,0.00),(47,1,'2023-03-29 12:17:26','','',0.00,0.00,50.00),(48,1,'2023-03-29 14:06:02','peed','',0.00,0.00,0.00),(49,1,'2023-03-29 14:41:49','peed,pooped','',0.00,0.00,0.00),(50,1,'2023-03-29 15:01:16','','',0.00,30.00,60.00),(51,1,'2023-03-29 15:16:37','','',0.00,0.00,55.00),(52,1,'2023-03-29 18:32:58','peed','',0.00,0.00,60.00),(53,1,'2023-03-29 19:02:46','peed,pooped','',0.00,35.00,0.00),(54,1,'2023-03-29 19:55:21','peed','',0.00,0.00,35.00),(55,1,'2023-03-29 23:02:31','peed','',20.00,0.00,0.00),(56,1,'2023-03-29 23:10:46','peed','',0.00,0.00,0.00),(57,1,'2023-03-29 23:18:37','','',0.00,0.00,70.00),(58,1,'2023-03-29 23:36:22','','',0.00,0.00,50.00),(59,1,'2023-03-30 00:05:41','peed','',0.00,0.00,0.00),(60,1,'2023-03-30 00:30:42','peed','',0.00,0.00,0.00),(61,1,'2023-03-30 06:50:13','','',0.00,0.00,80.00),(62,1,'2023-03-30 07:30:48','peed','',0.00,0.00,0.00),(63,1,'2023-03-30 09:47:49','','',0.00,80.00,0.00),(64,1,'2023-03-30 10:45:55','peed','',0.00,0.00,0.00),(65,1,'2023-03-30 14:08:42','peed','',0.00,0.00,0.00),(66,1,'2023-03-30 14:34:18','','',0.00,0.00,90.00),(67,1,'2023-03-30 17:44:41','peed','',0.00,55.00,0.00),(68,1,'2023-03-30 20:29:50','peed','',0.00,0.00,0.00),(69,1,'2023-03-31 06:00:00','peed','',0.00,0.00,0.00),(70,1,'2023-03-31 06:30:00','','',21.00,0.00,0.00),(71,1,'2023-03-30 23:45:00','','',0.00,55.00,0.00),(72,1,'2023-03-31 00:01:00','peed','',0.00,0.00,0.00),(73,1,'2023-03-30 22:35:00','peed','',0.00,0.00,0.00),(74,1,'2023-03-30 22:26:00','','',0.00,0.00,60.00),(75,1,'2023-03-30 22:20:00','peed','',0.00,0.00,0.00),(76,1,'2023-03-30 20:45:00','','',0.00,0.00,65.00),(77,1,'2023-03-30 12:09:00','','',0.00,0.00,90.00),(78,1,'2023-03-30 15:17:00','peed','',0.00,0.00,0.00),(79,1,'2023-03-30 15:28:00','','',0.00,0.00,65.00),(80,1,'2023-03-30 15:46:00','peed','',0.00,0.00,0.00),(81,1,'2023-03-30 04:15:00','peed','',0.00,0.00,0.00);
/*!40000 ALTER TABLE `baby_logs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-04-05 20:48:40

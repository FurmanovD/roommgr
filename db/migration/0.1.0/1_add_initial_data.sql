
INSERT INTO `companies` (`id`, `name`) VALUES
(1, 'Administration'),
(2, 'PEPSI'),
(3, 'COKE')
ON DUPLICATE KEY UPDATE id=id;


INSERT INTO `users` (`id`, `userid`, `company_id`, `password`, `group_id`, `first_name`, `last_name`) VALUES 
(1, 'admin', 1, 'somepass', 1, 'admin', 'admin'),
(2, 'Joe', 2, 'someanotherpass', 2, 'John', 'Doe'),
(3, 'Sam2002', 3, 'someevenmoreanotherpass', 2, 'Samantha', 'Fox')
ON DUPLICATE KEY UPDATE id=id;


INSERT INTO `rooms` (`id`, `name`, `company_id`) VALUES
(1,'P01', 2),
(2,'P02', 2),
(3,'P03', 2),
(4,'P04', 2),
(5,'P05', 2),
(6,'P06', 2),
(7,'P07', 2),
(8,'P08', 2),
(9,'P09', 2),
(10,'P10', 2),
(11,'C01', 3),
(12,'C02', 3),
(13,'C03', 3),
(14,'C04', 3),
(15,'C05', 3),
(16,'C06', 3),
(17,'C07', 3),
(18,'C08', 3),
(19,'C09', 3),
(20,'C10', 3)
ON DUPLICATE KEY UPDATE id=id;

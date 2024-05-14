CREATE TABLE `user` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `deleted_at` TIMESTAMP NULL,
                        `address` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `contract_address` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `name` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `symbol` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
                        `total_supply` bigint(20) NOT NULL,
                        `decimals` bigint(20) NOT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_address` (`address`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
# Day 4 - April 10, 2025

## ✅ What I Learned

-How to connect golang to database with gorm

## ✍️ Example Database

1.  Create Database
    Create database golang

2.  Create Tabel Inside Database
    CREATE TABLE `users` (

        `id` int(11) NOT NULL,
        `nama_user` varchar(50) NOT NULL,
        `last` varchar(50) NOT NULL,
        `handle` varchar(50) NOT NULL
         )

3.  Insert Data Inside Table
    INSERT INTO `users` (`id`, `nama_user`, `last`, `handle`) VALUES
    (1, 'anam', 'Otto', '@mdo'),
    (2, 'Budi', 'Thornton', '@fat'),
    (3, 'Larry', 'Bird', '@twitter');

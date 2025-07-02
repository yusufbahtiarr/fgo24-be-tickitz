INSERT INTO profile (fullname, phone)
VALUES
('yusuf bahtiar', '089553546443'),
('handi daniel', '089346446443'),
('gandi samuel', '089584578673'),
('rian kurniawan', '089678456768'),
('thomas zino', '089123435345'),
('fani utami', '089586657575'),
('rani mutiara', '089685678786'),
('bagas hendrawan', '089123412554'),
('cyntia agustin', '089474567456'),
('caca lunar', '089235346345');

INSERT INTO users (email, password, role, id_profile)
VALUES
('yusuf@mail.com','7b21848ac9af35be0ddb2d6b9fc3851934db8420','admin',1),
('handi@mail.com','1a9b9508b6003b68ddfe03a9c8cbc4bd4388339b','admin',2),
('gandi@mail.com','403d9917c3e950798601addf7ba82cd3c83f344b','user',3),
('rian@mail.com','3e511da7577d1864871b760ab30e05b56943c9b2','user',4),
('thomas@mail.com','69df79bef9287d3bcb8f104a408b06de6a108fd8','user',5),
('fani@mail.com','62a013722e0619aef2edee368768cd07b6529655','user',6),
('rani@mail.com','d559965849921585c1849af03b7a51638700d979','user',7),
('bagas@mail.com','9eab102e8f9431bb23016851d11e658e0b20b730','user',8),
('cyntia@mail.com','a045b7efa463c6ed195c644163f4168952fbd34a','user',9),
('caca@mail.com','6934105ad50010b814c933314b1da6841431bc8b','user',10);


INSERT INTO casts (cast_name)
VALUES
('Reza Rahadian'),
('Vino G. Bastian'),
('Nicholas Saputra'),
('Joe Taslim'),
('Iko Uwais'),
('Abimana Aryasatya'),
('Refal Hady'),
('Anya Geraldine'),
('Adhisty Zara'),
('Dian Sastrowardoyo'),
('Luna Maya'),
('Pevita Pearce'),
('Tatjana Saphira'),
('Prilly Latuconsina'),
('Mawar Eva de Jongh'),
('Aghniny Haque'),
('Tissa Biani'),
('Rachel Amanda'),
('Bella Graceva'),
('Raline Shah'),
('Iqbaal Ramadhan'),
('Ari Irham'),
('Morgan Oey'),
('Bryan Domani'),
('Angga Yunanda'),
('Umay Shahab');

INSERT INTO directors (director_name)
VALUES
('Joko Anwar'),
('Angga Dwimas Sasongko'),
('Riri Riza'),
('Hanung Bramantyo'),
('Mouly Surya'),
('Garin Nugroho'),
('Kamila Andini'),
('Timo Tjahjanto'),
('Gina S. Noer'),
('Ernest Prakasa'),
('Yandy Laurens'),
('Yosep Anggi Noen'),
('Wregas Bhanuteja'),
('Usmar Ismail'),
('Nia Dinata'),
('Upi Avianto'),
('Asrul Sani');

INSERT INTO genres (genre_name)
VALUES
('Horor'),
('Drama'),
('Komedi'),
('Romantis'),
('Aksi'),
('Kriminal'),
('Fantasi'),
('Keluarga'),
('Drama Musikal'),
('Indie'),
('Sejarah');

INSERT INTO cinemas (cinema_name, image_url)
VALUES 
('Ebv', '../assets/images/ebv.png'),
('Hiflix', '../assets/images/hiflix.png'),
('CineOne21', '../assets/images/cineone21.png'),
('XXI', '../assets/images/xxi.png');

INSERT INTO locations (location)
VALUES
('Jakarta'),
('Depok'),
('Bogor'),
('Bekasi'),
('Tangerang');

INSERT INTO times (time)
VALUES
('13:15'),
('15:20'),
('17:25'),
('19:30'),
('21:35');

INSERT INTO payment_methods (payment_method, image_url)
VALUES
('Google Pay', '../assets/images/gpay.png'),
('Visa', '../assets/images/visa.png'),
('Gopay', '../assets/images/gopay.png'),
('Paypal', '../assets/images/paypal.png'),
('Dana', '../assets/images/dana.png'),
('BCA', '../assets/images/bca.png'),
('BRI', '../assets/images/bri.png'),
('OVO', '../assets/images/ovo.png');

INSERT INTO movies (poster_url, backdrop_url, title, release_date, runtime, overview, rating, id_director)
VALUES
('/qZDjcqvgt3xWcAdc7wHpuqGZ3i.jpg','/lOje1iz4VYWELYWRkZAwI7oIJd0.jpg','Operasi Improvisasi', '2025-06-12', 110, 'Film komedi aksi bertempo cepat yang dibintangi Bryce Dallas Howard sebagai Kat, seorang guru improvisasi komedi yang mempertanyakan apakah peluang suksesnya terlewatkan. Saat polisi yang menyamar menawarkan peluang sangat langka, Kat merekrut dua muridnya untuk menyusup ke dunia gangster London dengan samaran penjahat berbahaya.',6.8, 2),
('/4akm5FQdZFfNpYGgVyOEXi72eby.jpg','/aQ5nvQGT6mM6TliOM5iSgrKVF4C.jpg','Echo Valley', '2025-06-12', 95, 'Julianne Moore adalah Kate, yang kehidupan tenangnya diganggu oleh putrinya, ketika dia muncul berlumuran darah orang. Ketika Kate mengungkap kebenaran yang mengejutkan, dia mendapati sejauh mana dia akan bertindak demi anaknya.',6.4, 3),
('/t3cmnXYtxJb9vVL1ThvT2CWSe1n.jpg','/wnnu8htEZBLtwrke9QYfLKx6zjp.jpg','STRAW', '2025-06-12', 115, 'Apa pemicu terakhirnya? Satu hari yang sangat buruk membuat seorang ibu tunggal pekerja keras akhirnya merasa muak—dan melakukan tindakan putus asa yang mengejutkan.',6.5, 3),
('/kRbpUTRNm6QbLQFPFWUcNC4czEm.jpg','/ozVwXlfxqNsariipatGwa5px3Pm.jpg','Lucy', '2014-07-25', 105, 'Seorang perempuan, yang tidak sengaja terjebak dalam urusan kelam, membalikkan keadaan dengan penyanderanya dan berubah menjadi pejuang tanpa ampun yang menjelma di luar logika manusia.',6.8, 2),
('/dq18nCTTLpy9PmtzZI6Y2yAgdw5.jpg','/qAZ0pzat24kLdO3o8ejmbLxyOac.jpg','Black Widow', '2021-07-07', 110, 'Dalam film ini akan diceritakan tentang Natasha Romanoff dan Yelena Belova yang merupakan saudara. Pada tahun 1995, saat Natasha dan Yelena masih muda, mereka dibawa pergi oleh orang tua pengganti mereka—Alexei Shostakov yang merupakan seorang tentara atau dikenal sebagai Red Guardian dan Melina Vostokoff sebagai agen rahasia serta mantan Black Widow. Keduanya merupakan agen rahasia yang menyamar sebagai pasangan suami istri untuk menjalankan misi. Mereka kabur dari rumahnya di Ohio setelah menyelesaikan misi mereka.',7.2, 4),
('/wQubJeuxeQRGSrrRCEmxnHMS1xh.jpg','/wQubJeuxeQRGSrrRCEmxnHMS1xh.jpg','Immaculate', '2024-03-20', 100, 'Cecilia, seorang wanita yang taat iman, disambut hangat di pedesaan Italia yang sempurna di mana dia ditawari peran baru di sebuah biara terkenal. Tetapi menjadi jelas bagi Cecilia bahwa rumah barunya menyimpan rahasia yang gelap dan mengerikan.',7.0, 5),
('/dmo6TYuuJgaYinXBPjrgG9mB5od.jpg','/uQ4lG7E7mlyKsGvbASftQ6Hu2IX.jpg','Yang Terakhir Dari Kita', '2023-01-15', 110, 'Dua puluh tahun setelah peradaban modern dihancurkan, Joel, seorang penyintas yang tangguh, disewa untuk menyelundupkan Ellie, seorang gadis berusia 14 tahun, keluar dari zona karantina yang menindas. Apa yang dimulai sebagai pekerjaan kecil segera menjadi perjalanan yang brutal dan memilukan, karena mereka berdua harus melintasi Amerika Serikat dan bergantung satu sama lain untuk bertahan hidup.',7.3, 3),
('/cQvc9N6JiMVKqol3wcYrGshsIdZ.jpg','/z7noaCJ4KtmhwHw7QcNtnMMo4Qy.jpg','The Equalizer 2', '2018-07-19', 100, 'Robert McCall seorang mantan agen rahasia selalu berusaha memberikan keadilan bagi orang yang membutuhkannya. Hingga suatu ketika ia mengetahui bahwa sahabat lamanya, Susan Plummer menjadi korban pembunuhan yang dilakukan musuhnya. Bagi McCall, ini adalah panggilan perang. Kini ia mencari tahu siapa yang telah membunuh temannya itu untuk membalaskan dendam.',7.0, 7),
('/eyWICPcxOuTcDDDbTMOZawoOn8d.jpg','/2uSCHUsmzb6KkQPFSxBQ7bgfJLE.jpg','The Meg', '2018-08-09', 105, 'Lima tahun lalu, seorang ahli penyelam laut dan juga mantan Kapten Angkatan Laut bernama Jonas Taylor (Jason Statham) menghadapi marabahaya yang tidak terduga di relung Palung Mariana yang memaksanya untuk membatalkan misinya dan juga meninggalkan crew nya.',6.5, 8),
('/M7SUK85sKjaStg4TKhlAVyGlz3.jpg','/70AV2Xx5FQYj20labp0EGdbjI6E.jpg','Wrath of Man', '2021-04-22', 115, 'Wrath of Man bercerita tentang seorang pria bernama H (Jason Statham), yang bekerja di perusahaan truk tunai dengan tanggung jawab memindahkan uang dengan jumlah banyak di sekitar Los Angeles, setiap minggu. H merupakan pria yang misterius.  H adalah anggota baru dalam perusahaan tersebut. Suatu hari, H berhadapan dengan para perampok yang menghadang mobil truk berisi uang dalam jumlah banyak. Karena kemampuan yang dimilikinya, H mampu mengalahkan para perampok tersebut seorang diri.',6.7, 7);

INSERT INTO movie_casts(id_movie,id_cast)
VALUES
(1, 2),
(1, 3),
(1, 5),
(2, 3),
(2, 7),
(2, 8),
(3, 7),
(3, 12),
(3, 17),
(4, 16),
(4, 12),
(4, 7),
(5, 14),
(5, 16),
(5, 6),
(6, 12),
(6, 9),
(6, 6),
(7, 9),
(7, 17),
(7, 15),
(8, 18),
(8, 7),
(8, 5),
(9, 12),
(9, 22),
(9, 15),
(10, 18),
(10, 21),
(10, 20);

INSERT INTO movie_genres(id_movie, id_genre)
VALUES
(1, 3),
(1, 6),
(1, 9),
(2, 5),
(2, 6),
(2, 10),
(3, 11),
(3, 10),
(3, 1),
(4, 10),
(4, 11),
(4, 7),
(5, 8),
(5, 9),
(5, 6),
(6, 10),
(6, 9),
(6, 6),
(7, 9),
(7, 7),
(7, 5),
(8, 8),
(8, 7),
(8, 5),
(9, 8),
(9, 10),
(9, 7),
(10, 8),
(10, 2),
(10, 6);

INSERT INTO transactions (name, email, phone, virtual_account, total_payment, movie_date, id_movie, id_cinema, id_time, id_location, id_payment_method, id_user)
VALUES
('Yusuf Bahtiar', 'yusuf@mail.com', '089533452345', '089735324534563456456345', 180000, '2025-06-24', 1, 2, 1, 1, 6, 1),
('Aisyah Sari', 'aisyah.sari@gmail.com', '081234567890', '123456789012345678901234', 60000, '2025-06-25', 1, 1, 2, 2, 1, 2),
('Bima Santoso', 'bima.santoso@yahoo.com', '087788990011', '456789012345678901234567', 120000, '2025-06-26', 3, 3, 3, 3, 2, 3),
('Citra Lestari', 'citra.lestari@outlook.com', '082199887766', '789012345678901234567890', 120000, '2025-06-27', 4, 4, 4, 4, 3, 4),
('Dedi Kurniawan', 'dedi.kurniawan@mail.com', '085566778899', '234567890123456789012345', 180000, '2025-06-28', 5, 2, 5, 5, 4, 5),
('Eka Putri', 'eka.putri@gmail.com', '081122334455', '567890123456789012345678', 120000, '2025-06-29', 6, 1, 1, 1, 5, 6),
('Fajar Nugroho', 'fajar.nugroho@yahoo.com', '089977665544', '890123456789012345678901', 120000, '2025-06-30', 7, 3, 2, 2, 7, 7),
('Gita Rahayu', 'gita.rahayu@outlook.com', '083344556677', '345678901234567890123456', 180000, '2025-06-24', 8, 4, 3, 3, 8, 8),
('Hadi Wibowo', 'hadi.wibowo@mail.com', '087711223344', '678901234567890123456789', 240000, '2025-06-25', 9, 2, 4, 4, 1, 9),
('Indah Permata', 'indah.permata@gmail.com', '081988776655', '901234567890123456789012', 120000, '2025-06-26', 10, 1, 5, 5, 2, 10),
('Joko Susilo', 'joko.susilo@yahoo.com', '085522334466', '123456789012345678901235', 180000, '2025-06-27', 1, 3, 1, 1, 3, 1),
('Kania Dewi', 'kania.dewi@outlook.com', '082211445566', '456789012345678901234568', 180000, '2025-06-28', 2, 4, 2, 2, 4, 2),
('Lukman Hakim', 'lukman.hakim@mail.com', '089633445577', '789012345678901234567891', 120000, '2025-06-29', 3, 2, 3, 3, 5, 3),
('Mira Amalia', 'mira.amalia@gmail.com', '081344556688', '234567890123456789012346', 180000, '2025-06-30', 4, 1, 4, 4, 6, 4),
('Nanda Pratama', 'nanda.pratama@yahoo.com', '087755667799', '567890123456789012345679', 120000, '2025-06-24', 5, 3, 5, 5, 7, 5),
('Oki Setiawan', 'oki.setiawan@outlook.com', '085588779900', '890123456789012345678902', 180000, '2025-06-25', 6, 4, 1, 1, 8, 6),
('Putri Anggraini', 'putri.anggraini@mail.com', '081677889011', '345678901234567890123457', 240000, '2025-06-26', 7, 2, 2, 2, 1, 7),
('Rudi Hartono', 'rudi.hartono@gmail.com', '082299001122', '678901234567890123456790', 120000, '2025-06-27', 8, 1, 3, 3, 2, 8),
('Sari Wulandari', 'sari.wulandari@yahoo.com', '089511223344', '901234567890123456789013', 180000, '2025-06-28', 9, 3, 4, 4, 3, 9),
('Teguh Santoso', 'teguh.santoso@outlook.com', '087822334455', '123456789012345678901236', 120000, '2025-06-29', 10, 4, 5, 5, 4, 10),
('Umi Kalsum', 'umi.kalsum@mail.com', '081933445566', '456789012345678901234569', 180000, '2025-06-30', 1, 2, 1, 1, 5, 1),
('Vino Ramadhan', 'vino.ramadhan@gmail.com', '085644556677', '789012345678901234567892', 120000, '2025-06-24', 2, 1, 2, 2, 6, 2),
('Wulan Sari', 'wulan.sari@yahoo.com', '082155667788', '234567890123456789012347', 180000, '2025-06-25', 3, 3, 3, 3, 7, 3),
('Xaverius Tan', 'xaverius.tan@outlook.com', '089766778899', '567890123456789012345680', 60000, '2025-06-26', 4, 4, 4, 4, 8, 4),
('Yanti Puspita', 'yanti.puspita@mail.com', '081377889900', '890123456789012345678903', 60000, '2025-06-27', 5, 2, 5, 5, 1, 5),
('Zainal Abidin', 'zainal.abidin@gmail.com', '087988990011', '345678901234567890123458', 120000, '2025-06-28', 6, 1, 1, 1, 2, 6),
('Amira Syifa', 'amira.syifa@yahoo.com', '085299001122', '678901234567890123456791', 120000, '2025-06-29', 7, 3, 2, 2, 3, 7),
('Budi Pranoto', 'budi.pranoto@outlook.com', '082411223344', '901234567890123456789014', 180000, '2025-06-30', 8, 4, 3, 3, 4, 8),
('Candra Wijaya', 'candra.wijaya@mail.com', '089522334455', '123456789012345678901237', 180000, '2025-06-24', 9, 2, 4, 4, 5, 9),
('Dewi Lestari', 'dewi.lestari@gmail.com', '081633445566', '456789012345678901234570', 120000, '2025-06-25', 10, 1, 5, 5, 6, 10);

INSERT INTO transaction_details (id_transaction, seat)
VALUES
(1, 'A1'), (1, 'A2'), (1, 'A3'),
(2, 'A4'),
(3, 'A5'), (3, 'A6'),
(4, 'A7'), (4, 'A8'),
(5, 'A9'), (5, 'A10'), (5, 'A11'),
(6, 'A12'), (6, 'A13'),
(7, 'A14'), (7, 'B1'),
(8, 'B2'), (8, 'B3'), (8, 'B4'),
(9, 'B5'), (9, 'B6'), (9, 'B7'), (9, 'B8'),
(10, 'B9'), (10, 'B10'),
(11, 'B11'), (11, 'B12'), (11, 'B13'),
(12, 'B14'), (12, 'C1'), (12, 'C2'),
(13, 'C3'), (13, 'C4'),
(14, 'C5'), (14, 'C6'), (14, 'C7'),
(15, 'C8'), (15, 'C9'),
(16, 'C10'), (16, 'C11'), (16, 'C12'),
(17, 'C13'), (17, 'C14'), (17, 'D1'), (17, 'D2'),
(18, 'D3'), (18, 'D4'),
(19, 'D5'), (19, 'D6'), (19, 'D7'),
(20, 'D8'), (20, 'D9'),
(21, 'D10'), (21, 'D11'), (21, 'D12'),
(22, 'D13'), (22, 'D14'),
(23, 'E1'), (23, 'E2'), (23, 'E3'),
(24, 'E4'),
(25, 'E5'),
(26, 'E6'), (26, 'E7'),
(27, 'E8'), (27, 'E9'),
(28, 'E10'), (28, 'E11'), (28, 'E12'),
(29, 'E13'), (29, 'E14'), (29, 'F1'),
(30, 'F2'), (30, 'F3');


--SELECT
SELECT t.id, m.title, t.name, to_char(t.movie_date, 'DD-MM-YYYY') AS date, to_char(tm.time,'HH24:MI') AS time, c.cinema_name, STRING_AGG(td.seat,', ') AS seats, t.total_payment from transactions t 
JOIN transaction_details td ON td.id_transaction = t.id
JOIN movies m ON m.id = t.id_movie
JOIN times tm ON tm.id = t.id_time
JOIN cinemas c ON c.id = t.id_cinema
GROUP BY t.id, m.title, t.name, t.transaction_date, tm.time, c.cinema_name, t.total_payment;

SELECT m.id, m.title,
  ( SELECT STRING_AGG(c.cast_name, ', ') 
    FROM movie_casts mc  
    JOIN casts c ON c.id = mc.id_cast 
    WHERE mc.id_movie = m.id) 
  AS cast_names,
  
  ( SELECT STRING_AGG(g.genre_name, ', ') 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre_names,
  m.overview,
  m.runtime
FROM movies m;


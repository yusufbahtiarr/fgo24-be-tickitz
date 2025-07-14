INSERT INTO profiles (fullname, phone)
VALUES
('admin', '089553546443'),
('yusuf', '089346446443'),
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
('admin@mail.com', '$2a$10$XH0U4WX9yQn8dJkzLb5rE.9qOY7w6vV1mWc1ZJd3Kf4rN5s6G7h8I', 'admin', 1),
('yusuf@mail.com', '$2a$10$YI1J2K3L4M5N6O7P8Q9R0S1T2U3V4W5X6Y7Z8A9B0C1D2E3F4G5H6', 'admin', 2),
('gandi@mail.com', '$2a$10$Z9A8B7C6D5E4F3G2H1I0J9K8L7M6N5O4P3Q2R1S0T9U8V7W6X5Y4Z3', 'user', 3),
('rian@mail.com', '$2a$10$1A2B3C4D5E6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7', 'user', 4),
('thomas@mail.com', '$2a$10$2B3C4D5E6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8', 'user', 5),
('fani@mail.com', '$2a$10$3C4D5E6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8B9', 'user', 6),
('rani@mail.com', '$2a$10$4D5E6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8B9C0', 'user', 7),
('bagas@mail.com', '$2a$10$5E6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8B9C0D1', 'user', 8),
('cyntia@mail.com', '$2a$10$6F7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8B9C0D1E2', 'user', 9),
('caca@mail.com', '$2a$10$7G8H9I0J1K2L3M4N5O6P7Q8R9S0T1U2V3W4X5Y6Z7A8B9C0D1E2F3', 'user', 10);

INSERT INTO casts (cast_name) VALUES
('Jackie Chan'),
('Ben Wang'),
('Ralph Macchio'),
('Joshua Jackson'),
('Leonardo DiCaprio'),
('Brad Pitt'),
('Tom Hanks'),
('Johnny Depp'),
('Robert Downey Jr.'),
('Chris Hemsworth'),
('Chris Evans'),
('Tom Holland'),
('Ryan Reynolds'),
('Dwayne Johnson'),
('Keanu Reeves'),
('Christian Bale'),
('Matt Damon'),
('Will Smith'),
('Hugh Jackman'),
('Cillian Murphy'),
('Joaquin Phoenix'),
('Daniel Radcliffe'),
('Benedict Cumberbatch'),
('Michael Fassbender'),
('Morgan Freeman'),
('Jake Gyllenhaal'),
('Robert Pattinson'),
('Timothée Chalamet'),
('Adam Driver'),
('Chadwick Boseman'),
('Mark Ruffalo'),
('Andrew Garfield'),
('Tobey Maguire'),
('Henry Cavill'),
('Paul Rudd'),
('Jason Momoa'),
('Tom Hardy'),
('Idris Elba'),
('Anthony Hopkins'),
('Sylvester Stallone'),
('Arnold Schwarzenegger'),
('Al Pacino'),
('Robert De Niro'),
('Javier Bardem'),
('Rami Malek'),
('Oscar Isaac'),
('Mahershala Ali'),
('Dev Patel'),
('Eddie Redmayne'),
('Matthew McConaughey'),
('Ewan McGregor'),
('Jeremy Renner'),
('Forest Whitaker'),
('Sam Worthington'),
('Meryl Streep'),
('Scarlett Johansson'),
('Natalie Portman'),
('Emma Watson'),
('Jennifer Lawrence'),
('Angelina Jolie'),
('Anne Hathaway'),
('Kate Winslet'),
('Nicole Kidman'),
('Sandra Bullock'),
('Charlize Theron'),
('Emma Stone'),
('Margot Robbie'),
('Gal Gadot'),
('Zendaya'),
('Florence Pugh'),
('Emily Blunt'),
('Viola Davis'),
('Amy Adams'),
('Saoirse Ronan'),
('Rachel McAdams'),
('Kristen Stewart'),
('Lupita Nyong'),
('Millie Bobby Brown'),
('Tilda Swinton'),
('Penélope Cruz'),
('Cate Blanchett');

INSERT INTO directors (director_name)
VALUES
('Steven Spielberg'),
('Christopher Nolan'),
('Martin Scorsese'),
('Quentin Tarantino'),
('James Cameron'),
('Peter Jackson'),
('Stanley Kubrick'),
('Ridley Scott'),
('Denis Villeneuve'),
('Alfonso Cuarón'),
('Guillermo del Toro'),
('Bong Joon-ho'),
('Wes Anderson'),
('Tim Burton'),
('Greta Gerwig'),
('Sofia Coppola'),
('David Fincher'),
('George Lucas'),
('Hayao Miyazaki'),
('Ang Lee'),
('Jonathan Entwistle'),
('Michael Bay'),
('Len Wiseman'),
('Robert Zemeckis'),
('Ron Howard'),
('Frank Darabont'),
('Spike Lee'),
('Mel Gibson'),
('Taika Waititi'),
('Jean-Pierre Jeunet'),
('Pedro Almodóvar'),
('Ken Loach'),
('Danny Boyle'),
('Lars von Trier'),
('Yorgos Lanthimos'),
('Akira Kurosawa'),
('Park Chan-wook'),
('Kim Jee-woon'),
('Hirokazu Kore-eda'),
('Satoshi Kon'),
('Luca Guadagnino'),
('Gabriele Muccino'),
('Jean-Luc Godard'),
('Luc Besson'),
('Michel Gondry'),
('Alejandro González Iñárritu'),
('Fernando Meirelles'),
('Walter Salles'),
('Andrei Tarkovsky'),
('Fatih Akin');

INSERT INTO genres (genre_name)
VALUES
  ('Action'),
  ('Adventure'),
  ('Animation'),
  ('Comedy'),
  ('Crime'),
  ('Documentary'),
  ('Drama'),
  ('Family'),
  ('Fantasy'),
  ('History'),
  ('Horror'),
  ('Music'),
  ('Mystery'),
  ('Romance'),
  ('Science Fiction'),
  ('TV Movie'),
  ('Thriller'),
  ('War'),
  ('Western');

INSERT INTO cinemas (cinema_name, image_url)
VALUES 
('Ebv', '/assets/ebv.png'),
('Hiflix', '/assets/hiflix.png'),
('CineOne21', '/assets/cineone21.png'),
('XXI', '/assets/xxi.png');

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
('Google Pay', '/assets/gpay.png'),
('Visa', '/assets/visa.png'),
('Gopay', '/assets/gopay.png'),
('Paypal', '/assets/paypal.png'),
('Dana', '/assets/dana.png'),
('BCA', '/assets/bca.png'),
('BRI', '/assets/bri.png'),
('OVO', '/assets/ovo.png');

INSERT INTO movies (poster_url, backdrop_url, title, release_date, runtime, overview, rating)
VALUES
('https://image.tmdb.org/t/p/w500/wqfu3bPLJaEWJVk3QOm0rKhxf1A.jpg', 'https://image.tmdb.org/t/p/original/x58Gk2ZGU5AEBp25MQe2nhZhd5z.jpg', 'The Old Guard 2', '2025-07-01', 121, 'Andy and her team of immortal warriors fight with renewed purpose as they face a powerful new foe threatening their mission to protect humanity.', 6.2),
('https://image.tmdb.org/t/p/w500/2VUmvqsHb6cEtdfscEA6fqqVzLg.jpg', 'https://image.tmdb.org/t/p/original/sItIskd5xpiE64bBWYwZintkGf3.jpg', 'Ballerina', '2025-06-04', 95, 'Taking place during the events of John Wick: Chapter 3 – Parabellum, Eve Macarro begins her training in the assassin traditions of the Ruska Roma.', 7.4),
('https://image.tmdb.org/t/p/w500/AEgggzRr1vZCLY86MAp93li43z.jpg', 'https://image.tmdb.org/t/p/original/nKyBbFSooRPTJVqjrDteD1lF733.jpg', 'Karate Kid: Legends', '2025-05-08', 123, 'After a family tragedy, kung fu prodigy Li Fong is uprooted from his home in Beijing and forced to move to New York City with his mother. When a new friend needs his help, Li enters a karate competition – but his skills alone aren''t enough. Li''s kung fu teacher Mr. Han enlists original Karate Kid Daniel LaRusso for help, and Li learns a new way to fight, merging their two styles into one for the ultimate martial arts showdown.', 7.2),
('https://image.tmdb.org/t/p/w500/2vHQBX5L4yoExTa55KmGIg2Q5s3.jpg', 'https://image.tmdb.org/t/p/original/962KXsr09uK8wrmUg9TjzmE7c4e.jpg', 'Ice Road: Vengeance', '2025-06-27', 112, 'Big rig ice road driver Mike McCann travels to Nepal to scatter his late brother’s ashes on Mt. Everest. While on a packed tour bus traversing the deadly 12,000 ft. terrain of the infamous Road to the Sky, McCann and his mountain guide encounter a group of mercenaries and must fight to save themselves, the busload of innocent travelers, and the local villagers’ homeland.', 7.0),
('https://image.tmdb.org/t/p/w500/lVgE5oLzf7ABmzyASEVcjYyHI41.jpg', 'https://image.tmdb.org/t/p/original/xABhldZaMb6wfCH5oigV333OYnb.jpg', 'Heads of State', '2025-07-02', 111, 'The UK Prime Minister and US President have a public rivalry that risks their countries'' alliance. But when they become targets of a powerful enemy, they''re forced to rely on each other as they go on a wild, multinational run. Allied with Noel, a brilliant MI6 agent, they must find a way to thwart a conspiracy that threatens the free world.', 7.0),
('https://image.tmdb.org/t/p/w500/q0fGCmjLu42MPlSO9OYWpI5w86I.jpg', 'https://image.tmdb.org/t/p/original/peAzdLKtT6VDWIfPQO9LJD1NCG4.jpg', 'Jurassic World Rebirth', '2025-07-01', 121, 'Five years after the events of Jurassic World Dominion, covert operations expert Zora Bennett is contracted to lead a skilled team on a top-secret mission to secure genetic material from the world''s three most massive dinosaurs. When Zora''s operation intersects with a civilian family whose boating expedition was capsized, they all find themselves stranded on an island where they come face-to-face with a sinister, shocking discovery that''s been hidden from the world for decades.', 6.5),
('https://image.tmdb.org/t/p/w500/ombsmhYUqR4qqOLOxAyr5V8hbyv.jpg', 'https://image.tmdb.org/t/p/original/g62G6aBcAcJv3ClCKmJgmHarHvq.jpg', 'Superman', '2025-07-09', 114, 'Superman, a journalist in Metropolis, embarks on a journey to reconcile his Kryptonian heritage with his human upbringing as Clark Kent.', 7.3),
('https://image.tmdb.org/t/p/w500/r3d6u2n7iPoWNsSWwlJJWrDblOH.jpg', 'https://image.tmdb.org/t/p/original/9A0wQG38VdEu3DYh8HzXKXKhA6g.jpg', 'Dora and the Search for Sol Dorado', '2025-07-02', 112, 'Dora, Diego, and their new friends trek through the perilous dangers of the Amazonian jungle in search of the ancient and powerful treasure of Sol Dorado to keep it out of enemy hands.', 7.1),
('https://image.tmdb.org/t/p/w500/tUae3mefrDVTgm5mRzqWnZK6fOP.jpg', 'https://image.tmdb.org/t/p/original/7Zx3wDG5bBtcfk8lcnCWDOLM4Y4.jpg', 'Lilo & Stitch', '2025-05-17', 110, 'The wildly funny and touching story of a lonely Hawaiian girl and the fugitive alien who helps to mend her broken family.', 7.1),
('https://image.tmdb.org/t/p/w500/tObSf1VzzHt9xB0csanFtb3DRjf.jpg', 'https://image.tmdb.org/t/p/original/5esDYWV0NoFwqPa1iC0g9akqZo9.jpg', 'Bring Her Back', '2025-05-28', 100, 'Following the death of their father, a brother and sister are introduced to their new sibling by their foster mother, only to learn that she has a terrifying secret.', 7.3),
('https://image.tmdb.org/t/p/w500/43c1efKzA1kigNzY0HBzeoXp8LR.jpg', 'https://image.tmdb.org/t/p/original/l3ycQYwWmbz7p8otwbomFDXIEhn.jpg', 'KPop Demon Hunters', '2025-06-20', 108, 'When K-pop superstars Rumi, Mira and Zoey aren''t selling out stadiums, they''re using their secret powers to protect their fans from supernatural threats.', 8.6),
('https://image.tmdb.org/t/p/w500/3lwlJL8aW6Wor9tKvME8VoMnBkn.jpg', 'https://image.tmdb.org/t/p/original/7HqLLVjdjhXS0Qoz1SgZofhkIpE.jpg', 'How to Train Your Dragon', '2025-06-06', 103, 'On the rugged isle of Berk, where Vikings and dragons have been bitter enemies for generations, Hiccup stands apart, defying centuries of tradition when he befriends Toothless, a feared Night Fury dragon. Their unlikely bond reveals the true nature of dragons, challenging the very foundations of Viking society.', 7.9),
('https://image.tmdb.org/t/p/w500/vqBmyAj0Xm9LnS1xe1MSlMAJyHq.jpg', 'https://image.tmdb.org/t/p/original/lkDYN0whyE82mcM20rwtwjbniKF.jpg', 'F1', '2025-06-25', 105, 'Racing legend Sonny Hayes is coaxed out of retirement to lead a struggling Formula 1 team—and mentor a young hotshot driver—while chasing one more chance at glory.', 7.7),
('https://image.tmdb.org/t/p/w500/ktqPs5QyuF8SpKnipvVHb3fwD8d.jpg', 'https://image.tmdb.org/t/p/original/yAqL0makiGke5yYiVWpmBDSKIVP.jpg', 'The Ritual', '2025-04-18', 98, 'Two priests, one in crisis with his faith and the other confronting a turbulent past, must overcome their differences to perform a risky exorcism.', 6.0),
('https://image.tmdb.org/t/p/w500/yqsCU5XOP2mkbFamzAqbqntmfav.jpg', 'https://image.tmdb.org/t/p/original/nAxGnGHOsfzufThz20zgmRwKur3.jpg', 'Sinners', '2025-04-16', 99, 'Trying to leave their troubled lives behind, twin brothers return to their hometown to start again, only to discover that an even greater evil is waiting to welcome them back.', 7.6),
('https://image.tmdb.org/t/p/w500/wnHUip9oKvDJEJUEk62L4rFSYGa.jpg', 'https://image.tmdb.org/t/p/original/6WqqEjiycNvDLjbEClM1zCwIbDD.jpg', '28 Years Later', '2025-06-18', 120, 'Twenty-eight years since the rage virus escaped a biological weapons laboratory, now, still in a ruthlessly enforced quarantine, some have found ways to exist amidst the infected...', 7.2),
('https://image.tmdb.org/t/p/w500/yHYj2t3NzsjkHVYa0bPri4sMBsg.jpg', 'https://image.tmdb.org/t/p/original/tdMbbFhqyEqOK1QzNTvJjHWKbZX.jpg', 'Ziam', '2025-07-08', 131, 'In a fight for survival against a horrifying army of zombies, a former Muay Thai fighter must use skill, speed and grit to save his girlfriend.', 6.7),
('https://image.tmdb.org/t/p/w500/uFQduVyYIinJy3eLjozgfl6Xtcn.jpg', 'https://image.tmdb.org/t/p/original/fPWJn5pqBr8n4h0YxW3QuasdvoI.jpg', 'Diablo', '2025-06-13', 113, 'Ex-con Kris Chaney seizes the daughter of a Colombian gangster to fulfill a noble promise to the young girl''s mother...', 7.4),
('https://image.tmdb.org/t/p/w500/4a63rQqIDTrYNdcnTXdPsQyxVLo.jpg', 'https://image.tmdb.org/t/p/original/kyBOGOBUMdGWOhzECuosPSzoMpi.jpg', 'M3GAN 2.0', '2025-06-25', 116, 'After the underlying tech for M3GAN is stolen and misused by a powerful defense contractor to create a military-grade weapon known as Amelia, M3GAN''s creator Gemma realizes that the only option is to resurrect M3GAN and give her a few upgrades...', 7.2),
('https://image.tmdb.org/t/p/w500/t3cmnXYtxJb9vVL1ThvT2CWSe1n.jpg', 'https://image.tmdb.org/t/p/original/wnnu8htEZBLtwrke9QYfLKx6zjp.jpg', 'STRAW', '2025-06-05', 117, 'What will be her last straw? A devastatingly bad day pushes a hardworking single mother to the breaking point — and into a shocking act of desperation.', 7.9),
('https://image.tmdb.org/t/p/w500/i1eFKvcUNdyNd53YeEOK2vNkBsL.jpg', 'https://image.tmdb.org/t/p/original/k3oqTinhaXFKxlY09oYxnSks9R0.jpg', 'Misericordia', '2024-10-16', 96, 'Jérémie returns to his hometown for the funeral of his former boss, the village baker. He decides to stay for a few days with Martine, the man''s widow. A mysterious disappearance, a threatening neighbor and a priest with strange intentions make Jérémie''s short stay in the village take an unexpected turn.', 6.3),
('https://image.tmdb.org/t/p/w500/x26MtUlwtWD26d0G0FXcppxCJio.jpg', 'https://image.tmdb.org/t/p/original/4b2mLlMMaMcZhvNvmPgwaFqme8Y.jpg', 'The Fantastic Four: First Steps', '2025-07-23', 118, 'Against the vibrant backdrop of a 1960s-inspired, retro-futuristic world, Marvel''s First Family is forced to balance their roles as heroes with the strength of their family bond, while defending Earth from a ravenous space god called Galactus and his enigmatic Herald, Silver Surfer.', 7.2),
('https://image.tmdb.org/t/p/w500/jtEqpy0K1iVuCebRwWqT6BZ6jLN.jpg', 'https://image.tmdb.org/t/p/original/w3RDV3pSpxN0C2DZ4Xpw4o5LWpI.jpg', 'The Phoenician Scheme', '2025-05-23', 105, 'Wealthy businessman Zsa-zsa Korda appoints his only daughter, a nun, as sole heir to his estate. As Korda embarks on a new enterprise, they soon become the target of scheming tycoons, foreign terrorists, and determined assassins.', 6.805),
('https://image.tmdb.org/t/p/w500/juA4IWO52Fecx8lhAsxmDgy3M3.jpg', 'https://image.tmdb.org/t/p/original/icFWIk1KfkWLZnugZAJEDauNZ94.jpg', 'Until Dawn', '2025-04-23', 97, 'One year after her sister Melanie mysteriously disappeared, Clover and her friends head into the remote valley where she vanished in search of answers. Exploring an abandoned visitor center, they find themselves stalked by a masked killer and horrifically murdered one by one...only to wake up and find themselves back at the beginning of the same evening.', 6.601),
('https://image.tmdb.org/t/p/w500/2VLnT5JtJ77ErG6GD3c1J628VLp.jpg', 'https://image.tmdb.org/t/p/original/21lM9hfYNDX9RnWzp4aq0iQTEbq.jpg', 'given the Movie: To the Sea', '2024-09-20', 101, 'Mafuyu doubts his place in the band, but a familiar face could be the spark that brings him back to the stage.', 8.333),
('https://image.tmdb.org/t/p/w500/wVujUVvY4qvKARAslItQ4ARKqtz.jpg', 'https://image.tmdb.org/t/p/original/dJ1JpHuTCRxtGpRq8bQzk2jBqTi.jpg', 'The Legend of Ochi', '2025-04-18', 108, 'In a remote village on the island of Carpathia, a shy farm girl named Yuri is raised to fear an elusive animal species known as ochi. But when Yuri discovers a wounded baby ochi has been left behind, she escapes on a quest to bring him home.', 6.4),
('https://image.tmdb.org/t/p/w500/eDo0pNruy0Qgj6BdTyHIR4cxHY8.jpg', 'https://image.tmdb.org/t/p/original/lqwwGkwJHtz9QgKtz4zeY19YgDg.jpg', 'Materialists', '2025-06-12', 100, 'A young, ambitious New York City matchmaker finds herself torn between the perfect match and her imperfect ex.', 7.0),
('https://image.tmdb.org/t/p/w500/q0bCG4NX32iIEsRFZqRtuvzNCyZ.jpg', 'https://image.tmdb.org/t/p/original/hwlyY7LJdEFbCPaGNXiskKKmJ5X.jpg', 'Flight Risk', '2025-01-22', 113, 'A U.S. Marshal escorts a government witness to trial after he''s accused of getting involved with a mob boss, only to discover that the pilot who is transporting them is also a hitman sent to assassinate the informant. After they subdue him, they''re forced to fly together after discovering that there are others attempting to eliminate them.', 6.053),
('https://image.tmdb.org/t/p/w500/8Vt6mWEReuy4Of61Lnj5Xj704m8.jpg', 'https://image.tmdb.org/t/p/original/kVd3a9YeLGkoeR50jGEXM6EqseS.jpg', 'Spider-Man: Across the Spider-Verse', '2023-05-31', 120, 'After reuniting with Gwen Stacy, Brooklyn’s full-time, friendly neighborhood Spider-Man is catapulted across the Multiverse, where he encounters the Spider Society, a team of Spider-People charged with protecting the Multiverse’s very existence. But when the heroes clash on how to handle a new threat, Miles finds himself pitted against the other Spiders and must set out on his own to save those he loves most.', 8.3),
('https://image.tmdb.org/t/p/w500/8o6lkhL32xQJeB52IIG1us5BVey.jpg', 'https://image.tmdb.org/t/p/original/nZFgO3BDxpLzTpDXhYbdXBGdUVm.jpg', 'Smurfs', '2025-06-30', 98, 'When Papa Smurf is mysteriously taken by evil wizards, Razamel and Gargamel, Smurfette leads the Smurfs on a mission into the real world to save him. With the help of new friends, the Smurfs must discover what defines their destiny to save the universe.', 7.0),
('https://image.tmdb.org/t/p/w500/aFRDH3P7TX61FVGpaLhKr6QiOC1.jpg', 'https://image.tmdb.org/t/p/original/63RnEJRNzRDvi0Fv4jhrT2eh0fl.jpg', 'Demon Slayer: Kimetsu no Yaiba Infinity Castle', '2025-07-18', 115, 'As the Demon Slayer Corps members and Hashira engaged in a group strength training program, the Hashira Training, in preparation for the forthcoming battle against the demons, Muzan Kibutsuji appears at the Ubuyashiki Mansion. With the head of the Demon Corps in danger, Tanjiro and the Hashira rush to the headquarters but are plunged into a deep descent to a mysterious space by the hands of Muzan Kibutsuji. The destination of where Tanjiro and Demon Slayer Corps have fallen is the demons'' stronghold – the Infinity Castle. And so, the battleground is set as the final battle between the Demon Slayer Corps and the demons ignites.', 6.8),
('https://image.tmdb.org/t/p/w500/xYduFGuch9OwbCOEUiamml18ZoB.jpg', 'https://image.tmdb.org/t/p/original/P82NAcEsLIYgQtrtn36tYsj41m.jpg', 'The Garfield Movie', '2024-04-30', 99, 'Garfield, the world-famous, Monday-hating, lasagna-loving indoor cat, is about to have a wild outdoor adventure! After an unexpected reunion with his long-lost father – scruffy street cat Vic – Garfield and his canine friend Odie are forced from their perfectly pampered life into joining Vic in a hilarious, high-stakes heist.', 7.058),
('https://image.tmdb.org/t/p/w500/lxM6kqilAdpdhqUl2biYp5frUxE.jpg', 'https://image.tmdb.org/t/p/original/xgU3KkqiME9pGe5gGCNpUYkoSWg.jpg', 'Jaws', '1975-06-20', 115, 'When the seaside community of Amity finds itself under attack by a dangerous great white shark, the town''s chief of police, a young marine biologist, and a grizzled hunter embark on a desperate quest to destroy the beast before it strikes again.', 7.673),
('https://image.tmdb.org/t/p/w500/iRCgqpdVE4wyLQvGYU3ZP7pAtUc.jpg', 'https://image.tmdb.org/t/p/original/zViRwl3ySscZnbXZJ2Q9wq3SeUG.jpg', 'Transformers One', '2024-09-11', 102, 'The untold origin story of Optimus Prime and Megatron, better known as sworn enemies, but once were friends bonded like brothers who changed the fate of Cybertron forever.', 8.1),
('https://image.tmdb.org/t/p/w500/mFVdwAnHCUg49S8Q1XFrJotPyIF.jpg', 'https://image.tmdb.org/t/p/original/qra80oA1Aw1KXj1JS46GLK4m3eE.jpg', 'The Bad Guys 2', '2025-07-25', 107, 'The now-reformed Bad Guys are trying (very, very hard) to be good, but instead find themselves hijacked into a high-stakes, globe-trotting heist, masterminded by a new team of criminals they never saw coming: The Bad Girls.', 7.0),
('https://image.tmdb.org/t/p/w500/woaN8CbloH0akyX0E72ayxlJAB4.jpg', 'https://image.tmdb.org/t/p/original/jl2YIADk391yc6Qjy9JhgCRkHJk.jpg', 'Memoir of a Snail', '2024-10-17', 118, 'Forcibly separated from her twin brother when they are orphaned, a melancholic misfit learns how to find confidence within herself amid the clutter of misfortunes and everyday life.', 8.069),
('https://image.tmdb.org/t/p/w500/9wV65OmsjLAqBfDnYTkMPutXH8j.jpg', 'https://image.tmdb.org/t/p/original/dcLV5rEXuQRW0ZlB7IMLArHMyWh.jpg', 'Freakier Friday', '2025-08-06', 107, 'Years after Tess and Anna endured an identity crisis, Anna now has a daughter of her own and a soon-to-be stepdaughter. As they navigate the myriad challenges that come when two families merge, Tess and Anna discover lightning might indeed strike twice.', 0.0),
('https://image.tmdb.org/t/p/w500/A06yXys3hrCWu8xiNoHCFLTG5SH.jpg', 'https://image.tmdb.org/t/p/original/djkxDBzBEL3xShrtpJP17L8V2U3.jpg', 'I Know What You Did Last Summer', '2025-07-16', 113, 'When five friends inadvertently cause a deadly car accident, they cover up their involvement and make a pact to keep it a secret rather than face the consequences. A year later, their past comes back to haunt them and they''re forced to confront a horrifying truth: someone knows what they did last summer…and is hell-bent on revenge.', 10.0),
('https://image.tmdb.org/t/p/w500/tjBLr7wAVFaOcdUpFRKX9wKQmfT.jpg', 'https://image.tmdb.org/t/p/original/7A3ru3bxmWtyF0ep0CJcLLzWT35.jpg', 'Legends of the Condor Heroes: The Gallants', '2025-01-29', 103, 'Under Genghis Khan, the Mongolian army pushes west to destroy the Jin Dynasty, setting its sights on the Song Dynasty next. Amid internal conflicts among martial arts schools, Guo Jing unites the Central Plains'' warriors to defend Xiangyang, embodying courage and loyalty in the fight for the nation.', 6.1),
('https://image.tmdb.org/t/p/w500/iGCtYxfuvXfy0BD5m6p7vKuPOxS.jpg', 'https://image.tmdb.org/t/p/original/A7KwHKB5faRRRRveVQC5TPI35kV.jpg', 'Young Hearts', '2024-10-26', 120, '14-year-old Elias increasingly feels like an outsider in his village. When he meets his new neighbour of the same age, Alexander, Elias is confronted with his burgeoning sexuality.', 8.35),
('https://image.tmdb.org/t/p/w500/v313aUGmMNj6yNveaiQXysBmjVS.jpg', 'https://image.tmdb.org/t/p/original/qSOMdbZ6AOdHR999HWwVAh6ALFI.jpg', 'Alarum', '2025-01-16', 95, 'Two married spies caught in the crosshairs of an international intelligence network will stop at nothing to obtain a critical asset. Joe and Lara are agents living off the grid whose quiet retreat at a winter resort is blown to shreds when members of the old guard suspect the two may have joined an elite team of rogue spies, known as Alarum.', 5.81),
('https://image.tmdb.org/t/p/w500/3RokBWEkQqJRZEVP3DPwGm5MYh6.jpg', 'https://image.tmdb.org/t/p/original/p5b5Y5l7qJyqzGqk1OhQ2qPX328.jpg', 'Guns Up', '2025-05-15', 110, 'When a job goes horribly wrong, an ex-cop and family man who moonlights as a mob henchman has one night to get his family out of the city.', 6.868),
('https://image.tmdb.org/t/p/w500/hZ8dTeBzigV5SVgwG1ikSROAFiS.jpg', 'https://image.tmdb.org/t/p/original/eT1L3IVcHUL58wLn48E5dxlea3Z.jpg', 'Presence', '2025-01-17', 99, 'A couple and their children move into a seemingly normal suburban home. When strange events occur, they begin to believe there is something else in the house with them. The presence is about to disrupt their lives in unimaginable ways.', 6.031),
('https://image.tmdb.org/t/p/w500/9tk3Si960hg4E49eMt81dS7Qe9Z.jpg', 'https://image.tmdb.org/t/p/original/baAFsaYx5dZvb4g8NyE7iDh8p6q.jpg', 'Dangerous Animals', '2025-06-05', 102, 'A savvy and free-spirited surfer is abducted by a shark-obsessed serial killer. Held captive on his boat, she must figure out how to escape before he carries out a ritualistic feeding to the sharks below.', 7.0),
('https://image.tmdb.org/t/p/w500/vRCXxDdAQPjPrJgxKyDyxkb0dDt.jpg', 'https://image.tmdb.org/t/p/original/wKVHfpDFlD5jYWbpptTQdKnr3mO.jpg', 'Tornado', '2025-05-30', 106, 'When her father''s puppet samurai show gets ambushed by a notorious gang, Tornado vows to seek vengeance and forge her own destiny by stealing their ill-gotten gold.', 6.1),
('https://image.tmdb.org/t/p/w500/u7Lp1Hi8aBS73jv4KRMIv5aK4ax.jpg', 'https://image.tmdb.org/t/p/original/uUJp5I4IbzuhdUiEx4R9OAoFpbz.jpg', 'Saw X', '2023-09-27', 117, 'Between the events of ''Saw'' and ''Saw II'', a sick and desperate John Kramer travels to Mexico for a risky and experimental medical procedure in hopes of a miracle cure for his cancer, only to discover the entire operation is a scam to defraud the most vulnerable. Armed with a newfound purpose, the infamous serial killer returns to his work, turning the tables on the con artists in his signature visceral way through devious, deranged, and ingenious traps.', 7.212),
('https://image.tmdb.org/t/p/w500/2C1yy7lPYAYvbtW8bDckGGdPhmt.jpg', 'https://image.tmdb.org/t/p/original/yIvQblXpyQ2qLfqQQpk1BY5MiTa.jpg', 'She''s Got No Name', '2025-06-14', 105, 'Set in a bustling alleyway during the Japanese occupation of Shanghai in the 1940s, the film centres on a wife, Zhan Zhou, who was being charged with the bloody dismemberment of her husband – a killing that seems impossible for her to have committed alone. The murder thrusts Zhan into the spotlight and the court of public opinion, forcing her towards a fate intertwined with that of her own country.', 6.0),
('https://image.tmdb.org/t/p/w500/5Z9tkiIDZ6u49VUeY9Fm3PuSxKT.jpg', 'https://image.tmdb.org/t/p/original/fUoHnP40Mpuo14jgAOMJMPXSVfG.jpg', 'The Naked Gun', '2025-07-30', 120, 'Only one man has the particular set of skills... to lead Police Squad and save the world: Lt. Frank Drebin Jr.', 7.10),
('https://image.tmdb.org/t/p/w500/uuitWHpJwxD1wruFl2nZHIb4UGN.jpg', 'https://image.tmdb.org/t/p/original/kjby7cr8JnIPkfqMGKutOFrM2at.jpg', 'Home Alone 2: Lost in New York', '1992-11-15', 113, 'Instead of flying to Florida with his folks, Kevin ends up alone in New York, where he gets a hotel room with his dad''s credit card—despite problems from a clerk and meddling bellboy. But when Kevin runs into his old nemeses, the Wet Bandits, he''s determined to foil their plans to rob a toy store on Christmas Eve.', 6.764),
('https://image.tmdb.org/t/p/w500/fuJWObGkdhAqAO34FgQx6sVHxtu.jpg', 'https://image.tmdb.org/t/p/original/diRS2bfX9KZG4ijX50G4FTNeQy9.jpg', 'A Cool Fish 2', '2025-06-26', 109, 'Zhang Sanjin sold Yiwu lanterns to Thai TV. Due to his delay in receiving the payment, the capital chain was broken. The disheartened Sanjin climbed to the rooftop to end his life, but he and Mei, who was also a hostage, were accidentally involved in a "carefully planned" kidnapping incident...', 7.5),
('https://image.tmdb.org/t/p/w500/oumprkO9bThExP8NwxBIBnvBu2v.jpg', 'https://image.tmdb.org/t/p/original/zVAA1hKxxXzk85j7nCP5ELlR8Ef.jpg', 'The Life of Chuck', '2025-06-05', 114, 'In this extraordinary story of an ordinary man, Charles ''Chuck'' Krantz experiences the wonder of love, the heartbreak of loss, and the multitudes contained in all of us.', 7.311),
('https://image.tmdb.org/t/p/w500/ccw7CCIAvcZV431CP7NhHAHuiHR.jpg', 'https://image.tmdb.org/t/p/original/mr2LvZydvor6CDZhrpJLRYMCcBc.jpg', 'The King of Kings', '2025-04-07', 100, 'Charles Dickens tells his young son Walter the greatest story ever told, and what begins as a bedtime tale becomes a life-changing journey. Through vivid imagination, the boy walks alongside Jesus, witnessing His miracles, facing His trials, and understanding His ultimate sacrifice.', 8.6),
('https://image.tmdb.org/t/p/w500/qs2AeWbLklk9I5UbI8bllZVfF2V.jpg', 'https://image.tmdb.org/t/p/original/jKMbP0rwEkJYxkNRFQXFkVIIzGP.jpg', 'On Swift Horses', '2025-04-24', 98, 'In the 1950s, a seemingly sensible newlywed and her wayward brother-in-law undertake parallel journeys of risk, romance, and self-discovery.', 6.0),
('https://image.tmdb.org/t/p/w500/3Z9c1tbUhP0QruRjczPHnbx3U2D.jpg', 'https://image.tmdb.org/t/p/original/6GhU4BJnqLSaAuz0yQBq3RfdzsF.jpg', 'Oddity', '2024-07-19', 116, 'After the brutal murder of her twin sister, Darcy goes after those responsible by using haunted items as her tools for revenge.', 6.727),
('https://image.tmdb.org/t/p/w500/gyh0eECE2IqrW8GWl3KoHBfc45j.jpg', 'https://image.tmdb.org/t/p/original/zYdTGqWIup2SMg6A8ZpeLuvCpiy.jpg', 'Teenage Mutant Ninja Turtles: Mutant Mayhem', '2023-07-31', 102, 'After years of being sheltered from the human world, the Turtle brothers set out to win the hearts of New Yorkers and be accepted as normal teenagers through heroic acts. Their new friend April O''Neil helps them take on a mysterious crime syndicate, but they soon get in over their heads when an army of mutants is unleashed upon them.', 7.219),
('https://image.tmdb.org/t/p/w500/cpf7vsRZ0MYRQcnLWteD5jK9ymT.jpg', 'https://image.tmdb.org/t/p/original/kyqM6padQzZ1eYxv84i9smNvZAG.jpg', 'Weapons', '2025-08-06', 105, 'When all but one child from the same class mysteriously vanish on the same night at exactly the same time, a community is left questioning who or what is behind their disappearance.', 6.4),
('https://image.tmdb.org/t/p/w500/iTyMhb3GeMUTzjb4NUBr8gqb2UN.jpg', 'https://image.tmdb.org/t/p/original/70eJ2fuUFFYjMmkF2b2m5iVb503.jpg', 'Together', '2025-07-30', 115, 'Years into their relationship, a couple find themselves at a crossroads as they move to the country, abandoning all that is familiar in their lives except each other. With tensions already flaring, a nightmarish encounter with a mysterious, unnatural force threatens to corrupt their lives, their love, and their flesh.', 6.8),
('https://image.tmdb.org/t/p/w500/4GIqZUgPZ146BhibsPHMHef2nXX.jpg', 'https://image.tmdb.org/t/p/original/znr3SYHph3FtTLpcFTeuMAAJ5mP.jpg', 'Eddington', '2025-07-16', 120, 'In May of 2020, a standoff between a small-town sheriff and mayor sparks a powder keg as neighbor is pitted against neighbor in Eddington, New Mexico.', 6.2),
('https://image.tmdb.org/t/p/w500/ynT06XivgBDkg7AtbDbX1dJeBGY.jpg', 'https://image.tmdb.org/t/p/original/m7vBHmPQSokvN4hMQDxdJBdqAtJ.jpg', 'Happy Gilmore 2', '2025-07-25', 95, 'Revisit Happy Gilmore''s golf career after his win in the Tour Championship.', 6.5),
('https://image.tmdb.org/t/p/w500/gA33EXemAhzXXTdWLOD9JOZkIS6.jpg', 'https://image.tmdb.org/t/p/original/jYz9lDmKl2f4mz1Ph03bikZ8TSe.jpg', 'Red Sonja', '2025-07-31', 120, 'A young girl rises from the ashes of tragedy to become the most feared warrior woman of all time: the She-Devil with a Sword.', 6.9),
('https://image.tmdb.org/t/p/w500/cy76B7VCiIvqUibGLTgdYlK8kx8.jpg', 'https://image.tmdb.org/t/p/original/pc3iV55bKQXhRJuSxEYGJZusXxp.jpg', 'Bambi: The Reckoning', '2025-07-11', 105, 'After a mother and son get in a car wreck, they soon become hunted by Bambi, a mutated grief-stricken deer on a deadly rampage seeking revenge for the death of his mother.', 7.3),
('https://image.tmdb.org/t/p/w500/eG7pnsX28JU0NrWswa0k3d0noiW.jpg', 'https://image.tmdb.org/t/p/original/zj48TAxLo6dEaTImaSTPBbLdgr6.jpg', 'Hipak', '2025-07-18', 110, 'Roy believes he''s found love and luck with three women who visited his vape shop. But as passionate romances unravel, hidden agendas force him to unleash the beast within.', 6.7),
('https://image.tmdb.org/t/p/w500/jrhXbIOFingzdLjkccjg9vZnqIp.jpg', 'https://image.tmdb.org/t/p/original/A466i5iATrpbVjX30clP1Zyfp31.jpg', 'My Oxford Year', '2025-07-31', 120, 'An ambitious American fulfilling her dream of studying at Oxford falls for a charming Brit hiding a secret that may upend her perfectly planned life.', 6.1),
('https://image.tmdb.org/t/p/w500/XBC1Oi5pXmlmvE1n2KqaS1qm7B.jpg', 'https://image.tmdb.org/t/p/original/fLecTFzJrYHYTakGH6jpWXrwN87.jpg', 'Oh, Hi!', '2025-07-25', 95, 'Iris has met her perfect guy, Isaac, and is enjoying their first romantic getaway together — what could go wrong?', 6.3);

INSERT INTO movie_casts(id_movie, id_cast)
VALUES
(1, 12), (1, 23), (1, 45), (1, 66), (1, 8),
(2, 3), (2, 19), (2, 42), (2, 55),
(3, 7), (3, 28), (3, 39), (3, 51), (3, 14),
(4, 5), (4, 17), (4, 31), (4, 44),
(5, 9), (5, 22), (5, 37), (5, 50), (5, 63),
(6, 2), (6, 15), (6, 29), (6, 43),
(7, 11), (7, 24), (7, 38), (7, 52), (7, 65),
(8, 4), (8, 18), (8, 33), (8, 47),
(9, 6), (9, 21), (9, 35), (9, 49), (9, 62),
(10, 13), (10, 26), (10, 40), (10, 54),
(11, 1), (11, 16), (11, 32), (11, 45), (11, 58),
(12, 8), (12, 20), (12, 34), (12, 48),
(13, 10), (13, 25), (13, 39), (13, 53), (13, 66),
(14, 7), (14, 19), (14, 36), (14, 50),
(15, 14), (15, 27), (15, 41), (15, 55), (15, 64),
(16, 3), (16, 17), (16, 30), (16, 44),
(17, 5), (17, 22), (17, 37), (17, 51), (17, 63),
(18, 9), (18, 23), (18, 38), (18, 52),
(19, 11), (19, 26), (19, 42), (19, 56), (19, 65),
(20, 4), (20, 18), (20, 33), (20, 47),
(21, 6), (21, 21), (21, 35), (21, 49), (21, 62),
(22, 12), (22, 28), (22, 40), (22, 54),
(23, 2), (23, 15), (23, 31), (23, 45), (23, 58),
(24, 8), (24, 20), (24, 34), (24, 48),
(25, 10), (25, 25), (25, 39), (25, 53), (25, 66),
(26, 7), (26, 19), (26, 36), (26, 50),
(27, 14), (27, 27), (27, 41), (27, 55), (27, 64),
(28, 3), (28, 17), (28, 30), (28, 44),
(29, 5), (29, 22), (29, 37), (29, 51), (29, 63),
(30, 9), (30, 23), (30, 38), (30, 52),
(31, 11), (31, 26), (31, 42), (31, 56), (31, 65),
(32, 4), (32, 18), (32, 33), (32, 47),
(33, 6), (33, 21), (33, 35), (33, 49), (33, 62),
(34, 12), (34, 28), (34, 40), (34, 54),
(35, 2), (35, 15), (35, 31), (35, 45), (35, 58),
(36, 8), (36, 20), (36, 34), (36, 48),
(37, 10), (37, 25), (37, 39), (37, 53), (37, 66),
(38, 7), (38, 19), (38, 36), (38, 50),
(39, 14), (39, 27), (39, 41), (39, 55), (39, 64),
(40, 3), (40, 17), (40, 30), (40, 44),
(41, 5), (41, 22), (41, 37), (41, 51), (41, 63),
(42, 9), (42, 23), (42, 38), (42, 52),
(43, 11), (43, 26), (43, 42), (43, 56), (43, 65),
(44, 4), (44, 18), (44, 33), (44, 47),
(45, 6), (45, 21), (45, 35), (45, 49), (45, 62),
(46, 12), (46, 28), (46, 40), (46, 54),
(47, 2), (47, 15), (47, 31), (47, 45), (47, 58),
(48, 8), (48, 20), (48, 34), (48, 48),
(49, 10), (49, 25), (49, 39), (49, 53), (49, 66),
(50, 7), (50, 19), (50, 36), (50, 50),
(51, 14), (51, 27), (51, 41), (51, 55), (51, 64),
(52, 3), (52, 17), (52, 30), (52, 44),
(53, 5), (53, 22), (53, 37), (53, 51), (53, 63),
(54, 9), (54, 23), (54, 38), (54, 52),
(55, 11), (55, 26), (55, 42), (55, 56), (55, 65),
(56, 5), (56, 18), (56, 32), (56, 47),
(57, 12), (57, 25), (57, 39),
(58, 3), (58, 21), (58, 35), (58, 50),
(59, 8), (59, 24), (59, 42),
(60, 15), (60, 28), (60, 44), (60, 53),
(61, 7), (61, 19), (61, 37),
(62, 11), (62, 26), (62, 45),
(63, 4), (63, 17), (63, 31), (63, 49),
(64, 9), (64, 23), (64, 38);


INSERT INTO movie_genres(id_movie, id_genre)
VALUES
(1, 3), (1, 6), (1, 9),
(2, 5), (2, 6), (2, 10),
(3, 1), (3, 7), (3, 12),
(4, 2), (4, 8), (4, 15),
(5, 4), (5, 11), (5, 16),
(6, 3), (6, 9), (6, 14),
(7, 5), (7, 10), (7, 17),
(8, 1), (8, 6), (8, 13),
(9, 2), (9, 7), (9, 18),
(10, 4), (10, 12), (10, 19),
(11, 3), (11, 8), (11, 15),
(12, 5), (12, 11), (12, 16),
(13, 1), (13, 9), (13, 14),
(14, 2), (14, 10), (14, 17),
(15, 4), (15, 6), (15, 13),
(16, 3), (16, 7), (16, 18),
(17, 5), (17, 12), (17, 19),
(18, 1), (18, 8), (18, 15),
(19, 2), (19, 11), (19, 16),
(20, 4), (20, 9), (20, 14),
(21, 3), (21, 10), (21, 17),
(22, 5), (22, 6), (22, 13),
(23, 1), (23, 7), (23, 18),
(24, 2), (24, 12), (24, 19),
(25, 4), (25, 8), (25, 15),
(26, 3), (26, 11), (26, 16),
(27, 5), (27, 9), (27, 14),
(28, 1), (28, 10), (28, 17),
(29, 2), (29, 6), (29, 13),
(30, 4), (30, 7), (30, 18),
(31, 3), (31, 12), (31, 19),
(32, 5), (32, 8), (32, 15),
(33, 1), (33, 11), (33, 16),
(34, 2), (34, 9), (34, 14),
(35, 4), (35, 10), (35, 17),
(36, 3), (36, 6), (36, 13),
(37, 5), (37, 7), (37, 18),
(38, 1), (38, 12), (38, 19),
(39, 2), (39, 8), (39, 15),
(40, 4), (40, 11), (40, 16),
(41, 3), (41, 9), (41, 14),
(42, 5), (42, 10), (42, 17),
(43, 1), (43, 6), (43, 13),
(44, 2), (44, 7), (44, 18),
(45, 4), (45, 12), (45, 19),
(46, 3), (46, 8), (46, 15),
(47, 5), (47, 11), (47, 16),
(48, 1), (48, 9), (48, 14),
(49, 2), (49, 10), (49, 17),
(50, 4), (50, 6), (50, 13),
(51, 3), (51, 7), (51, 18),
(52, 5), (52, 12), (52, 19),
(53, 1), (53, 8), (53, 15),
(54, 2), (54, 11), (54, 16),
(55, 4), (55, 9), (55, 14),
(56, 3), (56, 8), (56, 15),
(57, 5), (57, 12),
(58, 1), (58, 7), (58, 19),
(59, 4), (59, 11),
(60, 2), (60, 9), (60, 16),
(61, 6), (61, 14),
(62, 10), (62, 17),
(63, 13), (63, 18),
(64, 3), (64, 6), (64, 12);

INSERT INTO movie_directors(id_movie, id_director)
VALUES
(1, 3),
(2, 7),
(3, 12),
(4, 18),
(5, 23),
(6, 29),
(7, 35),
(8, 41),
(9, 47),
(10, 5),
(11, 10),
(12, 15),
(13, 20),
(14, 25),
(15, 30),
(16, 36),
(17, 42),
(18, 48),
(19, 4),
(20, 9),
(21, 14),
(22, 19),
(23, 24),
(24, 31),
(25, 37),
(26, 43),
(27, 49),
(28, 6),
(29, 11),
(30, 16),
(31, 22),
(32, 28),
(33, 34),
(34, 40),
(35, 46),
(36, 2),
(37, 8),
(38, 13),
(39, 17),
(40, 21),
(41, 26),
(42, 32),
(43, 38),
(44, 44),
(45, 50),
(46, 1),
(47, 7),
(48, 12),
(49, 18),
(50, 23),
(51, 27),
(52, 33),
(53, 39),
(54, 45),
(55, 3),
(56, 12),
(57, 27),
(58, 5),
(59, 33),
(60, 18),
(61, 41),
(62, 9),
(63, 22),
(64, 38);

INSERT INTO transactions (name, email, phone, total_payment, movie_date, id_movie, id_cinema, id_time, id_location, id_payment_method, id_user)
VALUES
('Budi Santoso', 'budi.santoso@mail.com', '081234567891', 100000, '2025-07-15', 5, 1, 3, 1, 2, 1),
('Budi Santoso', 'budi.santoso@mail.com', '081234567891', 150000, '2025-07-16', 12, 2, 1, 2, 3, 1),
('Budi Santoso', 'budi.santoso@mail.com', '081234567891', 200000, '2025-07-17', 23, 3, 4, 3, 1, 1),
('Budi Santoso', 'budi.santoso@mail.com', '081234567891', 50000, '2025-07-18', 34, 4, 2, 4, 5, 1),
('Budi Santoso', 'budi.santoso@mail.com', '081234567891', 250000, '2025-07-19', 45, 1, 5, 5, 4, 1),
('Ani Wijaya', 'ani.wijaya@gmail.com', '082345678901', 50000, '2025-07-15', 7, 2, 1, 2, 2, 2),
('Ani Wijaya', 'ani.wijaya@gmail.com', '082345678901', 100000, '2025-07-16', 15, 3, 3, 3, 3, 2),
('Ani Wijaya', 'ani.wijaya@gmail.com', '082345678901', 150000, '2025-07-17', 27, 4, 2, 4, 1, 2),
('Ani Wijaya', 'ani.wijaya@gmail.com', '082345678901', 200000, '2025-07-18', 38, 1, 4, 5, 6, 2),
('Ani Wijaya', 'ani.wijaya@gmail.com', '082345678901', 50000, '2025-07-19', 50, 2, 5, 1, 7, 2),
('Cahyo Pratama', 'cahyo.pratama@yahoo.com', '083456789012', 150000, '2025-07-15', 3, 3, 2, 3, 4, 3),
('Cahyo Pratama', 'cahyo.pratama@yahoo.com', '083456789012', 100000, '2025-07-16', 18, 4, 1, 4, 5, 3),
('Cahyo Pratama', 'cahyo.pratama@yahoo.com', '083456789012', 50000, '2025-07-17', 29, 1, 3, 5, 2, 3),
('Cahyo Pratama', 'cahyo.pratama@yahoo.com', '083456789012', 200000, '2025-07-18', 42, 2, 4, 1, 3, 3),
('Cahyo Pratama', 'cahyo.pratama@yahoo.com', '083456789012', 250000, '2025-07-19', 55, 3, 5, 2, 8, 3),
('Dewi Lestari', 'dewi.lestari@outlook.com', '084567890123', 200000, '2025-07-15', 8, 4, 3, 4, 6, 4),
('Dewi Lestari', 'dewi.lestari@outlook.com', '084567890123', 50000, '2025-07-16', 21, 1, 2, 5, 7, 4),
('Dewi Lestari', 'dewi.lestari@outlook.com', '084567890123', 150000, '2025-07-17', 33, 2, 1, 1, 1, 4),
('Dewi Lestari', 'dewi.lestari@outlook.com', '084567890123', 100000, '2025-07-18', 47, 3, 4, 2, 2, 4),
('Dewi Lestari', 'dewi.lestari@outlook.com', '084567890123', 250000, '2025-07-19', 40, 4, 5, 3, 3, 4),
('Eko Nugroho', 'eko.nugroho@mail.com', '085678901234', 50000, '2025-07-15', 11, 1, 1, 4, 4, 5),
('Eko Nugroho', 'eko.nugroho@mail.com', '085678901234', 100000, '2025-07-16', 24, 2, 3, 5, 5, 5),
('Eko Nugroho', 'eko.nugroho@mail.com', '085678901234', 200000, '2025-07-17', 37, 3, 2, 1, 6, 5),
('Eko Nugroho', 'eko.nugroho@mail.com', '085678901234', 150000, '2025-07-18', 52, 4, 4, 2, 7, 5),
('Eko Nugroho', 'eko.nugroho@mail.com', '085678901234', 250000, '2025-07-19', 46, 1, 5, 3, 8, 5),
('Fitriani Sari', 'fitri.sari@gmail.com', '086789012345', 150000, '2025-07-15', 14, 2, 2, 4, 1, 6),
('Fitriani Sari', 'fitri.sari@gmail.com', '086789012345', 200000, '2025-07-16', 28, 3, 1, 5, 2, 6),
('Fitriani Sari', 'fitri.sari@gmail.com', '086789012345', 50000, '2025-07-17', 41, 4, 3, 1, 3, 6),
('Fitriani Sari', 'fitri.sari@gmail.com', '086789012345', 100000, '2025-07-18', 46, 1, 4, 2, 4, 6),
('Fitriani Sari', 'fitri.sari@gmail.com', '086789012345', 250000, '2025-07-19', 49, 2, 5, 3, 5, 6),
('Gunawan Setiawan', 'gunawan.setia@yahoo.com', '087890123456', 200000, '2025-07-15', 17, 3, 3, 4, 6, 7),
('Gunawan Setiawan', 'gunawan.setia@yahoo.com', '087890123456', 50000, '2025-07-16', 32, 4, 2, 5, 7, 7),
('Gunawan Setiawan', 'gunawan.setia@yahoo.com', '087890123456', 150000, '2025-07-17', 46, 1, 1, 1, 8, 7),
('Gunawan Setiawan', 'gunawan.setia@yahoo.com', '087890123456', 100000, '2025-07-18', 41, 2, 4, 2, 1, 7),
('Gunawan Setiawan', 'gunawan.setia@yahoo.com', '087890123456', 250000, '2025-07-19', 45, 3, 5, 3, 2, 7),
('Hesti Rahayu', 'hesti.rahayu@outlook.com', '088901234567', 50000, '2025-07-15', 20, 4, 1, 4, 3, 8),
('Hesti Rahayu', 'hesti.rahayu@outlook.com', '088901234567', 100000, '2025-07-16', 35, 1, 3, 5, 4, 8),
('Hesti Rahayu', 'hesti.rahayu@outlook.com', '088901234567', 200000, '2025-07-17', 51, 2, 2, 1, 5, 8),
('Hesti Rahayu', 'hesti.rahayu@outlook.com', '088901234567', 150000, '2025-07-18', 45, 3, 4, 2, 6, 8),
('Hesti Rahayu', 'hesti.rahayu@outlook.com', '088901234567', 250000, '2025-07-19', 47, 4, 5, 3, 7, 8),
('Irfan Maulana', 'irfan.maulana@mail.com', '089012345678', 150000, '2025-07-15', 22, 1, 2, 4, 8, 9),
('Irfan Maulana', 'irfan.maulana@mail.com', '089012345678', 200000, '2025-07-16', 39, 2, 1, 5, 1, 9),
('Irfan Maulana', 'irfan.maulana@mail.com', '089012345678', 50000, '2025-07-17', 54, 3, 3, 1, 2, 9),
('Irfan Maulana', 'irfan.maulana@mail.com', '089012345678', 100000, '2025-07-18', 48, 4, 4, 2, 3, 9),
('Irfan Maulana', 'irfan.maulana@mail.com', '089012345678', 250000, '2025-07-19', 2, 1, 5, 3, 4, 9),
('Jihan Aulia', 'jihan.aulia@gmail.com', '081123456789', 200000, '2025-07-15', 25, 2, 3, 4, 5, 10),
('Jihan Aulia', 'jihan.aulia@gmail.com', '081123456789', 50000, '2025-07-16', 43, 3, 2, 5, 6, 10),
('Jihan Aulia', 'jihan.aulia@gmail.com', '081123456789', 150000, '2025-07-17', 48, 4, 1, 1, 7, 10),
('Jihan Aulia', 'jihan.aulia@gmail.com', '081123456789', 100000, '2025-07-18', 42, 1, 4, 2, 8, 10),
('Jihan Aulia', 'jihan.aulia@gmail.com', '081123456789', 250000, '2025-07-19', 9, 2, 5, 3, 1, 10);

INSERT INTO transaction_details (id_transaction, seat)
VALUES
(1, 'B5'), (1, 'B6'),
(2, 'C3'), (2, 'C4'), (2, 'C5'),
(3, 'D7'), (3, 'D8'), (3, 'D9'), (3, 'D10'),
(4, 'E12'),
(5, 'F1'), (5, 'F2'), (5, 'F3'), (5, 'F4'), (5, 'F5'),
(6, 'A8'),
(7, 'B9'), (7, 'B10'),
(8, 'C12'), (8, 'C13'), (8, 'C14'),
(9, 'D2'), (9, 'D3'), (9, 'D4'), (9, 'D5'),
(10, 'E7'),
(11, 'G3'), (11, 'G4'), (11, 'G5'),
(12, 'A10'), (12, 'A11'),
(13, 'B14'),
(14, 'C1'), (14, 'C2'), (14, 'C3'), (14, 'C4'),
(15, 'D8'), (15, 'D9'), (15, 'D10'), (15, 'D11'), (15, 'D12'),
(16, 'E5'),
(17, 'F7'), (17, 'F8'),
(18, 'G9'), (18, 'G10'), (18, 'G11'),
(19, 'A12'), (19, 'A13'), (19, 'A14'), (19, 'B1'),
(20, 'C6'),
(21, 'D14'), (21, 'E1'), (21, 'E2'),
(22, 'F3'), (22, 'F4'),
(23, 'G5'),
(24, 'A7'), (24, 'A8'), (24, 'A9'), (24, 'A10'),
(25, 'B11'), (25, 'B12'), (25, 'B13'), (25, 'B14'), (25, 'C1'),
(26, 'D4'),
(27, 'E6'), (27, 'E7'),
(28, 'F8'), (28, 'F9'), (28, 'F10'),
(29, 'G11'), (29, 'G12'), (29, 'G13'), (29, 'G14'),
(30, 'A3'),
(31, 'B5'), (31, 'B6'), (31, 'B7'),
(32, 'C8'), (32, 'C9'),
(33, 'D10'),
(34, 'E12'), (34, 'E13'), (34, 'E14'), (34, 'F1'),
(35, 'G2'), (35, 'G3'), (35, 'G4'), (35, 'G5'), (35, 'G6'),
(36, 'A4'),
(37, 'B7'), (37, 'B8'),
(38, 'C10'), (38, 'C11'), (38, 'C12'),
(39, 'D13'), (39, 'D14'), (39, 'E1'), (39, 'E2'),
(40, 'F3'),
(41, 'G7'), (41, 'G8'), (41, 'G9'),
(42, 'A5'), (42, 'A6'),
(43, 'B9'),
(44, 'C14'), (44, 'D1'), (44, 'D2'), (44, 'D3'),
(45, 'E4'), (45, 'E5'), (45, 'E6'), (45, 'E7'), (45, 'E8'),
(46, 'F10'),
(47, 'G12'), (47, 'G13'),
(48, 'A1'), (48, 'A2'), (48, 'A3'),
(49, 'B4'), (49, 'B5'), (49, 'B6'), (49, 'B7'),
(50, 'C8');


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

SELECT *	FROM movies m	WHERE title ILIKE 'aksi' AND (
		EXISTS (
				SELECT 1
				FROM movie_genres mg
				JOIN genres g on g.id = mg.id_genre
				WHERE mg.id_movie = m.id
				AND g.genre_name ILIKE 'aksi'
			)
);

SELECT m.id, m.poster_url, m.backdrop_url, m.title, m.release_date, m.runtime, m.overview, m.rating
		FROM movies m
		WHERE m.title ILIKE $1
		AND (
		$4 = '' OR
		EXISTS (
				SELECT 1
				FROM movie_genres mg
				JOIN genres g on g.id = mg.id_genre
				WHERE mg.id_movie = m.id
				AND g.genre_name ILIKE $4
			)
		
		)
		ORDER BY m.title
		LIMIT $2 OFFSET $3;


select * from transactions where id_user = 10;
select * from users where id = 10;

SELECT 
			m.title, 
			t.movie_date, 
			tm.time, 
			COUNT(td.seat) as seat_count, 
			array_agg(td.seat) as seats,
			t.total_payment
		FROM transactions t
		JOIN movies m ON m.id = t.id_movie
		JOIN times tm ON tm.id = t.id_time
		JOIN transaction_details td ON td.id_transaction = t.id
		WHERE t.id_user = 10 and t.id = 10
		GROUP BY m.title, t.movie_date, tm.time, t.total_payment;

		select * from transaction_details where id_transaction = 10;

		SELECT 
			m.title, 
			t.movie_date, 
			tm.time, 
			COUNT(td.seat) as seat_count, 
			ARRAY_AGG(td.seat) as seats,
			t.total_payment
		FROM transactions t
		JOIN movies m ON m.id = t.id_movie
		JOIN times tm ON tm.id = t.id_time
		JOIN transaction_details td ON td.id_transaction = t.id
		WHERE t.id_user = 10 and t.id = 10
		GROUP BY m.title, t.movie_date, tm.time, t.total_payment;

		SELECT t.name, t.email, t.phone, t.total_payment, t.movie_date, m.title, id_cinema, id_time, id_location, id_payment_method 
	FROM transactions t 
	JOIN movies m ON m.id = t.id_movie
	JOIN cinemas c ON c.id = t.id_cinema
	JOIN times tm ON tm.id = t.id_time
	JOIN locations l ON l.id = t.id_location
	JOIN payment_methods pm ON pm.id = t.id_payment_method
	WHERE t.id = 1;
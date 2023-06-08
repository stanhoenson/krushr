INSERT INTO roles (role)
VALUES
	("Administrator"),
	("Moderator"),
	("Creator");

INSERT INTO users (email, password, role_id)
VALUES
	("stan@hoenson.xyz", "stanaap2", 1);

INSERT INTO routes (title, status_id, user_id)
VALUES
	("De Boswandeling", 1, 1),
	("Het Mysterieuze Pad", 2, 1);

INSERT INTO statuses (status)
VALUES
	("Published"),
	("Unpublished"),
	("Unfinished");

INSERT INTO points_of_interest (title, longitude, latitude)
VALUES
	("De Boomhut van Tjeerd", 3.574934, 45.889978),
	("De Grote Glijbaan", 13.343422, 97.111223),
	("Het Verlaten Kasteel", 87.758234, 23.099231);

INSERT INTO types (type, icon)
VALUES
	("Facebook", "https://krushr.com/facebook.png"),
	("Instagram", "https://krushr.com/instagram.png"),
	("Text", "https://krushr.com/text.png"),
	("Image", "https://krushr.com/image.png");

INSERT INTO entries (content, hyperlink, type_id)
VALUES
	("100 jaar geleden was Tjeerd klaar met het maken van deze boomhut.", "https://krushr.com/", 2),
	("Tjeerd heeft een naam gegeven aan deze boomhut: Tjeerd.", "https://krushr.com/", 2),
	("Ontdek alle geheimen van het bos.", "https://krushr.com/", 2);

INSERT INTO categories (category, icon, weight, type_id)
VALUES
	("Geschiedenis en Cultuur", "https://krushr.com/geschiedenis-en-cultuur.png", 1, NULL),
	("Hoogtepunten en Evenementen", "https://krushr.com/hoogtepunten-en-evenementen.png", 2, NULL),
	("Horeca", "https://krushr.com/horeca.png", 1, NULL);

INSERT INTO routes_points_of_interest (route_id, point_of_interest_id)
VALUES
	(1, 1),
	(1, 2),
	(2, 3);

INSERT INTO entries_routes (entry_id, route_id)
VALUES
	(3, 1);

INSERT INTO entries_points_of_interest (entry_id, point_of_interest_id)
VALUES
	(1, 1),
	(2, 1);

INSERT INTO categories_routes (category_id, route_id)
VALUES
	(1, 1),
	(1, 2);

INSERT INTO categories_points_of_interest (category_id, point_of_interest_id)
VALUES
	(3, 1);

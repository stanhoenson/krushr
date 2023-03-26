SELECT routes.title, statuses.status, users.email, points_of_interest.title, points_of_interest.longitude, points_of_interest.latitude
FROM routes
JOIN statuses ON routes.status_id = statuses.id
JOIN users ON routes.user_id = users.id
JOIN routes_points_of_interest ON routes.id = routes_points_of_interest.route_id
JOIN points_of_interest ON points_of_interest.id = routes_points_of_interest.point_of_interest_id

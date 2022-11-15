CREATE OR REPLACE FUNCTION get_spots_in_square (longitude double precision, latitude double precision, radius double precision)
	RETURNS TABLE (id uuid, name varchar, website varchar, coordinates geography, description varchar, rating float) as
$BODY$
DECLARE
distanceFromCenter float := SQRT(POWER(radius, 2)*2);
	   	leftTop GEOGRAPHY := ST_Project(ST_Point(longitude, latitude)::geography, distanceFromCenter, radians(315.0));
 		leftBottom GEOGRAPHY := ST_Project(ST_Point(longitude, latitude)::geography, distanceFromCenter, radians(225.0));
 		rightTop GEOGRAPHY := ST_Project(ST_Point(longitude, latitude)::geography, distanceFromCenter, radians(45.0));
 		rightBottom GEOGRAPHY := ST_Project(ST_Point(longitude, latitude)::geography, distanceFromCenter, radians(135.0));

BEGIN
    RETURN QUERY
    SELECT points_tbl.id, points_tbl.name, points_tbl.website, points_tbl.coordinates, points_tbl.description, points_tbl.rating
    FROM "MY_TABLE" as points_tbl
    WHERE ST_X(points_tbl.coordinates::geometry) < ST_X(rightTop::geometry)
      AND ST_X(points_tbl.coordinates::geometry) > ST_X(leftTop::geometry)
      AND ST_Y(points_tbl.coordinates::geometry) < ST_Y(rightTop::geometry)
      AND ST_Y(points_tbl.coordinates::geometry) > ST_Y(rightBottom::geometry);
END;
$BODY$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_spots_in_circle (longitude double precision, latitude double precision, radius double precision)
	RETURNS TABLE (id uuid, name varchar, website varchar, coordinates geography, description varchar, rating float) as
$BODY$
BEGIN
    RETURN QUERY
    SELECT points_tbl.id, points_tbl.name, points_tbl.website, points_tbl.coordinates, points_tbl.description, points_tbl.rating
    FROM "MY_TABLE" as points_tbl
    WHERE ST_Distance(ST_Point(longitude, latitude)::geography, points_tbl.coordinates) < radius;
END;
$BODY$ LANGUAGE plpgsql;

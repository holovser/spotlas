CREATE OR REPLACE FUNCTION get_spots_with_domain_count()
	RETURNS TABLE (spot_name varchar, domain text, count bigint) as
$BODY$
BEGIN
	CREATE TEMP TABLE extract_domain_tbl AS
SELECT name, substring(website, '^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/?\n]+)') as domain
FROM  "MY_TABLE"
WHERE website is not NULL;

CREATE TEMP TABLE domain_count_tbl AS
SELECT extract_domain_tbl.domain, count(extract_domain_tbl.domain) AS count
FROM extract_domain_tbl
GROUP by extract_domain_tbl.domain
HAVING COUNT(extract_domain_tbl.domain)>1;


RETURN QUERY
SELECT name, extract_domain_tbl.domain, domain_count_tbl.count
FROM extract_domain_tbl
         RIGHT JOIN domain_count_tbl
                    ON extract_domain_tbl.domain=domain_count_tbl.domain;

DROP TABLE IF EXISTS extract_domain_tbl, domain_count_tbl;
END;
$BODY$ LANGUAGE plpgsql;


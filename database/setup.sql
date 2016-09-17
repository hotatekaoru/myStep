CREATE ROLE postgres SUPERUSER;
ALTER ROLE postgres WITH LOGIN;
CREATE ROLE postgres WITH createdb password='testrole';
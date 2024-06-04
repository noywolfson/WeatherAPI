CREATE TABLE IF NOT EXISTS weather (longitude REAL, latitude REAL, forecastTime TIMESTAMP, temperature REAL, precipitation REAL);
COPY weather FROM '/docker-entrypoint-initdb.d/File1.csv' DELIMITER ',' CSV HEADER;
COPY weather FROM '/docker-entrypoint-initdb.d/File2.csv' DELIMITER ',' CSV HEADER;
COPY weather FROM '/docker-entrypoint-initdb.d/File3.csv' DELIMITER ',' CSV HEADER;

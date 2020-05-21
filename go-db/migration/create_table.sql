CREATE TABLE experiment (
    id INT NOT NULL PRIMARY KEY,
    name varchar(20) NOT NULL,
    s2Ids varchar[] NOT NULL,
    timeStart timestamp NOT NULL,
    timeEnd timestamp NOT NULL,
    createdAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)

CREATE OR REPLACE FUNCTION update_updatedAt_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updatedAt = now();
   RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_experiment_updatedAt BEFORE UPDATE ON experiment FOR EACH ROW EXECUTE PROCEDURE update_updatedAt_column();

INSERT INTO experiment VALUES(1, 'exp1', '{123456, 234567, 345678}', '2020-03-13 12:00:00', '2020-03-13 18:00:00');
INSERT INTO experiment VALUES(2, 'exp1', '{123456, 234567, 345678}', '2020-03-14 12:00:00', '2020-03-14 18:00:00');
-------------------------------------------------------------------------------------------------------------------------------------------

-- WITH T as (
--     SELECT COUNT(1) as c
--     FROM experiment
--     WHERE s2Ids && ARRAY [$input_s2Ids_comma_seperated]
--       AND (TIMESTAMP $input_timeStart, $input_timeEnd) OVERLAPS (timeStart, timeEnd)
-- ) INSERT INTO experiment VALUES(123, 'exp2', '{$input_s2Ids_comma_seperated}', '$input_timeStart', '$input_timeEnd') ON CONFLICT (T.c > 0) DO NOTHING;
--
-- -- ex:
-- WITH T as (
--     SELECT COUNT(1) as c
--     FROM experiment
--     WHERE s2Ids && ARRAY [123456, 234567, 345678]
--       AND (TIMESTAMP '2020-03-15 12:00:00', TIMESTAMP '2020-03-15 12:00:00') OVERLAPS (timeStart, timeEnd)
-- ) INSERT INTO experiment VALUES(3, 'exp2', '{123456, 234567, 345678}', '2020-03-15 12:00:00', '2020-03-15 12:00:00') ON CONFLICT CHECK  (T.c > 0) DO NOTHING;
--
-- : DIDN'T Work

ALTER TABLE experiment ADD COLUMN enable bool;

CREATE OR REPLACE FUNCTION enable_if_unique()
RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'INSERT') THEN
    UPDATE experiment
    SET NEW.enable=true
    WHERE (SELECT count(1)
           FROM experiment
           WHERE enable= true AND location_Ids && OLD.location_ids AND (OLD.timeStart, OLD.timeEnd) OVERLAPS (timeStart, timeEnd)
        ) = 0;
    RETURN NEW;
    END IF;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER enable_if_unique_trigger INSTEAD OF INSERT ON experiment FOR EACH ROW EXECUTE PROCEDURE enable_if_unique();

-----------------------------------------------------------------------------------------------------------------------------
create table experiment (
                            id int not null primary key,
                            name varchar(20) not null,
                            locationids int[] not null,
                            timestart timestamp not null,
                            timeend timestamp not null,
                            createdat timestamp not null default current_timestamp,
                            updatedat timestamp not null default current_timestamp
);

-- this prevents overlaps in the locationids AND the time range
alter table experiment
    add constraint no_overlap
        exclude using gist (locationids with &&, tsrange(timestart, timeend) with &&);

insert into experiment (id, name, locationids, timestart, timeend)
values
(1, 'one', array[1,2], timestamp '2020-03-01 08:00:00', timestamp '2020-03-04 17:30:00'),
(2, 'two', array[3,4], timestamp '2020-03-01 08:00:00', timestamp '2020-03-04 17:30:00'),
(3, 'three', array[5,6,7], timestamp '2020-03-02 14:00:00', timestamp '2020-03-10 20:00:00');

-- this insert fails
insert into experiment (id, name, locationids, timestart, timeend)
values
(4, 'four', array[7,8,9], timestamp '2020-03-01 10:00:00', timestamp '2020-03-02 18:30:00');
ALTER TABLE manager."THEA_DETAILS_HOSTING"
    ADD COLUMN uuid character varying(36);

COMMENT ON COLUMN manager."THEA_DETAILS_HOSTING".uuid
    IS 'Identifica la sesión que se está usando en el momento';

CREATE UNIQUE INDEX THEA_DETAILS_HOSTING_uuid 
    ON manager."THEA_DETAILS_HOSTING" (uuid);

ALTER TABLE manager."THEA_DETAILS_HOSTING" 
    ADD CONSTRAINT unique_uuid 
    UNIQUE USING INDEX THEA_DETAILS_HOSTING_uuid;

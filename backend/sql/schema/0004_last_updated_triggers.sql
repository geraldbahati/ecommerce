-- +goose Up
CREATE OR REPLACE FUNCTION update_last_updated_column() RETURNS TRIGGER AS $$ BEGIN NEW.last_updated = NOW(); RETURN NEW; END; $$ LANGUAGE plpgsql;
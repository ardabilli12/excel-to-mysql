-- Script to insert cities for Auvergne-Rhône-Alpes state

-- First insert or update the state record
INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d14-61a5-71a9-847b-9e345b265ea9", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Auvergne-Rhône-Alpes", NOW(), NOW())
ON DUPLICATE KEY UPDATE updated_at = NOW();

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-7c8e-af6b-54f73f59400d", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Bourgogne-Franche-Comté", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-7d61-b05b-357717151e2e", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Bretagne", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-7a0a-97d3-350bfa77f5a4", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Centre-Val de Loire", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-72e0-ab7b-768646ce8a00", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Corse", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES (, "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Grand-Est", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-71c4-a055-46b0881661dd", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Hauts-de-France", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-7792-8c1c-a541176b8dc7", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Île-de-France", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-72c3-b557-8f9ae2f60d33", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Normandie", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-75df-b714-31ebc3c7d52d", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Nouvelle-Aquitaine", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-7d97-a0b9-40cf9f1b13ba", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Occitanie", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-77ba-8f98-ff1e00c09fd5", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Pays-de-la-Loire", NOW(), NOW());

INSERT INTO master_states (uuid, country_uuid, name, created_at, updated_at)
VALUES ("01975d2c-f07e-72c3-b557-8f9ae2f60d33", "05e397b2-fbf7-40ef-96d8-bfad81c01fbe", "Provence-Alpes-Côte-d’Azur", NOW(), NOW());
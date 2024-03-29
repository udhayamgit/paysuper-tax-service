ALTER TABLE public.taxes ALTER COLUMN rate TYPE numeric(20,8);

INSERT INTO  public.taxes (country, rate, created_at, updated_at) VALUES
('AU', 0.1, current_timestamp, current_timestamp),
('CA', 0.09975, current_timestamp, current_timestamp),
('KR', 0.1, current_timestamp, current_timestamp),
('RU', 0.2, current_timestamp, current_timestamp),
('TR', 0.18, current_timestamp, current_timestamp),
('AT', 0.2, current_timestamp, current_timestamp),
('BE', 0.21, current_timestamp, current_timestamp),
('BG', 0.2, current_timestamp, current_timestamp),
('GB', 0.2, current_timestamp, current_timestamp),
('HU', 0.27, current_timestamp, current_timestamp),
('DE', 0.19, current_timestamp, current_timestamp),
('GR', 0.24, current_timestamp, current_timestamp),
('DK', 0.25, current_timestamp, current_timestamp),
('IE', 0.23, current_timestamp, current_timestamp),
('ES', 0.21, current_timestamp, current_timestamp),
('IT', 0.22, current_timestamp, current_timestamp),
('CY', 0.19, current_timestamp, current_timestamp),
('LV', 0.21, current_timestamp, current_timestamp),
('LT', 0.21, current_timestamp, current_timestamp),
('LU', 0.17, current_timestamp, current_timestamp),
('MT', 0.18, current_timestamp, current_timestamp),
('NL', 0.21, current_timestamp, current_timestamp),
('PL', 0.23, current_timestamp, current_timestamp),
('PT', 0.23, current_timestamp, current_timestamp),
('RO', 0.19, current_timestamp, current_timestamp),
('SK', 0.2, current_timestamp, current_timestamp),
('SI', 0.22, current_timestamp, current_timestamp),
('FI', 0.24, current_timestamp, current_timestamp),
('FR', 0.2, current_timestamp, current_timestamp),
('HR', 0.25, current_timestamp, current_timestamp),
('CZ', 0.21, current_timestamp, current_timestamp),
('SE', 0.25, current_timestamp, current_timestamp),
('EE', 0.2, current_timestamp, current_timestamp),
('IS', 0.24, current_timestamp, current_timestamp),
('NZ', 0.15, current_timestamp, current_timestamp),
('NO', 0.25, current_timestamp, current_timestamp),
('SG', 0.07, current_timestamp, current_timestamp),
('JP', 0.08, current_timestamp, current_timestamp),
('LI', 0.077, current_timestamp, current_timestamp),
('CH', 0.077, current_timestamp, current_timestamp),
('AR', 0.21, current_timestamp, current_timestamp),
('BY', 0.2, current_timestamp, current_timestamp),
('AE', 0.05, current_timestamp, current_timestamp),
('SA', 0.05, current_timestamp, current_timestamp),
('RS', 0.2, current_timestamp, current_timestamp),
('TW', 0.05, current_timestamp, current_timestamp),
('ZA', 0.15, current_timestamp, current_timestamp),
('AM', 0.2, current_timestamp, current_timestamp),
('BH', 0.05, current_timestamp, current_timestamp),
('EG', 0.14, current_timestamp, current_timestamp),
('IN', 0.18, current_timestamp, current_timestamp),
('TZ', 0.18, current_timestamp, current_timestamp),
('UY', 0.22, current_timestamp, current_timestamp),
('AL', 0.2, current_timestamp, current_timestamp),
('BR', 0.17, current_timestamp, current_timestamp),
('GH', 0.125, current_timestamp, current_timestamp),
('KE', 0.16, current_timestamp, current_timestamp),
('CO', 0.19, current_timestamp, current_timestamp);
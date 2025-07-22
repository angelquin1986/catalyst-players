-- Script para limpiar datos de foto existentes que tienen el prefijo embebido
-- Este script decodifica el base64 que contiene el prefijo y extrae solo la imagen

-- Función para decodificar base64 (MySQL)
DELIMITER $$
CREATE FUNCTION IF NOT EXISTS BASE64_DECODE_SAFE(str TEXT) 
RETURNS LONGBLOB
READS SQL DATA
DETERMINISTIC
BEGIN
    DECLARE result LONGBLOB;
    SET result = FROM_BASE64(str);
    RETURN result;
END$$
DELIMITER ;

-- Limpiar datos de Team
UPDATE teams 
SET photo = BASE64_DECODE_SAFE(
    SUBSTRING(
        FROM_BASE64(photo), 
        LOCATE('base64,', FROM_BASE64(photo)) + 7
    )
)
WHERE photo IS NOT NULL 
AND LENGTH(photo) > 0
AND FROM_BASE64(photo) LIKE 'data:image/%';

-- Limpiar datos de Player  
UPDATE players 
SET photo = BASE64_DECODE_SAFE(
    SUBSTRING(
        FROM_BASE64(photo), 
        LOCATE('base64,', FROM_BASE64(photo)) + 7
    )
)
WHERE photo IS NOT NULL 
AND LENGTH(photo) > 0
AND FROM_BASE64(photo) LIKE 'data:image/%';

-- Limpiar la función temporal
DROP FUNCTION IF EXISTS BASE64_DECODE_SAFE; 
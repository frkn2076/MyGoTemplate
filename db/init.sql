INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('GlobalErrorMessage', 'Geçici süre işleminizi gerçekleştiremiyoruz.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('GlobalErrorMessage', 'We are temporarily unable to process your transaction.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-1', 'Kullanıcı bulunamadı.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-1', 'User not found.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-2', 'Kayıtlı kullanıcı.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-2', 'User already exists.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-3', 'Lütfen geçerli bir e-mail adresi giriniz.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-3', 'Please enter a valid e-mail.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-4', 'Yanlış şifre girdiniz.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-4', 'You entered the wrong password.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-5', 'Şifrenizin uzunluğu minimum 6 karakter olmalı.', 'TR', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);
INSERT IGNORE INTO Localization(resource, message, language, created_at, updated_at) VALUES('Err-5', 'Your password length must be at least 6 characters.', 'EN', CURDATE(), CURDATE()) ON DUPLICATE KEY UPDATE resource = VALUES(resource), message = VALUES(message), language = VALUES(language);

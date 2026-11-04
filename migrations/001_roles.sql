CREATE TABLE IF NOT EXISTS roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO roles (name, code, description) VALUES 
('Студент', 'student', 'Записывается на курсы и обучается на платформе(клиент)'),
('Ментор', 'mentor', 'Следит за курсами, выкладывает уроки'),
('Администратор', 'admin', 'Следит за платформой, имеет полный доступ');

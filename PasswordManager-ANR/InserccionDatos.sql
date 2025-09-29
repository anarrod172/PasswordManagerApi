-- Eliminamos la tabla para poder crearlas
drop table users;

--Insertamos algunos datos en la base de datos
insert into users (username, password, url, notes, path, width, height) values 
('Pepe', '123psd', 'http://ejemplo1.com', 'Cambiar nombre', 'icon1.png', 64, 64),
('Anita12', 'contraseña', 'http://ejemplo2.com', 'Cambiar contraseña', 'icon2.png', 128, 64),
('LuisaLopez', '1111', 'http://ejemplo3.com', 'Nuevo usuario', 'icon3.png', 128, 128),
('python8', '0kde2', 'http://ejemplo4.com', 'No se requiere', 'icon4.png', 64, 128),
('ventilador', 'seguridad31', 'http://ejemplo5.com', 'Modificar url', 'icon5.png', 32, 64);

# Ranking App Backend

## Descripción general
Este es el backend de la aplicación Ranking App, que permite gestionar jugadores y usuarios.

## Endpoints
A continuación se detallan los endpoints disponibles en la API:

### Jugadores (Players)

- `POST /players`: Crea un nuevo jugador.
- `GET /players`: Obtiene la lista de todos los jugadores.
- `GET /players/{id}`: Obtiene los detalles de un jugador por su ID.
- `PUT /players/{id}`: Actualiza el nombre de un jugador.
- `POST /players/{id}/double-points`: Agrega puntos dobles a la puntuación de un jugador.
- `POST /players/{id}/single-points`: Agrega puntos sencillos a la puntuación de un jugador.
- `DELETE /players/{id}`: Elimina un jugador.

### Usuarios (Users)

- `POST /signup`: Registra un nuevo usuario.
- `POST /login`: Inicia sesión y obtiene un token JWT.
- `POST /logout`: Cierra la sesión del usuario.
- `GET /validate`: Valida el estado de autenticación del usuario.

## Modelos de datos

### Jugador (Player)
- `ID`: Identificador único del jugador (int)
- `Name`: Nombre del jugador (string)
- `Points`: Puntuación total del jugador (int)

### Usuario (User)
- `ID`: Identificador único del usuario (int)
- `Username`: Nombre de usuario (string)
- `Password`: Contraseña del usuario (string)

## Autenticación y autorización
La API utiliza un sistema de autenticación basado en tokens JWT. Los usuarios deben iniciar sesión para obtener un token, y este token debe ser enviado en el encabezado `Authorization` de las solicitudes posteriores.

## Instalación y configuración
1. Clona el repositorio: `git clone https://github.com/cesart18/ranking_app.git`
2. Instala las dependencias: `go mod download`
3. Configura las variables de entorno necesarias (por ejemplo, la conexión a la base de datos).
4. Inicia la aplicación: `go run main.go`

## Contacto
Si tienes alguna duda o sugerencia, puedes comunicarte con nosotros a través de:

- Correo electrónico: info@rankingapp.com
- GitHub: https://github.com/cesart18/ranking_app
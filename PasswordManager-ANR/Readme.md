##### Instrucciones de instalación y ejecución.

Password Manager API, es una API REST desarrollada en lenguaje Go usando Gin como framework web y GORM como ORM para la gestión de base de datos con PostgresSQL. Esta API permite crear, consultar, actualizar y eliminar usuarios almacenados en la base de datos.



Las tecnologías utilizadas en el proyecto son:

-Go versión 1.25.1.

-Gin Gonic como framework web.

-GORM como ORM para Go.

-PostgresSQL como base de datos.

-Postman como herramienta para hacer pruebas de la API.



Descargamos e instalamos los siguientes programas.
-Visual Studio Code.
-DBeaver.
-Postman.

2\. Una vez instalados procederemos a configurarlos para preparar nuestro entorno de trabajo.

* En Visual Studio Code instalamos el plugin de Go.
* A continuación, creamos una carpeta con el nombre de nuestro proyecto y empezamos a instalar las dependencias. Para instalar las dependencias debemos ejecutar los siguientes comandos:

&nbsp;	-go get -u github.com/gin-gonic/gin (instala las dependencias del framework de Gin).

&nbsp;	-go get -u gorm.io/driver/postgres (descargamos el driver de PostgresSQL para GORM).   

&nbsp;	-go get -u gorm.io/gorm (descargamos el paquete principal de GORM).

* Creamos un "Hola mundo!" e iniciamos el servidor para comprobar que funciona ejecutando el comando go run main.go. 

Si todo está configurado correctamente, el servidor mostrará en la consola el mensaje "DB connected" indicando que la base de datos se ha conectado con éxito, seguido de una salida del tipo \[GIN-debug] Listening and serving HTTP on :8080, que confirma que el servidor está en funcionamiento y escuchando el puerto.

* Generamos el módulo de conexión a la base de datos en el que definimos una función de conexión y las variables globales que necesitaremos.

3\. Generamos la conexión local a la base de datos en DBeaver.

* Abrimos el DBeaver.
* En el menú superior seleccionamos la opción "Base de datos". En el desplegable que se nos abre seleccionamos la opción "crear conexión". 
* Elegimos PostgresSQL y configuramos los parámetros de la conexión (host,port,database,nombre de usuario y contraseña). Previamente, debemos tener instalado PostgresSQL en nuestro ordenador. A continuación, creamos la conexión, se genera la base de datos y podemos escribir nuestro script.

4\. Probar en Postman el funcionamiento de nuestra API.

* Desde nuestra cuenta creamos una petición haciendo click sobre "New -->HTTP Request.
* Seleccionamos el método HTTP(GET,POST,PUT,DELETE).
* Escribimos la URL, como por ejemplo, http://localhost:5433/users/newuser.
* Si hay una petición POST O PUT, desde la pestaña de Body --> seleccionamos raw -->elegimos JSON y escribimos el cuerpo de una petición.
* Finalmente, hacemos clic en Send y vemos la respuesta que nos devuelve el servidor. Postman nos mostrará códigos de estado (200, 201, 400, 404...) y el contenido de la respuesta que devuelve mi API 





##### Dependencias necesarias.



Mi proyecto ha sido realizado en Go con el uso de las librerías Gin, GORM y PostgresSQL. Las dependencias que he utilizado son las siguientes:



\-***Para crear la API REST.***

 github.com/gin-gonic/gin



\-***Driver de PostgresSQL para GORM.***

 gorm.io/driver/postgres

 

\-***ORM para trabajar con la base de datos.***

 gorm.io/gorm



\-***Dependencias del archivo go.mod.***

github.com/bytedance/gopkg v0.1.3 // indirect

 	github.com/bytedance/sonic v1.14.1 // indirect

 	github.com/bytedance/sonic/loader v0.3.0 // indirect

 	github.com/cloudwego/base64x v0.1.6 // indirect

 	github.com/gabriel-vasile/mimetype v1.4.10 // indirect

 	github.com/gin-contrib/sse v1.1.0 // indirect

 	github.com/gin-gonic/gin v1.11.0 // indirect

 	github.com/go-playground/locales v0.14.1 // indirect

 	github.com/go-playground/universal-translator v0.18.1 // indirect

 	github.com/go-playground/validator/v10 v10.27.0 // indirect

 	github.com/goccy/go-json v0.10.5 // indirect

 	github.com/goccy/go-yaml v1.18.0 // indirect

 	github.com/jackc/pgpassfile v1.0.0 // indirect

 	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect

 	github.com/jackc/pgx/v5 v5.7.6 // indirect

 	github.com/jackc/puddle/v2 v2.2.2 // indirect

 	github.com/jinzhu/inflection v1.0.0 // indirect

 	github.com/jinzhu/now v1.1.5 // indirect

 	github.com/json-iterator/go v1.1.12 // indirect

 	github.com/klauspost/cpuid/v2 v2.3.0 // indirect

 	github.com/leodido/go-urn v1.4.0 // indirect

 	github.com/mattn/go-isatty v0.0.20 // indirect

 	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect

 	github.com/modern-go/reflect2 v1.0.2 // indirect

 	github.com/pelletier/go-toml/v2 v2.2.4 // indirect

 	github.com/quic-go/qpack v0.5.1 // indirect

 	github.com/quic-go/quic-go v0.54.0 // indirect

 	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect

 	github.com/ugorji/go/codec v1.3.0 // indirect

 	go.uber.org/mock v0.6.0 // indirect

 	golang.org/x/arch v0.21.0 // indirect

 	golang.org/x/crypto v0.42.0 // indirect

 	golang.org/x/mod v0.28.0 // indirect

 	golang.org/x/net v0.44.0 // indirect

 	golang.org/x/sync v0.17.0 // indirect

 	golang.org/x/sys v0.36.0 // indirect

 	golang.org/x/text v0.29.0 // indirect

 	golang.org/x/tools v0.37.0 // indirect

 	google.golang.org/protobuf v1.36.9 // indirect

 	gorm.io/driver/postgres v1.6.0 // indirect

 	gorm.io/gorm v1.31.0 // indirect





##### Variables de entorno requeridas (ej: claves de cifrado, credenciales de la base de datos...).



-Credenciales para la base de datos.

DB\_HOST:localhost

DB\_USER:postgres

DB\_PASSWORD:pss123

DB\_NAME:postgres

DB\_PORT:5433





##### Ejemplos de endpoints y cómo probarlos.



Para comprobar el correcto funcionamiento de la API, podemos utilizar Postman, una herramienta que nos permite enviar peticiones HTTP como GET, POST, PUT y DELETE y visualizar las respuestas que devuelve nuestra aplicación.



A continuación, se muestran los endpoints disponibles junto con ejemplos de uso:

\-***GET http://localhost:8080/users?url=http://ejemplo1.com***

Endpoint que devuelve todos los usuarios cuyo campo url coincida con el valor porporcionado por la query string. Solo muestra tres campos, el username, la url y las notas.



\-***POST http://localhost:8080/users/newuser***

Al que le pasamos el siguiente body(JSON):

{

    "username": "lolaq2",

    "password": "inventado",

    "url": "http://ejemplo6.com",

    "notes": "El nuevo usuario necesita algunas modificaciones",

    "Icon": {

        "path": "/images/icon.png",

        "width": 64,

        "height": 64

    }

}



Endpoint que permite crear un nuevo usuario en la base de datos con la información enviada en el cuerpo de la petición. Muestra un usuario con todos sus campos completados y en la base de datos aparece la fecha de cuando se creó.



\-***PUT http://localhost:8080/users/id/3***

Al que le pasamos el siguiente body(JSON) con los campos que deseamos modificar:

{

    "notes":"Usuario modificado con éxito"

}



Endpoint que actualiza la información del campo notas del usuario con id=3. En el campo notas debe aparecer un mensaje diferente al que había cuando se creó la base de datos y se muestra la hora de modificación en el campo updated de la base de datos.



\-***DELETE http://localhost:8080/users/username/Pepe***

Elimina el usuario con nombre Pepe si existe en la base de datos. En la base de datos aparece completado el campo de deleted.


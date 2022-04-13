# SNIPPETBOX

_Proyecto desarrollado en GOLANG


## Descripción de la aplicación  

_Snippetbox es una aplicación desarrollada en Go desarrollada con fines de estudio a partir del libro Let's Go. Esta aplicación tiene como función crear "snippets", mostrar los 10 últimos creados y crear cuentas. Para poder crear un snippet se requiere de una cuenta, para visualizar el top 10 no se necesita estar registrado.

_Un snippet es un fragmento de texto que incluye; título, contenido y fecha de expiración. Al crearse, también se guarda la fecha de creación. Como se menciona anteriormente, solo personas con cuenta registrada pueden crear un snippet, pero para poder ver alguno a detalle o el listado de ellos no se requiere cuenta.

_El registro de cuenta es simple, la aplicación pide ingresar; nombre, email y una password. 
De igual manera, cuando se crea un nuevo usuario, el sistemta tambiên registra la fecha de creación y agrega un campo más para activar o desactivar una cuenta. En cuanto al password, el sistema cifra el strign que se ingresa por seguridad.
Una persoa puede; registrarse, iniciar sesión y cerrar sesión.

### Temas que se tocan

```
Estructura de un proyecto en Go
Templating y herencia
Funciones default y personalizadas dentro de los templates
Routing request
Base de datos
Processing forms
Despliegue de data dinámica
Cuentas de usuario
Configuración de  HTTPs server
Creación de un certificado TLS
Session management
Autenticación de usuario
Middleware
Logs para los request
Confiogurar los headers de seguridad
Mensajes flash
Variables de ambiente
Manejador de errores
Validación de datos
Testing (Unit test/ Table Driven Test)
Mocking dependecies
```

### Versiones y dependencias

_A continuación se enlistan las versiones y dependencias que fueron utilizadas para el desarrollo de esta aplicación

```
Go v1.17

github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f (matching routes - RESTful Routing)
github.com/go-sql-driver/mysql v1.6.0 (Database driver)
github.com/justinas/alice v1.2.0  (to manage middleware/hanlder chains)
github.com/golangcollege/sessions v1.2.0  (session manager/security considerations)
github.com/justinas/nosurf v1.1.1 (mitigate risk of CSRF using the double submit cookie pattern)
golang.org/x/crypto v0.0.0-20200317142112-1b76d66859c6 (hashing and cheking passwords mostly)
golang.org/x/sys v0.0.0-20190412213103-97732733099d (sub-dependecies of golangcollege/sessions)
```

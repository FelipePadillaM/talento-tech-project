# Talento Tech - API Go

Una API REST moderna construida en Go para el proyecto Talento Tech.

## 🚀 Características

- **API REST completa** con Gin framework
- **Autenticación de usuarios** con sesiones
- **Gestión de archivos** con subida y descarga
- **CORS configurado** para integración con frontend
- **Logging estructurado** con Logrus
- **Validación de datos** con Gin binding
- **Arquitectura modular** y escalable

## 📋 Prerrequisitos

- Go 1.21 o superior
- Git

## 🛠️ Instalación

1. **Clonar el repositorio**
   ```bash
   cd go-app
   ```

2. **Instalar dependencias**
   ```bash
   go mod tidy
   ```

3. **Configurar variables de entorno**
   ```bash
   cp env.example .env
   # Editar .env con tus configuraciones
   ```

4. **Ejecutar la aplicación**
   ```bash
   go run main.go
   ```

## 🌐 Endpoints de la API

### Autenticación

- `POST /api/v1/auth/register` - Registrar nuevo usuario
- `POST /api/v1/auth/login` - Iniciar sesión
- `POST /api/v1/auth/logout` - Cerrar sesión
- `GET /api/v1/auth/me` - Obtener usuario actual

### Usuarios

- `GET /api/v1/users/` - Obtener todos los usuarios
- `GET /api/v1/users/:id` - Obtener usuario por ID
- `POST /api/v1/users/` - Crear nuevo usuario
- `PUT /api/v1/users/:id` - Actualizar usuario
- `DELETE /api/v1/users/:id` - Eliminar usuario

### Archivos

- `GET /api/v1/files/` - Obtener todos los archivos
- `GET /api/v1/files/:id` - Obtener archivo por ID
- `GET /api/v1/files/:id/download` - Descargar archivo
- `POST /api/v1/files/` - Subir archivo
- `DELETE /api/v1/files/:id` - Eliminar archivo

### Utilidades

- `GET /` - Información de la API
- `GET /api/v1/health` - Health check

## 📝 Ejemplos de Uso

### 1. Desde cURL (línea de comandos)

#### Registrar un usuario
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@ejemplo.com",
    "password": "123456",
    "name": "Usuario Ejemplo"
  }'
```

#### Iniciar sesión
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "usuario@ejemplo.com",
    "password": "123456"
  }'
```

#### Subir un archivo
```bash
curl -X POST http://localhost:8080/api/v1/files/ \
  -F "file=@documento.pdf" \
  -F "name=Mi Documento" \
  -F "public=true"
```

### 2. Desde JavaScript/TypeScript

#### Con fetch API
```javascript
// Login
const response = await fetch('http://localhost:8080/api/v1/auth/login', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  credentials: 'include',
  body: JSON.stringify({
    email: 'usuario@ejemplo.com',
    password: '123456'
  })
});

const data = await response.json();
console.log(data);
```

#### Con Axios
```javascript
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  withCredentials: true
});

// Login
const login = async (email, password) => {
  const response = await api.post('/auth/login', { email, password });
  return response.data;
};
```

### 3. Desde Angular

```typescript
import { ApiService } from './services/api.service';

constructor(private apiService: ApiService) {}

// Login
this.apiService.login({ email: 'user@example.com', password: '123456' })
  .subscribe({
    next: (response) => {
      if (response.success) {
        console.log('Login exitoso:', response.data);
      }
    },
    error: (error) => {
      console.error('Error en login:', error);
    }
  });
```

### 4. Interfaz Web de Pruebas

Abre el archivo `examples/test-api.html` en tu navegador para probar todos los endpoints de forma visual.

### 5. Desde Postman/Insomnia

Importa esta colección de ejemplo:

```json
{
  "info": {
    "name": "Talento Tech API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "url": "http://localhost:8080/api/v1/health"
      }
    },
    {
      "name": "Login",
      "request": {
        "method": "POST",
        "url": "http://localhost:8080/api/v1/auth/login",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"email\": \"usuario@ejemplo.com\",\n  \"password\": \"123456\"\n}"
        }
      }
    }
  ]
}
```

## 🏗️ Estructura del Proyecto

```
go-app/
├── main.go              # Punto de entrada de la aplicación
├── go.mod               # Dependencias del módulo
├── env.example          # Variables de entorno de ejemplo
├── README.md            # Documentación
├── handlers/            # Controladores HTTP
│   ├── auth_handler.go  # Autenticación
│   ├── user_handler.go  # Gestión de usuarios
│   └── file_handler.go  # Gestión de archivos
├── models/              # Modelos de datos
│   ├── user.go          # Modelo de usuario
│   └── file.go          # Modelo de archivo
├── services/            # Lógica de negocio
├── middleware/          # Middleware personalizado
├── utils/               # Utilidades
└── uploads/             # Directorio de archivos subidos
```

## 🔧 Configuración

### Variables de Entorno

- `PORT`: Puerto del servidor (default: 8080)
- `GIN_MODE`: Modo de Gin (debug/release)
- `CORS_ALLOWED_ORIGINS`: Orígenes permitidos para CORS
- `UPLOAD_DIR`: Directorio para archivos subidos
- `MAX_FILE_SIZE`: Tamaño máximo de archivo en bytes

## 🚀 Despliegue

### Desarrollo
```bash
go run main.go
```

### Producción
```bash
go build -o app main.go
./app
```

### Docker (futuro)
```bash
docker build -t talento-tech-go .
docker run -p 8080:8080 talento-tech-go
```

## 🔒 Seguridad

- **CORS configurado** para orígenes específicos
- **Validación de datos** en todos los endpoints
- **Sesiones seguras** con cookies httpOnly
- **Sanitización de archivos** antes de guardar

## 📊 Monitoreo

- **Health check** en `/api/v1/health`
- **Logging estructurado** con Logrus
- **Manejo de errores** centralizado

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para detalles.

## 🆘 Soporte

Para soporte, email: soporte@talento-tech.com o crear un issue en el repositorio. 
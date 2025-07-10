// Ejemplos de consumo de la API de Go con JavaScript puro

const API_BASE_URL = 'http://localhost:8080/api/v1';

// Función helper para hacer requests
async function apiRequest(endpoint, options = {}) {
  const url = `${API_BASE_URL}${endpoint}`;
  const defaultOptions = {
    headers: {
      'Content-Type': 'application/json',
    },
    credentials: 'include', // Para incluir cookies
  };

  const response = await fetch(url, { ...defaultOptions, ...options });
  const data = await response.json();
  
  if (!response.ok) {
    throw new Error(data.message || 'Error en la petición');
  }
  
  return data;
}

// Ejemplos de uso:

// 1. Health check
async function checkHealth() {
  try {
    const result = await apiRequest('/health');
    console.log('API Status:', result);
  } catch (error) {
    console.error('Error checking health:', error);
  }
}

// 2. Registrar usuario
async function registerUser(email, password, name) {
  try {
    const result = await apiRequest('/auth/register', {
      method: 'POST',
      body: JSON.stringify({ email, password, name })
    });
    console.log('Usuario registrado:', result);
    return result;
  } catch (error) {
    console.error('Error en registro:', error);
    throw error;
  }
}

// 3. Login
async function loginUser(email, password) {
  try {
    const result = await apiRequest('/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password })
    });
    console.log('Login exitoso:', result);
    return result;
  } catch (error) {
    console.error('Error en login:', error);
    throw error;
  }
}

// 4. Obtener usuario actual
async function getCurrentUser() {
  try {
    const result = await apiRequest('/auth/me');
    console.log('Usuario actual:', result);
    return result;
  } catch (error) {
    console.error('Error obteniendo usuario:', error);
    throw error;
  }
}

// 5. Obtener todos los usuarios
async function getUsers() {
  try {
    const result = await apiRequest('/users/');
    console.log('Usuarios:', result);
    return result;
  } catch (error) {
    console.error('Error obteniendo usuarios:', error);
    throw error;
  }
}

// 6. Subir archivo
async function uploadFile(file, name, isPublic = false) {
  try {
    const formData = new FormData();
    formData.append('file', file);
    if (name) {
      formData.append('name', name);
    }
    formData.append('public', isPublic.toString());

    const result = await fetch(`${API_BASE_URL}/files/`, {
      method: 'POST',
      body: formData,
      credentials: 'include'
    });

    const data = await result.json();
    console.log('Archivo subido:', data);
    return data;
  } catch (error) {
    console.error('Error subiendo archivo:', error);
    throw error;
  }
}

// 7. Obtener archivos
async function getFiles() {
  try {
    const result = await apiRequest('/files/');
    console.log('Archivos:', result);
    return result;
  } catch (error) {
    console.error('Error obteniendo archivos:', error);
    throw error;
  }
}

// 8. Descargar archivo
async function downloadFile(fileId) {
  try {
    const response = await fetch(`${API_BASE_URL}/files/${fileId}/download`, {
      credentials: 'include'
    });
    
    if (!response.ok) {
      throw new Error('Error descargando archivo');
    }
    
    const blob = await response.blob();
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'archivo-descargado';
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
    
    console.log('Archivo descargado exitosamente');
  } catch (error) {
    console.error('Error descargando archivo:', error);
    throw error;
  }
}

// Ejemplo de uso completo
async function ejemploCompleto() {
  try {
    // 1. Verificar que la API esté funcionando
    await checkHealth();
    
    // 2. Registrar un usuario
    const registerResult = await registerUser('test@ejemplo.com', '123456', 'Usuario Test');
    
    // 3. Hacer login
    const loginResult = await loginUser('test@ejemplo.com', '123456');
    
    // 4. Obtener usuario actual
    const currentUser = await getCurrentUser();
    
    // 5. Obtener todos los usuarios
    const users = await getUsers();
    
    // 6. Obtener archivos
    const files = await getFiles();
    
    console.log('Ejemplo completado exitosamente');
  } catch (error) {
    console.error('Error en el ejemplo:', error);
  }
}

// Exportar funciones para uso en otros archivos
if (typeof module !== 'undefined' && module.exports) {
  module.exports = {
    checkHealth,
    registerUser,
    loginUser,
    getCurrentUser,
    getUsers,
    uploadFile,
    getFiles,
    downloadFile,
    ejemploCompleto
  };
}

// Para uso en navegador, las funciones estarán disponibles globalmente 
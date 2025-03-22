// ola_mundo_http.js

// Cria um servidor HTTP simples
const http = require('http');

// Configura o servidor para responder "Olá Mundo"
const server = http.createServer((req, res) => {
  res.statusCode = 200; // Código de status HTTP 200 (OK)
  res.setHeader('Content-Type', 'text/plain');
  res.end('Olá Mundo\n');
});

// Define a porta que o servidor irá escutar
const PORT = env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Servidor rodando em http://localhost:${PORT}/`);
});

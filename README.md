# Golang: Autenticação com JWT
### Rodar o Projeto

Para rodar o projeto, siga os passos abaixo:

1. Abra o terminal na raiz do projeto.
2. Execute o comando para iniciar a aplicação:

```sh
go run main.go
```

### Testar

1. **Registrar Usuário**:
   - Método: `POST`
   - URL: `http://localhost:8080/register`
   - Body (JSON):
   ```json
   {
     "username": "testuser",
     "password": "testpassword"
   }
   ```

2. **Login de Usuário**:
   - Método: `POST`
   - URL: `http://localhost:8080/login`
   - Body (JSON):
   ```json
   {
     "username": "testuser",
     "password": "testpassword"
   }
   ```
   - Copie o token JWT da resposta.

3. **Teste de Autenticação**:
   - Método: `POST`
   - URL: `http://localhost:8080/protected/welcome`
   - Headers:
     - Key: `Authorization`
     - Value: `YOUR_JWT_TOKEN` (substitua `YOUR_JWT_TOKEN` pelo token copiado)


4. **Obter Todos os Usuários**:
   - Método: `GET`
   - URL: `http://localhost:8080/protected/users`
   - Headers:
     - Key: `Authorization`
     - Value: `Bearer YOUR_JWT_TOKEN` (substitua `YOUR_JWT_TOKEN` pelo token copiado)
# Autenticação com JWT em Go

Este projeto é uma aplicação web simples em Go que utiliza o framework [Gin](https://github.com/gin-gonic/gin) para criar uma API RESTful com autenticação baseada em JWT (JSON Web Token). A aplicação permite registro de usuários, login e acesso a rotas protegidas.

## Funcionalidades

- **Registro de Usuário**: Permite criar um novo usuário com um nome de usuário e uma senha.
- **Login de Usuário**: Autentica um usuário e gera um token JWT.
- **Rota Protegida**: `GET /protected/dashboard` - Rota protegida que retorna uma mensagem de boas-vindas e o nome de usuário.
- **Obter Todos os Usuários**: `GET /protected/users` - Rota protegida que retorna todos os usuários registrados.

## Tecnologias

Este projeto utiliza as seguintes tecnologias e bibliotecas:

- [Gin](https://github.com/gin-gonic/gin) - Framework web para Go.
- [GORM](https://gorm.io) - ORM para Go.
- [JWT-Go](https://github.com/dgrijalva/jwt-go) - Biblioteca para manipulação de JWT.
- [bcrypt](https://golang.org/x/crypto/bcrypt) - Biblioteca para hashing de senhas.

## Instalação

### Requisitos

- [Go](https://golang.org/doc/install) instalado na sua máquina.
- [SQLite](https://www.sqlite.org/download.html) instalado para o banco de dados.

### Passos

1. **Clone o repositório**:

   ```sh
   git clone https://github.com/seu-usuario/seu-repositorio.git
   ```

2. **Navegue para o diretório do projeto**:

   ```sh
   cd seu-repositorio
   ```

3. **Instale as dependências**:

   Execute o seguinte comando para baixar todas as dependências do projeto:

   ```sh
   go mod tidy
   ```

## Configuração

### Configuração do Banco de Dados

A configuração do banco de dados está definida no arquivo `config/database.go`. O projeto está configurado para usar o SQLite e cria um arquivo `test.db` no diretório raiz do projeto. Você pode alterar a configuração do banco de dados conforme necessário.

### Ambiente de Desenvolvimento

Certifique-se de que o ambiente de desenvolvimento está configurado para usar o Go e as bibliotecas necessárias.

## Uso

### Inicie a Aplicação

Para iniciar o servidor, execute o seguinte comando:

```sh
go run main.go
```

O servidor estará disponível em `http://localhost:8080`.

### Endpoints

- **Registrar Usuário**

  - **Método:** `POST`
  - **URL:** `http://localhost:8080/register`
  - **Body (JSON):**

    ```json
    {
      "username": "testuser",
      "password": "testpassword"
    }
    ```

- **Login de Usuário**

  - **Método:** `POST`
  - **URL:** `http://localhost:8080/login`
  - **Body (JSON):**

    ```json
    {
      "username": "testuser",
      "password": "testpassword"
    }
    ```

  - **Resposta (JSON):**

    ```json
    {
      "token": "YOUR_JWT_TOKEN"
    }
    ```

- **Acessar Rota Protegida**

  - **Método:** `GET`
  - **URL:** `http://localhost:8080/protected/welcome`
  - **Headers:**

    - Key: `Authorization`
    - Value: `YOUR_JWT_TOKEN`

  - **Resposta (JSON):**

    ```json
    {
      "message": "Welcome testuser",
    }
    ```

- **Obter Todos os Usuários**

  - **Método:** `GET`
  - **URL:** `http://localhost:8080/protected/users`
  - **Headers:**

    - Key: `Authorization`
    - Value: `YOUR_JWT_TOKEN`

  - **Resposta (JSON):**

    ```json
    [
      {
        "ID": 1,
        "CreatedAt": "2024-07-14T00:00:00Z",
        "UpdatedAt": "2024-07-14T00:00:00Z",
        "DeletedAt": null,
        "Username": "testuser",
        "Password": "$2a$10$..."
      }
    ]
    ```

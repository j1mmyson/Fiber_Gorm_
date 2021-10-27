# Fiber + Gorm

## Setting environment variables

Creating `.env` file to compose environment variables.  
`.env`  

```
DB_HOST= <host of db>
DB_NAME= <name of db>
DB_USER= <user of db>
DB_PASSWORD= <password of db>
DB_PORT= <port number>
```

## How to run

- Make sure your DB server is running (MySQL)

- build binary file  

  ```bash
  $ go build -o server
  ```

- run binary file  

  ```bash
  $ ./server
  ```

  
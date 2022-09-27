<h1>📚 Manga Store Project</h1>
<h3>🧑🏼‍💻 Established by</h3>
<ul>
    <li><strong>Ruangyot Nanchiang</strong></li>
</ul>

<h3>🚀 Version</h3>
<ul>
    <li><strong>1.0.0</strong></li>
</ul>

<h3>🌱 Status</h3>
<ul>
    <li><strong>In progress about 10%</strong></li>
</ul>

<h3>📃 Introduction</h3>
<ul>
    <li>This project is to build a REST API with Golang and GO Fiber using the clean architecture of Uncle Bob.</li>
</ul>

<h3>📝 Requirements</h3>
<ul>
    <li>🐳 Dokcer</li>
    <li>⚡ GO v1.18+</li>
    <li>🐘 PostgreSQL</li>
    <li>📕 Redis</li>
    <li>🧑‍🚀 Postman</li>
</ul>

<h3>Install the postgreSQL on Docker</h3>
<ul>
<li>

<strong>Pull Image -> <a href="https://hub.docker.com/_/postgres" target="_blank">PostgreSQL Docker Image</a></strong>

```
docker pull postgres:alpine
```
</li>

<li>

<strong>Run the container</strong>

```
docker run --name manga-store-db-<your-stage> -e POSTGRES_PASSWORD=123456 -p <your-port>:5432 -d postgres:alpine
```
</li>
<li>

<strong>Config the postgres</strong>

```
docker exec -it manga-store-db-<your-stage> bash
```
```
psql -U postgres
```
```
create database manga_store_db;
```
</li>
<li>

<strong>Check the database that created or not</strong>

```
\l
```
```
      Name      |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
----------------+----------+----------+------------+------------+-----------------------
 manga_store_db_<your_stage> | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
```
</li>
</ul>

<h3>Install the Redis on Docker</h3>
<ul>
<li>

<strong>Pull Image -> <a href="https://hub.docker.com/_/redis" target="_blank">Redis Docker Image</a></strong>

```
docker pull redis:alpine
```
</li>

<li>

<strong>Run the container</strong>

```
docker run --name manga-store-redis -p 6379:6379 -d redis
```
</li>
</ul>

<h3>💾 Database Schema</h3>
<img src="./assets/screenshots/schema_v7.png">

<h3>🪜 Database Migration</h3>
<ul>
<li>

<strong>Migrate Up</strong>

```
migrate -database 'postgres://postgres:123456@0.0.0.0:<your_port>/manga_store_db_<your_stage>?sslmode=disable' -source file://to-the-path -verbose up
```
</li>

<li>

<strong>Migrate Down</strong>

```
migrate -database 'postgres://postgres:123456@0.0.0.0:<your_port>/manga_store_db_<your_stage>?sslmode=disable' -source file://to-the-path -verbose down
```
</li>
</ul>

<h3>🔩 Initialize The Project (Just an example)</h3>
<ul>

```zsh
📂 app/
├─ 📄 main.go
📂 assets/
├─ 📂 logs/
│  ├─ 📄 2022-08-20T08:18:00.log
│  ├─ 📄 2022-08-21T08:18:00.log
│  ├─ 📄 2022-08-22T08:18:00.log
📂 configs/
├─ 📄 configs.go
📂 internals/
├─ 📂 servers/
│  ├─ 📄 server.go
│  ├─ 📄 handler.go
├─ 📂 entities/
│  ├─ 📄 books.go
│  ├─ 📄 monitors.go
│  ├─ 📄 response.go
├─ 📂 monitors/
│  ├─ 📂 controllers/http/
│  │  ├─ 📄 monitor_controllers.go
│  ├─ 📂 usecases/
│  │  ├─ 📄 monitor_usecases.go
├─ 📂 books/
│  ├─ 📂 controllers/http/
│  │  ├─ 📄 book_controllers.go
│  ├─ 📂 usecases/
│  │  ├─ 📄 book_usecases.go
│  ├─ 📂 repositories/
│  │  ├─ 📄 book_repositories.go
📂 pkg/
├─ 📂 databases/
│  ├─ 📂 migrations/
│  │  ├─ 📄 manga_store_db_001_up.sql
│  │  ├─ 📄 manga_store_db_001_down.sql
│  ├─ 📄 postgresql.go
├─ 📂 cache/
│  ├─ 📄 redis.go
├─ 📂 middlewares/
├─ 📂 utils/
📂 tests/
├─ 📂 books/
│  ├─ 📄 book_test.go
📄 .env.dev
📄 .env.prod
📄 .env.test
```
</ul>
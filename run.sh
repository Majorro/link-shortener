export POSTGRES_PASSWORD=postgres
export SHORT_LINK_BASE=http://localhost:15001/
export POSTGRES_CONN_STRING=postgres://postgres:postgres@db:5432/postgres?sslmode=disable
export ALLOW_CORS_FROM=http://localhost:8080
export REACT_APP_API_BASE_URL=http://localhost:15001/
docker compose up --build
package queries

const CreateUserQuery = `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

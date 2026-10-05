package queries

const CreateUserQuery = `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

const ListUsersQuery = `SELECT * FROM users`

const ReadUserQuery = `SELECT id, nickname FROM users WHERE id = ?`

const UpdateUserQuery = `UPDATE users SET nickname = ? WHERE id = ?`

const DeleteUserQuery = `DELETE FROM users WHERE id = ?`
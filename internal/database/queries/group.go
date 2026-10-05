package queries

const CreateGroupQuery = `CREATE TABLE IF NOT EXISTS group (id INTEGER PRIMARY KEY AUTOINCREMENT,name TEXT NOT NULL,description TEXT, created_by INTEGER NOT NULL, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, FOREIGN KEY (created_by) REFERENCES users(id));`

const UpdateGroupQuery = `UPDATE group SET name = ?, description = ? WHERE id = ?;`

const ReadGroupQuery = `SELECT id, name, description, created_by, created_at FROM group WHERE id = ?;`

const DeleteGroupQuery = `DELETE FROM group WHERE id = ?;`

const ListByUserIDQuery = `SELECT g.id, g.name, g.description, g.created_by, g.created_at FROM group g JOIN group_user gu ON g.id = gu.group_id WHERE gu.user_id = ?;`

const AddUserQuery = `INSERT INTO group_user (group_id, user_id) VALUES (?, ?);`
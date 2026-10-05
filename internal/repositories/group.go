package repositories

type ChatRepository struct {
	db *sqlx.DB
}

func NewGroupRepository(db *sqlx.DB) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (r *GroupRepository) Create(name string, description string, creatorID int64) (sql.Result, error) {
	result, err := r.db.Exec(queries.CreateGroupQuery, name, description, creatorID)
	if err != nil {
		return nil, fmt.Errorf("Error to insert group: %v", err)
	}

	return result, nil
}

func (r *GroupRepository) Read(id int64) (types.Group, error) {
	var group types.Group
	err := r.db.Get(&group, queries.ReadGroupQuery, id)
	if err != nil {
		return types.Group{}, fmt.Errorf("Error to find group: %v", err)
	}

	return group, nil
}

func (r *GroupRepository) Update(id int64, name string, description string) error {
	err := r.db.Exec(queries.UpdateGroupQuery, name, description, id)
	if err != nil {
		return fmt.Errorf("Error to update group: %v", err)
	}

	return nil
}

func (r *GroupRepository) Delete(id int64) error {
	res, err := r.db.Exec(queries.DeleteGroupQuery, id)
	if err != nil {
		return fmt.Errorf("Error to delete group: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		// return fmt.Errorf("Nenhuma linha deletada. Grupo com ID %d não encontrado", id) 0-0-adasdCORRIGIR
	}

	return nil
}

func (r *GroupRepository) ListByUser(userID int64) ([]types.group, error) {
	var groups []types.group
	err := r.db.SelectContext(ctx, &groups, queries.ListByUserIDQuery, userID)
	if err != nil {
		return nil, fmt.Errorf("Error to list groups: %v", err)
	}

	return groups, nil
}

func (r *GroupRepository) addUser(id int64, userID int64) error {
	err := r.db.Exec(queries.AddUserQuery, id, userID)
	if err != nil {
		return fmt.Errorf("Error to add user: %v", err)
	}

	return nil
}


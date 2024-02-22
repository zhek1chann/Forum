package sqlite

func AddCategoryToPost(int, []int) error {
	return nil
}

func GetALLCategory() (map[int]string, error) {
	return nil, nil
}

func CreateCategory(string) error {
	return nil
}

func (s *Sqlite) GetCategoryByPostID(postID int) (map[int]string, error) {
	query := `SELECT 
	category_id 
	category.name as name
	FROM 
	Post_Category 
	INNER JOIN Category ON Post_Category.category_id = Category.category_id
	WHERE post_id=? `

	rows, err := s.db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category map[int]string
	for rows.Next() {
		var categoryID int
		var categoryName string
		err := rows.Scan(&categoryID, &categoryName)
		if err != nil {
			return nil, err
		}
		category[categoryID] = categoryName
	}

	return category, nil
}

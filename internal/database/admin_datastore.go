package database

func (db *ShinkanDatabase) PostCircle(input PostCircleInput) error {
	db.Map.AddTableWithName(PostCircle{}, tableCircles)
	db.Map.AddTableWithName(PostCirclesCircleTypes{}, tableCirclesCircleTypes)
	db.Map.AddTableWithName(PostCircleImages{}, tableCircleImages)
	tx, err := db.Map.Begin()
	if err != nil {
		return err
	}
	circle := PostCircle{
		Name:        input.Name,
		About:       input.About,
		CatchCopy:   input.CatchCopy,
		Description: input.Description,
		EyeCatch:    input.EyeCatch,
		Email:       input.Email,
		Twitter:     input.Twitter,
		URL:         input.URL,
		CategoryID:  input.CategoryID,
	}

	if err := tx.Insert(&circle); err != nil {
		tx.Rollback()
		return err
	}

	//TODO
	for _, v := range input.Types {
		content := PostCirclesCircleTypes{
			CircleID:     circle.ID,
			CircleTypeID: v.ID,
		}
		if err := tx.Insert(&content); err != nil {
			tx.Rollback()
			return err
		}
	}

	//TODO
	for _, v := range input.Images {
		content := PostCircleImages{
			URL:      v.URL,
			CircleID: circle.ID,
		}
		if err := tx.Insert(&content); err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

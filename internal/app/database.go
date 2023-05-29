package app

// Initialize database data

func InitDatabaseData() error {
	if err := InitTableProperties(); err != nil {
		return err
	}
	if err := InitTableSysctl(); err != nil {
		return err
	}

	return nil
}

func InitTableProperties() error {

	return nil
}

func InitTableSysctl() error {

	return nil
}

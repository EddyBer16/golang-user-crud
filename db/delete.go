package db

func DeleteUser(id string) error {
	_, err := con.Query("DELETE FROM users WHERE ID=? AND NOT ID=0", id)
	if err != nil {
		return err
	}
	return nil
}

package db

// always look into cache first, go DB if no cache is found.
func Get(model interface{}) error {

	err := getCache(model)
	if err != nil {
		return err
	}

	return nil
}

func CachedInsert() {

}

func CachedUpdate() {

}

func DeleteWithCache() {

}

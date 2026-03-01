package initializer

func init() {
	InitViper()
	InitMySQL()
	err := InitConsul()
	if err != nil {
		return
	}
}

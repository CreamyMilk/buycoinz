package storage

var FakeStore map[string]string

func InitalizeDB() {
	FakeStore = make(map[string]string)
}

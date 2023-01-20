package types

type Dog struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogStoreRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

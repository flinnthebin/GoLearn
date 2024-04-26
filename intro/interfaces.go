package main

import "fmt"

// idomatic Go interface naming convention appends (er)
type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresNumberStore struct {
	// some postgres values
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7}, nil
}

func (p PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into the Postgres database")
	return nil
}

type MongoDBNumberStore struct {
	// some values
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into the MongoDB")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: PostgresNumberStore{},
	}

	if err := apiServer.numberStore.Put(1); err != nil {
		panic(err)
	}

	nums, err := apiServer.numberStore.GetAll()
	if err != nil {
		panic(err)
	}
	fmt.Println(nums)
}

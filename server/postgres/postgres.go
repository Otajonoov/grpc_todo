package postgres

import (
	pb "GRPC-TODO/genproto/store"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "quvonchbek"
	password = "0101"
	database = "store"
)

type Store struct {
	ID          int64
	Name        string
	Discription string
	Addresses   []string
	IsOpen      bool
}

func ConDB() *sql.DB {
	connDB := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", connDB)
	if err != nil {
		fmt.Println("Error to connect DB : ", err)
	}
	return db
}

func CreateStore(store *pb.Store) (*pb.Store, error) {
	db := ConDB()
	query := `
	INSERT INTO stores
		(name, description, is_open, addresses)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, description, is_open, addresses	
	`
	newStore := pb.Store{}
	err := db.QueryRow(query, store.Name, store.Discription, store.IsOpen,
		pq.Array(store.Addresses)).Scan(
		&newStore.Id,
		&newStore.Name,
		&newStore.Discription,
		&newStore.IsOpen,
		pq.Array(&newStore.Addresses),
	)

	if err != nil {
		return nil, err
	}

	return &newStore, nil
}

func GetStore(id int64) (*pb.Store, error) {
	db := ConDB()

	store := pb.Store{}
	query := `
	SELECT
		id, name, description, is_open, addresses
	FROM
		stores
	WHERE id=$1`
	row := db.QueryRow(query, id)
	err := row.Scan(
		&store.Id,
		&store.Name,
		&store.Discription,
		&store.IsOpen,
		pq.Array(&store.Addresses),
	)
	if err != nil {
		return nil, err
	}

	return &store, nil
}

func UpdateStore(store *pb.Store) error {
	db := ConDB()
	defer db.Close()
	upInfo := Store{
		ID:          1,
		Name:        "Davron",
		Discription: ".NET",
		Addresses:   []string{"Tashkent", "Chilonzor"},
		IsOpen:      false,
	}

	query := `
	UPDATE stores SET
		name=$1, description=$2, is_open=$3, addresses=$4
	WHERE id=$5	`
	_, err := db.Exec(query, upInfo.Name, upInfo.Discription, upInfo.IsOpen, pq.Array(upInfo.Addresses))
	if err != nil {
		fmt.Println("Error to update : ", err)
	}
	return nil
}

func DeleteStore(id int64) error {
	db := ConDB()
	defer db.Close()
	_, err := db.Exec("DELETE FROM stores WHERE id=$1", id)
	if err != nil {
		fmt.Println("Error to delete : ", err)
	}
	return nil
}

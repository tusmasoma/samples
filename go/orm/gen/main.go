package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	infra "github.com/tusmasoma/samples/go/orm/gen/database/mysql"
	"github.com/tusmasoma/samples/go/orm/gen/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "go_orm_gen_setup_db"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_NAME)
	db, _ := gorm.Open(mysql.Open(dsn))
	repo := infra.NewUser(db)

	ctx := context.Background()
	userID := uuid.NewString()
	user, err := entity.NewUser(userID, "user", "email@gmail.com", "password")
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
		return
	}
	if err := repo.Create(ctx, user); err != nil {
		log.Fatalf("failed to create user: %v", err)
		return
	}
	user, err = repo.Get(ctx, userID)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
		return
	}
	fmt.Printf("user: %v\n", user)
	user.Name = "updated_user"
	if err := repo.Update(ctx, user); err != nil {
		log.Fatalf("failed to update user: %v", err)
		return
	}
	user, err = repo.Get(ctx, userID)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
		return
	}
	fmt.Printf("user: %v\n", user)
	if err := repo.Delete(ctx, userID); err != nil {
		log.Fatalf("failed to delete user: %v", err)
		return
	}
	user, err = repo.Get(ctx, userID)
	if err == nil {
		log.Fatalf("user should be deleted: %v", user)
		return
	}
	fmt.Println("user deleted successfully")
}

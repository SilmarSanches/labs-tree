package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateAuction_MockClosesAuctionAutomatically(t *testing.T) {
	os.Setenv("AUCTION_INTERVAL", "1s")

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("should not error on create and auto-close auction", func(mt *mtest.T) {
		repo := &AuctionRepository{
			Collection: mt.Coll,
		}

		now := time.Now()
		auction := &auction_entity.Auction{
			Id:          uuid.New().String(),
			ProductName: "Mocked Product",
			Category:    "Mocked Category",
			Description: "Mocked Description",
			Condition:   auction_entity.New,
			Status:      auction_entity.Active,
			Timestamp:   now,
			EndAuction:  now.Add(getAuctionInterval()),
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := repo.CreateAuction(context.TODO(), auction)
		if err != nil {
			t.Fatalf("Erro ao criar leilão com mock: %v", err)
		}

		time.Sleep(2 * time.Second)

		mt.AddMockResponses(
			mtest.CreateSuccessResponse(),
			mtest.CreateSuccessResponse(),
		)
	})
}

func TestAuctionIsAutomaticallyClosedMongoReal(t *testing.T) {
	os.Setenv("AUCTION_INTERVAL", "3s")
	ctx := context.Background()

	db, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		t.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	repo := NewAuctionRepository(db)

	now := time.Now()
	auction := &auction_entity.Auction{
		Id:          uuid.New().String(),
		ProductName: "Produto Teste",
		Category:    "Categoria Teste",
		Description: "Descrição",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   now,
		EndAuction:  now.Add(getAuctionInterval()),
	}

	errCreate := repo.CreateAuction(ctx, auction)
	if errCreate != nil {
		t.Fatalf("Erro ao criar leilão: %v", errCreate)
	}

	time.Sleep(5 * time.Second)

	var result AuctionEntityMongo
	err = repo.Collection.FindOne(ctx, bson.M{"_id": auction.Id}).Decode(&result)
	if err != nil {
		t.Fatalf("Erro ao buscar o leilão: %v", err)
	}

	if result.Status != auction_entity.Completed {
		t.Errorf("Esperava status 'completed', mas encontrou '%v'", result.Status)
	}
}

package database

import ( 
	"context" 
	"fmt" 
	"os" 
	"time"

	"github.com/joho/godotenv" 
	"go.mongodb.org/mongo-driver/mongo" 
	"go.mongodb.org/mongo-driver/mongo/options"
) 

func ConexaoBD() (*mongo.Collection, error) {
	// Localizar .env
    err := godotenv.Load("../../../.env") 
    if err != nil {
		return nil, fmt.Errorf("erro ao carregar .env: %w", err)
	}

	// Obter valor de .env
    MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		return nil, fmt.Errorf("MONGODB_URI não definido no .env")
	}

	// Conectar ao banco de dados.
	clientOptions := options.Client().ApplyURI(MONGODB_URI) 
	
	// Conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao MongoDB: %w", err)
	}

	// Verifica a conexão
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("erro ao verificar conexão com MongoDB: %w", err)
	}

	// Retorna a coleção "test" do banco "testdb"
	collection := client.Database("testdb").Collection("test")
	return collection, nil
}
package main

import ( 
	"context" 
	"fmt" 
	"log" 
	"os" 

	"github.com/joho/godotenv" 
	"go.mongodb.org/mongo-driver/mongo" 
	"go.mongodb.org/mongo-driver/mongo/options"
) 

func main() {
	Connect();
}

func Connect() *mongo.Collection {
	// Localizar .evn
    err := godotenv.Load("../../../../.env") 
    if err != nil { 
        log.Fatalf( "Erro ao carregar arquivo .env: %s" , err) 
    }

	// Obter valor de .env
    MONGODB_URI := os.Getenv("MONGODB_URI")

	// Conectar ao banco de dados.
	clientOption := options.Client().ApplyURI(MONGODB_URI) 
	client, err := mongo.Connect(context.Background(), clientOption) 
	if err != nil { 
		log.Fatal(err) 
	} 

	// Verificar a conexão.
	err = client.Ping(context.Background(), nil ) 
	if err != nil { 
		log.Fatal(err) 
	} 

	// Criar teste
	collection := client.Database("testdb").Collection("test")
	if err != nil {
		log.Fatal(err)
	}

	// Se tudo der certo, Conectado.
	fmt.Println("Connected to db")

	return collection
}
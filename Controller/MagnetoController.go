package Controller

import (
	"github.com/magneto/Repository"
	"github.com/magneto/core"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
)

type MagnetoController struct {
}

func (p *MagnetoController) IsMutant (chain string) bool  {
	board := p.MakeBoard(chain)
	var dna = new(core.Dna)
	dna.InitDna(board)
	return dna.IsMutant()
}

func (p *MagnetoController) Save(data string, isMutant bool, repository Repository.Repository) int {
	var stats core.Stats
	stats.Dna = data
	stats.IsMutant = isMutant
	var ctx = repository.CreateContext()
	var client = repository.CreateClientDatabase(ctx)
	collection := client.Database("magneto").Collection("stats")
	filterCursor, err := collection.Find(ctx, bson.M{"_dna": data })
	if err != nil {
		log.Fatal(err)
	}
	var statsFiltered []bson.M
	if err = filterCursor.All(ctx, &statsFiltered); err != nil {
		log.Fatal(err)
	}
	if len(statsFiltered) == 0 {
		_, _ = collection.InsertOne(ctx, stats)
	}
	return 0
}

func (p *MagnetoController) MakeBoard(chain string) [][]string {
	array := strings.Split(chain, ",")
	var board [][]string
	board = make([][]string, len(array),len(array))
	for row := 0; row < len(array); row ++ {
		sequence := array[row]
		adn := strings.Split(sequence, "")
		board[row] = make([]string, 0, len(array))
		for column := 0; column < len(adn); column++ {
			board[row] = append(board[row], adn[column])
		}
	}
	return board
}

func (p *MagnetoController) FindAll(repository Repository.Repository) core.StatsResponse {
	var ctx = repository.CreateContext()
	var client = repository.CreateClientDatabase(ctx)
	collection := client.Database("magneto").Collection("stats")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return core.StatsResponse{}
	}
	defer cursor.Close(ctx)
	var countMutantDna float64
	var countHumanDna float64
	for cursor.Next(ctx) {
		var stats core.Stats
		_ = cursor.Decode(&stats)
		if stats.IsMutant {
			countMutantDna = countMutantDna + 1
		}else{
			countHumanDna = countHumanDna + 1
		}
	}
	var ratio = countMutantDna / countHumanDna
	var response = core.StatsResponse{
		CountMutantDna: countMutantDna,
		CountHumanDna:  countHumanDna,
		Ratio:          ratio,
	}
	return response
}
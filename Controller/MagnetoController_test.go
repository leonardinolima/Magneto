package Controller_test

import (
	"github.com/magneto/Controller"
	"github.com/magneto/Repository"
	"github.com/magneto/core"
	"testing"
)

var magneto Controller.MagnetoController

func TestHaveNotArrowConcordance_returnFalse(t *testing.T){
	data := "ZTGCZA,CAGTGC,TTATGT,AGAAGG,ZCCCTA,TCACTG"
	var dna core.Dna
	board := magneto.MakeBoard(data)
	if dna.HaveArrowConcordance(board) {
		t.Error("Expected false "," and received ", true)
		t.Fail()
	}else{
		t.Log("Success")
	}
}

func TestHaveArrowHorizontalConcordance_returnTrue(t *testing.T){
	data := "ATGCGA,CAGTGC,TTATGT,AGAAGG,CCCCTA,TCACTG"
	var dna core.Dna
	board := magneto.MakeBoard(data)
	if !dna.HaveArrowConcordance(board) {
		t.Error("Expected true "," and received ", false)
		t.Fail()
	}else{
		t.Log("Success")
	}
}

func TestHaveArrowVerticalConcordance_returnTrue(t *testing.T){
	data := "ZTGCGA,CAGTGC,TTATGT,AGAAGG,CCCZTA,TCACTG"
	var dna core.Dna
	board := magneto.MakeBoard(data)
	if !dna.HaveArrowConcordance(board) {
		t.Error("Expected true "," and received ", false)
		t.Fail()
	}else{
		t.Log("Success")
	}
}

func TestHaveDescendingDiagonalConcordance_returnTrue(t *testing.T){
	data := "ATGCZA,CAGTGC,TTATGT,AGAAGG,CCCZTA,TCACTG"
	var dna core.Dna
	board := magneto.MakeBoard(data)
	if !dna.HaveDescendingDiagonalConcordance(board) {
		t.Error("Expected true "," and received ", false)
		t.Fail()
	}else{
		t.Log("Success")
	}
}

func TestHaveAscendingDiagonalConcordance_returnTrue(t *testing.T){
	data := "ZTGCZA,CAGTGC,TTATGT,AGTAGG,CTCZTA,TCACTG"
	var dna core.Dna
	board := magneto.MakeBoard(data)
	if !dna.HaveAscendingDiagonalConcordance(board) {
		t.Error("Expected true "," and received ", false)
		t.Fail()
	}else{
		t.Log("Success")
	}
}

func TestMagnetoController_FindAll(t *testing.T) {
	var repo = Repository.MongoMutantRepository{}
	magneto.FindAll(&repo)
}

func TestMagnetoController_Save(t *testing.T) {
	var repo = Repository.MongoMutantRepository{}
	chain := "ZTGCZA,CAGTGC,TTATGT,AGAAGG,ZCCCTA,TCACTG"

	if magneto.Save(chain, false,&repo) != 0{
		t.Error("Expected 0 "," and received ", -1)
		t.Fail()
	}else{
		t.Log("Success")
	}
}
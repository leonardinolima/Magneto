package core

type Dna struct {
	board [][]string
}

func (p *Dna) InitDna(board [][]string){
	p.board = board
}

func (p *Dna) IsMutant() bool {
	if p.HaveArrowConcordance(p.board){
		return true
	}
	if p.HaveDescendingDiagonalConcordance(p.board){
		return true
	}
	if p.HaveAscendingDiagonalConcordance (p.board){
		return true
	}
	return false
}

func (p *Dna) HaveArrowConcordance(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		horizontalConcordance := 1
		verticalConcordance := 1
		for j := 0; j < len(board) -1; j++ {
			if board[i][j] == board[i][j+1] {
				horizontalConcordance = horizontalConcordance + 1
				if horizontalConcordance == 4 {
					return true
				}
			}else{
				horizontalConcordance = 1
			}
			if board[j][i] == board[j+1][i] {
				verticalConcordance = verticalConcordance + 1
				if verticalConcordance == 4{
					return true
				}
			}else{
				verticalConcordance = 1
			}
		}
	}
	return false
}

func (p *Dna) HaveDescendingDiagonalConcordance(board [][]string) bool {
	descendingConcordance := 1
	for row := 0; row < len(board) -1; row++ {
		for column := 0; column < len(board) -1; column++ {
			next := column
			for i := row; i < len(board) -1; i++ {
				if board[i][next] == board[i+1][next+1] {
					descendingConcordance = descendingConcordance + 1
					if descendingConcordance == 4 {
						return true
					}
				}else{
					descendingConcordance = 1
				}
				next = next + 1
				if next == len(board) -1 {
					break
				}
			}
		}
	}
	return false
}

func (p *Dna) HaveAscendingDiagonalConcordance(board [][]string) bool {
	ascendingConcordance := 1
	for row := len(board) -1; row > 0; row-- {
		for column := 0; column < len(board) -1; column++ {
			next := column
			for i := row; i > 0; i-- {
				if board[i][next] == board[i-1][next + 1] {
					ascendingConcordance = ascendingConcordance + 1
					if ascendingConcordance == 4 {
						return true
					}
				}else{
					ascendingConcordance = 1
				}
				next = next + 1
				if next == len(board) -1 {
					break
				}
			}
		}
	}
	return false
}


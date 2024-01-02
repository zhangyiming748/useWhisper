package sql

import "testing"

// go test v -cgo -run TestInit
func TestInit(t *testing.T) {

	SetEngine()
	y := new(Ytdlp)
	y.ID = 1
	y.Status = "success"
	result := y.UpdateStatusById()
	t.Logf("1.%v\n2.%v\n", result.Statement, result.Error)
}

func TestFindByName(t *testing.T) {
	SetEngine()
	y := new(Ytdlp)
	y.Name = "Taboo Big Tits Threesome Virtual Sex"
	Success, Failure, Skip := y.FindDupByName()
	t.Logf("同名文件 成功%d\t失败%d\t跳过%d\n", len(Success), len(Failure), len(Skip))
}

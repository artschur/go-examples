package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func acessarConta(idUsuario int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Acessando conta do usuario %d \n", idUsuario)
	time.Sleep(time.Second * 2) //acesso ao banco de dados
	fmt.Printf("Usuario %d: 3000reais \n", idUsuario)

}

func runWg() {
	var wg sync.WaitGroup
	for user := range 10 {
		wg.Add(1)
		go acessarConta(user, &wg)
	}
	wg.Wait()
}

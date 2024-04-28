package main

import (
	"fmt"
	"sync"
)

func sum(a int, b int) int {
	fmt.Println(a + b)
	return a + b
}

// produceNotifications создает новые уведомления и отправляет их в канал
func produceNotifications(notificationsCh chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 10; i++ {
		notification := fmt.Sprintf("Уведомление записано #%d", i)

		notificationsCh <- notification

		//time.Sleep(1 * time.Second)
	}

	close(notificationsCh)
}

// displayNotifications читает уведы из каналы и печатает их
func displayNotifications(notificationsCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	//for i := 0; i < 11; i++ {
	//notificationsCh <- strconv.Itoa(5)
	//}

	for notificationFromChannel := range notificationsCh {
		fmt.Println("Уведомление прочитано: ", notificationFromChannel)
	}
}

type interface1 interface {
	Puk1()
}
type interface2 interface {
	Puk2()
}
type TestStruct struct {
}

func (ts *TestStruct) Puk1() {
	fmt.Println("puk1")
}
func (ts *TestStruct) Puk2() {
	fmt.Println("puk2")
}

func main() {
	/*u := model.NewUser(2, "Hello")

	go sum(5, 2)
	time.Sleep(1 * time.Second)

	//fmt.Println(sum(5, 2))

	fmt.Printf("%p", &u)
	*/

	/*	var wg sync.WaitGroup

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fmt.Println(i)
			}()
		}

		wg.Wait()*/

	/*	var wg sync.WaitGroup
		var mutex sync.Map
	*/
	/*notifCh := make(chan string)

	wg.Add(2)

	go produceNotifications(notifCh, &wg)
	go displayNotifications(notifCh, &wg)

	wg.Wait()
	fmt.Println("программа завершила работу")
	*/

	/*	for i := 0; i < 2; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				mutex.Lock()
				dict[i] = i
				mutex.Unlock()
			}()
		}

		wg.Wait()
	*/

	var a interface1 = &TestStruct{}
	a.Puk1()
}

package main

//func main() {
//	w := sync.WaitGroup{}
//	m := sync.Mutex{}
//	sum := 0
//	start := make([]int, 10)
//	for i := 0; i < 10; i++ {
//		start[i] = i * 1000
//	}
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			w.Add(1)
//			for k := 0; k <= 1000; k++ {
//				m.Lock()
//				sum += start[i] + k
//				m.Unlock()
//			}
//			w.Done()
//		}(i)
//	}
//	w.Wait()
//	fmt.Println(sum)
//}

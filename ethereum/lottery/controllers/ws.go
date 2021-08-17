package controllers

func init() {
	go SendWsData()
}

func SendWsData() {
	for {
		select { // 保证一次协程调用完成
		case accounts := <-accountsChan:
			for client := range connects {
				err := client.WriteJSON(accounts)
				if err != nil {
					client.Close()
					delete(connects, client)
				}
			}
		}
	}
}

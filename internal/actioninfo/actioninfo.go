package actioninfo
import "log"
type DataParser interface {
	// TODO: добавить методы
	Parse(data string) error
   ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
for _, data := range dataset {

		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных: %v", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации об активности: %v", err)
			continue
		}

		log.Println(info)
	}
}
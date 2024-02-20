build_pkl:
	~/pkl eval -f json main.pkl > cyan.json


run: build_pkl
	go run cyan/main.go
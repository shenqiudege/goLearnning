package main

func main() {
	var way = 0
	var j int = 0
	for {
		switch way {
		case 0:
			{
				print(way)
				way++
			}
			break
		case 1:
			{
				print(way)
				way++
			}
			break
		case 2:
			{
				print(way)
				way = 0
			}
			break
		}
		j++
		if j == 9 {
			break
		}
	}
}

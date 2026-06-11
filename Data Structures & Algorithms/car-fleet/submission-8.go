func carFleet(target int, position []int, speed []int) int {
	positionStack := []int{}

	if len(position) == 1 {
		return 1
	}

	
	for index, _ := range position {
		positionStack = append(positionStack, index)
	}

	// sort the positionStack
	for i:=1; i<len(positionStack);i++ {
		for j:=i;j>0;j--{
			if position[positionStack[j-1]] > position[positionStack[j]] {
				break
			}
			positionStack[j], positionStack[j-1] = positionStack[j-1], positionStack[j] 
		}
	}

	count := 0
	var maxTime float64
	for _, i := range positionStack {
		time := float64(target-position[i]) / float64(speed[i])
		if time > maxTime {
			maxTime = time
			count++
		}
	}

	// result := []float64{}

	// for {
	// 	if len(positionStack)>0 && len(result) > 0 {
	// 		lastElement := positionStack[len(positionStack)-1]
			
	// 		lastResultTime := result[len(result)-1]

	// 		position := position[lastElement]
	// 		speed := speed[lastElement]
	// 		var time float64 = 0.0
	// 		if speed > 0.0 {
	// 			time = float64((target-position)) / float64(speed)
	// 		}

	// 		if lastResultTime < time {
	// 			result = append(result, time)
	// 		}

	// 		positionStack = positionStack[:len(positionStack)-1]
			
	// 	} else if len(result) == 0 {
	// 		lastElement := positionStack[len(positionStack)-1]

	// 		position := position[lastElement]
	// 		speed := speed[lastElement]
	// 		var time float64 = 0.0
	// 		if speed > 0.0 {
	// 			time = float64((target-position)) / float64(speed)
	// 		}
	// 		result = append(result, time)
	// 		positionStack = positionStack[:len(positionStack)-1]
	// 	} else {
	// 		break
	// 	}
	// }

	return count
	
}




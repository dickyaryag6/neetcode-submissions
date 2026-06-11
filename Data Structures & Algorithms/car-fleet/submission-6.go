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
			if position[positionStack[j-1]] < position[positionStack[j]] {
				break
			}
			positionStack[j], positionStack[j-1] = positionStack[j-1], positionStack[j] 
		}
	}

	result := []float64{}

	// if len(positionStack) >= 2 {
		// lastElement := positionStack[len(positionStack)-1]
		// secondLastElement := positionStack[len(positionStack)-2]

		// currentPosition := position[lastElement]
		// currentSpeed := speed[lastElement]

		// aheadPosition := position[secondLastElement]
		// aheadSpeed := speed[secondLastElement]

		// var currentTime float64 = 0.0
		// if currentSpeed > 0 {
		// 	currentTime = float64((target-currentPosition)) / float64(currentSpeed)
		// }

		// var aheadTime float64 = 0.0
		// if aheadSpeed > 0 {
		// 	aheadTime = float64((target-aheadPosition)) / float64(aheadSpeed)
		// }

		// if aheadTime >= currentTime {
		// 	result = append(result, aheadTime)
		// } else {
		// 	result = append(result, currentTime, aheadTime)
		// }
		
		// positionStack = positionStack[:len(positionStack)-2]
	// }

	// fmt.Println(result)

	for {
		if len(positionStack)>0 && len(result) > 0 {
			lastElement := positionStack[len(positionStack)-1]
			
			lastResultTime := result[len(result)-1]

			position := position[lastElement]
			speed := speed[lastElement]
			var time float64 = 0.0
			if speed > 0.0 {
				time = float64((target-position)) / float64(speed)
			}

			// fmt.Println(lastResultTime, " ", time)

			if lastResultTime < time {
				result = append(result, time)
			}

			positionStack = positionStack[:len(positionStack)-1]
			
		} else if len(result) == 0 {
			lastElement := positionStack[len(positionStack)-1]

			position := position[lastElement]
			speed := speed[lastElement]
			var time float64 = 0.0
			if speed > 0.0 {
				time = float64((target-position)) / float64(speed)
			}
			result = append(result, time)
			positionStack = positionStack[:len(positionStack)-1]
		} else {
			break
		}
	}

	return len(result)
	
}




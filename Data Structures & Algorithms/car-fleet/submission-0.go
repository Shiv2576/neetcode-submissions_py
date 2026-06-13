func carFleet(target int, position []int, speed []int) int {

    n := len(position)


    type Car struct {
        pos  int
        time  float64
    }

    cars := make([]Car , n)

    for i := 0 ; i< n ; i++ {
        cars[i] = Car{
            pos : position[i],
            time : float64(target - position[i]) / float64(speed[i]),
        }
    }


    sort.Slice(cars , func(a, b int) bool {
        return cars[a].pos < cars[b].pos
    })


    stack := []float64{}

    for i:= n-1 ; i >= 0 ; i -- {
        t := cars[i].time

        if len(stack) == 0 || t > stack[len(stack)-1] {
            stack = append(stack , t)
        }
    }


    return len(stack)


}

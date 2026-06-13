

func topKFrequent(nums []int, k int) []int {

    frqmap := make(map[int]int)

    for _ , num := range nums {
        frqmap[num] ++
    }

    bucket := make([][]int , len(nums)+ 1)

    for k , v := range frqmap {
        bucket[v] = append(bucket[v] , k)
    }


    res := []int{}

    for i := len(bucket) - 1 ; i >= 0 && k > 0 ; i -- {
        for _ , num := range bucket[i] {
            res = append(res , num)
            k --
            if k == 0 {

                break
            }
        }
    }

    return res




}

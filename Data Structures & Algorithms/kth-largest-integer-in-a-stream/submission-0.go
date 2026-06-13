type KthLargest struct {
    key int
    set []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest {
        key : k,
        set : nums,
    }
}


func (this *KthLargest) Add(val int) int {
    this.set = append(this.set , val)

    n := len(this.set)

    quicksort(this.set ,0 , n-1)

    return this.set[n-this.key]
}


func quicksort(arr []int, low, high int) {
    if low < high {
        p := partitionAsc(arr , low , high)
        quicksort(arr , low , p-1)
        quicksort(arr , p+1 , high)
    }
}

func partitionAsc(arr []int , low, high int) int {
    pivot := arr[high]
    i := low 

    for j := low ; j < high ; j++ {
        if arr[j] <= pivot {
            arr[i] , arr[j] = arr[j] , arr[i]
            i++
        }

    }

    arr[i] , arr[high] = arr[high], arr[i]

    return i
}
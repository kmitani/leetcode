func twoSum(nums []int, target int) []int {
    var i, j int
    var m map[int]int = make(map[int]int) //https://qiita.com/kotobuki5991/items/1f31b95223976e2b5f8b
    for i = 0; i < len(nums); i++{
        diff := target - nums[i]
        j, ok := m[diff]
        if ok && j != i{
            return []int{i, j}
        }
        m[nums[i]] = i
    }
    return []int{i, j}
}

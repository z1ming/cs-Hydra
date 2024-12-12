func compareVersion(version1 string, version2 string) int {
    arr1 := strings.Split(version1, ".")
    arr2 := strings.Split(version2, ".")
    for i := 0; i < len(arr1) || i < len(arr2); i++ {
        x, y := 0, 0
        if i < len(arr1) {
            x, _ = strconv.Atoi(arr1[i])
        }
        if i < len(arr2) {
            y, _ = strconv.Atoi(arr2[i])
        }
        if x > y {
            return 1
        }
        if x < y {
            return -1
        }
    }
    return 0
}

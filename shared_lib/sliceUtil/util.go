package sliceUtil


func DropElementAtIndex (slice [][]string, index int) [][]string {

	if len(slice) > 1 {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return [][]string{}
	}
}
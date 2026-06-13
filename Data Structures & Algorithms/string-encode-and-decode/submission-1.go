type Solution struct{
}

func (s *Solution) Encode(strs []string) string {
    encoded := ""

    for _ , ch := range strs {
        encoded += strconv.Itoa(len(ch)) + "#" + ch
    }


    return encoded
}

func (s *Solution) Decode(encoded string) []string {

    decoded := []string{}
    i := 0

    for i < len(encoded) {
        j := i

        for encoded[j] != '#' {
            j++
        }


        length , _ := strconv.Atoi(encoded[i:j])
        j++

        decoded = append(decoded , encoded[j:j+length])

        i = j+length

    }

    return decoded

}

package mypckg

// // read line by line into memory
// // all file contents is stores in lines[]
// func ReadLines(path string) ([]string, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)

// 	//scanner := bufio.NewScanner(strings.NewReader(file))

// 	scanner.Split(bufio.ScanWords)

// 	for scanner.Scan() {
// 		//lines = append(lines, scanner.Text())
// 		println(scanner.Text())
// 	}
// 	//return lines, scanner.Err()
// 	return lines, scanner.Err()
// }

// func DoEet(layoutNamesFile string) {
// 	// open file for reading
// 	// read line by line
// 	lines, err := ReadLines(layoutNamesFile)
// 	if err != nil {
// 		log.Fatalf("readLines: %s", err)
// 	}
// 	// print file contents
// 	for i, line := range lines {
// 		fmt.Println(i, line)
// 	}
// }

// func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

//     // Return nothing if at end of file and no data passed
//     if atEOF && len(data) == 0 {
//         return 0, nil, nil
//     }

//     // Find the index of the input of a newline followed by a
//     // pound sign.
//     if i := strings.Index(string(data), "\n#"); i >= 0 {
//         return i + 1, data[0:i], nil
//     }

//     // If at end of file with data return the data
//     if atEOF {
//         return len(data), data, nil
//     }

//     return
// }

// func ReadLayoutNames(layoutNamesFile string) {
// 	b, err := ioutil.ReadFile(layoutNamesFile)
// 	// can file be opened?
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	// convert bytes to string
// 	str := string(b)

// 	// show file datas
// 	fmt.Println(str)
// }

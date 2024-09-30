package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Author struct {
	Id   int
	Name string
}

type Book struct {
	Id     int
	Title  string
	Author Author
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	//run struct
	//structBook()
	//run max gap of binay string
	//maxGapBinary(1041)
	//test marshal
	//testMarshal()
	//
	//cyclicRotation()
	//findMinLetter("axxz", "yzwy")
	//append item in array
	//tryAppend()
	////////////////////////
	//ABC()
	//Seat()
	//DICE()
	//Conform(876, 678, 786)
	//SmallestNumber(16)
	//AB("ABBBBBABBA")
	//conLetters("baaaa")
	//CountMove("><^v")
	//Count min Rollers

	// CountMinRollers(X, Y, 2)
	//
	//Task1("zyzyzyz")
	// P := []int{4, 4, 2, 4, 1}
	// S := []int{5, 5, 2, 5, 2}
	// Task2(P, S)
	//Task3("..xxx...xx..xxxx.xxxx", 9)
	//fmt.Println(int('0'), int('9'), int('A'), int('Z'))
	//Hero_Task1("abcd", "ca")
	connections := []string{"fred:joe", "joe:mary", "mary:fred", "mary:june", "june:bill", "june:fred"}
	Hero_Task2(connections, "fred", "june")

}

func Hero_Task2(connections []string, name1 string, name2 string) int {
	iResult := 0
	//
	name1 = strings.ToLower(name1)
	name2 = strings.ToLower(name2)
	if name1 == "" || name2 == "" {
		return 0
	}

	//
	iResult += fnCountConnections(connections, 0, name1, name2, 0)
	//
	fmt.Println(iResult)
	//
	return iResult

}

func fnCountConnections(con []string, iIndex int, p1 string, p2 string, iResult int) int {
	str := con[iIndex]
	iRecur := iIndex
	if iIndex >= len(con)-1 {
		return iResult
	}
	//
	iName := ""
	if strings.Index(str, p1) == 0 {
		//fmt.Println(str, "-", p1)
		iName = str[strings.Index(str, ":")+1 : len(str)]
		iResult += 1
		if iName == p2 {
			return iResult
		} else {
			iRecur += 1
			return fnCountConnections(con, iRecur, iName, p2, iResult)
		}
	} else {
		iRecur += 1
		return fnCountConnections(con, iRecur, p1, p2, iResult)
	}

	return iResult
}

func Hero_Task1(X string, Y string) bool {
	iResult := true
	X = strings.ToUpper(X)
	Y = strings.ToUpper(Y)
	lenX := len(X)
	lenY := len(Y)

	for i := 0; i < lenX; i++ {
		if int(X[i]) < 48 || int(X[i]) > 90 {
			return false
		}
	}
	for i := 0; i < lenY; i++ {
		if int(Y[i]) < 48 || int(Y[i]) > 90 {
			return false
		}
	}
	////
	firstIndex := 0
	//lastIndex := 0

	for i := 0; i < lenY; i++ {
		for j := i; j < lenX; j++ {
			fmt.Println(fsubstring(X, j, lenX-i))
			firstIndex = strings.Index(fsubstring(X, j, lenX-i), string(Y[i]))
			//if Y[i] == X[j] {
			fmt.Println(firstIndex)
			if firstIndex > -1 {
				iResult = true
				fmt.Println(iResult)
				break
			} else {
				iResult = false
				fmt.Println(iResult)
			}
		}

	}
	////
	//fmt.Println(iResult)
	return iResult
}

func NAB_Task3(S string, B int) int {

	//validate
	if B < 0 || B > 200000 {
		return 0
	}
	N := len(S)
	if N < 1 || N > 100000 {
		return 0
	}

	S = strings.ToLower(S)
	for i := 0; i < N; i++ {
		//fmt.Println(string(S[i]))
		if string(S[i]) != "." && string(S[i]) != "x" {
			return 0
		}
	}

	//

	arrHole := []int{}
	iC := 0
	for i := 0; i < N; i++ {
		if string(S[i]) == "x" {
			iC += 1
			if i == N-1 {
				arrHole = append(arrHole, iC)
			}
		} else {
			arrHole = append(arrHole, iC)
			iC = 0
		}
	}
	//

	arrHole = sortArrayLonNho(arrHole)
	fmt.Println(arrHole)
	iResult := 0
	B1 := B //cost
	fmt.Println(B1)
	for i := 0; i < len(arrHole); i++ {
		if B1-1 > arrHole[i] {
			iResult += arrHole[i]
			B1 = B1 - arrHole[i] - 1
		} else if B1-1 == arrHole[i] {
			fmt.Println(B1)
			iResult += arrHole[i]
			break
		} else {
			iResult += B1 - 1
			break
		}
	}
	fmt.Println(iResult)
	//
	return iResult
}

func NAB_Task2(P []int, S []int) int {

	//
	N1 := len(P)
	N2 := len(S)
	if N1 < 0 || N1 > 100000 {
		return 0
	}
	if N2 < 0 || N2 > 100000 {
		return 0
	}
	if N1 != N2 {
		return 0
	}
	for i := 0; i < N1; i++ {
		if P[i] < 1 || P[i] > 9 {
			return 0
		}
		if S[i] < 1 || S[i] > 9 {
			return 0
		}
	}

	//
	iResult := 0
	P = sortArrayLonNho(P)
	S = sortArrayLonNho(S)
	//people
	pp := 0
	for i := 0; i < N1; i++ {
		pp += P[i]
	}

	ss := 0
	for j := 0; j < N2; j++ {
		ss += S[j]
		iResult += 1
		if ss >= pp {
			break
		}
	}
	fmt.Println(iResult)
	//
	return iResult
}

func fsubstring(S string, begin int, length int) string {
	asRunes := []rune(S)

	if begin >= len(asRunes) {
		return ""
	}

	if begin+length > len(asRunes) {
		length = len(asRunes) - begin
	}

	return string(asRunes[begin : begin+length])
}

func NAB_Task1(S string) int {
	S = strings.ToLower(S)
	N := len(S)
	if N <= 0 || N > 200 {
		return 0
	}
	for i := 0; i < N; i++ {
		if S[i] < 97 || S[i] > 122 {
			return 0
		}
	}
	//
	iResult := N
	firstIndex := 0
	lastIndex := 0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			//fmt.Println(i, j)
			S1 := fsubstring(S, i, (j - i))
			fmt.Println(S1)
			firstIndex = strings.Index(S, S1)
			lastIndex = strings.LastIndex(S, S1)
			fmt.Println(firstIndex, lastIndex)
			if firstIndex == lastIndex {
				if len(S1) < iResult {
					iResult = len(S1)
				}
			}
		}
	}
	fmt.Println(iResult)
	//
	return iResult
}

func CountMinRollers(X []int, Y []int, W int) int {
	// Implement your solution here
	iResult := 0
	//validate
	lenX := len(X)
	lenY := len(Y)
	if lenX < 1 || lenX > 100000 {
		iResult = 0
	}
	if lenY < 1 || lenY > 100000 {
		iResult = 0
	}
	if W <= 0 || W > 1000000000 {
		iResult = 0
	}
	if lenX != lenY {
		iResult = 0
	}
	//
	maxX := 0
	maxY := 0
	for i := 0; i < lenX; i++ {
		if X[i] < 0 || X[i] > 1000000000 {
			return 0
		}
		if Y[i] < 0 || Y[i] > 1000000000 {
			return 0
		}
		if maxX < X[i] {
			maxX = X[i]
		}
		if maxY < Y[i] {
			maxY = Y[i]
		}
	}
	//
	countRollers := 0
	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(W)

	for i := 0; i < maxX/W; i++ {
		fmt.Println("W[", i, "]=", i)
		for j := 0; j < lenX; j++ {
			fmt.Println("X", j, ":", X[j])
			if X[j] > W*i && X[j] <= (i+1)*W {
				countRollers += 1
				fmt.Println("C:", countRollers)
				break
			}
		}
	}
	iResult = countRollers
	fmt.Println(countRollers)
	//
	return iResult
}

func CountMove(S string) int {
	iResult := 0
	//
	LenS := len(S)
	//validate
	if LenS > 50 || LenS == 0 {
		iResult = 0
	}
	//
	for i := 0; i < LenS; i++ {
		if S[i] != '^' && S[i] != 'v' && S[i] != '>' && S[i] != '<' {
			iResult = 0
		}
	}
	//
	iCount := 0
	if S[0] == '<' || S[0] == 'v' || S[0] == '^' {
		iCount += 1
	}
	if S[LenS-1] == '>' || S[LenS-1] == 'v' || S[LenS-1] == '^' {
		iCount += 1
	}
	for i := 1; i < LenS-1; i++ {
		if S[i] == 'v' || S[i] == '^' {
			fmt.Println(i, iCount)
			iCount += 1
		}
		if (S[i-1] == '<' || S[i-1] == 'v' || S[i-1] == '^') && S[i] == '<' && S[0] != '>' {
			fmt.Println(i, "-", iCount)
			iCount += 1
		}
	}
	//
	iResult = iResult + iCount
	fmt.Println(iResult)
	//
	return iResult
}

func conLetters(S string) int {
	S = strings.ToLower(S)
	N := len(S)
	iResult := 0
	//validate
	if N > 200000 {
		iResult = 0
	}
	for i := 0; i < len(S); i++ {
		if S[i] != 'a' && S[i] != 'b' {
			iResult = 0
		}
	}
	//3 first letters
	if N < 3 {
		return 0
	} else {
		if S[0] == 'a' && S[1] == 'a' && S[2] == 'a' {
			iResult += 1
		}
		if S[0] == 'b' && S[1] == 'b' && S[2] == 'b' {
			iResult += 1
		}
	}

	if N >= 3 {
		for i := 3; i < N; i += 3 {
			//
			iRemind := 0
			iCountA := 0
			iCountB := 0
			//
			if i >= 1 {
				//fmt.Println(i)
				if S[i-2] == 'a' && S[i-1] == 'a' {
					iCountA += 2
					iCountB = 0
				}
				if S[i-2] == 'b' && S[i-1] == 'b' {
					iCountB += 2
					iCountA = 0
				}
				if S[i-2] != S[i-1] {
					if S[i-1] == 'a' {
						iCountA += 1
						iCountB = 0
					} else {
						iCountB += 1
						iCountA = 0
					}
				}
				//

				if iCountA == 2 || iCountB == 2 {
					if S[i] == S[i-1] {
						iRemind = 2
					}
				}
				if iCountA == 1 || iCountB == 1 {
					if S[i] == S[i-1] {
						iRemind = 1
					}
				}
				//

			}
			//
			if iRemind == 2 {
				if iCountA == 2 && S[i] == 'a' {
					//
					iResult += 1
				}
				if iCountB == 2 && S[i] == 'b' {
					//
					iResult += 1
				}
			}
			if iRemind == 1 {
				if iCountA == 1 && S[i] == 'a' && S[i+1] == 'a' {
					//
					iResult += 1
				}
				if iCountB == 1 && S[i] == 'b' && S[i+1] == 'b' {
					//
					iResult += 1
				}
			}

		}
	}

	fmt.Println(iResult)
	//
	return iResult
}

func AB(S string) int {
	iResult := 0
	S = strings.ToUpper(S)
	//
	iLen := len(S)
	if iLen > 100000 {
		iResult = 0
	}

	for i := 0; i < iLen; i++ {
		if S[i] != 'A' || S[i] != 'B' {
			iResult = 0
			break
		}
	}
	countA := 0
	countB := 0
	for i := 0; i < iLen; i++ {
		if S[i] == 'A' {
			countA += 1
		} else {
			countB += 1
		}
	}
	if countA == iLen || countB == iLen {
		iResult = 0
	}
	//
	//firstIndexB := strings.IndexAny(S, "B")
	lastIndexA := strings.LastIndexAny(S, "A")

	for i := 0; i < lastIndexA; i++ {
		if S[i] != 'A' {
			iResult += 1
		}
	}
	for i := lastIndexA + 1; i < len(S); i++ {
		if S[i] != 'B' {
			iResult += 1
		}
	}
	fmt.Println(iResult)
	//
	return iResult
}

func SmallestNumber(N int) int {
	if N < 0 {
		return 0
	} else if N > 50 {
		return 50
	} else if N < 10 {
		return N
	} else {

		//iBool := false
		//iNum := 10
		for iNum := 10; iNum < 1000000; iNum++ {
			str := strconv.Itoa(iNum)
			fmt.Println(str)
			arrNum := []int{}
			for i := 0; i < len(str); i++ {
				if s, err := strconv.Atoi(string(str[i])); err == nil {
					arrNum = append(arrNum, s)
					//fmt.Println(str[i])
					fmt.Println(s)
				}
			}

			fmt.Println(arrNum)
			iCheck := 0
			for j := 0; j < len(arrNum); j++ {
				iCheck += arrNum[j]
			}
			//fmt.Println(iCheck)

			if iCheck == N {
				return iNum
			}
		}

	}
	return 0
}

func Conform(A int, B int, C int) int {
	def := "000000000000000000000000000000"

	if A > 1073741823 || A < 0 {
		return 0
	}
	if B > 1073741823 || B < 0 {
		return 0
	}
	if C > 1073741823 || C < 0 {
		return 0
	}
	biA := fmt.Sprintf("%b", A)
	biB := fmt.Sprintf("%b", B)
	biC := fmt.Sprintf("%b", C)
	if len(biA) < 30 {
		biA = def[0:len(def)-len(biA)] + string(biA)
	}
	if len(biB) < 30 {
		biB = def[0:len(def)-len(biB)] + string(biB)
	}
	if len(biC) < 30 {
		biC = def[0:len(def)-len(biC)] + string(biC)
	}

	fmt.Println(biA)
	fmt.Println(biB)
	fmt.Println(biC)
	diffCount := 0
	iTotal := 0
	if biA == biB && biA == biC {
		diffCount = 30
	} else {
		for i := 0; i < len(biA); i++ {
			if biA[i] != biB[i] || biA[i] != biC[i] {
				diffCount += 1
			}
		}
	}
	if diffCount == 30 {
		iTotal = 1073741824
	} else {
		iTotal = diffCount * 2
	}

	fmt.Println(iTotal)
	return iTotal
}

func DICE() {
	iResult := []int{}
	A := []int{6, 1}
	F := 1
	M := 1 //iResult
	iDivA := len(A) + F
	iTotalA := 0
	for i := 0; i < len(A); i++ {
		iTotalA += A[i]
	}
	X := iDivA*M - iTotalA
	//
	if X <= 0 {
		fmt.Println(X, M)
		iResult = append(iResult, 0)
		fmt.Println(iResult)
	}
	if int(X/F) > 6 {
		fmt.Println(X, M)
		iResult = append(iResult, 0)
		fmt.Println(iResult)
	}
	//
	for k := 1; k <= F; k++ {
		if X >= 6 && (X-6) >= F-k {
			iResult = append(iResult, 6)
			X = X - 6
		} else if X >= 5 && (X-5) >= F-k {
			iResult = append(iResult, 5)
			X = X - 5
		} else if X >= 4 && (X-4) >= F-k {
			iResult = append(iResult, 4)
			X = X - 4
		} else if X >= 3 && (X-3) >= F-k {
			iResult = append(iResult, 3)
			X = X - 3
		} else if X >= 2 && (X-2) >= F-k {
			iResult = append(iResult, 2)
			X = X - 2
		} else {
			iResult = append(iResult, 1)
			X = X - 1
		}
	}
	fmt.Println(X)
	fmt.Println(iResult)
}

func Seat() {
	P := []int{1, 4, 1}
	S := []int{1, 5, 1}
	totalPeople := 0
	for i := 0; i < len(P); i++ {
		totalPeople += P[i]
	}

	fmt.Println(S)
	iCount := 0
	iAvai := totalPeople
	for j := 0; j < len(S); j++ {
		iAvai = iAvai - S[j]

		if iAvai < 0 {
			break
		} else {
			iCount += 1
		}
	}
	fmt.Println(iCount)
}

func ABC() {
	A := []int{1, 1, 3, 4, 4, 4}
	// Implement your solution here
	iTotal := len(A)
	if iTotal%2 != 0 {
		iTotal = (iTotal - 1) / 2
	} else {
		iTotal = iTotal / 2
	}
	fmt.Println(iTotal)
}

func tryAppend() {
	a := []int{1, 2, 3}
	a = append(a, 9, 2)
	fmt.Println(a)
}

func CountLetters(S string) map[string]int {
	mapS := make(map[string]int)
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := 0; i < len(alphabet); i++ {
		iCountS := strings.Count(S, alphabet[i])
		if iCountS > 0 {
			mapS[alphabet[i]] = iCountS
		}
	}
	return mapS
}

func findMinLetter(P string, Q string) int {
	P = strings.ToLower(P)
	Q = strings.ToLower(Q)
	//validate
	if len(P) >= 2000000 || len(Q) > 2000000 {
		return 0
	}
	if len(P) != len(Q) {
		return 0
	}
	mapP := CountLetters(P)
	mapQ := CountLetters(Q)

	fmt.Println(P)
	// fmt.Println(mapP)
	fmt.Println(Q)
	// fmt.Println(mapQ)
	maxLetters := 0
	//
	if len(mapP) > len(mapQ) {
		maxLetters = len(mapQ)
		for i := 0; i < len(Q); i++ {
			if Q[i] != P[i] {
				Q = strings.Replace(Q, string(Q[i]), string(P[i]), 1)
			}
			tmpMap := CountLetters(Q)
			tmpCount := len(tmpMap)
			if tmpCount < maxLetters {
				maxLetters = tmpCount
			}
			T := Q
			for j := i + 1; j < len(T); j++ {
				if T[j] != P[i] {
					T = strings.Replace(T, string(T[j]), string(P[j]), 1)
				}
				tmpMap := CountLetters(T)
				tmpCount := len(tmpMap)
				if tmpCount < maxLetters {
					maxLetters = tmpCount
				}
				fmt.Println(T, ":", maxLetters, "vs", tmpCount)
			}
			fmt.Println(Q, ":", maxLetters, "vs", tmpCount)
		}
	} else {
		maxLetters = len(mapP)
		fmt.Println("-------")
		for i := 0; i < len(P); i++ {
			if P[i] != Q[i] {
				T := strings.Replace(P, string(P[i]), string(Q[i]), 1)
				fmt.Println(T)
			}
			// tmpMap := make(map[string]int)
			// tmpCount := 0

			// for j := i + 1; j < len(T1); j++ {
			// 	if T[j] != Q[i] {
			// 		T = strings.Replace(T, string(T[j]), string(Q[j]), 1)
			// 	}
			// 	tmpMap = CountLetters(T)
			// 	tmpCount = len(tmpMap)
			// 	if tmpCount < maxLetters {
			// 		maxLetters = tmpCount
			// 	}
			// 	fmt.Println(T, ":", maxLetters, "vs", tmpCount)
			// }
			// fmt.Println(P, ":", maxLetters, "vs", tmpCount)
		}
	}
	fmt.Println("max:", maxLetters)
	return maxLetters
	//
}

func Copy_findMinLetter(P string, Q string) int {
	P = strings.ToLower(P)
	Q = strings.ToLower(Q)
	//validate
	if len(P) >= 2000000 || len(Q) > 2000000 {
		return 0
	}
	if len(P) != len(Q) {
		return 0
	}
	//
	T := ""
	if P == Q {
		T = P
	} else {
		T = P + Q
	}
	//
	arrDist := []string{}
	arrValue := []int{}
	for i := 0; i < len(T); i++ {
		//check if letter is not 'a' or 'b'
		if int(T[i]) < 97 || int(T[i]) > 122 {
			return 0
		}
		for j := i + 1; j < len(T); j++ {
			if T[i] == T[j] {
				break
			}
			if j == len(T)-1 {
				arrDist = append(arrDist, string(T[i]))
			}
		}
		if i == len(T)-1 {
			arrDist = append(arrDist, string(T[i]))
		}
	}
	fmt.Println(P, "-", Q, "-", T)
	fmt.Println(arrDist)
	//
	if len(arrDist) > 20 {
		return 0
	}
	//
	for i := 0; i < len(arrDist); i++ {
		iCount := 0
		for j := 0; j < len(T); j++ {
			if string(arrDist[i]) == string(T[j]) {
				iCount += 1
			}
		}
		arrValue = append(arrValue, iCount)
	}
	//
	sortArrayLonNho(arrValue)
	fmt.Println(arrValue)

	iLen := 0
	iCount := 0
	for i := 0; i < len(arrValue); i++ {
		if iLen+arrValue[i] <= len(P) {
			iLen += arrValue[i]
			iCount += 1
		}

		if iLen >= len(P) {
			break
		}
		//fmt.Println(iLen)
	}
	//

	fmt.Println(iCount)
	return iCount
	//
}

func sortArrayNhoLon(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//sort tu nho > lon
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func sortArrayLonNho(arr []int) []int {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			//sort tu nho > lon
			//if arr[j] > arr[j+1] {
			//sort tu lon > nho
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// cyclic rotation of array of N integer
func cyclicRotation() {
	arr := []int{3, 2, 6, 7, 8, 0, 6, 9, 4}
	//fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		tmp := arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[len(arr)-2-i]
		arr[len(arr)-2-i] = tmp

	}
	//fmt.Println(arr)
}

// desc: marshal and unmarshal
func testMarshal() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

// desc: struct and unmarshal
func structBook() {
	data_Book := `[{"Id":1,"Title":"Title 1","Author":{"Id":1,"Name":"Jacky 1"}},{"Id":3,"Title":"Title 3","Author":{"Id":3,"Name":"Jacky 3"}},{"Id":2,"Title":"Title 2","Author":{"Id":2,"Name":"Jacky 2"}}]`

	var books []Book

	json.Unmarshal([]byte(data_Book), &books)

	fmt.Println(books)
	fmt.Println("--------")
	for i := 0; i < len(books); i++ {
		fmt.Println(books[i].Id)
	}
	fmt.Println("--------")
	sortBook(books)
	fmt.Println(books)
}

func aboutInterface() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func aboutMap() {
	m := make(map[int]string)
	m[1] = "3"
	m[2] = "5"
	m[3] = "7"
	m[4] = "9"
	fmt.Println(m[1])
}

func aboutString() {
	str := "tan lam ngan went to table tennis club"
	fmt.Println(strings.ToLower(str))
}

// desc: sort array
func sortArray() {
	// //
	sl := []string{"mumbai", "london", "tokyo", "seattle"}
	num := []int{1, 2, 6, 3, 2, 5, 6, 9}

	fmt.Println(sl)
	sort.Sort(sort.StringSlice(sl))
	fmt.Println(sl)
	fmt.Println(num)
	// sort.Sort(sort.IntSlice(num))
	slices.Sort(num)
	fmt.Println(num)
	slices.Reverse(num)
	fmt.Println(num)
}

// sort struct by 1 element in struct
func sortBook(b []Book) {
	sort.Slice(b, func(i, j int) bool {
		return b[i].Id < b[j].Id
	})
}

// desc: list prime numbers of N integer
func printPrimeNumber() {
	// print prime number less than a number
	var n = 100
	fmt.Println("The prime numbers smaller than ", n)
	fmt.Println(1)
	for i := 3; i <= n; i++ {
		for j := 2; j <= i; j++ {
			if (i % j) != 0 {
				continue
			} else if i == j {
				fmt.Println(i)
			} else {
				break
			}

		}
	}
}

// desc: max length of gap of binary string
func maxGapBinary(N int) int {
	strBinary := fmt.Sprintf("%b", N)
	fmt.Println(string(strBinary))
	maxrCnt := 0
	currCnt := 0
	for i := 0; i < len(strBinary); i++ {
		if string(strBinary[i]) == "0" {
			currCnt += 1
			//fmt.Println(currCnt)
		} else {
			if currCnt >= maxrCnt {
				maxrCnt = currCnt

			}
			currCnt = 0
		}

	}
	fmt.Println(maxrCnt)
	return maxrCnt
}

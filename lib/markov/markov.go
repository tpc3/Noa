package markov

import (
	crand "crypto/rand"
	"log"
	"math"
	"math/big"
	"math/rand"

	"github.com/bluele/mecab-golang"
)

func ParseToNode(m *mecab.MeCab, input string) []string {
	words := []string{}
	tg, err := m.NewTagger()
	if err != nil {
		log.Fatal("New tagger error: ", err)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(input)
	if err != nil {
		log.Fatal("New Lattice error: ", err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		if node.Surface() != "" {
			words = append(words, node.Surface())
		}
		if node.Next() != nil {
			break
		}
	}

	return words
}

func GetMarkovBlocks(words []string) [][]string {
	res := [][]string{}
	resHead := []string{}
	resEnd := []string{}

	if len(words) < 3 {
		return res
	}

	resHead = []string{"#This is empty#", words[0], words[1]}
	res = append(res, resHead)

	for i := 1; i < len(words)-2; i++ {
		markovBlock := []string{words[i], words[i+1], words[i+2]}
		res = append(res, markovBlock)
	}

	resEnd = []string{words[len(words)-2], words[len(words)-1], "#This is empty#"}
	res = append(res, resEnd)

	return res
}

func FindBlocks(array [][]string, target string) [][]string {
	blocks := [][]string{}
	for _, s := range array {
		if s[0] == target {
			blocks = append(blocks, s)
		}
	}

	return blocks
}

func ConnentBlocks(array [][]string, dist []string) []string {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	i := 0

	if len(array) > 0 {
		for _, word := range array[rand.Intn(len(array))] {
			if i != 0 {
				dist = append(dist, word)
			}
			i += 1
		}
	}

	return dist
}

func MarkovChainExec(array [][]string) []string {
	ret := []string{}
	block := [][]string{}
	count := 0

	block = FindBlocks(array, "#This is empty#")
	ret = ConnentBlocks(block, ret)

	for len(ret) > 0 && ret[len(ret)-1] != "#This is empty#" {
		block = FindBlocks(array, ret[len(ret)-1])
		if len(block) == 0 {
			break
		}
		ret = ConnentBlocks(block, ret)

		count++
		if count == 150 {
			break
		}
	}

	return ret
}

func TextGenerate(array []string) string {
	ret := ""
	for _, s := range array {
		if s == "#This is empty#" {
			continue
		}

		if len([]rune(ret)) >= 90 {
			break
		}

		ret += s
	}

	return ret
}

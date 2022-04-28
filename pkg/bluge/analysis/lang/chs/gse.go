package chs

import (
	"github.com/blugelabs/bluge/analysis"
	"github.com/go-ego/gse"

	"github.com/zinclabs/zinc/pkg/bluge/analysis/lang/chs/analyzer"
	"github.com/zinclabs/zinc/pkg/bluge/analysis/lang/chs/token"
	"github.com/zinclabs/zinc/pkg/bluge/analysis/lang/chs/tokenizer"
	"github.com/zinclabs/zinc/pkg/zutils"
)

func NewGseStandardAnalyzer() *analysis.Analyzer {
	return analyzer.NewStandardAnalyzer(seg)
}

func NewGseSearchAnalyzer() *analysis.Analyzer {
	return analyzer.NewSearchAnalyzer(seg)
}

func NewGseStandardTokenizer() analysis.Tokenizer {
	return tokenizer.NewStandardTokenizer(seg)
}
func NewGseSearchTokenizer() analysis.Tokenizer {
	return tokenizer.NewSearchTokenizer(seg)
}

func NewGseStopTokenFilter() analysis.TokenFilter {
	return token.NewStopTokenFilter(seg, nil)
}

var seg *gse.Segmenter

func init() {
	seg = new(gse.Segmenter)
	enable := zutils.GetEnvToBool("ZINC_PLUGIN_GSE_ENABLE", "FALSE")     // false / true
	embed := zutils.GetEnvToUpper("ZINC_PLUGIN_GSE_DICT_EMBED", "SMALL") // small / big
	if enable {
		if embed == "BIG" {
			seg.LoadDictEmbed("zh_s")
			seg.LoadStopEmbed()
		} else {
			seg.LoadDictStr(_dictCHS)
			seg.LoadStopStr(_dictStop)
		}
	} else {
		seg.LoadDictStr(`zinc`)
		seg.LoadStopStr(_dictStop)
	}
	seg.Load = true
	seg.SkipLog = true

	// load user dict
	dataPath := zutils.GetEnv("ZINC_PLUGIN_GSE_DICT_PATH", "./plugins/gse/dict")
	userDict := dataPath + "/user.txt"
	if ok, _ := zutils.IsExist(userDict); ok {
		seg.LoadDict(userDict)
	}
	stopDict := dataPath + "/stop.txt"
	if ok, _ := zutils.IsExist(stopDict); ok {
		seg.LoadStop(stopDict)
	}
}

package search

import (
	"fmt"
	"log"

	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
)

var (
	// searcher is coroutine safe
	searcher = riot.Engine{}

	text  = "Google Is Experimenting With Virtual Reality Advertising"
	text1 = `Google accidentally pushed Bluetooth update for Home
	speaker early`
	text2 = `Google is testing another Search results layout with 
	rounded cards, new colors, and the 4 mysterious colored dots again`

	opts = types.EngineOpts{
		Using: 1,
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.DocIdsIndex,
		},
		UseStore:    true,
		StoreFolder: "../data/store", //存储路径
		StoreEngine: "bg",            //存储引擎 bg: badger, lbd: leveldb, bolt: bolt
		// GseDict: "../../data/dict/dictionary.txt", //字典
		GseDict:       "../../testdata/test_dict.txt",
		StopTokenFile: "../../data/dict/stop_tokens.txt", //停用词
	}
)

func initEngine() {
	// gob.Register(MyAttriStruct{})

	// var path = "./riot-index"

	// 搜索引擎初始化
	searcher.Init(opts)
	defer searcher.Close()
	// os.MkdirAll(path, 0777)

	// 添加文档索引
	// Add the document to the index, docId starts at 1
	searcher.Index("1", types.DocData{Content: text})
	searcher.Index("2", types.DocData{Content: text1})
	searcher.Index("3", types.DocData{Content: text2})
	searcher.Index("5", types.DocData{Content: text2})

	//  删除文档索引
	searcher.RemoveDoc("5")

	// 等待索引刷新
	// Wait for the index to refresh
	searcher.Flush()

	log.Println("Created index number: ", searcher.NumDocsIndexed())
}

// 从持久化的数据中恢复
func restoreIndex() {
	// var path = "./riot-index"
	// 搜索引擎初始化
	searcher.Init(opts)
	defer searcher.Close()
	// os.MkdirAll(path, 0777)

	// Wait for the index to refresh
	// 等待索引刷新
	searcher.Flush()

	log.Println("recover index number: ", searcher.NumDocsIndexed())
}

func Search() {
	initEngine()
	// restoreIndex()

	sea := searcher.Search(types.SearchReq{
		Text: "google testing",
		RankOpts: &types.RankOpts{
			OutputOffset: 0,
			MaxOutputs:   100,
		}})

	fmt.Println("search response: ", sea, "; docs = ", sea.Docs)

	// os.RemoveAll("riot-index")
}

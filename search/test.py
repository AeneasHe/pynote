from whoosh.qparser import QueryParser
from whoosh.index import create_in
from whoosh.index import open_dir
from whoosh.fields import *
from jieba.analyse import ChineseAnalyzer
from get_comment import SQL
from whoosh.sorting import FieldFacet

analyser = ChineseAnalyzer()  # 导入中文分词工具

# 创建索引结构
schema = Schema(
    phone_name=TEXT(stored=True, analyzer=analyser),
    price=NUMERIC(stored=True),
    phoneid=ID(stored=True),
)


# path 为索引创建的地址，indexname为索引名称
ix = create_in("path", schema=schema, indexname="indexname")
writer = ix.writer()
writer.add_document(phone_name="name", price="price", phoneid="id")  #  此处为添加的内容
print("建立完成一个索引")
writer.commit()


# 以上为建立索引的过程
new_list = []
index = open_dir("indexpath", indexname="comment")  # 读取建立好的索引
with index.searcher() as searcher:
    parser = QueryParser("要搜索的项目，比如“phone_name", index.schema)
    myquery = parser.parse("搜索的关键字")

    # 按序排列搜索结果
    facet = FieldFacet("price", reverse=True)

    # limit为搜索结果的限制，默认为10，详见博客开头的官方文档
    results = searcher.search(myquery, limit=None, sortedby=facet)

    for result1 in results:
        print(dict(result1))
        new_list.append(dict(result1))

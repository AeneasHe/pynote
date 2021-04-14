import logging
import os
import shutil
from django.conf import settings

from whoosh.fields import Schema, ID, TEXT, NUMERIC
from whoosh.index import create_in, open_dir
from whoosh.qparser import MultifieldParser
from jieba.analyse import ChineseAnalyzer

from .models import Article

log = logging.getLogger(__name__)

# 索引目录
index_dir = os.path.join(settings.BASE_DIR, "whoosh_index")

# 索引
indexer = open_dir(index_dir)


# 重建索引
def rebuild():
    if os.path.exists(index_dir):
        shutil.rmtree(index_dir)
    os.makedirs(index_dir)

    analyzer = ChineseAnalyzer()

    # 索引描述，类似数据库的字段
    schema = Schema(
        id=ID(stored=True, unique=True),
        slug=TEXT(stored=True),
        title=TEXT(),
        content=TEXT(analyzer=analyzer),
    )
    indexer = create_in(index_dir, schema)

    index_all_articles(indexer)


# 从数据库取出数据并创建索引
def index_all_articles(indexer):
    writer = indexer.writer()
    published_articles = Article.objects.exclude(is_draft=True)
    for article in published_articles:
        writer.add_document(
            id=str(article.id),
            slug=article.slug,
            title=article.title,
            content=article.content,
        )
    writer.commit()


# 更新索引
def article_update_index(article):
    """
  updating an article to indexer, adding if not.
  """
    writer = indexer.writer()
    writer.update_document(
        id=str(article.id),
        slug=article.slug,
        title=article.title,
        content=article.content,
    )

    writer.commit()


# 删除索引
def article_delete_index(article):
    writer = indexer.writer()
    writer.delete_by_term("id", str(article.id))

    writer.commit()


# 搜索文章
def articles_search(keyword):

    parser = MultifieldParser(
        ["content", "title"], schema=indexer.schema, fieldboosts={"title": 5.0}
    )
    query = parser.parse(keyword)

    with indexer.searcher() as searcher:
        results = searcher.search(query, limit=15)

        articles = []
        for hit in results:
            log.debug(hit)
            articles.append(
                {"id": hit["id"], "slug": hit["slug"],}
            )

    return articles

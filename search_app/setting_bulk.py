# search_app/setting_bulk.py

from elasticsearch import Elasticsearch

es = Elasticsearch()

# 리인덱스 파이프라인 생성
es.ingest.put_pipeline(
    id='reindex_pipeline',
    body={
        "description": "Reindexing pipeline with new analyzer",
        "processors": [
            {
                "set": {
                    "field": "_source",
                    "value": {
                        "Title": "{{Title}}"
                        # Add other fields 
                    }
                }
            }
        ]
    }
)

# 리인덱스 파이프라인 실행 
response = es.reindex(
    body={
        "source": {
            "index": "book_all"
        },
        "dest": {
            "index": "book_ngram",
            "pipeline": "reindex_pipeline"
        }
    },
    wait_for_completion=True,
    request_timeout=300
)

# 분석기 정의 및 인덱스 생성 
es.indices.create(
    index='book_ngram',
    body={
        "settings": {
            "analysis": {
                "tokenizer": {
                    "nori_tokenizer": {
                        "type": "nori_tokenizer"
                    }
                },
                "filter": {
                    "ngram_filter": {
                        "type": "ngram",
                        "min_gram": 1,
                        "max_gram": 20
                    }
                },
                "analyzer": {
                    "combined_analyzer": {
                        "type": "custom",
                        "tokenizer": "nori_tokenizer",
                        "filter": ["ngram_filter"]
                    }
                }
            }
        },
        "mappings": {
            "properties": {
                "Title": {
                    "type": "text",
                    "analyzer": "combined_analyzer"
                }
            }
        }
    }
)
